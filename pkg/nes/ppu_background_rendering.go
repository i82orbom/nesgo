package nes

// backgroundTileInfo holds information of the background
// tile being rendered
type backgroundNextTileInfo struct {
	ID        uint8
	Attribute uint8
	ValLSB    uint8
	ValMSB    uint8
}

// background shifters holds the pattern and attribute shifters for background rendering
type backgroundShifters struct {
	patternLO uint16
	patternHI uint16

	attributeLO uint16
	attributeHI uint16
}

func (s *backgroundShifters) shift() {
	s.patternLO <<= 1
	s.patternHI <<= 1

	s.attributeLO <<= 1
	s.attributeHI <<= 1
}

func (ppu *PPU) backgroundRenderingCycle() {
	if ppu.cycle >= 2 && ppu.cycle < 258 || (ppu.cycle >= 321 && ppu.cycle < 338) {
		ppu.UpdateShifters()

		switch (ppu.cycle - 1) % 8 {
		case 0:
			// Prepare the background shifters to render the next tile and read the background tile ID
			ppu.LoadBackgroundShifters()
			ppu.bgNextTile.ID = ppu.PPURead((0x2000 | ppu.vramAddress.Value()&0x0FFF), false)
		case 2:
			// Get the background tile attribute
			address := 0x23C0 |
				(ppu.vramAddress.NameTableY << 11) |
				(ppu.vramAddress.NameTableX << 10) |
				((ppu.vramAddress.CoarseY >> 2) << 3) |
				(ppu.vramAddress.CoarseX >> 2)
			ppu.bgNextTile.Attribute = ppu.PPURead(address, false)
			if ppu.vramAddress.CoarseY&0x02 != 0 {
				ppu.bgNextTile.Attribute >>= 4
			}
			if ppu.vramAddress.CoarseX&0x02 != 0 {
				ppu.bgNextTile.Attribute >>= 2
			}
			ppu.bgNextTile.Attribute &= 0x03
		case 4:
			// Get the LSB Value of the tile
			address := uint16(ppu.controlRegister.PatternBackground)<<12 +
				(uint16(ppu.bgNextTile.ID) << 4) +
				ppu.vramAddress.FineY
			ppu.bgNextTile.ValLSB = ppu.PPURead(address, false)
		case 6:
			// Get the MSB Value of the tile
			address := uint16(ppu.controlRegister.PatternBackground)<<12 +
				(uint16(ppu.bgNextTile.ID) << 4) +
				ppu.vramAddress.FineY + 8 // Notice here the + 8
			ppu.bgNextTile.ValMSB = ppu.PPURead(address, false)
		case 7:
			ppu.IncrementScrollX()
		}
	}
	// END OF VISIBLE SCANLINE
	if ppu.cycle == 256 {
		ppu.IncrementScrollY()
	}

	// RESET X position
	if ppu.cycle == 257 {
		ppu.LoadBackgroundShifters()
		ppu.TransferAddressX()
	}
	// Read next background tile ID
	if ppu.cycle == 338 || ppu.cycle == 340 {
		ppu.bgNextTile.ID = ppu.PPURead(0x2000|(ppu.vramAddress.Value()&0x0FFF), false)
	}

	if ppu.scanline == -1 && ppu.cycle >= 280 && ppu.cycle < 305 {
		ppu.TransferAddressY()
	}
}

// renderBackgroundPixel will return the correct pixel and paletteID for the current background tile
func (ppu *PPU) renderBackgroundPixel() (uint8, uint8) {

	backgroundPixel := uint8(0x00)
	backgroundPalette := uint8(0x00)
	if ppu.isBackgroundRendering() {
		bitMux := uint16(0x8000) >> ppu.fineX

		p0Pixel := uint8(0x0)
		if (ppu.bgShifter.patternLO & bitMux) > 0 {
			p0Pixel = 0x01
		}
		p1Pixel := uint8(0x0)
		if (ppu.bgShifter.patternHI & bitMux) > 0 {
			p1Pixel = 0x01
		}
		backgroundPixel = (p1Pixel << 1) | p0Pixel

		bgPal0 := uint8(0x0)
		if (ppu.bgShifter.attributeLO & bitMux) > 0 {
			bgPal0 = 0x01
		}
		bgPal1 := uint8(0x0)
		if (ppu.bgShifter.attributeHI & bitMux) > 0 {
			bgPal1 = 0x01
		}
		backgroundPalette = (bgPal1 << 1) | bgPal0
	}
	return backgroundPixel, backgroundPalette
}

// UpdateShifters shifts one position the background shifters
func (ppu *PPU) UpdateShifters() {
	if ppu.maskRegister.RenderBackground != 0 {
		ppu.bgShifter.shift()
	}
}

// LoadBackgroundShifters prepares the shifters to render the next background tile
func (ppu *PPU) LoadBackgroundShifters() {
	ppu.bgShifter.patternLO = (ppu.bgShifter.patternLO & 0xFF00) | uint16(ppu.bgNextTile.ValLSB)
	ppu.bgShifter.patternHI = (ppu.bgShifter.patternHI & 0xFF00) | uint16(ppu.bgNextTile.ValMSB)

	tileAttribLSB := uint16(0x0000)
	if ppu.bgNextTile.Attribute&0b01 != 0 {
		tileAttribLSB = 0x00FF
	}
	tileAttribMSB := uint16(0x0000)
	if ppu.bgNextTile.Attribute&0b10 != 0 {
		tileAttribMSB = 0x00FF
	}

	ppu.bgShifter.attributeLO = (ppu.bgShifter.attributeLO & 0xFF00) | tileAttribLSB
	ppu.bgShifter.attributeHI = (ppu.bgShifter.attributeHI & 0xFF00) | tileAttribMSB
}

// IncrementScrollX updates variables for X-scrolling
func (ppu *PPU) IncrementScrollX() {
	if ppu.isRendering() {
		if ppu.vramAddress.CoarseX == 31 {
			ppu.vramAddress.CoarseX = 0

			ppu.vramAddress.FlipNameTableX()
		} else {
			ppu.vramAddress.CoarseX++
		}
	}
}

// IncrementScrollY update variables for Y-scrolling
func (ppu *PPU) IncrementScrollY() {
	if ppu.isRendering() {
		if ppu.vramAddress.FineY < 7 {
			ppu.vramAddress.FineY++
		} else {
			ppu.vramAddress.FineY = 0
			if ppu.vramAddress.CoarseY == 29 {
				ppu.vramAddress.CoarseY = 0

				ppu.vramAddress.FlipNameTableY()
			} else if ppu.vramAddress.CoarseY == 31 {
				ppu.vramAddress.CoarseY = 0
			} else {
				ppu.vramAddress.CoarseY++
			}
		}
	}
}

// TransferAddressX copies from the tempVRAM the nametable and coarseX variables
func (ppu *PPU) TransferAddressX() {
	if ppu.isRendering() {
		ppu.vramAddress.NameTableX = ppu.tramAddress.NameTableX
		ppu.vramAddress.CoarseX = ppu.tramAddress.CoarseX
	}
}

// TransferAddressY copies from the tempVRAM the nametable and coarseY variables and fineY
func (ppu *PPU) TransferAddressY() {
	if ppu.isRendering() {
		ppu.vramAddress.FineY = ppu.tramAddress.FineY
		ppu.vramAddress.NameTableY = ppu.tramAddress.NameTableY
		ppu.vramAddress.CoarseY = ppu.tramAddress.CoarseY
	}
}

func (ppu *PPU) isBackgroundRendering() bool {
	return ppu.maskRegister.RenderBackground != 0
}
