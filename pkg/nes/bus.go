package nes

type BUS struct {
	// Console RAM
	ram [2048]uint8

	// Devices connected to the BUS
	cpu  *CPU
	cart *Cartridge
}

func NewBUS() *BUS {
	cpu := NewCPU()
	bus := &BUS{
		cpu: cpu,
	}

	// Connect devices
	cpu.ConnectBus(bus)

	return bus
}

func (b *BUS) InsertCartridge(c *Cartridge) {
	b.cart = c
}

func (b *BUS) Reset() {
	b.cpu.Reset()
}

func (b *BUS) Step() {
	for {
		b.cpu.Step()
		// The CPU will clock until no more cycles need to be executed
		if b.cpu.Complete() {
			break
		}
	}
}

func (b *BUS) CPURead(address uint16, readOnly bool) uint8 {
	data := uint8(0x00)
	// The cartridge can veto reads to other devices
	if b.cart.CPURead(address, &data) {
		return data
	}

	if address >= 0x0000 && address <= 0x1FFF {
		data = b.ram[address&0x07FF] // Mask for mirroring
	}

	// Read from other devices
	return data
}

func (b *BUS) CPUWrite(address uint16, data uint8) {
	// The cartridge can veto writes to other devices
	if b.cart.CPUWrite(address, data) {
		return
	}

	if address >= 0x0000 && address <= 0x1FFF {
		b.ram[address&0x07FF] = data // Mask for mirroring
	}
}
