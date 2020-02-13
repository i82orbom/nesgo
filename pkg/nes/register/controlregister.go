package register

type ControlRegister struct {
	EnableNMI         uint8 // 1000 0000
	SlaveMode         uint8 // 0100 0000
	SpriteSize        uint8 // 0010 0000
	PatternBackground uint8 // 0001 0000
	PatternSprite     uint8 // 0000 1000
	IncrementMode     uint8 // 0000 0100
	NameTableY        uint8 // 0000 0010
	NameTableX        uint8 // 0000 0001
}

func (r *ControlRegister) Value() uint8 {
	res := uint8(0x0)
	res |= r.EnableNMI << 7
	res |= (r.SlaveMode << 6)
	res |= (r.SpriteSize << 5)
	res |= (r.PatternBackground << 4)
	res |= (r.PatternSprite << 3)
	res |= (r.IncrementMode << 2)
	res |= (r.NameTableY << 1)
	res |= r.NameTableX
	return res
}

func (r *ControlRegister) Set(val uint8) {
	r.EnableNMI = bit(val, 7)
	r.SlaveMode = bit(val, 6)
	r.SpriteSize = bit(val, 5)
	r.PatternBackground = bit(val, 4)
	r.PatternSprite = bit(val, 3)
	r.IncrementMode = bit(val, 2)
	r.NameTableY = bit(val, 1)
	r.NameTableX = bit(val, 0)
}

func bit(n uint8, pos uint8) uint8 {
	val := n & (1 << pos)
	if val > 0 {
		return 1
	}
	return 0
}
