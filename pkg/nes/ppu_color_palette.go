package nes

import (
	"image/color"
)

func colorPalette() []color.RGBA {
	colors := make([]color.RGBA, 0x40)
	colors[0x00] = color.RGBA{84, 84, 84, 255}
	colors[0x01] = color.RGBA{0, 30, 116, 255}
	colors[0x02] = color.RGBA{8, 16, 144, 255}
	colors[0x03] = color.RGBA{48, 0, 136, 255}
	colors[0x04] = color.RGBA{68, 0, 100, 255}
	colors[0x05] = color.RGBA{92, 0, 48, 255}
	colors[0x06] = color.RGBA{84, 4, 0, 255}
	colors[0x07] = color.RGBA{60, 24, 0, 255}
	colors[0x08] = color.RGBA{32, 42, 0, 255}
	colors[0x09] = color.RGBA{8, 58, 0, 255}
	colors[0x0A] = color.RGBA{0, 64, 0, 255}
	colors[0x0B] = color.RGBA{0, 60, 0, 255}
	colors[0x0C] = color.RGBA{0, 50, 60, 255}
	colors[0x0D] = color.RGBA{0, 0, 0, 255}
	colors[0x0E] = color.RGBA{0, 0, 0, 255}
	colors[0x0F] = color.RGBA{0, 0, 0, 255}

	colors[0x10] = color.RGBA{152, 150, 152, 255}
	colors[0x11] = color.RGBA{8, 76, 196, 255}
	colors[0x12] = color.RGBA{48, 50, 236, 255}
	colors[0x13] = color.RGBA{92, 30, 228, 255}
	colors[0x14] = color.RGBA{136, 20, 176, 255}
	colors[0x15] = color.RGBA{160, 20, 100, 255}
	colors[0x16] = color.RGBA{152, 34, 32, 255}
	colors[0x17] = color.RGBA{120, 60, 0, 255}
	colors[0x18] = color.RGBA{84, 90, 0, 255}
	colors[0x19] = color.RGBA{40, 114, 0, 255}
	colors[0x1A] = color.RGBA{8, 124, 0, 255}
	colors[0x1B] = color.RGBA{0, 118, 40, 255}
	colors[0x1C] = color.RGBA{0, 102, 120, 255}
	colors[0x1D] = color.RGBA{0, 0, 0, 255}
	colors[0x1E] = color.RGBA{0, 0, 0, 255}
	colors[0x1F] = color.RGBA{0, 0, 0, 255}

	colors[0x20] = color.RGBA{236, 238, 236, 255}
	colors[0x21] = color.RGBA{76, 154, 236, 255}
	colors[0x22] = color.RGBA{120, 124, 236, 255}
	colors[0x23] = color.RGBA{176, 98, 236, 255}
	colors[0x24] = color.RGBA{228, 84, 236, 255}
	colors[0x25] = color.RGBA{236, 88, 180, 255}
	colors[0x26] = color.RGBA{236, 106, 100, 255}
	colors[0x27] = color.RGBA{212, 136, 32, 255}
	colors[0x28] = color.RGBA{160, 170, 0, 255}
	colors[0x29] = color.RGBA{116, 196, 0, 255}
	colors[0x2A] = color.RGBA{76, 208, 32, 255}
	colors[0x2B] = color.RGBA{56, 204, 108, 255}
	colors[0x2C] = color.RGBA{56, 180, 204, 255}
	colors[0x2D] = color.RGBA{60, 60, 60, 255}
	colors[0x2E] = color.RGBA{0, 0, 0, 255}
	colors[0x2F] = color.RGBA{0, 0, 0, 255}

	colors[0x30] = color.RGBA{236, 238, 236, 255}
	colors[0x31] = color.RGBA{168, 204, 236, 255}
	colors[0x32] = color.RGBA{188, 188, 236, 255}
	colors[0x33] = color.RGBA{212, 178, 236, 255}
	colors[0x34] = color.RGBA{236, 174, 236, 255}
	colors[0x35] = color.RGBA{236, 174, 212, 255}
	colors[0x36] = color.RGBA{236, 180, 176, 255}
	colors[0x37] = color.RGBA{228, 196, 144, 255}
	colors[0x38] = color.RGBA{204, 210, 120, 255}
	colors[0x39] = color.RGBA{180, 222, 120, 255}
	colors[0x3A] = color.RGBA{168, 226, 144, 255}
	colors[0x3B] = color.RGBA{152, 226, 180, 255}
	colors[0x3C] = color.RGBA{160, 214, 228, 255}
	colors[0x3D] = color.RGBA{160, 162, 160, 255}
	colors[0x3E] = color.RGBA{0, 0, 0, 255}
	colors[0x3F] = color.RGBA{0, 0, 0, 255}

	return colors
}
