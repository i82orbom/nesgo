package nes

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

func createLookupTable(cpu *CPU) []Instruction {
	return []Instruction{
		{"BRK", cpu.BRK, AddressingFunction{cpu.IMP, "IMP"}, 7}, {"ORA", cpu.ORA, AddressingFunction{cpu.IZX, "IZX"}, 6}, {"???", cpu.XXX, AddressingFunction{cpu.IMP, "IMP"}, 2}, {"???", cpu.XXX, AddressingFunction{cpu.IMP, "IMP"}, 8}, {"???", cpu.NOP, AddressingFunction{cpu.IMP, "IMP"}, 3}, {"ORA", cpu.ORA, AddressingFunction{cpu.ZP0, "ZP0"}, 3}, {"ASL", cpu.ASL, AddressingFunction{cpu.ZP0, "ZP0"}, 5}, {"???", cpu.XXX, AddressingFunction{cpu.IMP, "IMP"}, 5}, {"PHP", cpu.PHP, AddressingFunction{cpu.IMP, "IMP"}, 3}, {"ORA", cpu.ORA, AddressingFunction{cpu.IMM, "IMM"}, 2}, {"ASL", cpu.ASL, AddressingFunction{cpu.IMP, "IMP"}, 2}, {"???", cpu.XXX, AddressingFunction{cpu.IMP, "IMP"}, 2}, {"???", cpu.NOP, AddressingFunction{cpu.IMP, "IMP"}, 4}, {"ORA", cpu.ORA, AddressingFunction{cpu.ABS, "ABS"}, 4}, {"ASL", cpu.ASL, AddressingFunction{cpu.ABS, "ABS"}, 6}, {"???", cpu.XXX, AddressingFunction{cpu.IMP, "IMP"}, 6},
		{"BPL", cpu.BPL, AddressingFunction{cpu.REL, "REL"}, 2}, {"ORA", cpu.ORA, AddressingFunction{cpu.IZY, "IZY"}, 5}, {"???", cpu.XXX, AddressingFunction{cpu.IMP, "IMP"}, 2}, {"???", cpu.XXX, AddressingFunction{cpu.IMP, "IMP"}, 8}, {"???", cpu.NOP, AddressingFunction{cpu.IMP, "IMP"}, 4}, {"ORA", cpu.ORA, AddressingFunction{cpu.ZPX, "ZPX"}, 4}, {"ASL", cpu.ASL, AddressingFunction{cpu.ZPX, "ZPX"}, 6}, {"???", cpu.XXX, AddressingFunction{cpu.IMP, "IMP"}, 6}, {"CLC", cpu.CLC, AddressingFunction{cpu.IMP, "IMP"}, 2}, {"ORA", cpu.ORA, AddressingFunction{cpu.ABY, "ABY"}, 4}, {"???", cpu.NOP, AddressingFunction{cpu.IMP, "IMP"}, 2}, {"???", cpu.XXX, AddressingFunction{cpu.IMP, "IMP"}, 7}, {"???", cpu.NOP, AddressingFunction{cpu.IMP, "IMP"}, 4}, {"ORA", cpu.ORA, AddressingFunction{cpu.ABX, "ABX"}, 4}, {"ASL", cpu.ASL, AddressingFunction{cpu.ABX, "ABX"}, 7}, {"???", cpu.XXX, AddressingFunction{cpu.IMP, "IMP"}, 7},
		{"JSR", cpu.JSR, AddressingFunction{cpu.ABS, "ABS"}, 6}, {"AND", cpu.AND, AddressingFunction{cpu.IZX, "IZX"}, 6}, {"???", cpu.XXX, AddressingFunction{cpu.IMP, "IMP"}, 2}, {"???", cpu.XXX, AddressingFunction{cpu.IMP, "IMP"}, 8}, {"BIT", cpu.BIT, AddressingFunction{cpu.ZP0, "ZP0"}, 3}, {"AND", cpu.AND, AddressingFunction{cpu.ZP0, "ZP0"}, 3}, {"ROL", cpu.ROL, AddressingFunction{cpu.ZP0, "ZP0"}, 5}, {"???", cpu.XXX, AddressingFunction{cpu.IMP, "IMP"}, 5}, {"PLP", cpu.PLP, AddressingFunction{cpu.IMP, "IMP"}, 4}, {"AND", cpu.AND, AddressingFunction{cpu.IMM, "IMM"}, 2}, {"ROL", cpu.ROL, AddressingFunction{cpu.IMP, "IMP"}, 2}, {"???", cpu.XXX, AddressingFunction{cpu.IMP, "IMP"}, 2}, {"BIT", cpu.BIT, AddressingFunction{cpu.ABS, "ABS"}, 4}, {"AND", cpu.AND, AddressingFunction{cpu.ABS, "ABS"}, 4}, {"ROL", cpu.ROL, AddressingFunction{cpu.ABS, "ABS"}, 6}, {"???", cpu.XXX, AddressingFunction{cpu.IMP, "IMP"}, 6},
		{"BMI", cpu.BMI, AddressingFunction{cpu.REL, "REL"}, 2}, {"AND", cpu.AND, AddressingFunction{cpu.IZY, "IZY"}, 5}, {"???", cpu.XXX, AddressingFunction{cpu.IMP, "IMP"}, 2}, {"???", cpu.XXX, AddressingFunction{cpu.IMP, "IMP"}, 8}, {"???", cpu.NOP, AddressingFunction{cpu.IMP, "IMP"}, 4}, {"AND", cpu.AND, AddressingFunction{cpu.ZPX, "ZPX"}, 4}, {"ROL", cpu.ROL, AddressingFunction{cpu.ZPX, "ZPX"}, 6}, {"???", cpu.XXX, AddressingFunction{cpu.IMP, "IMP"}, 6}, {"SEC", cpu.SEC, AddressingFunction{cpu.IMP, "IMP"}, 2}, {"AND", cpu.AND, AddressingFunction{cpu.ABY, "ABY"}, 4}, {"???", cpu.NOP, AddressingFunction{cpu.IMP, "IMP"}, 2}, {"???", cpu.XXX, AddressingFunction{cpu.IMP, "IMP"}, 7}, {"???", cpu.NOP, AddressingFunction{cpu.IMP, "IMP"}, 4}, {"AND", cpu.AND, AddressingFunction{cpu.ABX, "ABX"}, 4}, {"ROL", cpu.ROL, AddressingFunction{cpu.ABX, "ABX"}, 7}, {"???", cpu.XXX, AddressingFunction{cpu.IMP, "IMP"}, 7},
		{"RTI", cpu.RTI, AddressingFunction{cpu.IMP, "IMP"}, 6}, {"EOR", cpu.EOR, AddressingFunction{cpu.IZX, "IZX"}, 6}, {"???", cpu.XXX, AddressingFunction{cpu.IMP, "IMP"}, 2}, {"???", cpu.XXX, AddressingFunction{cpu.IMP, "IMP"}, 8}, {"???", cpu.NOP, AddressingFunction{cpu.IMP, "IMP"}, 3}, {"EOR", cpu.EOR, AddressingFunction{cpu.ZP0, "ZP0"}, 3}, {"LSR", cpu.LSR, AddressingFunction{cpu.ZP0, "ZP0"}, 5}, {"???", cpu.XXX, AddressingFunction{cpu.IMP, "IMP"}, 5}, {"PHA", cpu.PHA, AddressingFunction{cpu.IMP, "IMP"}, 3}, {"EOR", cpu.EOR, AddressingFunction{cpu.IMM, "IMM"}, 2}, {"LSR", cpu.LSR, AddressingFunction{cpu.IMP, "IMP"}, 2}, {"???", cpu.XXX, AddressingFunction{cpu.IMP, "IMP"}, 2}, {"JMP", cpu.JMP, AddressingFunction{cpu.ABS, "ABS"}, 3}, {"EOR", cpu.EOR, AddressingFunction{cpu.ABS, "ABS"}, 4}, {"LSR", cpu.LSR, AddressingFunction{cpu.ABS, "ABS"}, 6}, {"???", cpu.XXX, AddressingFunction{cpu.IMP, "IMP"}, 6},
		{"BVC", cpu.BVC, AddressingFunction{cpu.REL, "REL"}, 2}, {"EOR", cpu.EOR, AddressingFunction{cpu.IZY, "IZY"}, 5}, {"???", cpu.XXX, AddressingFunction{cpu.IMP, "IMP"}, 2}, {"???", cpu.XXX, AddressingFunction{cpu.IMP, "IMP"}, 8}, {"???", cpu.NOP, AddressingFunction{cpu.IMP, "IMP"}, 4}, {"EOR", cpu.EOR, AddressingFunction{cpu.ZPX, "ZPX"}, 4}, {"LSR", cpu.LSR, AddressingFunction{cpu.ZPX, "ZPX"}, 6}, {"???", cpu.XXX, AddressingFunction{cpu.IMP, "IMP"}, 6}, {"CLI", cpu.CLI, AddressingFunction{cpu.IMP, "IMP"}, 2}, {"EOR", cpu.EOR, AddressingFunction{cpu.ABY, "ABY"}, 4}, {"???", cpu.NOP, AddressingFunction{cpu.IMP, "IMP"}, 2}, {"???", cpu.XXX, AddressingFunction{cpu.IMP, "IMP"}, 7}, {"???", cpu.NOP, AddressingFunction{cpu.IMP, "IMP"}, 4}, {"EOR", cpu.EOR, AddressingFunction{cpu.ABX, "ABX"}, 4}, {"LSR", cpu.LSR, AddressingFunction{cpu.ABX, "ABX"}, 7}, {"???", cpu.XXX, AddressingFunction{cpu.IMP, "IMP"}, 7},
		{"RTS", cpu.RTS, AddressingFunction{cpu.IMP, "IMP"}, 6}, {"ADC", cpu.ADC, AddressingFunction{cpu.IZX, "IZX"}, 6}, {"???", cpu.XXX, AddressingFunction{cpu.IMP, "IMP"}, 2}, {"???", cpu.XXX, AddressingFunction{cpu.IMP, "IMP"}, 8}, {"???", cpu.NOP, AddressingFunction{cpu.IMP, "IMP"}, 3}, {"ADC", cpu.ADC, AddressingFunction{cpu.ZP0, "ZP0"}, 3}, {"ROR", cpu.ROR, AddressingFunction{cpu.ZP0, "ZP0"}, 5}, {"???", cpu.XXX, AddressingFunction{cpu.IMP, "IMP"}, 5}, {"PLA", cpu.PLA, AddressingFunction{cpu.IMP, "IMP"}, 4}, {"ADC", cpu.ADC, AddressingFunction{cpu.IMM, "IMM"}, 2}, {"ROR", cpu.ROR, AddressingFunction{cpu.IMP, "IMP"}, 2}, {"???", cpu.XXX, AddressingFunction{cpu.IMP, "IMP"}, 2}, {"JMP", cpu.JMP, AddressingFunction{cpu.IND, "IND"}, 5}, {"ADC", cpu.ADC, AddressingFunction{cpu.ABS, "ABS"}, 4}, {"ROR", cpu.ROR, AddressingFunction{cpu.ABS, "ABS"}, 6}, {"???", cpu.XXX, AddressingFunction{cpu.IMP, "IMP"}, 6},
		{"BVS", cpu.BVS, AddressingFunction{cpu.REL, "REL"}, 2}, {"ADC", cpu.ADC, AddressingFunction{cpu.IZY, "IZY"}, 5}, {"???", cpu.XXX, AddressingFunction{cpu.IMP, "IMP"}, 2}, {"???", cpu.XXX, AddressingFunction{cpu.IMP, "IMP"}, 8}, {"???", cpu.NOP, AddressingFunction{cpu.IMP, "IMP"}, 4}, {"ADC", cpu.ADC, AddressingFunction{cpu.ZPX, "ZPX"}, 4}, {"ROR", cpu.ROR, AddressingFunction{cpu.ZPX, "ZPX"}, 6}, {"???", cpu.XXX, AddressingFunction{cpu.IMP, "IMP"}, 6}, {"SEI", cpu.SEI, AddressingFunction{cpu.IMP, "IMP"}, 2}, {"ADC", cpu.ADC, AddressingFunction{cpu.ABY, "ABY"}, 4}, {"???", cpu.NOP, AddressingFunction{cpu.IMP, "IMP"}, 2}, {"???", cpu.XXX, AddressingFunction{cpu.IMP, "IMP"}, 7}, {"???", cpu.NOP, AddressingFunction{cpu.IMP, "IMP"}, 4}, {"ADC", cpu.ADC, AddressingFunction{cpu.ABX, "ABX"}, 4}, {"ROR", cpu.ROR, AddressingFunction{cpu.ABX, "ABX"}, 7}, {"???", cpu.XXX, AddressingFunction{cpu.IMP, "IMP"}, 7},
		{"???", cpu.NOP, AddressingFunction{cpu.IMP, "IMP"}, 2}, {"STA", cpu.STA, AddressingFunction{cpu.IZX, "IZX"}, 6}, {"???", cpu.NOP, AddressingFunction{cpu.IMP, "IMP"}, 2}, {"???", cpu.XXX, AddressingFunction{cpu.IMP, "IMP"}, 6}, {"STY", cpu.STY, AddressingFunction{cpu.ZP0, "ZP0"}, 3}, {"STA", cpu.STA, AddressingFunction{cpu.ZP0, "ZP0"}, 3}, {"STX", cpu.STX, AddressingFunction{cpu.ZP0, "ZP0"}, 3}, {"???", cpu.XXX, AddressingFunction{cpu.IMP, "IMP"}, 3}, {"DEY", cpu.DEY, AddressingFunction{cpu.IMP, "IMP"}, 2}, {"???", cpu.NOP, AddressingFunction{cpu.IMP, "IMP"}, 2}, {"TXA", cpu.TXA, AddressingFunction{cpu.IMP, "IMP"}, 2}, {"???", cpu.XXX, AddressingFunction{cpu.IMP, "IMP"}, 2}, {"STY", cpu.STY, AddressingFunction{cpu.ABS, "ABS"}, 4}, {"STA", cpu.STA, AddressingFunction{cpu.ABS, "ABS"}, 4}, {"STX", cpu.STX, AddressingFunction{cpu.ABS, "ABS"}, 4}, {"???", cpu.XXX, AddressingFunction{cpu.IMP, "IMP"}, 4},
		{"BCC", cpu.BCC, AddressingFunction{cpu.REL, "REL"}, 2}, {"STA", cpu.STA, AddressingFunction{cpu.IZY, "IZY"}, 6}, {"???", cpu.XXX, AddressingFunction{cpu.IMP, "IMP"}, 2}, {"???", cpu.XXX, AddressingFunction{cpu.IMP, "IMP"}, 6}, {"STY", cpu.STY, AddressingFunction{cpu.ZPX, "ZPX"}, 4}, {"STA", cpu.STA, AddressingFunction{cpu.ZPX, "ZPX"}, 4}, {"STX", cpu.STX, AddressingFunction{cpu.ZPY, "ZPY"}, 4}, {"???", cpu.XXX, AddressingFunction{cpu.IMP, "IMP"}, 4}, {"TYA", cpu.TYA, AddressingFunction{cpu.IMP, "IMP"}, 2}, {"STA", cpu.STA, AddressingFunction{cpu.ABY, "ABY"}, 5}, {"TXS", cpu.TXS, AddressingFunction{cpu.IMP, "IMP"}, 2}, {"???", cpu.XXX, AddressingFunction{cpu.IMP, "IMP"}, 5}, {"???", cpu.NOP, AddressingFunction{cpu.IMP, "IMP"}, 5}, {"STA", cpu.STA, AddressingFunction{cpu.ABX, "ABX"}, 5}, {"???", cpu.XXX, AddressingFunction{cpu.IMP, "IMP"}, 5}, {"???", cpu.XXX, AddressingFunction{cpu.IMP, "IMP"}, 5},
		{"LDY", cpu.LDY, AddressingFunction{cpu.IMM, "IMM"}, 2}, {"LDA", cpu.LDA, AddressingFunction{cpu.IZX, "IZX"}, 6}, {"LDX", cpu.LDX, AddressingFunction{cpu.IMM, "IMM"}, 2}, {"???", cpu.XXX, AddressingFunction{cpu.IMP, "IMP"}, 6}, {"LDY", cpu.LDY, AddressingFunction{cpu.ZP0, "ZP0"}, 3}, {"LDA", cpu.LDA, AddressingFunction{cpu.ZP0, "ZP0"}, 3}, {"LDX", cpu.LDX, AddressingFunction{cpu.ZP0, "ZP0"}, 3}, {"???", cpu.XXX, AddressingFunction{cpu.IMP, "IMP"}, 3}, {"TAY", cpu.TAY, AddressingFunction{cpu.IMP, "IMP"}, 2}, {"LDA", cpu.LDA, AddressingFunction{cpu.IMM, "IMM"}, 2}, {"TAX", cpu.TAX, AddressingFunction{cpu.IMP, "IMP"}, 2}, {"???", cpu.XXX, AddressingFunction{cpu.IMP, "IMP"}, 2}, {"LDY", cpu.LDY, AddressingFunction{cpu.ABS, "ABS"}, 4}, {"LDA", cpu.LDA, AddressingFunction{cpu.ABS, "ABS"}, 4}, {"LDX", cpu.LDX, AddressingFunction{cpu.ABS, "ABS"}, 4}, {"???", cpu.XXX, AddressingFunction{cpu.IMP, "IMP"}, 4},
		{"BCS", cpu.BCS, AddressingFunction{cpu.REL, "REL"}, 2}, {"LDA", cpu.LDA, AddressingFunction{cpu.IZY, "IZY"}, 5}, {"???", cpu.XXX, AddressingFunction{cpu.IMP, "IMP"}, 2}, {"???", cpu.XXX, AddressingFunction{cpu.IMP, "IMP"}, 5}, {"LDY", cpu.LDY, AddressingFunction{cpu.ZPX, "ZPX"}, 4}, {"LDA", cpu.LDA, AddressingFunction{cpu.ZPX, "ZPX"}, 4}, {"LDX", cpu.LDX, AddressingFunction{cpu.ZPY, "ZPY"}, 4}, {"???", cpu.XXX, AddressingFunction{cpu.IMP, "IMP"}, 4}, {"CLV", cpu.CLV, AddressingFunction{cpu.IMP, "IMP"}, 2}, {"LDA", cpu.LDA, AddressingFunction{cpu.ABY, "ABY"}, 4}, {"TSX", cpu.TSX, AddressingFunction{cpu.IMP, "IMP"}, 2}, {"???", cpu.XXX, AddressingFunction{cpu.IMP, "IMP"}, 4}, {"LDY", cpu.LDY, AddressingFunction{cpu.ABX, "ABX"}, 4}, {"LDA", cpu.LDA, AddressingFunction{cpu.ABX, "ABX"}, 4}, {"LDX", cpu.LDX, AddressingFunction{cpu.ABY, "ABY"}, 4}, {"???", cpu.XXX, AddressingFunction{cpu.IMP, "IMP"}, 4},
		{"CPY", cpu.CPY, AddressingFunction{cpu.IMM, "IMM"}, 2}, {"CMP", cpu.CMP, AddressingFunction{cpu.IZX, "IZX"}, 6}, {"???", cpu.NOP, AddressingFunction{cpu.IMP, "IMP"}, 2}, {"???", cpu.XXX, AddressingFunction{cpu.IMP, "IMP"}, 8}, {"CPY", cpu.CPY, AddressingFunction{cpu.ZP0, "ZP0"}, 3}, {"CMP", cpu.CMP, AddressingFunction{cpu.ZP0, "ZP0"}, 3}, {"DEC", cpu.DEC, AddressingFunction{cpu.ZP0, "ZP0"}, 5}, {"???", cpu.XXX, AddressingFunction{cpu.IMP, "IMP"}, 5}, {"INY", cpu.INY, AddressingFunction{cpu.IMP, "IMP"}, 2}, {"CMP", cpu.CMP, AddressingFunction{cpu.IMM, "IMM"}, 2}, {"DEX", cpu.DEX, AddressingFunction{cpu.IMP, "IMP"}, 2}, {"???", cpu.XXX, AddressingFunction{cpu.IMP, "IMP"}, 2}, {"CPY", cpu.CPY, AddressingFunction{cpu.ABS, "ABS"}, 4}, {"CMP", cpu.CMP, AddressingFunction{cpu.ABS, "ABS"}, 4}, {"DEC", cpu.DEC, AddressingFunction{cpu.ABS, "ABS"}, 6}, {"???", cpu.XXX, AddressingFunction{cpu.IMP, "IMP"}, 6},
		{"BNE", cpu.BNE, AddressingFunction{cpu.REL, "REL"}, 2}, {"CMP", cpu.CMP, AddressingFunction{cpu.IZY, "IZY"}, 5}, {"???", cpu.XXX, AddressingFunction{cpu.IMP, "IMP"}, 2}, {"???", cpu.XXX, AddressingFunction{cpu.IMP, "IMP"}, 8}, {"???", cpu.NOP, AddressingFunction{cpu.IMP, "IMP"}, 4}, {"CMP", cpu.CMP, AddressingFunction{cpu.ZPX, "ZPX"}, 4}, {"DEC", cpu.DEC, AddressingFunction{cpu.ZPX, "ZPX"}, 6}, {"???", cpu.XXX, AddressingFunction{cpu.IMP, "IMP"}, 6}, {"CLD", cpu.CLD, AddressingFunction{cpu.IMP, "IMP"}, 2}, {"CMP", cpu.CMP, AddressingFunction{cpu.ABY, "ABY"}, 4}, {"NOP", cpu.NOP, AddressingFunction{cpu.IMP, "IMP"}, 2}, {"???", cpu.XXX, AddressingFunction{cpu.IMP, "IMP"}, 7}, {"???", cpu.NOP, AddressingFunction{cpu.IMP, "IMP"}, 4}, {"CMP", cpu.CMP, AddressingFunction{cpu.ABX, "ABX"}, 4}, {"DEC", cpu.DEC, AddressingFunction{cpu.ABX, "ABX"}, 7}, {"???", cpu.XXX, AddressingFunction{cpu.IMP, "IMP"}, 7},
		{"CPX", cpu.CPX, AddressingFunction{cpu.IMM, "IMM"}, 2}, {"SBC", cpu.SBC, AddressingFunction{cpu.IZX, "IZX"}, 6}, {"???", cpu.NOP, AddressingFunction{cpu.IMP, "IMP"}, 2}, {"???", cpu.XXX, AddressingFunction{cpu.IMP, "IMP"}, 8}, {"CPX", cpu.CPX, AddressingFunction{cpu.ZP0, "ZP0"}, 3}, {"SBC", cpu.SBC, AddressingFunction{cpu.ZP0, "ZP0"}, 3}, {"INC", cpu.INC, AddressingFunction{cpu.ZP0, "ZP0"}, 5}, {"???", cpu.XXX, AddressingFunction{cpu.IMP, "IMP"}, 5}, {"INX", cpu.INX, AddressingFunction{cpu.IMP, "IMP"}, 2}, {"SBC", cpu.SBC, AddressingFunction{cpu.IMM, "IMM"}, 2}, {"NOP", cpu.NOP, AddressingFunction{cpu.IMP, "IMP"}, 2}, {"???", cpu.SBC, AddressingFunction{cpu.IMP, "IMP"}, 2}, {"CPX", cpu.CPX, AddressingFunction{cpu.ABS, "ABS"}, 4}, {"SBC", cpu.SBC, AddressingFunction{cpu.ABS, "ABS"}, 4}, {"INC", cpu.INC, AddressingFunction{cpu.ABS, "ABS"}, 6}, {"???", cpu.XXX, AddressingFunction{cpu.IMP, "IMP"}, 6},
		{"BEQ", cpu.BEQ, AddressingFunction{cpu.REL, "REL"}, 2}, {"SBC", cpu.SBC, AddressingFunction{cpu.IZY, "IZY"}, 5}, {"???", cpu.XXX, AddressingFunction{cpu.IMP, "IMP"}, 2}, {"???", cpu.XXX, AddressingFunction{cpu.IMP, "IMP"}, 8}, {"???", cpu.NOP, AddressingFunction{cpu.IMP, "IMP"}, 4}, {"SBC", cpu.SBC, AddressingFunction{cpu.ZPX, "ZPX"}, 4}, {"INC", cpu.INC, AddressingFunction{cpu.ZPX, "ZPX"}, 6}, {"???", cpu.XXX, AddressingFunction{cpu.IMP, "IMP"}, 6}, {"SED", cpu.SED, AddressingFunction{cpu.IMP, "IMP"}, 2}, {"SBC", cpu.SBC, AddressingFunction{cpu.ABY, "ABY"}, 4}, {"NOP", cpu.NOP, AddressingFunction{cpu.IMP, "IMP"}, 2}, {"???", cpu.XXX, AddressingFunction{cpu.IMP, "IMP"}, 7}, {"???", cpu.NOP, AddressingFunction{cpu.IMP, "IMP"}, 4}, {"SBC", cpu.SBC, AddressingFunction{cpu.ABX, "ABX"}, 4}, {"INC", cpu.INC, AddressingFunction{cpu.ABX, "ABX"}, 7}, {"???", cpu.XXX, AddressingFunction{cpu.IMP, "IMP"}, 7},
	}
}

