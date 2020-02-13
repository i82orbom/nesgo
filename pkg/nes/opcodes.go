package nes

type instructionSet []instruction

type instruction struct {
	name            string
	operate         func() uint8
	adressingModeFn addressingFunction
	cycles          uint8
}

type addressingFunction struct {
	fn   func() uint8
	Type string
}

func createLookupTable(cpu *CPU) []instruction {
	return []instruction{
		{"BRK", cpu.brk, addressingFunction{cpu.imp, "IMP"}, 7}, {"ORA", cpu.ora, addressingFunction{cpu.izx, "IZX"}, 6}, {"???", cpu.xxx, addressingFunction{cpu.imp, "IMP"}, 2}, {"???", cpu.xxx, addressingFunction{cpu.imp, "IMP"}, 8}, {"???", cpu.nop, addressingFunction{cpu.imp, "IMP"}, 3}, {"ORA", cpu.ora, addressingFunction{cpu.zp0, "ZP0"}, 3}, {"ASL", cpu.asl, addressingFunction{cpu.zp0, "ZP0"}, 5}, {"???", cpu.xxx, addressingFunction{cpu.imp, "IMP"}, 5}, {"PHP", cpu.php, addressingFunction{cpu.imp, "IMP"}, 3}, {"ORA", cpu.ora, addressingFunction{cpu.imm, "IMM"}, 2}, {"ASL", cpu.asl, addressingFunction{cpu.imp, "IMP"}, 2}, {"???", cpu.xxx, addressingFunction{cpu.imp, "IMP"}, 2}, {"???", cpu.nop, addressingFunction{cpu.imp, "IMP"}, 4}, {"ORA", cpu.ora, addressingFunction{cpu.abs, "ABS"}, 4}, {"ASL", cpu.asl, addressingFunction{cpu.abs, "ABS"}, 6}, {"???", cpu.xxx, addressingFunction{cpu.imp, "IMP"}, 6},
		{"BPL", cpu.bpl, addressingFunction{cpu.rel, "REL"}, 2}, {"ORA", cpu.ora, addressingFunction{cpu.izy, "IZY"}, 5}, {"???", cpu.xxx, addressingFunction{cpu.imp, "IMP"}, 2}, {"???", cpu.xxx, addressingFunction{cpu.imp, "IMP"}, 8}, {"???", cpu.nop, addressingFunction{cpu.imp, "IMP"}, 4}, {"ORA", cpu.ora, addressingFunction{cpu.zpx, "ZPX"}, 4}, {"ASL", cpu.asl, addressingFunction{cpu.zpx, "ZPX"}, 6}, {"???", cpu.xxx, addressingFunction{cpu.imp, "IMP"}, 6}, {"CLC", cpu.clc, addressingFunction{cpu.imp, "IMP"}, 2}, {"ORA", cpu.ora, addressingFunction{cpu.aby, "ABY"}, 4}, {"???", cpu.nop, addressingFunction{cpu.imp, "IMP"}, 2}, {"???", cpu.xxx, addressingFunction{cpu.imp, "IMP"}, 7}, {"???", cpu.nop, addressingFunction{cpu.imp, "IMP"}, 4}, {"ORA", cpu.ora, addressingFunction{cpu.abx, "ABX"}, 4}, {"ASL", cpu.asl, addressingFunction{cpu.abx, "ABX"}, 7}, {"???", cpu.xxx, addressingFunction{cpu.imp, "IMP"}, 7},
		{"JSR", cpu.jsr, addressingFunction{cpu.abs, "ABS"}, 6}, {"AND", cpu.and, addressingFunction{cpu.izx, "IZX"}, 6}, {"???", cpu.xxx, addressingFunction{cpu.imp, "IMP"}, 2}, {"???", cpu.xxx, addressingFunction{cpu.imp, "IMP"}, 8}, {"BIT", cpu.bit, addressingFunction{cpu.zp0, "ZP0"}, 3}, {"AND", cpu.and, addressingFunction{cpu.zp0, "ZP0"}, 3}, {"ROL", cpu.rol, addressingFunction{cpu.zp0, "ZP0"}, 5}, {"???", cpu.xxx, addressingFunction{cpu.imp, "IMP"}, 5}, {"PLP", cpu.plp, addressingFunction{cpu.imp, "IMP"}, 4}, {"AND", cpu.and, addressingFunction{cpu.imm, "IMM"}, 2}, {"ROL", cpu.rol, addressingFunction{cpu.imp, "IMP"}, 2}, {"???", cpu.xxx, addressingFunction{cpu.imp, "IMP"}, 2}, {"BIT", cpu.bit, addressingFunction{cpu.abs, "ABS"}, 4}, {"AND", cpu.and, addressingFunction{cpu.abs, "ABS"}, 4}, {"ROL", cpu.rol, addressingFunction{cpu.abs, "ABS"}, 6}, {"???", cpu.xxx, addressingFunction{cpu.imp, "IMP"}, 6},
		{"BMI", cpu.bmi, addressingFunction{cpu.rel, "REL"}, 2}, {"AND", cpu.and, addressingFunction{cpu.izy, "IZY"}, 5}, {"???", cpu.xxx, addressingFunction{cpu.imp, "IMP"}, 2}, {"???", cpu.xxx, addressingFunction{cpu.imp, "IMP"}, 8}, {"???", cpu.nop, addressingFunction{cpu.imp, "IMP"}, 4}, {"AND", cpu.and, addressingFunction{cpu.zpx, "ZPX"}, 4}, {"ROL", cpu.rol, addressingFunction{cpu.zpx, "ZPX"}, 6}, {"???", cpu.xxx, addressingFunction{cpu.imp, "IMP"}, 6}, {"SEC", cpu.sec, addressingFunction{cpu.imp, "IMP"}, 2}, {"AND", cpu.and, addressingFunction{cpu.aby, "ABY"}, 4}, {"???", cpu.nop, addressingFunction{cpu.imp, "IMP"}, 2}, {"???", cpu.xxx, addressingFunction{cpu.imp, "IMP"}, 7}, {"???", cpu.nop, addressingFunction{cpu.imp, "IMP"}, 4}, {"AND", cpu.and, addressingFunction{cpu.abx, "ABX"}, 4}, {"ROL", cpu.rol, addressingFunction{cpu.abx, "ABX"}, 7}, {"???", cpu.xxx, addressingFunction{cpu.imp, "IMP"}, 7},
		{"RTI", cpu.rti, addressingFunction{cpu.imp, "IMP"}, 6}, {"EOR", cpu.eor, addressingFunction{cpu.izx, "IZX"}, 6}, {"???", cpu.xxx, addressingFunction{cpu.imp, "IMP"}, 2}, {"???", cpu.xxx, addressingFunction{cpu.imp, "IMP"}, 8}, {"???", cpu.nop, addressingFunction{cpu.imp, "IMP"}, 3}, {"EOR", cpu.eor, addressingFunction{cpu.zp0, "ZP0"}, 3}, {"LSR", cpu.lsr, addressingFunction{cpu.zp0, "ZP0"}, 5}, {"???", cpu.xxx, addressingFunction{cpu.imp, "IMP"}, 5}, {"PHA", cpu.pha, addressingFunction{cpu.imp, "IMP"}, 3}, {"EOR", cpu.eor, addressingFunction{cpu.imm, "IMM"}, 2}, {"LSR", cpu.lsr, addressingFunction{cpu.imp, "IMP"}, 2}, {"???", cpu.xxx, addressingFunction{cpu.imp, "IMP"}, 2}, {"JMP", cpu.jmp, addressingFunction{cpu.abs, "ABS"}, 3}, {"EOR", cpu.eor, addressingFunction{cpu.abs, "ABS"}, 4}, {"LSR", cpu.lsr, addressingFunction{cpu.abs, "ABS"}, 6}, {"???", cpu.xxx, addressingFunction{cpu.imp, "IMP"}, 6},
		{"BVC", cpu.bvc, addressingFunction{cpu.rel, "REL"}, 2}, {"EOR", cpu.eor, addressingFunction{cpu.izy, "IZY"}, 5}, {"???", cpu.xxx, addressingFunction{cpu.imp, "IMP"}, 2}, {"???", cpu.xxx, addressingFunction{cpu.imp, "IMP"}, 8}, {"???", cpu.nop, addressingFunction{cpu.imp, "IMP"}, 4}, {"EOR", cpu.eor, addressingFunction{cpu.zpx, "ZPX"}, 4}, {"LSR", cpu.lsr, addressingFunction{cpu.zpx, "ZPX"}, 6}, {"???", cpu.xxx, addressingFunction{cpu.imp, "IMP"}, 6}, {"CLI", cpu.cli, addressingFunction{cpu.imp, "IMP"}, 2}, {"EOR", cpu.eor, addressingFunction{cpu.aby, "ABY"}, 4}, {"???", cpu.nop, addressingFunction{cpu.imp, "IMP"}, 2}, {"???", cpu.xxx, addressingFunction{cpu.imp, "IMP"}, 7}, {"???", cpu.nop, addressingFunction{cpu.imp, "IMP"}, 4}, {"EOR", cpu.eor, addressingFunction{cpu.abx, "ABX"}, 4}, {"LSR", cpu.lsr, addressingFunction{cpu.abx, "ABX"}, 7}, {"???", cpu.xxx, addressingFunction{cpu.imp, "IMP"}, 7},
		{"RTS", cpu.rts, addressingFunction{cpu.imp, "IMP"}, 6}, {"ADC", cpu.adc, addressingFunction{cpu.izx, "IZX"}, 6}, {"???", cpu.xxx, addressingFunction{cpu.imp, "IMP"}, 2}, {"???", cpu.xxx, addressingFunction{cpu.imp, "IMP"}, 8}, {"???", cpu.nop, addressingFunction{cpu.imp, "IMP"}, 3}, {"ADC", cpu.adc, addressingFunction{cpu.zp0, "ZP0"}, 3}, {"ROR", cpu.ror, addressingFunction{cpu.zp0, "ZP0"}, 5}, {"???", cpu.xxx, addressingFunction{cpu.imp, "IMP"}, 5}, {"PLA", cpu.pla, addressingFunction{cpu.imp, "IMP"}, 4}, {"ADC", cpu.adc, addressingFunction{cpu.imm, "IMM"}, 2}, {"ROR", cpu.ror, addressingFunction{cpu.imp, "IMP"}, 2}, {"???", cpu.xxx, addressingFunction{cpu.imp, "IMP"}, 2}, {"JMP", cpu.jmp, addressingFunction{cpu.ind, "IND"}, 5}, {"ADC", cpu.adc, addressingFunction{cpu.abs, "ABS"}, 4}, {"ROR", cpu.ror, addressingFunction{cpu.abs, "ABS"}, 6}, {"???", cpu.xxx, addressingFunction{cpu.imp, "IMP"}, 6},
		{"BVS", cpu.bvs, addressingFunction{cpu.rel, "REL"}, 2}, {"ADC", cpu.adc, addressingFunction{cpu.izy, "IZY"}, 5}, {"???", cpu.xxx, addressingFunction{cpu.imp, "IMP"}, 2}, {"???", cpu.xxx, addressingFunction{cpu.imp, "IMP"}, 8}, {"???", cpu.nop, addressingFunction{cpu.imp, "IMP"}, 4}, {"ADC", cpu.adc, addressingFunction{cpu.zpx, "ZPX"}, 4}, {"ROR", cpu.ror, addressingFunction{cpu.zpx, "ZPX"}, 6}, {"???", cpu.xxx, addressingFunction{cpu.imp, "IMP"}, 6}, {"SEI", cpu.sei, addressingFunction{cpu.imp, "IMP"}, 2}, {"ADC", cpu.adc, addressingFunction{cpu.aby, "ABY"}, 4}, {"???", cpu.nop, addressingFunction{cpu.imp, "IMP"}, 2}, {"???", cpu.xxx, addressingFunction{cpu.imp, "IMP"}, 7}, {"???", cpu.nop, addressingFunction{cpu.imp, "IMP"}, 4}, {"ADC", cpu.adc, addressingFunction{cpu.abx, "ABX"}, 4}, {"ROR", cpu.ror, addressingFunction{cpu.abx, "ABX"}, 7}, {"???", cpu.xxx, addressingFunction{cpu.imp, "IMP"}, 7},
		{"???", cpu.nop, addressingFunction{cpu.imp, "IMP"}, 2}, {"STA", cpu.sta, addressingFunction{cpu.izx, "IZX"}, 6}, {"???", cpu.nop, addressingFunction{cpu.imp, "IMP"}, 2}, {"???", cpu.xxx, addressingFunction{cpu.imp, "IMP"}, 6}, {"STY", cpu.sty, addressingFunction{cpu.zp0, "ZP0"}, 3}, {"STA", cpu.sta, addressingFunction{cpu.zp0, "ZP0"}, 3}, {"STX", cpu.stx, addressingFunction{cpu.zp0, "ZP0"}, 3}, {"???", cpu.xxx, addressingFunction{cpu.imp, "IMP"}, 3}, {"DEY", cpu.dey, addressingFunction{cpu.imp, "IMP"}, 2}, {"???", cpu.nop, addressingFunction{cpu.imp, "IMP"}, 2}, {"TXA", cpu.txa, addressingFunction{cpu.imp, "IMP"}, 2}, {"???", cpu.xxx, addressingFunction{cpu.imp, "IMP"}, 2}, {"STY", cpu.sty, addressingFunction{cpu.abs, "ABS"}, 4}, {"STA", cpu.sta, addressingFunction{cpu.abs, "ABS"}, 4}, {"STX", cpu.stx, addressingFunction{cpu.abs, "ABS"}, 4}, {"???", cpu.xxx, addressingFunction{cpu.imp, "IMP"}, 4},
		{"BCC", cpu.bcc, addressingFunction{cpu.rel, "REL"}, 2}, {"STA", cpu.sta, addressingFunction{cpu.izy, "IZY"}, 6}, {"???", cpu.xxx, addressingFunction{cpu.imp, "IMP"}, 2}, {"???", cpu.xxx, addressingFunction{cpu.imp, "IMP"}, 6}, {"STY", cpu.sty, addressingFunction{cpu.zpx, "ZPX"}, 4}, {"STA", cpu.sta, addressingFunction{cpu.zpx, "ZPX"}, 4}, {"STX", cpu.stx, addressingFunction{cpu.zpy, "ZPY"}, 4}, {"???", cpu.xxx, addressingFunction{cpu.imp, "IMP"}, 4}, {"TYA", cpu.tya, addressingFunction{cpu.imp, "IMP"}, 2}, {"STA", cpu.sta, addressingFunction{cpu.aby, "ABY"}, 5}, {"TXS", cpu.txs, addressingFunction{cpu.imp, "IMP"}, 2}, {"???", cpu.xxx, addressingFunction{cpu.imp, "IMP"}, 5}, {"???", cpu.nop, addressingFunction{cpu.imp, "IMP"}, 5}, {"STA", cpu.sta, addressingFunction{cpu.abx, "ABX"}, 5}, {"???", cpu.xxx, addressingFunction{cpu.imp, "IMP"}, 5}, {"???", cpu.xxx, addressingFunction{cpu.imp, "IMP"}, 5},
		{"LDY", cpu.ldy, addressingFunction{cpu.imm, "IMM"}, 2}, {"LDA", cpu.lda, addressingFunction{cpu.izx, "IZX"}, 6}, {"LDX", cpu.ldx, addressingFunction{cpu.imm, "IMM"}, 2}, {"???", cpu.xxx, addressingFunction{cpu.imp, "IMP"}, 6}, {"LDY", cpu.ldy, addressingFunction{cpu.zp0, "ZP0"}, 3}, {"LDA", cpu.lda, addressingFunction{cpu.zp0, "ZP0"}, 3}, {"LDX", cpu.ldx, addressingFunction{cpu.zp0, "ZP0"}, 3}, {"???", cpu.xxx, addressingFunction{cpu.imp, "IMP"}, 3}, {"TAY", cpu.tay, addressingFunction{cpu.imp, "IMP"}, 2}, {"LDA", cpu.lda, addressingFunction{cpu.imm, "IMM"}, 2}, {"TAX", cpu.tax, addressingFunction{cpu.imp, "IMP"}, 2}, {"???", cpu.xxx, addressingFunction{cpu.imp, "IMP"}, 2}, {"LDY", cpu.ldy, addressingFunction{cpu.abs, "ABS"}, 4}, {"LDA", cpu.lda, addressingFunction{cpu.abs, "ABS"}, 4}, {"LDX", cpu.ldx, addressingFunction{cpu.abs, "ABS"}, 4}, {"???", cpu.xxx, addressingFunction{cpu.imp, "IMP"}, 4},
		{"BCS", cpu.bcs, addressingFunction{cpu.rel, "REL"}, 2}, {"LDA", cpu.lda, addressingFunction{cpu.izy, "IZY"}, 5}, {"???", cpu.xxx, addressingFunction{cpu.imp, "IMP"}, 2}, {"???", cpu.xxx, addressingFunction{cpu.imp, "IMP"}, 5}, {"LDY", cpu.ldy, addressingFunction{cpu.zpx, "ZPX"}, 4}, {"LDA", cpu.lda, addressingFunction{cpu.zpx, "ZPX"}, 4}, {"LDX", cpu.ldx, addressingFunction{cpu.zpy, "ZPY"}, 4}, {"???", cpu.xxx, addressingFunction{cpu.imp, "IMP"}, 4}, {"CLV", cpu.clv, addressingFunction{cpu.imp, "IMP"}, 2}, {"LDA", cpu.lda, addressingFunction{cpu.aby, "ABY"}, 4}, {"TSX", cpu.tsx, addressingFunction{cpu.imp, "IMP"}, 2}, {"???", cpu.xxx, addressingFunction{cpu.imp, "IMP"}, 4}, {"LDY", cpu.ldy, addressingFunction{cpu.abx, "ABX"}, 4}, {"LDA", cpu.lda, addressingFunction{cpu.abx, "ABX"}, 4}, {"LDX", cpu.ldx, addressingFunction{cpu.aby, "ABY"}, 4}, {"???", cpu.xxx, addressingFunction{cpu.imp, "IMP"}, 4},
		{"CPY", cpu.cpy, addressingFunction{cpu.imm, "IMM"}, 2}, {"CMP", cpu.cmp, addressingFunction{cpu.izx, "IZX"}, 6}, {"???", cpu.nop, addressingFunction{cpu.imp, "IMP"}, 2}, {"???", cpu.xxx, addressingFunction{cpu.imp, "IMP"}, 8}, {"CPY", cpu.cpy, addressingFunction{cpu.zp0, "ZP0"}, 3}, {"CMP", cpu.cmp, addressingFunction{cpu.zp0, "ZP0"}, 3}, {"DEC", cpu.dec, addressingFunction{cpu.zp0, "ZP0"}, 5}, {"???", cpu.xxx, addressingFunction{cpu.imp, "IMP"}, 5}, {"INY", cpu.iny, addressingFunction{cpu.imp, "IMP"}, 2}, {"CMP", cpu.cmp, addressingFunction{cpu.imm, "IMM"}, 2}, {"DEX", cpu.dex, addressingFunction{cpu.imp, "IMP"}, 2}, {"???", cpu.xxx, addressingFunction{cpu.imp, "IMP"}, 2}, {"CPY", cpu.cpy, addressingFunction{cpu.abs, "ABS"}, 4}, {"CMP", cpu.cmp, addressingFunction{cpu.abs, "ABS"}, 4}, {"DEC", cpu.dec, addressingFunction{cpu.abs, "ABS"}, 6}, {"???", cpu.xxx, addressingFunction{cpu.imp, "IMP"}, 6},
		{"BNE", cpu.bne, addressingFunction{cpu.rel, "REL"}, 2}, {"CMP", cpu.cmp, addressingFunction{cpu.izy, "IZY"}, 5}, {"???", cpu.xxx, addressingFunction{cpu.imp, "IMP"}, 2}, {"???", cpu.xxx, addressingFunction{cpu.imp, "IMP"}, 8}, {"???", cpu.nop, addressingFunction{cpu.imp, "IMP"}, 4}, {"CMP", cpu.cmp, addressingFunction{cpu.zpx, "ZPX"}, 4}, {"DEC", cpu.dec, addressingFunction{cpu.zpx, "ZPX"}, 6}, {"???", cpu.xxx, addressingFunction{cpu.imp, "IMP"}, 6}, {"CLD", cpu.cld, addressingFunction{cpu.imp, "IMP"}, 2}, {"CMP", cpu.cmp, addressingFunction{cpu.aby, "ABY"}, 4}, {"NOP", cpu.nop, addressingFunction{cpu.imp, "IMP"}, 2}, {"???", cpu.xxx, addressingFunction{cpu.imp, "IMP"}, 7}, {"???", cpu.nop, addressingFunction{cpu.imp, "IMP"}, 4}, {"CMP", cpu.cmp, addressingFunction{cpu.abx, "ABX"}, 4}, {"DEC", cpu.dec, addressingFunction{cpu.abx, "ABX"}, 7}, {"???", cpu.xxx, addressingFunction{cpu.imp, "IMP"}, 7},
		{"CPX", cpu.cpx, addressingFunction{cpu.imm, "IMM"}, 2}, {"SBC", cpu.sbc, addressingFunction{cpu.izx, "IZX"}, 6}, {"???", cpu.nop, addressingFunction{cpu.imp, "IMP"}, 2}, {"???", cpu.xxx, addressingFunction{cpu.imp, "IMP"}, 8}, {"CPX", cpu.cpx, addressingFunction{cpu.zp0, "ZP0"}, 3}, {"SBC", cpu.sbc, addressingFunction{cpu.zp0, "ZP0"}, 3}, {"INC", cpu.inc, addressingFunction{cpu.zp0, "ZP0"}, 5}, {"???", cpu.xxx, addressingFunction{cpu.imp, "IMP"}, 5}, {"INX", cpu.inx, addressingFunction{cpu.imp, "IMP"}, 2}, {"SBC", cpu.sbc, addressingFunction{cpu.imm, "IMM"}, 2}, {"NOP", cpu.nop, addressingFunction{cpu.imp, "IMP"}, 2}, {"???", cpu.sbc, addressingFunction{cpu.imp, "IMP"}, 2}, {"CPX", cpu.cpx, addressingFunction{cpu.abs, "ABS"}, 4}, {"SBC", cpu.sbc, addressingFunction{cpu.abs, "ABS"}, 4}, {"INC", cpu.inc, addressingFunction{cpu.abs, "ABS"}, 6}, {"???", cpu.xxx, addressingFunction{cpu.imp, "IMP"}, 6},
		{"BEQ", cpu.beq, addressingFunction{cpu.rel, "REL"}, 2}, {"SBC", cpu.sbc, addressingFunction{cpu.izy, "IZY"}, 5}, {"???", cpu.xxx, addressingFunction{cpu.imp, "IMP"}, 2}, {"???", cpu.xxx, addressingFunction{cpu.imp, "IMP"}, 8}, {"???", cpu.nop, addressingFunction{cpu.imp, "IMP"}, 4}, {"SBC", cpu.sbc, addressingFunction{cpu.zpx, "ZPX"}, 4}, {"INC", cpu.inc, addressingFunction{cpu.zpx, "ZPX"}, 6}, {"???", cpu.xxx, addressingFunction{cpu.imp, "IMP"}, 6}, {"SED", cpu.sed, addressingFunction{cpu.imp, "IMP"}, 2}, {"SBC", cpu.sbc, addressingFunction{cpu.aby, "ABY"}, 4}, {"NOP", cpu.nop, addressingFunction{cpu.imp, "IMP"}, 2}, {"???", cpu.xxx, addressingFunction{cpu.imp, "IMP"}, 7}, {"???", cpu.nop, addressingFunction{cpu.imp, "IMP"}, 4}, {"SBC", cpu.sbc, addressingFunction{cpu.abx, "ABX"}, 4}, {"INC", cpu.inc, addressingFunction{cpu.abx, "ABX"}, 7}, {"???", cpu.xxx, addressingFunction{cpu.imp, "IMP"}, 7},
	}
}

