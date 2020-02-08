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
	bus             *BUS
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

type InstructionSet []Instruction

type Instruction struct {
	name         string
	operate      func() uint8
	address_mode AddressingFunction
	cycles       uint8
}

type AddressingFunction struct {
	fn   func() uint8
	Type string
}

func NewCPU() *CPU {
	cpu := &CPU{}
	cpu.Lookup = []Instruction{}

	return cpu
}

func (c *CPU) ConnectBus(bus *BUS) {
	c.bus = bus
}