// ADD
func (c *CPU) ADC() uint8 {
	c.fetch()
	flagC := c.GetFlag(C)
	temp := uint16(c.a) + uint16(c.fetched) + uint16(flagC)
	c.SetFlag(C, temp > 255)
	c.SetFlag(Z, (temp&0x00FF) == 0)
	c.SetFlagU(N, int(temp&0x80))
	c.SetFlagU(V, int(^(uint16(c.a)^uint16(c.fetched))&(uint16(c.a)^temp))&0x0080)
	c.a = uint8(temp & 0x00FF)
	return 1
}

// AND operation
func (c *CPU) AND() uint8 {
	c.fetch()
	c.a = c.a & c.fetched
	c.SetFlag(Z, c.a == 0x00)
	c.SetFlagU(N, int(c.a&0x80)) // If bit 7 is 1
	return 1
}

// Arithmetic shift left
func (c *CPU) ASL() uint8 {
	c.fetch()
	temp := uint16(c.fetched) << 1
	c.SetFlag(C, (temp&0xFF00) > 0)
	c.SetFlag(Z, (temp&0x00FF) == 0x00)
	c.SetFlagU(N, int(temp&0x80))
	if c.Lookup[c.opcode].address_mode.Type == "IMP" {
		c.a = uint8(temp & 0x00FF)
	} else {
		c.Write(c.addressAbsolute, uint8(temp&0x00FF))
	}
	return 0
}

