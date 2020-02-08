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