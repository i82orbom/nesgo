package mappers

// Mapper represents a NES cartridge mapper
type Mapper interface {
	// Maps the CPU bus addressed into PRG range, sometimes it may return data
	CPUMapRead(address uint16) (*uint32, *uint8)
	// Returns mapped address, and whether it should veto the write
	CPUMapWrite(address uint16, data uint8) (*uint32, bool)

	// Maps the PPU bus addresses into CHR range
	PPUMapRead(address uint16) *uint32
	PPUMapWrite(address uint16, data uint8) *uint32

	// Resets the mapper
	Reset()

	// Returns the current mirroring type if it is given by the mapper
	MirroringType() MirroringType

	// IRQState returns whether a IRQ was signaled by the mapper to the CPU
	IRQState() bool

	// IRQClear clears the irq flag
	IRQClear()

	// NotifyScanline notifies the mapper that a scanline has occured
	NotifyScanline()
}

// MirroringType represents the type of mirroring for the cartridge
type MirroringType uint8

const (
	Hardware MirroringType = iota
	Horizontal
	Vertical
	OneScreenLO
	OneScreenHI
)