// Branch if carry clear
func (c *CPU) BCC() uint8 {
	if c.GetFlag(C) == 0 {
		c.cycles++
		c.addressAbsolute = c.pc + c.addressRelative

		if (c.addressAbsolute & 0xFF00) != (c.pc & 0xFF00) {
			c.cycles++
		}
		c.pc = c.addressAbsolute // New location
	}
	return 0
}

// Branch if carry status is set
func (c *CPU) BCS() uint8 {
	if c.GetFlag(C) == 1 {
		c.cycles++
		c.addressAbsolute = c.pc + c.addressRelative

		if (c.addressAbsolute & 0xFF00) != (c.pc & 0xFF00) {
			c.cycles++
		}

		c.pc = c.addressAbsolute // New location
	}
	return 0
}

// Branch if equal
func (c *CPU) BEQ() uint8 {
	if c.GetFlag(Z) != 0 {
		c.cycles++
		c.addressAbsolute = c.pc + c.addressRelative

		if (c.addressAbsolute & 0xFF00) != (c.pc & 0xFF00) {
			c.cycles++
		}

		c.pc = c.addressAbsolute // New location
	}
	return 0
}

func (c *CPU) BIT() uint8 {
	c.fetch()
	temp := c.a & c.fetched
	c.SetFlag(Z, (temp&0x00FF) == 0x00)
	c.SetFlagU(N, int(c.fetched&(1<<7)))
	c.SetFlagU(V, int(c.fetched&(1<<6)))
	return 0
}

