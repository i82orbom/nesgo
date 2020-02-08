package mapper0

import "github.com/i82orbom/nesgo/pkg/nes/mappers"

var _ mappers.Mapper = &Mapper{}

type Mapper struct {
	prgBanks uint8
	chrBanks uint8
}

func NewMapper(prgBanks uint8, chrBanks uint8) mappers.Mapper {
	return &Mapper{
		prgBanks: prgBanks,
		chrBanks: chrBanks,
	}
}

func (m *Mapper) CPUMapRead(address uint16) *uint32 {
	if address >= 0x8000 && address <= 0xFFFF {
		temp := uint32(0x3FFF)
		if m.prgBanks > 1 {
			temp = 0x7FFF
		}
		newAddress := uint32(address) & temp
		return &newAddress
	}
	return nil
}

func (m *Mapper) CPUMapWrite(address uint16, data uint8) (*uint32, bool) {
	if address >= 0x8000 && address <= 0xFFFF {
		temp := uint32(0x3FFF)
		if m.prgBanks > 1 {
			temp = 0x7FFF
		}
		newAddress := uint32(address) & temp
		return &newAddress, false
	}
	return nil, false
}

func (m *Mapper) PPUMapRead(address uint16) *uint32 {
	if address < 0x2000 {
		return uint32Ptr(address)
	}
	return nil
}

func (m *Mapper) PPUMapWrite(address uint16, data uint8) *uint32 {
	if address >= 0x0000 && address <= 0x1FFF {
		if m.chrBanks == 0 {
			return uint32Ptr(address)
		}
	}
	return nil
}

func uint32Ptr(value uint16) *uint32 {
	v := uint32(value)
	return &v
}
