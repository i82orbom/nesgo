package nes

import "math/bits"

const (
	maxSpriteCount = 8
)

type oamMemory struct {
	address uint8 // Used when DMA (not used anymore later)
	data    [256]uint8
}

type spriteRenderInfo struct {
	spriteZeroRendered    bool
	spriteZeroHitPossible bool

	spriteCount      uint8
	spriteScanline   [maxSpriteCount]*oamMemoryEntry
	shifterPatternLO [maxSpriteCount]uint8
	shifterPatternHI [maxSpriteCount]uint8
}

// Copies the OAM data from the specified offset (not safe)
func (sr *spriteRenderInfo) copySpriteDataFromOAM(oam *oamMemory, offset int) {
	sr.spriteScanline[sr.spriteCount].y = oam.data[offset]
	sr.spriteScanline[sr.spriteCount].id = oam.data[offset+1]
	sr.spriteScanline[sr.spriteCount].attribute = oam.data[offset+2]
	sr.spriteScanline[sr.spriteCount].x = oam.data[offset+3]

	sr.spriteCount++
}

func newSpriteRenderInfo() *spriteRenderInfo {
	sprScanline := [maxSpriteCount]*oamMemoryEntry{}
	for i := 0; i < maxSpriteCount; i++ {
		sprScanline[i] = &oamMemoryEntry{}
	}
	return &spriteRenderInfo{
		spriteZeroRendered:    false,
		spriteZeroHitPossible: false,
		spriteScanline:        sprScanline,
		spriteCount:           0,
		shifterPatternHI:      [maxSpriteCount]uint8{},
		shifterPatternLO:      [maxSpriteCount]uint8{},
	}

}

func (sr *spriteRenderInfo) reset() {
	for i := 0; i < maxSpriteCount; i++ {
		sr.spriteScanline[i].reset()
		sr.shifterPatternLO[i] = 0
		sr.shifterPatternHI[i] = 0
	}
	sr.spriteZeroHitPossible = false
	sr.spriteCount = 0
}

type oamMemoryEntry struct {
	y         uint8
	id        uint8
	attribute uint8
	x         uint8
}

func newOamEntry() *oamMemoryEntry {
	return &oamMemoryEntry{
		y:         0xFF,
		id:        0xFF,
		attribute: 0xFF,
		x:         0xFF,
	}
}

func (e *oamMemoryEntry) reset() {
	e.y = 0xFF
	e.id = 0xFF
	e.attribute = 0xFF
	e.x = 0xFF
}

func (ppu *PPU) calculateSpritePatternAddressLO(currentSprite *oamMemoryEntry) uint16 {
	spritePatternAddressLO := uint16(0)

	// cast attributes to uint16 here to be cleaner
	spriteID := uint16(currentSprite.id)
	spriteY := uint16(currentSprite.y)
	spriteAttribute := uint16(currentSprite.attribute)
	currentScanline := uint16(ppu.scanline)

	isSpriteFlippedVertically := spriteAttribute&0x80 == 0

	// 8x8 sprite mode - the control register determines the pattern table
	if ppu.controlRegister.SpriteSize == 0 {

		if !isSpriteFlippedVertically {
			// Sprite is not flipped vertically
			spritePatternAddressLO =
				(uint16(ppu.controlRegister.PatternSprite) << 12) |
					(spriteID << 4) |
					(currentScanline - spriteY)
		} else {
			// Is flipped vertically
			spritePatternAddressLO =
				(uint16(ppu.controlRegister.PatternSprite) << 12) |
					(spriteID << 4) |
					(7 - (currentScanline - spriteY))
		}

	} else {
		// 8x16 sprite mode - The sprite attribute determines the pattern table
		if !isSpriteFlippedVertically {
			if currentScanline-spriteY < 8 {
				// Top half tile
				spritePatternAddressLO =
					((spriteID & 0x01) << 12) |
						((spriteID & 0xFE) << 4) |
						((currentScanline - spriteY) & 0x07)
			} else {
				// Bottom half tile
				spritePatternAddressLO =
					((spriteID & 0x01) << 12) |
						(((spriteID & 0xFE) + 1) << 4) |
						((currentScanline - spriteY) & 0x07)
			}
		} else {
			// Sprite is flipped vertically
			if currentScanline-spriteY < 8 {
				// Top half tile
				spritePatternAddressLO =
					((spriteID & 0x01) << 12) |
						(((spriteID & 0xFE) + 1) << 4) |
						(7-(currentScanline-spriteY))&0x07
			} else {
				// Bottom half tile
				spritePatternAddressLO =
					((spriteID & 0x01) << 12) |
						((spriteID & 0xFE) << 4) |
						(7-(currentScanline-spriteY))&0x07
			}
		}
	}
	return spritePatternAddressLO
}