// Branch if negative
func (c *CPU) BMI() uint8 {
	if c.GetFlag(N) == 1 {
		c.cycles++
		c.addressAbsolute = c.pc + c.addressRelative

		if (c.addressAbsolute & 0xFF00) != (c.pc & 0xFF00) {
			c.cycles++
		}

		c.pc = c.addressAbsolute // New location
	}
	return 0
}

// Branch if not equal
func (c *CPU) BNE() uint8 {
	if c.GetFlag(Z) == 0 {
		c.cycles++
		c.addressAbsolute = c.pc + c.addressRelative

		if (c.addressAbsolute & 0xFF00) != (c.pc & 0xFF00) {
			c.cycles++
		}

		c.pc = c.addressAbsolute // New location
	}
	return 0
}

// Branch if positive
func (c *CPU) BPL() uint8 {
	if c.GetFlag(N) == 0 {
		c.cycles++
		c.addressAbsolute = c.pc + c.addressRelative

		if (c.addressAbsolute & 0xFF00) != (c.pc & 0xFF00) {
			c.cycles++
		}
		c.pc = c.addressAbsolute // New location
	}
	return 0
}

// Break (software interrupt)
func (c *CPU) BRK() uint8 {
	c.pc++

	c.SetFlagU(I, 1)
	c.Write(0x0100+uint16(c.stackPointer), uint8((c.pc>>8)&0x00FF))
	c.stackPointer--
	c.Write(0x0100+uint16(c.stackPointer), uint8(c.pc&0x00FF))
	c.stackPointer--

	c.SetFlagU(B, 1)
	c.Write(0x0100+uint16(c.stackPointer), c.status)
	c.stackPointer--
	c.SetFlagU(B, 0)

	c.pc = uint16(c.Read(0xFFFE)) | (uint16(c.Read(0xFFFF)) << 8)
	return 0
}