// ADD
func (c *CPU) adc() uint8 {
	c.fetch()
	flagCValue := c.getFlag(flagC)
	temp := uint16(c.a) + uint16(c.fetched) + uint16(flagCValue)
	c.setFlag(flagC, temp > 255)
	c.setFlag(flagZ, (temp&0x00FF) == 0)
	c.setFlagN(flagN, int(temp&0x80))
	c.setFlagN(flagV, int(^(uint16(c.a)^uint16(c.fetched))&(uint16(c.a)^temp))&0x0080)
	c.a = uint8(temp & 0x00FF)
	return 1
}

// AND operation
func (c *CPU) and() uint8 {
	c.fetch()
	c.a = c.a & c.fetched
	c.setFlag(flagZ, c.a == 0x00)
	c.setFlagN(flagN, int(c.a&0x80)) // If bit 7 is 1
	return 1
}

// Arithmetic shift left
func (c *CPU) asl() uint8 {
	c.fetch()
	temp := uint16(c.fetched) << 1
	c.setFlag(flagC, (temp&0xFF00) > 0)
	c.setFlag(flagZ, (temp&0x00FF) == 0x00)
	c.setFlagN(flagN, int(temp&0x80))
	if c.Lookup[c.opcode].adressingModeFn.Type == "IMP" {
		c.a = uint8(temp & 0x00FF)
	} else {
		c.Write(c.addressAbsolute, uint8(temp&0x00FF))
	}
	return 0
}

