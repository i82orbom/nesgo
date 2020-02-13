package nes

import (
	"image"
	"image/color"

	"github.com/labstack/gommon/log"
)

const (
	width  = 240
	height = 256
)

// PPU represents the Picture Processing Unit of the NES
type PPU struct {
	cart *Cartridge

	renderedTexture            *image.RGBA
	spritePatternTableTextures [2]*image.RGBA

	// Color Palette
	colorPalette []color.RGBA

	// PPU internal structures
	tableName    [2][1024]uint8
	tablePalette [32]uint8
	tablePattern [2][4096]uint8 // Not really needed
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
	ppu.spritePatternTableTextures = [2]*image.RGBA{
		image.NewRGBA(image.Rect(0, 0, 128, 128)),
		image.NewRGBA(image.Rect(0, 0, 128, 128)),
	}
	// Color palette
	ppu.colorPalette = colorPalette()
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
	data := uint8(0x0)
	address &= 0x3FFF // Guard address

	if ppu.cart.PPURead(address, &data) { // Guard Cartridge
		return data
	} else if address >= 0x0000 && address <= 0x1FFF { // Pattern memory
		data = ppu.tablePattern[(address&0x1000)>>12][address&0x0FFF]
	} else if address >= 0x2000 && address <= 0x3EFF { // Nametable memory
		address &= 0x0FFF
		if ppu.cart.MirroringType == Vertical {
			if address >= 0x0000 && address <= 0x03FF {
				data = ppu.tableName[0][address&0x03FF]
			} else if address >= 0x0400 && address <= 0x07FF {
				data = ppu.tableName[1][address&0x03FF]
			} else if address >= 0x0800 && address <= 0x0BFF {
				data = ppu.tableName[0][address&0x03FF]
			} else if address >= 0x0C00 && address <= 0x0FFF {
				data = ppu.tableName[1][address&0x03FF]
			}
		} else if ppu.cart.MirroringType == Horizontal {
			if address >= 0x0000 && address <= 0x03FF {
				data = ppu.tableName[0][address&0x03FF]
			} else if address >= 0x0400 && address <= 0x07FF {
				data = ppu.tableName[0][address&0x03FF]
			} else if address >= 0x0800 && address <= 0x0BFF {
				data = ppu.tableName[1][address&0x03FF]
			} else if address >= 0x0C00 && address <= 0x0FFF {
				data = ppu.tableName[1][address&0x03FF]
			}
		}
	} else if address >= 0x3F00 && address <= 0x3FFF { // Palette memory
		address &= 0x001F
		// Mirroring
		if address == 0x0010 {
			address = 0x0000
		} else if address == 0x0014 {
			address = 0x0004
		} else if address == 0x0018 {
			address = 0x0008
		} else if address == 0x001C {
			address = 0x000C
		}

		data = ppu.tablePalette[address] & 0x3F
	}

	return data
}

// PPUWrite reads a value trigger by the PPU (CHR data)
func (ppu *PPU) PPUWrite(address uint16, data uint8) {
	address &= 0x3FFF // Guard address

	if ppu.cart.PPUWrite(address, data) { // Guard Cartridge
		return
	} else if address >= 0x0000 && address <= 0x1FFF { // Pattern memory
		ppu.tablePattern[(address&0x1000)>>12][address&0x0FFF] = data
	} else if address >= 0x2000 && address <= 0x3EFF { // Nametable memory
		address &= 0x0FFF
		if ppu.cart.MirroringType == Vertical {
			if address >= 0x0000 && address <= 0x03FF {
				ppu.tableName[0][address&0x03FF] = data
			} else if address >= 0x0400 && address <= 0x07FF {
				ppu.tableName[1][address&0x03FF] = data
			} else if address >= 0x0800 && address <= 0x0BFF {
				ppu.tableName[0][address&0x03FF] = data
			} else if address >= 0x0C00 && address <= 0x0FFF {
				ppu.tableName[1][address&0x03FF] = data
			}
		} else if ppu.cart.MirroringType == Horizontal {
			if address >= 0x0000 && address <= 0x03FF {
				ppu.tableName[0][address&0x03FF] = data
			} else if address >= 0x0400 && address <= 0x07FF {
				ppu.tableName[0][address&0x03FF] = data
			} else if address >= 0x0800 && address <= 0x0BFF {
				ppu.tableName[1][address&0x03FF] = data
			} else if address >= 0x0C00 && address <= 0x0FFF {
				ppu.tableName[1][address&0x03FF] = data
			}
		}
	} else if address >= 0x3F00 && address <= 0x3FFF { // Palette memory
		address &= 0x001F
		// Mirroring
		if address == 0x0010 {
			address = 0x0000
		} else if address == 0x0014 {
			address = 0x0004
		} else if address == 0x0018 {
			address = 0x0008
		} else if address == 0x001C {
			address = 0x000C
		}

		ppu.tablePalette[address] = data
	}
}

// InsertCartridge sets the current cartridge
func (ppu *PPU) InsertCartridge(c *Cartridge) {
	ppu.cart = c
}

// Texture returns the current texture rendered by the PPU
// This implements gui.TextureProvider
func (ppu *PPU) Texture(idx int) *image.RGBA {
	switch idx {
	case 0:
		return ppu.renderedTexture
	case 1:
		return ppu.spritePatternTableTextures[0]
	case 2:
		return ppu.spritePatternTableTextures[1]
	default:
		log.Infof("Invalid texture index selected: %v, defaulting to 0", idx)
		return ppu.renderedTexture
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
					color := ppu.getColorFromPalette(paletteID, pixel)
					ppu.spritePatternTableTextures[idx].Set(
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

func (ppu *PPU) getColorFromPalette(paletteID uint8, pixel uint8) *color.RGBA {
	address := uint16(0x3F00) + uint16(paletteID)<<2 + uint16(pixel)
	index := ppu.PPURead(address, false)
	return &ppu.colorPalette[index&0x3F]
}