// Branch if overflow
func (c *CPU) BVC() uint8 {
	if c.GetFlag(V) == 0 {
		c.cycles++
		c.addressAbsolute = c.pc + c.addressRelative

		if (c.addressAbsolute & 0xFF00) != (c.pc & 0xFF00) {
			c.cycles++
		}

		c.pc = c.addressAbsolute // New location
	}
	return 0
}

// Branch if not-overflow
func (c *CPU) BVS() uint8 {
	if c.GetFlag(V) == 1 {
		c.cycles++
		c.addressAbsolute = c.pc + c.addressRelative

		if (c.addressAbsolute & 0xFF00) != (c.pc & 0xFF00) {
			c.cycles++
		}
		c.pc = c.addressAbsolute // New location
	}
	return 0
}

func (c *CPU) CLC() uint8 {
	c.SetFlag(C, false)
	return 0
}

func (c *CPU) CLD() uint8 {
	c.SetFlag(D, false)
	return 0
}

func (c *CPU) CLI() uint8 {
	c.SetFlag(I, false)
	return 0
}

func (c *CPU) CLV() uint8 {
	c.SetFlag(V, false)
	return 0
}

// Compare Accumulator
func (c *CPU) CMP() uint8 {
	c.fetch()
	temp := uint16(c.a) - uint16(c.fetched)
	c.SetFlag(C, c.a >= c.fetched)
	c.SetFlag(Z, (temp&0x00FF) == 0x0000)
	c.SetFlagU(N, int(temp&0x0080))
	return 1
}

