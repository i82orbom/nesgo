package nes

import (
	"image"
	"image/color"
)

type Console struct {
	// Console RAM
	ram [2048]uint8

	// Temp texture
	texture *image.RGBA

	// Devices connected to the Console
	cpu  *CPU
	cart *Cartridge
}

func NewConsole() *Console {
	cpu := NewCPU()
	bus := &Console{
		cpu:     cpu,
		texture: image.NewRGBA(image.Rect(0, 0, 256, 240)),
	}

	// Connect devices
	cpu.ConnectBus(bus)

	return bus
}

func (b *Console) InsertCartridge(c *Cartridge) {
	b.cart = c
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

func (b *Console) Texture() *image.RGBA {
	c1 := color.RGBA{
		A: 255,
		B: 255,
		G: 255,
		R: 255,
	}
	c2 := color.RGBA{
		A: 0,
		B: 0,
		G: 0,
		R: 255,
	}
	latch := true
	for x := 0; x < 256; x++ {
		for y := 0; y < 240; y++ {
			if latch {
				b.texture.SetRGBA(x, y, c1)
				latch = !latch
			} else {
				b.texture.SetRGBA(x, y, c2)
				latch = !latch
			}
		}
	}

	return b.texture
}
