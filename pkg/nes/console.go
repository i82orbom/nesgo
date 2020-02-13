package nes

import (
	"os"

	"github.com/i82orbom/nesgo/pkg/gui"
)

// Console represents the NES
type Console struct {
	// Console RAM
	ram [2048]uint8

	// Status
	clockCounter int

	// Devices connected to the Console
	cpu  *CPU
	ppu  *PPU
	cart *Cartridge
}

// NewConsole creates an instance of a NES
func NewConsole() *Console {
	cpu := NewCPU()
	ppu := NewPPU()

	bus := &Console{
		cpu: cpu,
		ppu: ppu,
	}

	// Connect devices
	cpu.ConnectBus(bus)
	return bus
}

// Reset resets the console
func (b *Console) Reset() {
	b.cpu.reset()
	b.ppu.reset()
}

// InsertCartridge connects the cartridge to the console
func (b *Console) InsertCartridge(c *Cartridge) {
	b.cart = c
	b.ppu.InsertCartridge(c)
}

// Step steps the console a single cycle
func (b *Console) Step() {
	b.ppu.Step()

	if (b.clockCounter % 3) == 0 {
		b.cpu.Step()
	}
}

// StepFrame steps the console enough to generate one frame
func (b *Console) StepFrame() {
	for {
		b.Step()
		if b.ppu.frameComplete {
			b.ppu.frameComplete = false
			break
		}
	}
}

// Disassemble shows the currently executed code
func (b *Console) Disassemble() {
	b.cpu.DissasembleCurrentPC(os.Stdout)
}

func (b *Console) cpuRead(address uint16, readOnly bool) uint8 {
	data := uint8(0x00)
	// The cartridge can veto reads to other devices
	if b.cart.CPURead(address, &data) {
		return data
	}

	if address >= 0x0000 && address <= 0x1FFF {
		data = b.ram[address&0x07FF] // Mask for mirroring
	} else if address >= 0x2000 && address <= 0x3FFF { // PPU
		data = b.ppu.CPURead(address&0x0007, readOnly)
	}

	// Read from other devices
	return data
}

func (b *Console) cpuWrite(address uint16, data uint8) {
	// The cartridge can veto writes to other devices
	if b.cart.CPUWrite(address, data) {
		return
	}

	if address >= 0x0000 && address <= 0x1FFF {
		b.ram[address&0x07FF] = data // Mask for mirroring
	} else if address >= 0x2000 && address <= 0x3FFF { // PPU
		b.ppu.CPUWrite(address&0x0007, data)
	}
}

// TextureProvider returns the texture provider, in this case the PPU
func (b *Console) TextureProvider() gui.TextureProvider {
	return b.ppu
}
