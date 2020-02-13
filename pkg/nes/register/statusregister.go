package register

type StatusRegister struct {
	VerticalBlank  uint8 // 1000 0000
	SpriteZeroHit  uint8 // 0100 0000
	SpriteOverflow uint8 // 0010 0000
	Unused         uint8 // 0001 1111
}

func (r *StatusRegister) Value() uint8 {
	res := uint8(0x0)
	res |= (r.VerticalBlank << 7)
	res |= (r.SpriteZeroHit << 6)
	res |= (r.SpriteOverflow << 5)
	res |= (r.Unused)
	return res
}

func (r *StatusRegister) Set(val uint8) {
	r.VerticalBlank = bit(val, 7)
	r.SpriteZeroHit = bit(val, 6)
	r.SpriteOverflow = bit(val, 5)
	r.Unused = (0x1F & val)
}
