package nes

type FLAGS uint8

const (
	C FLAGS = (1 << 0)
	Z FLAGS = (1 << 1)
	I FLAGS = (1 << 2)
	D FLAGS = (1 << 3)
	B FLAGS = (1 << 4)
	U FLAGS = (1 << 5)
	V FLAGS = (1 << 6)
	N FLAGS = (1 << 7)
)

type CPU struct {
	bus             *Console
	a               uint8
	x               uint8
	y               uint8
	fetched         uint8
	status          uint8
	cycles          uint8
	opcode          uint8
	pc              uint16
	stackPointer    uint8
	addressAbsolute uint16
	addressRelative uint16
	Lookup          InstructionSet
	cyclesCounter   int
}

func NewCPU() *CPU {
	cpu := &CPU{}
	cpu.Lookup = createLookupTable(cpu)
	return cpu
}

func (c *CPU) ConnectBus(bus *Console) {
	c.bus = bus
}

func (c *CPU) Read(address uint16) uint8 {
	return c.bus.CPURead(address, false)
}

func (c *CPU) Write(address uint16, data uint8) {
	c.bus.CPUWrite(address, data)
}
