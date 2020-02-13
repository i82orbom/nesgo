package nes

import (
	"fmt"
	"io"
)

// Dissasemble returns
func (c *CPU) dissasembleState(nStart uint16, nStop uint16) ([]uint16, map[uint16]string) {
	addr := uint32(nStart)
	value := uint8(0x00)
	lo := uint8(0x00)
	hi := uint8(0x00)
	mapLines := map[uint16]string{}
	keys := []uint16{}
	lineAddress := uint16(0)

	for addr <= uint32(nStop) {
		lineAddress = uint16(addr)

		sInst := fmt.Sprintf("$0x%04x: ", addr)
		// Read instruction, and get its readable name
		opcode := c.bus.cpuRead(uint16(addr), true)
		addr++
		sInst += c.Lookup[opcode].name + " "

		if c.Lookup[opcode].adressingModeFn.Type == "IMP" {
			sInst += " {IMP}"
		} else if c.Lookup[opcode].adressingModeFn.Type == "IMM" {
			value = c.bus.cpuRead(uint16(addr), true)
			addr++
			sInst += "#$" + fmt.Sprintf("0x%02x", value) + " {IMM}"
		} else if c.Lookup[opcode].adressingModeFn.Type == "ZP0" {
			lo = c.bus.cpuRead(uint16(addr), true)
			addr++
			hi = 0x00
			sInst += "$" + fmt.Sprintf("0x%02x", lo) + " {ZP0}"
		} else if c.Lookup[opcode].adressingModeFn.Type == "ZPX" {
			lo = c.bus.cpuRead(uint16(addr), true)
			addr++
			hi = 0x00
			sInst += "$" + fmt.Sprintf("0x%02x", lo) + ", X {ZPX}"
		} else if c.Lookup[opcode].adressingModeFn.Type == "ZPY" {
			lo = c.bus.cpuRead(uint16(addr), true)
			addr++
			hi = 0x00
			sInst += "$" + fmt.Sprintf("0x%02x", lo) + ", Y {ZPY}"
		} else if c.Lookup[opcode].adressingModeFn.Type == "IZX" {
			lo = c.bus.cpuRead(uint16(addr), true)
			addr++
			hi = 0x00
			sInst += "($" + fmt.Sprintf("0x%02x", lo) + ", X) {IZX}"
		} else if c.Lookup[opcode].adressingModeFn.Type == "IZY" {
			lo = c.bus.cpuRead(uint16(addr), true)
			addr++
			hi = 0x00
			sInst += "($" + fmt.Sprintf("0x%02x", lo) + "), Y {IZY}"
		} else if c.Lookup[opcode].adressingModeFn.Type == "ABS" {
			lo = c.bus.cpuRead(uint16(addr), true)
			addr++
			hi = c.bus.cpuRead(uint16(addr), true)
			addr++
			sInst += "$" + fmt.Sprintf("0x%04x", ((uint16(hi)<<8)|uint16(lo))) + " {ABS}"
		} else if c.Lookup[opcode].adressingModeFn.Type == "ABX" {
			lo = c.bus.cpuRead(uint16(addr), true)
			addr++
			hi = c.bus.cpuRead(uint16(addr), true)
			addr++
			sInst += "$" + fmt.Sprintf("0x%04x", ((uint16(hi)<<8)|uint16(lo))) + ", X {ABX}"
		} else if c.Lookup[opcode].adressingModeFn.Type == "ABY" {
			lo = c.bus.cpuRead(uint16(addr), true)
			addr++
			hi = c.bus.cpuRead(uint16(addr), true)
			addr++
			sInst += "$" + fmt.Sprintf("0x%04x", ((uint16(hi)<<8)|uint16(lo))) + ", Y {ABY}"
		} else if c.Lookup[opcode].adressingModeFn.Type == "IND" {
			lo = c.bus.cpuRead(uint16(addr), true)
			addr++
			hi = c.bus.cpuRead(uint16(addr), true)
			addr++
			sInst += "($" + fmt.Sprintf("0x%04x", ((uint16(hi)<<8)|uint16(lo))) + ") {IND}"
		} else if c.Lookup[opcode].adressingModeFn.Type == "REL" {
			value = c.bus.cpuRead(uint16(addr), true)
			addr++
			sInst += "$" + fmt.Sprintf("0x%02x", value) + " [$" + fmt.Sprintf("0x%04x", int8(addr)+int8(value)) + "] {REL}"
		}

		if addr > uint32(c.pc-20) && uint32(addr) < uint32(c.pc+20) {
			mapLines[lineAddress] = sInst
			keys = append(keys, lineAddress)
		}
	}

	return keys, mapLines
}

// DissasembleCurrentPC writes to the writer the dissasembed currently executed code
func (c *CPU) DissasembleCurrentPC(writer io.Writer) {
	keys, data := c.dissasembleState(0x0000, 0xFFFF)
	for _, key := range keys {
		pc := c.pc
		if pc == key {
			writer.Write([]byte(fmt.Sprintf(">> [0x%04x] = %s\n", key, data[key])))

		} else {
			writer.Write([]byte(fmt.Sprintf("[0x%04x] = %s\n", key, data[key])))
		}
	}
}
