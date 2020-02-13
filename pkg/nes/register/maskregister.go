package register

type MaskRegister struct {
	EnhanceBlue          uint8 // 1000 0000
	EnhanceGreen         uint8 // 0100 0000
	EnhanceRed           uint8 // 0010 0000
	RenderSprites        uint8 // 0001 0000
	RenderBackground     uint8 // 0000 1000
	RenderSpritesLeft    uint8 // 0000 0100
	RenderBackgroundLeft uint8 // 0000 0010
	Grayscale            uint8 // 0000 0001
}

func (r *MaskRegister) Value() uint8 {
	res := uint8(0x0)
	res |= (r.EnhanceBlue << 7)
	res |= (r.EnhanceGreen << 6)
	res |= (r.EnhanceRed << 5)
	res |= (r.RenderSprites << 4)
	res |= (r.RenderBackground << 3)
	res |= (r.RenderSpritesLeft << 2)
	res |= (r.RenderBackgroundLeft << 1)
	res |= (r.Grayscale)
	return res
}

func (r *MaskRegister) Set(val uint8) {
	r.EnhanceBlue = bit(val, 7)
	r.EnhanceGreen = bit(val, 6)
	r.EnhanceRed = bit(val, 5)
	r.RenderSprites = bit(val, 4)
	r.RenderBackground = bit(val, 3)
	r.RenderSpritesLeft = bit(val, 2)
	r.RenderBackgroundLeft = bit(val, 1)
	r.Grayscale = bit(val, 0)
}
