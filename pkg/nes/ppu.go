package nes

import "image"

const (
	width  = 240
	height = 256
)

// PPU represents the Picture Processing Unit of the NES
type PPU struct {
	cart *Cartridge

	renderedTexture *image.RGBA
}

// NewPPU creates a new Picture Processing Unit instance
func NewPPU() *PPU {
	ppu := &PPU{}
	ppu.initialise()
	return ppu
}

func (ppu *PPU) initialise() {
	// Init texture
	ppu.renderedTexture = image.NewRGBA(image.Rect(0, 0, height, width))
}

// CPURead reads a value triggered by the CPU
func (ppu *PPU) CPURead(address uint16, readOnly bool) uint8 {
	return 0
}

// CPUWrite writes a value triggered by the CPU
func (ppu *PPU) CPUWrite(address uint16, data uint8) {

}

// PPURead reads a value trigger by the PPU (CHR data)
func (ppu *PPU) PPURead(address uint16, readOnly bool) uint8 {
	return 0
}

// PPUWrite reads a value trigger by the PPU (CHR data)
func (ppu *PPU) PPUWrite(address uint16, data uint8) {

}

// InsertCartridge sets the current cartridge
func (ppu *PPU) InsertCartridge(c *Cartridge) {
	ppu.cart = c
}

// Texture returns the current texture rendered by the PPU
// This implements gui.TextureProvider
func (ppu *PPU) Texture() *image.RGBA {
	return ppu.renderedTexture
}
