package nes

import (
	"image"
	"image/color"

	"github.com/i82orbom/nesgo/pkg/nes/register"
	"github.com/labstack/gommon/log"
)

const (
	width  = 240
	height = 256
)

// PPU represents the Picture Processing Unit of the NES
type PPU struct {
	cart *Cartridge

	nmiSignaled bool

	// Rendering flags
	cycle         int
	scanline      int
	frameComplete bool

	// Registers and data structures
	controlRegister *register.ControlRegister
	maskRegister    *register.MaskRegister
	statusRegister  *register.StatusRegister

	addressLatch  uint8
	ppuDataBuffer uint8
	fineX         uint8
	vramAddress   *register.VRAMRegister
	tramAddress   *register.VRAMRegister

	// Background variables
	bgShifter  backgroundShifters
	bgNextTile backgroundNextTileInfo

	// Foreground variables
	oamMemory        [256]uint8
	oamAddress       uint8
	spriteRenderInfo *spriteRenderInfo

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
	return ppu
}

func (ppu *PPU) reset() {
	// Init texture
	ppu.renderedTexture = image.NewRGBA(image.Rect(0, 0, height, width))
	ppu.spritePatternTableTextures = [2]*image.RGBA{
		image.NewRGBA(image.Rect(0, 0, 128, 128)),
		image.NewRGBA(image.Rect(0, 0, 128, 128)),
	}
	// Color palette
	ppu.colorPalette = colorPalette()

	// Registers
	ppu.controlRegister = &register.ControlRegister{}
	ppu.maskRegister = &register.MaskRegister{}
	ppu.statusRegister = &register.StatusRegister{}
	ppu.controlRegister.Set(0x00)
	ppu.maskRegister.Set(0x00)
	ppu.statusRegister.Set(0x00)

	ppu.addressLatch = 0x0
	ppu.ppuDataBuffer = 0x0
	ppu.vramAddress = &register.VRAMRegister{}
	ppu.tramAddress = &register.VRAMRegister{}
	ppu.vramAddress.Set(0x0000)
	ppu.tramAddress.Set(0x0000)

	ppu.oamMemory = [256]uint8{}
	ppu.oamAddress = 0x00
	ppu.spriteRenderInfo = newSpriteRenderInfo()
}

// CPURead reads a value triggered by the CPU
func (ppu *PPU) CPURead(address uint16, readOnly bool) uint8 {
	data := uint8(0)

	if readOnly {
		switch address {
		case 0x0000: // Control
			data = ppu.controlRegister.Value()
		case 0x0001: // Mask
			data = ppu.maskRegister.Value()
		case 0x0002: // Status
			data = ppu.statusRegister.Value()
		}
		// The rest is not readable (0x0003 - 0x0007)
		return data
	}

	switch address {
	case 0x0000: // Control - not readable
	case 0x0001: // Mask - not readable
	case 0x0002: // Status
		data = (ppu.statusRegister.Value() & 0xE0) | (ppu.ppuDataBuffer & 0x1F)
		ppu.statusRegister.VerticalBlank = 0
		ppu.addressLatch = 0
	case 0x0003: // OAM Address - not readable
	case 0x0004: // OAM Data
		data = ppu.oamMemory[ppu.oamAddress]
	case 0x0005: // Scroll
	case 0x0006: // PPU Address - not readable
	case 0x0007: // PPU Data
		data = ppu.ppuDataBuffer
		ppu.ppuDataBuffer = ppu.PPURead(ppu.vramAddress.Value(), false)

		if ppu.vramAddress.Value() >= 0x3F00 {
			data = ppu.ppuDataBuffer
		}

		if ppu.controlRegister.IncrementMode != 0 {
			ppu.vramAddress.Increment(32)
		} else {
			ppu.vramAddress.Increment(1)
		}
	}

	return data
}

// CPUWrite writes a value triggered by the CPU
func (ppu *PPU) CPUWrite(address uint16, data uint8) {
	switch address {
	case 0x0000: // Control
		ppu.controlRegister.Set(data)
		ppu.tramAddress.NameTableX = uint16(ppu.controlRegister.NameTableX)
		ppu.tramAddress.NameTableY = uint16(ppu.controlRegister.NameTableY)
	case 0x0001: // Mask
		ppu.maskRegister.Set(data)
	case 0x0002: // Status - not writable
	case 0x0003: // OAM Address
		ppu.oamAddress = data
	case 0x0004: // OAM Data
		ppu.oamMemory[ppu.oamAddress] = data
	case 0x0005: // Scroll
		if ppu.addressLatch == 0 {
			ppu.fineX = data & 0x07
			ppu.tramAddress.CoarseX = uint16(data) >> 3
			ppu.addressLatch = 1
		} else {
			ppu.tramAddress.FineY = uint16(data & 0x07)
			ppu.tramAddress.CoarseY = uint16(data) >> 3
			ppu.addressLatch = 0
		}
	case 0x0006: // PPU Address
		if ppu.addressLatch == 0 {
			ppu.tramAddress.Set((uint16(data&0x3F) << 8) | ppu.tramAddress.Value()&0x00FF)
			ppu.addressLatch = 1
		} else {
			ppu.tramAddress.Set(ppu.tramAddress.Value()&0xFF00 | uint16(data))
			ppu.vramAddress.Copy(ppu.tramAddress)
			ppu.addressLatch = 0
		}
	case 0x0007: // PPU Data
		ppu.PPUWrite(ppu.vramAddress.Value(), data)
		if ppu.controlRegister.IncrementMode != 0 {
			ppu.vramAddress.Increment(32)
		} else {
			ppu.vramAddress.Increment(1)
		}
	}
}

// PPURead reads a value triggered by the PPU (CHR data)
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

		if ppu.maskRegister.Grayscale == 1 {
			data = ppu.tablePalette[address] & 0x30
		} else {
			data = ppu.tablePalette[address] & 0x3F
		}
	}

	return data
}

// PPUWrite reads a value triggered by the PPU (CHR data)
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
func (ppu *PPU) Texture(idx ...int) *image.RGBA {
	paletteID := uint8(0)
	if len(idx) == 2 {
		paletteID = uint8(idx[1])
	}
	switch idx[0] {
	case 0:
		return ppu.renderedTexture
	case 1:
		// Refresh pattern table

		return ppu.patternTable(0, paletteID)
	case 2:
		// Refresh pattern table
		return ppu.patternTable(1, paletteID)
	default:
		log.Infof("Invalid texture index selected: %v, defaulting to 0", idx)
		return ppu.renderedTexture
	}
}