func (ppu *PPU) foregroundRenderingCycle() {

	if ppu.cycle == 257 && ppu.scanline >= 0 {
		ppu.spriteRenderInfo.reset()

		currentOAMEntry := 0
		for currentOAMEntry < 256 && ppu.spriteRenderInfo.spriteCount < maxSpriteCount {
			diff := int(ppu.scanline) - int(ppu.oamMemory.data[currentOAMEntry]) // Y
			compareWith := 8
			if ppu.controlRegister.SpriteSize != 0 { // If the sprite mode is 8x16
				compareWith = 16
			}

			if diff >= 0 && diff < compareWith {
				if ppu.spriteRenderInfo.spriteCount < maxSpriteCount {
					if currentOAMEntry == 0 {
						ppu.spriteRenderInfo.spriteZeroHitPossible = true
					}
					ppu.spriteRenderInfo.copySpriteDataFromOAM(ppu.oamMemory, currentOAMEntry)
				}
			}
			currentOAMEntry += 4 // Every OAM entry is 4 bytes
		}

		if ppu.spriteRenderInfo.spriteCount > maxSpriteCount {
			ppu.statusRegister.SpriteOverflow = 1
		} else {
			ppu.statusRegister.SpriteOverflow = 0
		}
	}

	// Process found sprites
	if ppu.cycle == 340 {
		for i := uint8(0); i < ppu.spriteRenderInfo.spriteCount; i++ {

			spritePatternAddressLO := ppu.calculateSpritePatternAddressLO(ppu.spriteRenderInfo.spriteScanline[i])
			spritePatternAddressHI := spritePatternAddressLO + 8

			spritePatternBitsLO := ppu.PPURead(spritePatternAddressLO, false)
			spritePatternBitsHI := ppu.PPURead(spritePatternAddressHI, false)

			if ppu.spriteRenderInfo.spriteScanline[i].attribute&0x0040 != 0 {
				// Flip pattern horizontally
				spritePatternBitsLO = bits.Reverse8(spritePatternBitsLO)
				spritePatternBitsHI = bits.Reverse8(spritePatternBitsHI)
			}

			// Store calculated bits!
			ppu.spriteRenderInfo.shifterPatternLO[i] = spritePatternBitsLO
			ppu.spriteRenderInfo.shifterPatternHI[i] = spritePatternBitsHI
		}
	}

}

// renderBackgroundPixel returns pixel and paletteID for the current sprite, and a whether the foreground has priority
func (ppu *PPU) renderForegroundPixel() (uint8, uint8, bool) {
	foregroundPixel := uint8(0)
	foregroundPalette := uint8(0)
	foregroundPriority := false

	if ppu.isRenderingSprites() {
		ppu.spriteRenderInfo.spriteZeroRendered = false

		for i := uint8(0); i < ppu.spriteRenderInfo.spriteCount; i++ {
			spriteScanline := ppu.spriteRenderInfo.spriteScanline[i]

			if spriteScanline.x == 0 {
				foregroundPixelLO := uint8(0)
				if (ppu.spriteRenderInfo.shifterPatternLO[i] & 0x80) > 0 {
					foregroundPixelLO = 0x01
				}

				foregroundPixelHI := uint8(0)
				if (ppu.spriteRenderInfo.shifterPatternHI[i] & 0x80) > 0 {
					foregroundPixelHI = 0x01
				}

				foregroundPixel = (foregroundPixelHI << 1) | foregroundPixelLO

				foregroundPalette = (spriteScanline.attribute & 0x03) + 0x04
				foregroundPriority = spriteScanline.attribute&0x20 == 0

				if foregroundPixel != 0 {

					if i == 0 {
						ppu.spriteRenderInfo.spriteZeroRendered = true
					}
					break
				}
			}
		}
	}
	return foregroundPixel, foregroundPalette, foregroundPriority
}

// isRendering returns true if background and sprites should be rendered
func (ppu *PPU) isRenderingSprites() bool {
	return ppu.maskRegister.RenderSprites != 0
}
