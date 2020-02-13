package register

type VRAMRegister struct {
	Unused     uint16 // 1000 0000 0000 0000 // 1
	FineY      uint16 // 0111 0000 0000 0000 // 3
	NameTableY uint16 // 0000 1000 0000 0000 // 1
	NameTableX uint16 // 0000 0100 0000 0000 // 1
	CoarseY    uint16 // 0000 0011 1110 0000 // 5
	CoarseX    uint16 // 0000 0000 0001 1111 // 5
}

func (r *VRAMRegister) Value() uint16 {
	res := uint16(0x0)
	res |= (r.Unused << 15)
	res |= (r.FineY << 12)
	res |= (r.NameTableY << 11)
	res |= (r.NameTableX << 10)
	res |= (r.CoarseY << 5)
	res |= (r.CoarseX)
	return res
}

func (r *VRAMRegister) Set(val uint16) {
	r.Unused = bit16(val, 15)
	r.FineY = (val & (0x7000)) >> 12
	r.NameTableY = bit16(val, 11)
	r.NameTableX = bit16(val, 10)
	r.CoarseY = (val & (0x03E0)) >> 5
	r.CoarseX = val & 0x001F
}

func (r *VRAMRegister) Increment(val int) {
	r.Set(r.Value() + uint16(val))
}

func (r *VRAMRegister) And(val uint16) {
	r.Set(r.Value() & val)
}

func (r *VRAMRegister) Copy(src *VRAMRegister) {
	r.Unused = src.Unused
	r.FineY = src.FineY
	r.NameTableY = src.NameTableY
	r.NameTableX = src.NameTableX
	r.CoarseY = src.CoarseY
	r.CoarseX = src.CoarseX
}

func bit16(n uint16, pos uint8) uint16 {
	val := n & (1 << pos)
	if val > 0 {
		return 1
	}
	return 0
}
