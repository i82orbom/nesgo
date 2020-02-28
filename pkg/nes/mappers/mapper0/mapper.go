package mapper0

import "github.com/i82orbom/nesgo/pkg/nes/mappers"

var _ mappers.Mapper = &mapper{}

type mapper struct {
	prgBanks uint8
	chrBanks uint8
}

// NewMapper creates a new Mapper 0
func NewMapper(prgBanks uint8, chrBanks uint8) mappers.Mapper {
	return &mapper{
		prgBanks: prgBanks,
		chrBanks: chrBanks,
	}
}

func (m *mapper) CPUMapRead(address uint16) (*uint32, *uint8) {
	if address >= 0x8000 && address <= 0xFFFF {
		temp := uint32(0x3FFF)
		if m.prgBanks > 1 {
			temp = 0x7FFF
		}
		newAddress := uint32(address) & temp
		return &newAddress, nil
	}
	return nil, nil
}

func (m *mapper) CPUMapWrite(address uint16, data uint8) (*uint32, bool) {
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

func (m *mapper) PPUMapRead(address uint16) *uint32 {
	if address < 0x2000 {
		return uint32Ptr(address)
	}
	return nil
}

func (m *mapper) PPUMapWrite(address uint16, data uint8) *uint32 {
	if address >= 0x0000 && address <= 0x1FFF {
		if m.chrBanks == 0 {
			return uint32Ptr(address)
		}
	}
	return nil
}

// Reset Resets the mapper
func (m *mapper) Reset() {
	// Does nothing here
}

// MirroringType Returns the current mirroring type if it is given by the mapper
func (m *mapper) MirroringType() mappers.MirroringType {
	return mappers.Hardware
}

// IRQState returns whether a IRQ was signaled by the mapper to the CPU
func (m *mapper) IRQState() bool {
	return false
}

// IRQClear clears the irq flag
func (m *mapper) IRQClear() {
	// Does nothing here
}

// NotifyScanline notifies the mapper that a scanline has occured
func (m *mapper) NotifyScanline() {
	// Does nothing here
}

func uint32Ptr(value uint16) *uint32 {
	v := uint32(value)
	return &v
}
