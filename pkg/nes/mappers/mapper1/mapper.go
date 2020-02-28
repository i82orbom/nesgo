package mapper1

import (
	"github.com/i82orbom/nesgo/pkg/nes/mappers"
)

var _ mappers.Mapper = &mapper{}

type mapper struct {
	mirroringMode mappers.MirroringType

	// PRG bank select (16k or 32k mode)
	prgBankSelect16LO uint8
	prgBankSelect16HI uint8
	prgBankSelect32   uint8

	// CHR bank select (4k or 8k mode)
	chrBankSelect4LO uint8
	chrBankSelect4HI uint8
	chrBankSelect8   uint8

	// Register
	loadRegister      uint8
	loadRegisterCount uint8
	controlRegister   uint8

	ram []uint8

	prgBanks uint8
	chrBanks uint8
}

// NewMapper creates a new Mapper 1
func NewMapper(prgBanks uint8, chrBanks uint8) mappers.Mapper {
	return &mapper{
		prgBanks: prgBanks,
		chrBanks: chrBanks,
		ram:      make([]uint8, 32*1024),
	}
}

func (m *mapper) CPUMapRead(address uint16) (*uint32, *uint8) {
	if address >= 0x6000 && address <= 0x7FFF {
		// Read is from static ram on cartridge
		newAddress := uint32(0xFFFFFFFF)

		// Read data from ram
		data := m.ram[newAddress&0x1FFF]

		return &newAddress, &data
	}
	if address >= 0x8000 {
		if m.controlRegister&0b01000 != 0 {

			// 16k mode
			if address >= 0x8000 && address <= 0xBFFF {
				newAddress := uint32(m.prgBankSelect16LO)*uint32(0x4000) + uint32(address&0x3FFF)
				return &newAddress, nil
			}

			if address >= 0xC000 && address <= 0xFFFF {
				newAddress := uint32(m.prgBankSelect16HI)*uint32(0x4000) + uint32(address&0x3FFF)
				return &newAddress, nil
			}
		} else {
			// 32K mode
			newAddress := uint32(m.prgBankSelect32)*uint32(0x8000) + uint32(address&0x7FFF)
			return &newAddress, nil
		}
	}

	return nil, nil
}

func (m *mapper) CPUMapWrite(address uint16, data uint8) (*uint32, bool) {
	if address >= 0x6000 && address <= 0x7FFF {
		newAddress := uint32(0xFFFFFFFF)

		// Write data to ram
		m.ram[address&0x1FFF] = data

		return &newAddress, true
	}

	if address >= 0x8000 {
		if data&0x80 != 0 {
			m.loadRegister = 0x00
			m.loadRegisterCount = 0
			m.controlRegister = m.controlRegister | 0x0C
		} else {
			m.loadRegister >>= 1
			m.loadRegister |= (data & 0x01) << 4
			m.loadRegisterCount++

			if m.loadRegisterCount == 5 {
				targetRegister := (address >> 13) & 0x03

				if targetRegister == 0 { // 0x8000 - 0x9FFF
					m.controlRegister = m.loadRegister & 0x1F

					switch m.controlRegister & 0x03 {
					case 0:
						m.mirroringMode = mappers.OneScreenLO
					case 1:
						m.mirroringMode = mappers.OneScreenHI
					case 2:
						m.mirroringMode = mappers.Vertical
					case 3:
						m.mirroringMode = mappers.Horizontal
					}
				} else if targetRegister == 1 { // 0xA000 - 0xBFFF

					if m.controlRegister&0b10000 != 0 {
						m.chrBankSelect4LO = m.loadRegister & 0x1F
					} else {
						m.chrBankSelect8 = m.loadRegister & 0x1E
					}
				} else if targetRegister == 2 { // 0xC000 - 0xDFFF
					if m.controlRegister&0b10000 != 0 {
						m.chrBankSelect4HI = m.loadRegister & 0x1F
					}
				} else if targetRegister == 3 { // 0xE000 - 0xFFFF
					prgMode := (m.controlRegister >> 2) & 0x03
					if prgMode == 0 || prgMode == 1 {
						// Set 32K prg bank at cpu 0x8000
						m.prgBankSelect32 = (m.loadRegister & 0x0E) >> 1
					} else if prgMode == 2 {
						m.prgBankSelect16LO = 0
						m.prgBankSelect16HI = m.loadRegister & 0x0F
					} else if prgMode == 3 {
						m.prgBankSelect16LO = m.loadRegister & 0x0F
						m.prgBankSelect16HI = m.prgBanks - 1
					}
				}

				m.loadRegister = 0x00
				m.loadRegisterCount = 0
			}
		}
		return nil, true
	}
	return nil, false
}

func (m *mapper) PPUMapRead(address uint16) *uint32 {
	if address < 0x2000 {
		if m.chrBanks == 0 {
			return uint32Ptr(address)
		}
		if m.controlRegister&0b10000 != 0 {
			// 4k CHR bank mode
			if address >= 0x0000 && address <= 0x0FFF {
				newAddress := uint32(m.chrBankSelect4LO)*uint32(0x1000) + uint32(address&0x0FFF)
				return &newAddress
			}

			if address >= 0x1000 && address <= 0x1FFF {
				newAddress := uint32(m.chrBankSelect4HI)*uint32(0x1000) + uint32(address&0x0FFF)
				return &newAddress
			}
		} else {
			newAddress := uint32(m.chrBankSelect8)&uint32(0x2000) + uint32(address&0x1FFF)
			return &newAddress
		}
	}
	return nil
}

func (m *mapper) PPUMapWrite(address uint16, data uint8) *uint32 {
	if address < 0x2000 {
		if m.chrBanks == 0 {
			return uint32Ptr(address)
		}
		return uint32Ptr(0x0000)
	}
	return nil
}

// Reset Resets the mapper
func (m *mapper) Reset() {
	m.controlRegister = 0x1C
	m.loadRegister = 0x00
	m.loadRegisterCount = 0x0

	m.chrBankSelect8 = 0
	m.chrBankSelect4LO = 0
	m.chrBankSelect4HI = 0

	m.prgBankSelect32 = 0
	m.prgBankSelect16LO = 0
	m.prgBankSelect16HI = m.prgBanks - 1
}

// MirroringType Returns the current mirroring type if it is given by the mapper
func (m *mapper) MirroringType() mappers.MirroringType {
	return m.mirroringMode
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
