package mappers

type Mapper interface {
	// Maps the CPU bus addressed into PRG range
	CPUMapRead(address uint16) *uint32
	// Returns mapped address, and whether it should veto the write
	CPUMapWrite(address uint16, data uint8) (*uint32, bool)

	// Maps the PPU bus addresses into CHR range
	PPUMapRead(address uint16) *uint32
	PPUMapWrite(address uint16, data uint8) *uint32
}
