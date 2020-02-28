package nes

import (
	"fmt"
	"io/ioutil"

	"github.com/i82orbom/nesgo/pkg/nes/mappers"
	"github.com/i82orbom/nesgo/pkg/nes/mappers/mapper0"
	"github.com/i82orbom/nesgo/pkg/nes/mappers/mapper1"
)

// Cartridge represents a NES cartridge
type Cartridge struct {
	prgMemory []uint8
	chrMemory []uint8

	mirrorType mappers.MirroringType
	mapper     mappers.Mapper

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

// NewCartridge creates a new cartridge from the specified file
func NewCartridge(filePath string) (*Cartridge, error) {
	rom, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	header, offset := readHeader(rom)
	c := &Cartridge{}

	// mapperID
	c.mapperID = ((header.mapper2 >> 4) << 4) | (header.mapper1 >> 4)
	fileType := 1
	if header.mapper2&0x0C == 0x080 {
		fileType = 2
	}

	// mirroring type
	c.mirrorType = mappers.Horizontal
	if header.mapper1&0x01 != 0 {
		c.mirrorType = mappers.Vertical
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
			c.chrMemory = rom[offset : len(rom)-1]
		}
	} else if fileType == 2 { // Not tested yet
		c.prgBankCount = uint8(uint16(header.prgRAMSize&0x07)<<8) | header.prgChunks
		prgMemorySize := int(c.prgBankCount) * 16384
		// Copy PRG
		c.prgMemory = rom[offset : prgMemorySize+offset]

		// Offset is now at
		offset = prgMemorySize + offset

		c.chrBankCount = uint8(uint16(header.prgRAMSize&0x38)<<8) | header.chrChunks
		// Copy CHR
		c.chrMemory = rom[offset : len(rom)-1]
	}

	switch c.mapperID {
	case 0:
		c.mapper = mapper0.NewMapper(c.prgBankCount, c.chrBankCount)
	case 1:
		c.mapper = mapper1.NewMapper(c.prgBankCount, c.chrBankCount)
	default:
		return nil, fmt.Errorf("Mapper %v not implemented", c.mapperID)
	}

	return c, nil
}

// CPURead makes a PRG read using a mapper as intermediary
func (p *Cartridge) CPURead(address uint16, data *uint8) bool {
	if mappedAddress, readData := p.mapper.CPUMapRead(address); mappedAddress != nil {
		if readData != nil {
			*data = *readData
		}

		if *mappedAddress == 0xFFFFFFFF {
			return true
		}
		*data = p.prgMemory[*mappedAddress]
		return true
	}
	return false
}

// CPUWrite makes a PRG write using a mapper as intermediary
func (p *Cartridge) CPUWrite(address uint16, data uint8) bool {
	if mappedAddress, veto := p.mapper.CPUMapWrite(address, data); mappedAddress != nil {
		if *mappedAddress == 0xFFFFFFFF {
			return true
		}

		if !veto {
			p.prgMemory[*mappedAddress] = data
			return true
		}
	}
	return false
}

// PPUWrite makes a CHR write using a mapper as intermediary
func (p *Cartridge) PPUWrite(address uint16, data uint8) bool {
	if mappedAddress := p.mapper.PPUMapWrite(address, data); mappedAddress != nil {
		p.chrMemory[*mappedAddress] = data
		return true
	}
	return false
}

// PPURead makes a CHR read using a mapper as intermediary
func (p *Cartridge) PPURead(address uint16, data *uint8) bool {
	if mappedAddress := p.mapper.PPUMapRead(address); mappedAddress != nil {
		*data = p.chrMemory[*mappedAddress]
		return true
	}
	return false
}

// MirroringType returns the current mirroring type
func (p *Cartridge) MirroringType() mappers.MirroringType {
	if p.mapper.MirroringType() == mappers.Hardware {
		return p.mirrorType
	}
	return p.mapper.MirroringType()
}

// Reset resets the cartridge
func (p *Cartridge) Reset() {
	p.mapper.Reset()
}
