package nes

import (
	"image"

	"github.com/i82orbom/nesgo/pkg/gui"
)

type Console struct {
	// Console RAM
	ram [2048]uint8

	// Temp texture
	texture *image.RGBA

	// Devices connected to the Console
	cpu  *CPU
	ppu  *PPU
	cart *Cartridge
}

func NewConsole() *Console {
	cpu := NewCPU()
	ppu := NewPPU()

	bus := &Console{
		cpu:     cpu,
		ppu:     ppu,
		texture: image.NewRGBA(image.Rect(0, 0, 256, 240)),
	}

	// Connect devices
	cpu.ConnectBus(bus)

	return bus
}

func (b *Console) InsertCartridge(c *Cartridge) {
	b.cart = c
	b.ppu.InsertCartridge(c)
}

func (b *Console) Reset() {
	b.cpu.Reset()
}

func (b *Console) Step() {
	for {
		b.cpu.Step()
		// The CPU will clock until no more cycles need to be executed
		if b.cpu.Complete() {
			break
		}
	}
}

func (b *Console) CPURead(address uint16, readOnly bool) uint8 {
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

func (b *Console) CPUWrite(address uint16, data uint8) {
	// The cartridge can veto writes to other devices
	if b.cart.CPUWrite(address, data) {
		return
	}

	if address >= 0x0000 && address <= 0x1FFF {
		b.ram[address&0x07FF] = data // Mask for mirroring
	}
}

func (b *Console) TextureProvider() gui.TextureProvider {
	return b.ppu
}