// Compare X Register
func (c *CPU) CPX() uint8 {
	c.fetch()
	temp := uint16(c.x) - uint16(c.fetched)
	c.SetFlag(C, c.x >= c.fetched)
	c.SetFlag(Z, (temp&0x00FF) == 0x0000)
	c.SetFlagU(N, int(temp&0x0080))
	return 0
}

// Compare Y Register
func (c *CPU) CPY() uint8 {
	c.fetch()
	temp := uint16(c.y) - uint16(c.fetched)
	c.SetFlag(C, c.y >= c.fetched)
	c.SetFlag(Z, (temp&0x00FF) == 0x0000)
	c.SetFlagU(N, int(temp&0x0080))
	return 0

}

// Decrement value at memory location
func (c *CPU) DEC() uint8 {
	c.fetch()
	temp := uint16(c.fetched) - 1
	c.Write(c.addressAbsolute, uint8(temp&0x00FF))
	c.SetFlag(Z, (temp&0x00FF) == 0x0000)
	c.SetFlagU(N, int(temp&0x0080))
	return 0
}

// Decrement X
func (c *CPU) DEX() uint8 {
	c.x--
	c.SetFlag(Z, c.x == 0x00)
	c.SetFlagU(N, int(c.x&0x80))
	return 0
}

// Decrement Y
func (c *CPU) DEY() uint8 {
	c.y--
	c.SetFlag(Z, c.y == 0x00)
	c.SetFlagU(N, int(c.y&0x80))
	return 0
}

// Logic XOR
func (c *CPU) EOR() uint8 {
	c.fetch()
	c.a = c.a ^ c.fetched
	c.SetFlag(Z, c.a == 0x00)
	c.SetFlagU(N, int(c.a&0x80))
	return 1
}

// Increment value at memory location
func (c *CPU) INC() uint8 {
	c.fetch()
	temp := uint16(c.fetched) + 1
	c.Write(c.addressAbsolute, uint8(temp&0x00FF))
	c.SetFlag(Z, (temp&0x00FF) == 0x0000)
	c.SetFlagU(N, int(temp&0x0080))
	return 0
}

func (c *CPU) INX() uint8 {
	c.x++
	c.SetFlag(Z, c.x == 0x0)
	c.SetFlagU(N, int(c.x&0x80))
	return 0
}

func (c *CPU) INY() uint8 {
	c.y++
	c.SetFlag(Z, c.y == 0x0)
	c.SetFlagU(N, int(c.y&0x80))
	return 0
}

func (c *CPU) JMP() uint8 {
	c.pc = c.addressAbsolute
	return 0
}

// Jump to subroutine
func (c *CPU) JSR() uint8 {
	c.pc--
	c.Write(0x0100+uint16(c.stackPointer), uint8((c.pc>>8)&0x00FF))
	c.stackPointer--
	c.Write(0x0100+uint16(c.stackPointer), uint8(c.pc&0x00FF))
	c.stackPointer--

	c.pc = c.addressAbsolute
	return 0
}

func (c *CPU) LDA() uint8 {
	c.fetch()
	c.a = c.fetched
	c.SetFlag(Z, c.a == 0x00)
	c.SetFlagU(N, int(c.a&0x80))
	return 1
}

func (c *CPU) LDX() uint8 {
	c.fetch()
	c.x = c.fetched
	c.SetFlag(Z, c.x == 0x00)
	c.SetFlagU(N, int(c.x&0x80))
	return 1
}

func (c *CPU) LDY() uint8 {
	c.fetch()
	c.y = c.fetched
	c.SetFlag(Z, c.y == 0x00)
	c.SetFlagU(N, int(c.y&0x80))
	return 1
}

func (c *CPU) LSR() uint8 {
	c.fetch()
	c.SetFlagU(C, int(c.fetched&0x0001))
	temp := uint16(c.fetched) >> 1
	c.SetFlag(Z, (temp&0x00FF) == 0x0000)
	c.SetFlagU(N, int(temp&0x0080))
	if c.Lookup[c.opcode].address_mode.Type == "IMP" {
		c.a = uint8(temp & 0x00FF)
	} else {
		c.Write(c.addressAbsolute, uint8(temp&0x00FF))
	}
	return 0
}

func (c *CPU) NOP() uint8 {
	// Not all Nops are the same
	switch c.opcode {
	case 0x1C:
	case 0x3C:
	case 0x5C:
	case 0x7C:
	case 0xDC:
	case 0xFC:
		return 1
	}
	return 0
}

// Bitwise OR
func (c *CPU) ORA() uint8 {
	c.fetch()
	c.a = c.a | c.fetched
	c.SetFlag(Z, c.a == 0x00)
	c.SetFlagU(N, int(c.a&0x80))
	return 1
}

// Push A to stack
func (c *CPU) PHA() uint8 {
	c.Write(0x0100+uint16(c.stackPointer), c.a)
	c.stackPointer--
	return 0
}

// Push status register to stack
func (c *CPU) PHP() uint8 {
	c.Write(0x0100+uint16(c.stackPointer), c.status|uint8(B)|uint8(U))
	c.SetFlagU(B, 0)
	c.SetFlagU(U, 0)
	c.stackPointer--
	return 0
}

// Pop A from stack
func (c *CPU) PLA() uint8 {
	c.stackPointer++
	c.a = c.Read(0x0100 + uint16(c.stackPointer))
	c.SetFlag(Z, c.a == 0x00)
	c.SetFlagU(N, int(c.a&0x80))
	return 0
}

