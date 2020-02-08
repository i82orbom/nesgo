package nes

type BUS struct {
	// Devices connected to the BUS
	cpu *CPU
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

func (b *BUS) CPURead(address uint16, readOnly bool) uint8 {
	return 0x0
}

func (b *BUS) CPUWrite(address uint16, data uint8) {

}
