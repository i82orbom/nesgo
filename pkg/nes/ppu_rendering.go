package nes

import (
	"image"
	"image/color"
)

// Step is in charge of rendering the output of the PPU
func (ppu *PPU) Step() {
	// PREPARE FOR RENDERING
	if ppu.scanline >= -1 && ppu.scanline < 240 {

		if ppu.scanline == 0 && ppu.cycle == 0 {
			ppu.cycle = 1
		}

		// RESET VARIABLES TO START THE RENDERING
		if ppu.scanline == -1 && ppu.cycle == 1 {
			ppu.statusRegister.VerticalBlank = 0
		}

		// BackgroundRenderingCycle gets&updates the variables to render the next background tile
		ppu.backgroundRenderingCycle()

		// TODO: Sprites
	}

	if ppu.scanline >= 241 && ppu.scanline < 261 {
		if ppu.scanline == 241 && ppu.cycle == 1 {
			ppu.statusRegister.VerticalBlank = 1
			if ppu.controlRegister.EnableNMI != 0 {
				ppu.nmiSignaled = true
			}
		}
	}

	// Render background pixel & palette
	bgPixel, bgPalette := ppu.renderBackgroundPixel()

	// Draw it!
	color := ppu.colorFromPalette(bgPalette, bgPixel)
	ppu.renderedTexture.SetRGBA(int(ppu.cycle-1), int(ppu.scanline), *color)

	ppu.cycle++
	if ppu.cycle >= 341 {
		ppu.cycle = 0
		ppu.scanline++
		if ppu.scanline >= 261 {
			ppu.scanline = -1
			ppu.frameComplete = true
		}
	}
}

func (ppu *PPU) patternTable(idx uint8, paletteID uint8) *image.RGBA {
	for yTile := uint16(0); yTile < 16; yTile++ {
		for xTile := uint16(0); xTile < 16; xTile++ {
			offset := (yTile * 256) + (xTile * 16)

			for row := uint16(0); row < 8; row++ {
				tileLSB := ppu.PPURead(uint16(idx)*0x1000+offset+row+0, false)
				tileMSB := ppu.PPURead(uint16(idx)*0x1000+offset+row+8, false)

				for col := uint16(0); col < 8; col++ {
					pixel := (tileLSB & 0x01) | (tileMSB&0x01)<<1
					tileLSB >>= 1
					tileMSB >>= 1

					x := (xTile * 8) + (7 - col)
					y := (yTile * 8) + row
					color := ppu.colorFromPalette(paletteID, pixel)
					ppu.spritePatternTableTextures[idx].SetRGBA(
						int(x),
						int(y),
						*color,
					)
				}
			}
		}
	}
	return ppu.spritePatternTableTextures[idx]
}

func (ppu *PPU) colorFromPalette(paletteID uint8, pixel uint8) *color.RGBA {
	address := uint16(0x3F00) + uint16(paletteID)<<2 + uint16(pixel)
	idx := ppu.PPURead(address, false)
	return &ppu.colorPalette[idx&0x3F]
}

// isRendering returns true if background and sprites should be rendered
func (ppu *PPU) isRendering() bool {
	return ppu.maskRegister.RenderBackground != 0 || ppu.maskRegister.RenderSprites != 0
}