// Pop accumulator off Stack
func (c *CPU) PLP() uint8 {
	c.stackPointer++
	c.status = c.Read(0x0100 + uint16(c.stackPointer))
	c.SetFlagU(U, 1)
	return 0
}
func (c *CPU) ROL() uint8 {
	c.fetch()
	temp := (uint16(c.fetched) << 1) | uint16(c.GetFlag(C))
	c.SetFlagU(C, int(temp&0xFF00))
	c.SetFlag(Z, (temp&0x00FF) == 0x0000)
	c.SetFlagU(N, int(temp&0x0080))
	if c.Lookup[c.opcode].address_mode.Type == "IMP" {
		c.a = uint8(temp & 0x00FF)
	} else {
		c.Write(c.addressAbsolute, uint8(temp&0x00FF))
	}
	return 0
}

func (c *CPU) ROR() uint8 {
	c.fetch()
	temp := (uint16(c.GetFlag(C)) << 7) | (uint16(c.fetched) >> 1)
	c.SetFlagU(C, int(c.fetched&0x01))
	c.SetFlag(Z, (temp&0x00FF) == 0x00)
	c.SetFlagU(N, int(temp&0x0080))
	if c.Lookup[c.opcode].address_mode.Type == "IMP" {
		c.a = uint8(temp & 0x00FF)
	} else {
		c.Write(c.addressAbsolute, uint8(temp&0x00FF))
	}
	return 0
}

func (c *CPU) RTI() uint8 {
	c.stackPointer++
	c.status = c.Read(0x0100 + uint16(c.stackPointer))
	c.status &= uint8(^B)
	c.status &= uint8(^U)

	c.stackPointer++
	c.pc = uint16(c.Read(0x0100 + uint16(c.stackPointer)))
	c.stackPointer++
	readValue := c.Read(0x0100 + uint16(c.stackPointer))
	c.pc |= uint16(readValue) << 8
	return 0
}

func (c *CPU) RTS() uint8 {
	c.stackPointer++
	c.pc = uint16(c.Read(0x0100 + uint16(c.stackPointer)))
	c.stackPointer++
	readValue := c.Read(0x0100 + uint16(c.stackPointer))
	c.pc |= uint16(readValue) << 8
	c.pc++
	return 0
}

func (c *CPU) SBC() uint8 {
	c.fetch()
	value := (uint16(c.fetched)) ^ 0x00FF // Same as ADC but inverting the fetched

	temp := uint16(c.a) + value + uint16(c.GetFlag(C))
	c.SetFlagU(C, int(temp&0xFF00))
	c.SetFlag(Z, (temp&0x00FF) == 0)
	c.SetFlagU(V, int((temp^(uint16(c.a)))&(temp^value)&0x0080))
	c.SetFlagU(N, int(temp&0x0080))
	c.a = uint8(temp & 0x00FF)
	return 1
}

func (c *CPU) SEC() uint8 {
	c.SetFlag(C, true)
	return 0
}

func (c *CPU) SED() uint8 {
	c.SetFlag(D, true)
	return 0
}

func (c *CPU) SEI() uint8 {
	c.SetFlag(I, true)
	return 0
}

func (c *CPU) STA() uint8 {
	c.Write(c.addressAbsolute, c.a)
	return 0
}

func (c *CPU) STX() uint8 {
	c.Write(c.addressAbsolute, c.x)
	return 0
}

func (c *CPU) STY() uint8 {
	c.Write(c.addressAbsolute, c.y)
	return 0
}

func (c *CPU) TAX() uint8 {
	c.x = c.a
	c.SetFlag(Z, c.x == 0x00)
	c.SetFlagU(N, int(c.x&0x80))
	return 0
}

func (c *CPU) TAY() uint8 {
	c.y = c.a
	c.SetFlag(Z, c.y == 0x00)
	c.SetFlagU(N, int(c.y&0x80))
	return 0
}

func (c *CPU) TSX() uint8 {
	c.x = c.stackPointer
	c.SetFlag(Z, c.x == 0x00)
	c.SetFlagU(N, int(c.x&0x80))
	return 0
}

func (c *CPU) TXA() uint8 {
	c.a = c.x
	c.SetFlag(Z, c.a == 0x00)
	c.SetFlagU(N, int(c.a&0x80))
	return 0
}

func (c *CPU) TXS() uint8 {
	c.stackPointer = c.x
	return 0
}

func (c *CPU) TYA() uint8 {
	c.a = c.y
	c.SetFlag(Z, c.a == 0x00)
	c.SetFlagU(N, int(c.a&0x80))
	return 0
}

func (c *CPU) XXX() uint8 {
	return 0
}

// Addressing modes
func (c *CPU) IMP() uint8 {
	c.fetched = c.a
	return 0
}

func (c *CPU) IMM() uint8 { // Immediate addressing
	c.addressAbsolute = c.pc
	c.pc++
	return 0
}

func (c *CPU) ZP0() uint8 { // Zero page addressing
	c.addressAbsolute = uint16(c.Read(c.pc))
	c.pc++
	c.addressAbsolute &= 0x00FF
	return 0
}

func (c *CPU) ZPX() uint8 { // Zero page addressing X offset
	c.addressAbsolute = uint16(c.Read(c.pc)) + uint16(c.x)
	c.pc++
	c.addressAbsolute &= 0x00FF
	return 0
}

func (c *CPU) ZPY() uint8 { // Zero page addressing Y offset
	c.addressAbsolute = uint16(c.Read(c.pc)) + uint16(c.y)
	c.pc++
	c.addressAbsolute &= 0x00FF
	return 0
}

func (c *CPU) ABS() uint8 {
	// Low byte
	lo := uint16(c.Read(c.pc))
	c.pc++
	// High byte
	hi := uint16(c.Read(c.pc))
	c.pc++

	c.addressAbsolute = (hi << 8) | lo

	return 0
}

func (c *CPU) ABX() uint8 { // Same as absolute but offset with X
	// Low byte
	lo := uint16(c.Read(c.pc))
	c.pc++
	// High byte
	hi := uint16(c.Read(c.pc))
	c.pc++

	c.addressAbsolute = ((hi << 8) | lo) + uint16(c.x)

	if (c.addressAbsolute & 0xFF00) != (hi << 8) { // If it is changed to a different page... we may need an additional clock cycle
		return 1
	}

	return 0
}