// Branch if carry clear
func (c *CPU) bcc() uint8 {
	if c.getFlag(flagC) == 0 {
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
func (c *CPU) bcs() uint8 {
	if c.getFlag(flagC) == 1 {
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
func (c *CPU) beq() uint8 {
	if c.getFlag(flagZ) != 0 {
		c.cycles++
		c.addressAbsolute = c.pc + c.addressRelative

		if (c.addressAbsolute & 0xFF00) != (c.pc & 0xFF00) {
			c.cycles++
		}

		c.pc = c.addressAbsolute // New location
	}
	return 0
}

func (c *CPU) bit() uint8 {
	c.fetch()
	temp := c.a & c.fetched
	c.setFlag(flagZ, (temp&0x00FF) == 0x00)
	c.setFlagN(flagN, int(c.fetched&(1<<7)))
	c.setFlagN(flagV, int(c.fetched&(1<<6)))
	return 0
}

// Branch if negative
func (c *CPU) bmi() uint8 {
	if c.getFlag(flagN) == 1 {
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
func (c *CPU) bne() uint8 {
	if c.getFlag(flagZ) == 0 {
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
func (c *CPU) bpl() uint8 {
	if c.getFlag(flagN) == 0 {
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
func (c *CPU) brk() uint8 {
	c.pc++

	c.setFlagN(flagI, 1)
	c.Write(0x0100+uint16(c.stackPointer), uint8((c.pc>>8)&0x00FF))
	c.stackPointer--
	c.Write(0x0100+uint16(c.stackPointer), uint8(c.pc&0x00FF))
	c.stackPointer--

	c.setFlagN(flagB, 1)
	c.Write(0x0100+uint16(c.stackPointer), c.status)
	c.stackPointer--
	c.setFlagN(flagB, 0)

	c.pc = uint16(c.Read(0xFFFE)) | (uint16(c.Read(0xFFFF)) << 8)
	return 0
}

// Branch if overflow
func (c *CPU) bvc() uint8 {
	if c.getFlag(flagV) == 0 {
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
func (c *CPU) bvs() uint8 {
	if c.getFlag(flagV) == 1 {
		c.cycles++
		c.addressAbsolute = c.pc + c.addressRelative

		if (c.addressAbsolute & 0xFF00) != (c.pc & 0xFF00) {
			c.cycles++
		}
		c.pc = c.addressAbsolute // New location
	}
	return 0
}

func (c *CPU) clc() uint8 {
	c.setFlag(flagC, false)
	return 0
}

func (c *CPU) cld() uint8 {
	c.setFlag(flagD, false)
	return 0
}

func (c *CPU) cli() uint8 {
	c.setFlag(flagI, false)
	return 0
}

func (c *CPU) clv() uint8 {
	c.setFlag(flagV, false)
	return 0
}

// Compare Accumulator
func (c *CPU) cmp() uint8 {
	c.fetch()
	temp := uint16(c.a) - uint16(c.fetched)
	c.setFlag(flagC, c.a >= c.fetched)
	c.setFlag(flagZ, (temp&0x00FF) == 0x0000)
	c.setFlagN(flagN, int(temp&0x0080))
	return 1
}

// Compare X Register
func (c *CPU) cpx() uint8 {
	c.fetch()
	temp := uint16(c.x) - uint16(c.fetched)
	c.setFlag(flagC, c.x >= c.fetched)
	c.setFlag(flagZ, (temp&0x00FF) == 0x0000)
	c.setFlagN(flagN, int(temp&0x0080))
	return 0
}

// Compare Y Register
func (c *CPU) cpy() uint8 {
	c.fetch()
	temp := uint16(c.y) - uint16(c.fetched)
	c.setFlag(flagC, c.y >= c.fetched)
	c.setFlag(flagZ, (temp&0x00FF) == 0x0000)
	c.setFlagN(flagN, int(temp&0x0080))
	return 0

}

// Decrement value at memory location
func (c *CPU) dec() uint8 {
	c.fetch()
	temp := uint16(c.fetched) - 1
	c.Write(c.addressAbsolute, uint8(temp&0x00FF))
	c.setFlag(flagZ, (temp&0x00FF) == 0x0000)
	c.setFlagN(flagN, int(temp&0x0080))
	return 0
}

// Decrement X
func (c *CPU) dex() uint8 {
	c.x--
	c.setFlag(flagZ, c.x == 0x00)
	c.setFlagN(flagN, int(c.x&0x80))
	return 0
}

// Decrement Y
func (c *CPU) dey() uint8 {
	c.y--
	c.setFlag(flagZ, c.y == 0x00)
	c.setFlagN(flagN, int(c.y&0x80))
	return 0
}

// Logic XOR
func (c *CPU) eor() uint8 {
	c.fetch()
	c.a = c.a ^ c.fetched
	c.setFlag(flagZ, c.a == 0x00)
	c.setFlagN(flagN, int(c.a&0x80))
	return 1
}

// Increment value at memory location
func (c *CPU) inc() uint8 {
	c.fetch()
	temp := uint16(c.fetched) + 1
	c.Write(c.addressAbsolute, uint8(temp&0x00FF))
	c.setFlag(flagZ, (temp&0x00FF) == 0x0000)
	c.setFlagN(flagN, int(temp&0x0080))
	return 0
}

func (c *CPU) inx() uint8 {
	c.x++
	c.setFlag(flagZ, c.x == 0x0)
	c.setFlagN(flagN, int(c.x&0x80))
	return 0
}

func (c *CPU) iny() uint8 {
	c.y++
	c.setFlag(flagZ, c.y == 0x0)
	c.setFlagN(flagN, int(c.y&0x80))
	return 0
}

func (c *CPU) jmp() uint8 {
	c.pc = c.addressAbsolute
	return 0
}

// Jump to subroutine
func (c *CPU) jsr() uint8 {
	c.pc--
	c.Write(0x0100+uint16(c.stackPointer), uint8((c.pc>>8)&0x00FF))
	c.stackPointer--
	c.Write(0x0100+uint16(c.stackPointer), uint8(c.pc&0x00FF))
	c.stackPointer--

	c.pc = c.addressAbsolute
	return 0
}

func (c *CPU) lda() uint8 {
	c.fetch()
	c.a = c.fetched
	c.setFlag(flagZ, c.a == 0x00)
	c.setFlagN(flagN, int(c.a&0x80))
	return 1
}

func (c *CPU) ldx() uint8 {
	c.fetch()
	c.x = c.fetched
	c.setFlag(flagZ, c.x == 0x00)
	c.setFlagN(flagN, int(c.x&0x80))
	return 1
}

func (c *CPU) ldy() uint8 {
	c.fetch()
	c.y = c.fetched
	c.setFlag(flagZ, c.y == 0x00)
	c.setFlagN(flagN, int(c.y&0x80))
	return 1
}

func (c *CPU) lsr() uint8 {
	c.fetch()
	c.setFlagN(flagC, int(c.fetched&0x0001))
	temp := uint16(c.fetched) >> 1
	c.setFlag(flagZ, (temp&0x00FF) == 0x0000)
	c.setFlagN(flagN, int(temp&0x0080))
	if c.Lookup[c.opcode].adressingModeFn.Type == "IMP" {
		c.a = uint8(temp & 0x00FF)
	} else {
		c.Write(c.addressAbsolute, uint8(temp&0x00FF))
	}
	return 0
}

func (c *CPU) nop() uint8 {
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
func (c *CPU) ora() uint8 {
	c.fetch()
	c.a = c.a | c.fetched
	c.setFlag(flagZ, c.a == 0x00)
	c.setFlagN(flagN, int(c.a&0x80))
	return 1
}

// Push A to stack
func (c *CPU) pha() uint8 {
	c.Write(0x0100+uint16(c.stackPointer), c.a)
	c.stackPointer--
	return 0
}

// Push status register to stack
func (c *CPU) php() uint8 {
	c.Write(0x0100+uint16(c.stackPointer), c.status|uint8(flagB)|uint8(flagU))
	c.setFlagN(flagB, 0)
	c.setFlagN(flagU, 0)
	c.stackPointer--
	return 0
}

// Pop A from stack
func (c *CPU) pla() uint8 {
	c.stackPointer++
	c.a = c.Read(0x0100 + uint16(c.stackPointer))
	c.setFlag(flagZ, c.a == 0x00)
	c.setFlagN(flagN, int(c.a&0x80))
	return 0
}

// Pop accumulator off Stack
func (c *CPU) plp() uint8 {
	c.stackPointer++
	c.status = c.Read(0x0100 + uint16(c.stackPointer))
	c.setFlagN(flagU, 1)
	return 0
}

func (c *CPU) rol() uint8 {
	c.fetch()
	temp := (uint16(c.fetched) << 1) | uint16(c.getFlag(flagC))
	c.setFlagN(flagC, int(temp&0xFF00))
	c.setFlag(flagZ, (temp&0x00FF) == 0x0000)
	c.setFlagN(flagN, int(temp&0x0080))
	if c.Lookup[c.opcode].adressingModeFn.Type == "IMP" {
		c.a = uint8(temp & 0x00FF)
	} else {
		c.Write(c.addressAbsolute, uint8(temp&0x00FF))
	}
	return 0
}

func (c *CPU) ror() uint8 {
	c.fetch()
	temp := (uint16(c.getFlag(flagC)) << 7) | (uint16(c.fetched) >> 1)
	c.setFlagN(flagC, int(c.fetched&0x01))
	c.setFlag(flagZ, (temp&0x00FF) == 0x00)
	c.setFlagN(flagN, int(temp&0x0080))
	if c.Lookup[c.opcode].adressingModeFn.Type == "IMP" {
		c.a = uint8(temp & 0x00FF)
	} else {
		c.Write(c.addressAbsolute, uint8(temp&0x00FF))
	}
	return 0
}

func (c *CPU) rti() uint8 {
	c.stackPointer++
	c.status = c.Read(0x0100 + uint16(c.stackPointer))
	c.status &= uint8(^flagB)
	c.status &= uint8(^flagU)

	c.stackPointer++
	c.pc = uint16(c.Read(0x0100 + uint16(c.stackPointer)))
	c.stackPointer++
	readValue := c.Read(0x0100 + uint16(c.stackPointer))
	c.pc |= uint16(readValue) << 8
	return 0
}

func (c *CPU) rts() uint8 {
	c.stackPointer++
	c.pc = uint16(c.Read(0x0100 + uint16(c.stackPointer)))
	c.stackPointer++
	readValue := c.Read(0x0100 + uint16(c.stackPointer))
	c.pc |= uint16(readValue) << 8
	c.pc++
	return 0
}

func (c *CPU) sbc() uint8 {
	c.fetch()
	value := (uint16(c.fetched)) ^ 0x00FF // Same as ADC but inverting the fetched

	temp := uint16(c.a) + value + uint16(c.getFlag(flagC))
	c.setFlagN(flagC, int(temp&0xFF00))
	c.setFlag(flagZ, (temp&0x00FF) == 0)
	c.setFlagN(flagV, int((temp^(uint16(c.a)))&(temp^value)&0x0080))
	c.setFlagN(flagN, int(temp&0x0080))
	c.a = uint8(temp & 0x00FF)
	return 1
}

func (c *CPU) sec() uint8 {
	c.setFlag(flagC, true)
	return 0
}

func (c *CPU) sed() uint8 {
	c.setFlag(flagD, true)
	return 0
}

func (c *CPU) sei() uint8 {
	c.setFlag(flagI, true)
	return 0
}

func (c *CPU) sta() uint8 {
	c.Write(c.addressAbsolute, c.a)
	return 0
}

func (c *CPU) stx() uint8 {
	c.Write(c.addressAbsolute, c.x)
	return 0
}

func (c *CPU) sty() uint8 {
	c.Write(c.addressAbsolute, c.y)
	return 0
}

func (c *CPU) tax() uint8 {
	c.x = c.a
	c.setFlag(flagZ, c.x == 0x00)
	c.setFlagN(flagN, int(c.x&0x80))
	return 0
}

func (c *CPU) tay() uint8 {
	c.y = c.a
	c.setFlag(flagZ, c.y == 0x00)
	c.setFlagN(flagN, int(c.y&0x80))
	return 0
}

func (c *CPU) tsx() uint8 {
	c.x = c.stackPointer
	c.setFlag(flagZ, c.x == 0x00)
	c.setFlagN(flagN, int(c.x&0x80))
	return 0
}

func (c *CPU) txa() uint8 {
	c.a = c.x
	c.setFlag(flagZ, c.a == 0x00)
	c.setFlagN(flagN, int(c.a&0x80))
	return 0
}

func (c *CPU) txs() uint8 {
	c.stackPointer = c.x
	return 0
}

func (c *CPU) tya() uint8 {
	c.a = c.y
	c.setFlag(flagZ, c.a == 0x00)
	c.setFlagN(flagN, int(c.a&0x80))
	return 0
}

func (c *CPU) xxx() uint8 {
	return 0
}

// Addressing modes
func (c *CPU) imp() uint8 {
	c.fetched = c.a
	return 0
}

func (c *CPU) imm() uint8 { // Immediate addressing
	c.addressAbsolute = c.pc
	c.pc++
	return 0
}

func (c *CPU) zp0() uint8 { // Zero page addressing
	c.addressAbsolute = uint16(c.Read(c.pc))
	c.pc++
	c.addressAbsolute &= 0x00FF
	return 0
}

func (c *CPU) zpx() uint8 { // Zero page addressing X offset
	c.addressAbsolute = uint16(c.Read(c.pc)) + uint16(c.x)
	c.pc++
	c.addressAbsolute &= 0x00FF
	return 0
}

func (c *CPU) zpy() uint8 { // Zero page addressing Y offset
	c.addressAbsolute = uint16(c.Read(c.pc)) + uint16(c.y)
	c.pc++
	c.addressAbsolute &= 0x00FF
	return 0
}

func (c *CPU) abs() uint8 {
	// Low byte
	lo := uint16(c.Read(c.pc))
	c.pc++
	// High byte
	hi := uint16(c.Read(c.pc))
	c.pc++

	c.addressAbsolute = (hi << 8) | lo

	return 0
}

func (c *CPU) abx() uint8 { // Same as absolute but offset with X
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

func (c *CPU) aby() uint8 { // Same as ABX but with Y
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

func (c *CPU) ind() uint8 {
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

func (c *CPU) izx() uint8 { // Indirect zero page X
	t := uint16(c.Read(c.pc))
	c.pc++

	lo := uint16(c.Read(uint16(t+uint16(c.x)) & 0x00FF))
	hi := uint16(c.Read(uint16(t+uint16(c.x)+1) & 0x00FF))

	c.addressAbsolute = (hi << 8) | lo

	return 0
}

// Indirect zero page Y, different than X as the memory addresses are c.read and y is added to the resolved address
// Rather than adding X to the memory address and then resolving
func (c *CPU) izy() uint8 {
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

func (c *CPU) rel() uint8 {
	// c.addressRelative = uint16(0xFF00) + uint16(c.Read(c.pc))
	c.addressRelative = uint16(c.Read(c.pc))

	c.pc++
	if (c.addressRelative & 0x80) != 0x0 {
		c.addressRelative |= 0xFF00
	}
	return 0
}

// Step steps the CPU a single micro-cycle
func (c *CPU) Step() {
	if c.cycles == 0 { // Not clock cycle accurate
		c.opcode = c.Read(c.pc)
		c.setFlag(flagU, true)

		c.pc++

		c.cycles = c.Lookup[c.opcode].cycles

		additionalCyclesAddrMode := (c.Lookup[c.opcode].adressingModeFn.fn)()
		additionalCyclesOperand := (c.Lookup[c.opcode].operate)()

		c.cycles += additionalCyclesAddrMode + additionalCyclesOperand

		c.setFlag(flagU, true)
	}
	c.cyclesCounter++
	c.cycles--
}

// Reset resets the CPU to the initial state
func (c *CPU) reset() { // Reset signal

	c.addressAbsolute = 0xFFFC
	lo := uint16(c.Read(c.addressAbsolute + 0))
	hi := uint16(c.Read(c.addressAbsolute + 1))

	c.pc = (hi << 8) | lo
	c.a = 0
	c.x = 0
	c.y = 0
	c.stackPointer = 0xFD
	c.status = 0x00 | uint8(flagU)

	c.addressRelative = 0x0000
	c.addressAbsolute = 0x0000
	c.fetched = 0x0

	c.cycles = 8
}

func (c *CPU) irq() { // Interrupt request
	if c.getFlag(flagI) == 0 {
		c.Write(0x0100+uint16(c.stackPointer), uint8((c.pc>>8)&0x00FF))
		c.stackPointer--

		c.Write(0x0100+uint16(c.stackPointer), uint8(c.pc&0x00FF))
		c.stackPointer--

		c.setFlagN(flagB, 0)
		c.setFlagN(flagU, 1)
		c.setFlagN(flagI, 1)

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

	c.setFlagN(flagB, 0)
	c.setFlagN(flagU, 1)
	c.setFlagN(flagI, 1)
	c.Write(0x0100+uint16(c.stackPointer), c.status)
	c.stackPointer--

	c.addressAbsolute = 0xFFFA
	lo := uint16(c.Read(c.addressAbsolute + 0))
	hi := uint16(c.Read(c.addressAbsolute + 1))
	c.pc = (hi << 8) | lo

	c.cycles = 8
}

func (c *CPU) fetch() uint8 {
	if c.Lookup[c.opcode].adressingModeFn.Type != "IMP" {
		c.fetched = c.Read(c.addressAbsolute)
	}
	return c.fetched
}

func (c *CPU) getFlag(f cpuFlag) uint8 {
	if (c.status & uint8(f)) > 0 {
		return 1
	}
	return 0
}

func (c *CPU) setFlag(f cpuFlag, v bool) {
	if v {
		c.status |= uint8(f)
	} else {
		c.status &= uint8(^f)
	}
}

// setFlagN sets a flag with an number as input
func (c *CPU) setFlagN(flagf cpuFlag, v int) {
	c.setFlag(flagf, v != 0)
}

// Complete indicates if the current instruction cycles have been consumed
func (c *CPU) Complete() bool {
	return c.cycles == 0
}
