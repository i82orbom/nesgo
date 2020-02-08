package nes

import (
	"fmt"
	"io/ioutil"

	"github.com/i82orbom/nesgo/pkg/nes/mappers"
	"github.com/i82orbom/nesgo/pkg/nes/mappers/mapper0"
)

type MirroringType uint8

const (
	OneScreenLO MirroringType = iota
	OneScreenHI
	Vertical
	Horizontal
)

type Cartridge struct {
	prgMemory []uint8
	chrMemory []uint8

	MirroringType
	mapper mappers.Mapper

	mapperID     uint8
	prgBankCount uint8
	chrBankCount uint8
}

type header struct {
	name       [4]byte
	prgChunks  uint8
	chrChunks  uint8
	mapper1    uint8
	mapper2    uint8
	prgRAMSize uint8
	tvSystem1  uint8
	tvSystem2  uint8
	unused     [5]byte
}

// Reads the header, returns it and the offset to skip training data if present
func readHeader(rom []byte) (header, int) {
	name := [4]byte{}
	for idx := range name {
		name[idx] = rom[idx]
	}
	h := header{
		name:       name,
		prgChunks:  uint8(rom[4]),
		chrChunks:  uint8(rom[5]),
		mapper1:    uint8(rom[6]),
		mapper2:    uint8(rom[7]),
		prgRAMSize: uint8(rom[8]),
		tvSystem1:  uint8(rom[9]),
		tvSystem2:  uint8(rom[10]),
	}
	unused := [5]byte{}
	offset := 0
	for idx := 11; idx <= 15; idx++ {
		unused[offset] = rom[idx]
		offset++
	}
	h.unused = unused
	offset = 16
	if (h.mapper1 & 0x04) == 1 {
		offset += 512
	}
	return h, offset
}

func NewCartridge(filePath string) (*Cartridge, error) {
	rom, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	header, offset := readHeader(rom)
	c := &Cartridge{}

	// mapperID
	c.mapperID = ((header.mapper2 >> 4) << 4) | (header.mapper1 >> 4)
	fileType := 1 // Asume type 1 always for now

	// mirroring type
	c.MirroringType = Horizontal
	if header.mapper1&0x01 != 0 {
		c.MirroringType = Vertical
	}

	if fileType == 1 {
		c.prgBankCount = header.prgChunks
		prgMemorySize := int(c.prgBankCount) * 16384
		// Copy PRG
		c.prgMemory = rom[offset : prgMemorySize+offset]

		// Update current offset
		offset = prgMemorySize + offset

		c.chrBankCount = header.chrChunks
		// Copy CHR
		chrMemorySize := int(c.chrBankCount) * 8192
		if chrMemorySize == 0 {
			c.chrMemory = make([]uint8, 8192)
		} else {
			c.chrMemory = rom[offset : chrMemorySize+offset]
		}
	}

	switch c.mapperID {
	case 0:
		c.mapper = mapper0.NewMapper(c.prgBankCount, c.chrBankCount)
	default:
		return nil, fmt.Errorf("Mapper %v not implemented", c.mapperID)
	}

	return c, nil
}

func (p *Cartridge) CPURead(address uint16, data *uint8) bool {
	if mappedAddress := p.mapper.CPUMapRead(address); mappedAddress != nil {
		*data = p.prgMemory[*mappedAddress]
		return true
	}
	return false
}

func (p *Cartridge) CPUWrite(address uint16, data uint8) bool {
	if mappedAddress, veto := p.mapper.CPUMapWrite(address, data); mappedAddress != nil {
		if !veto {
			p.prgMemory[*mappedAddress] = data
			return true
		}
	}
	return false
}

func (p *Cartridge) PPUWrite(address uint16, data uint8) bool {
	if mappedAddress := p.mapper.PPUMapWrite(address, data); mappedAddress != nil {
		p.chrMemory[*mappedAddress] = data
		return true
	}
	return false
}
func (p *Cartridge) PPURead(address uint16, data *uint8) bool {
	if mappedAddress := p.mapper.PPUMapRead(address); mappedAddress != nil {
		*data = p.chrMemory[*mappedAddress]
		return true
	}
	return false
}