func (c *CPU) ABY() uint8 { // Same as ABX but with Y
	// Low byte
	lo := uint16(c.Read(c.pc))
	c.pc++
	// High byte
	hi := uint16(c.Read(c.pc))
	c.pc++

	c.addressAbsolute = uint16(((hi << 8) | lo) + uint16(c.y))

	if (c.addressAbsolute & 0xFF00) != (hi << 8) { // If it is changed to a different page... we may need an additional clock cycle
		return 1
	}

	return 0
}

func (c *CPU) IND() uint8 {
	ptrLO := uint16(c.Read(c.pc))
	c.pc++
	ptrHI := uint16(c.Read(c.pc))
	c.pc++

	ptr := (ptrHI << 8) | ptrLO

	if ptrLO == 0x00FF { // Simulate page boundary hw bug
		readValue := uint16(c.Read(ptr&0xFF00)) << 8
		c.addressAbsolute = uint16(readValue | uint16(c.Read(ptr+0)))
	} else { // Normal flow
		readValue := uint16(c.Read(ptr+1)) << 8
		c.addressAbsolute = uint16(readValue | uint16(c.Read(ptr+0)))
	}
	return 0
}

func (c *CPU) IZX() uint8 { // Indirect zero page X
	t := uint16(c.Read(c.pc))
	c.pc++

	lo := uint16(c.Read(uint16(t+uint16(c.x)) & 0x00FF))
	hi := uint16(c.Read(uint16(t+uint16(c.x)+1) & 0x00FF))

	c.addressAbsolute = (hi << 8) | lo

	return 0
}

// Indirect zero page Y, different than X as the memory addresses are c.read and y is added to the resolved address
// Rather than adding X to the memory address and then resolving
func (c *CPU) IZY() uint8 {
	t := uint16(c.Read(c.pc))
	c.pc++

	lo := uint16(c.Read(t & 0x00FF))
	hi := uint16(c.Read((t + 1) & 0x00FF))

	c.addressAbsolute = (hi << 8) | lo
	c.addressAbsolute += uint16(c.y)

	if (c.addressAbsolute & 0xFF00) != (hi << 8) {
		return 1
	}
	return 0
}

func (c *CPU) REL() uint8 {
	// c.addressRelative = uint16(0xFF00) + uint16(c.Read(c.pc))
	c.addressRelative = uint16(c.Read(c.pc))

	c.pc++
	if (c.addressRelative & 0x80) != 0x0 {
		c.addressRelative |= 0xFF00
	}
	return 0
}

func (c *CPU) Step() {

	if c.cycles == 0 { // Not clock cycle accurate
		c.opcode = c.Read(c.pc)
		c.SetFlag(U, true)

		c.pc++

		c.cycles = c.Lookup[c.opcode].cycles

		additionalCyclesAddrMode := (c.Lookup[c.opcode].address_mode.fn)()
		additionalCyclesOperand := (c.Lookup[c.opcode].operate)()

		c.cycles += additionalCyclesAddrMode + additionalCyclesOperand

		c.SetFlag(U, true)
	}
	c.cyclesCounter++
	c.cycles--
}

func (c *CPU) Reset() { // Reset signal

	c.addressAbsolute = 0xFFFC
	lo := uint16(c.Read(c.addressAbsolute + 0))
	hi := uint16(c.Read(c.addressAbsolute + 1))

	c.pc = (hi << 8) | lo
	c.a = 0
	c.x = 0
	c.y = 0
	c.stackPointer = 0xFD
	c.status = 0x00 | uint8(U)

	c.addressRelative = 0x0000
	c.addressAbsolute = 0x0000
	c.fetched = 0x0

	c.cycles = 8
}

func (c *CPU) irq() { // Interrupt request
	if c.GetFlag(I) == 0 {
		c.Write(0x0100+uint16(c.stackPointer), uint8((c.pc>>8)&0x00FF))
		c.stackPointer--

		c.Write(0x0100+uint16(c.stackPointer), uint8(c.pc&0x00FF))
		c.stackPointer--

		c.SetFlagU(B, 0)
		c.SetFlagU(U, 1)
		c.SetFlagU(I, 1)

		c.Write(0x0100+uint16(c.stackPointer), c.status)
		c.stackPointer--

		c.addressAbsolute = 0xFFFE
		lo := uint16(c.Read(c.addressAbsolute + 0))
		hi := uint16(c.Read(c.addressAbsolute + 1))
		c.pc = (hi << 8) | lo

		c.cycles = 7
	}
}
func (c *CPU) nmi() { // Non maskable interrupt request, nothing can't stop this
	c.Write(0x0100+uint16(c.stackPointer), uint8((c.pc>>8)&0x00FF))
	c.stackPointer--

	c.Write(0x0100+uint16(c.stackPointer), uint8(c.pc&0x00FF))
	c.stackPointer--

	c.SetFlagU(B, 0)
	c.SetFlagU(U, 1)
	c.SetFlagU(I, 1)
	c.Write(0x0100+uint16(c.stackPointer), c.status)
	c.stackPointer--

	c.addressAbsolute = 0xFFFA
	lo := uint16(c.Read(c.addressAbsolute + 0))
	hi := uint16(c.Read(c.addressAbsolute + 1))
	c.pc = (hi << 8) | lo

	c.cycles = 8
}

func (c *CPU) fetch() uint8 {
	if c.Lookup[c.opcode].address_mode.Type != "IMP" {
		c.fetched = c.Read(c.addressAbsolute)
	}
	return c.fetched
}

func (c *CPU) GetFlag(f FLAGS) uint8 {
	if (c.status & uint8(f)) > 0 {
		return 1
	}
	return 0
}

func (c *CPU) SetFlag(f FLAGS, v bool) {
	if v {
		c.status |= uint8(f)
	} else {
		c.status &= uint8(^f)
	}
}

func (c *CPU) SetFlagU(f FLAGS, v int) {
	if v != 0 {
		c.status |= uint8(f)
	} else {
		c.status &= uint8(^f)
	}
}

func (c *CPU) Complete() bool {
	return c.cycles == 0
}
