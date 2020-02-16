package nes

// Button represents a controller button
type Button int

const (
	// ButtonA is the NES A button
	ButtonA Button = iota
	// ButtonB is the NES B button
	ButtonB
	// ButtonStart is the NES Start button
	ButtonStart
	// ButtonSelect is the NES Select button
	ButtonSelect
	// ButtonUP is the NES Up button
	ButtonUP
	// ButtonDOWN is the NES Down button
	ButtonDOWN
	// ButtonLEFT is the NES Left button
	ButtonLEFT
	// ButtonRIGHT is the NES Right button
	ButtonRIGHT
)

// InputController represents an input controller that acts as an interface to
// the internal nes controller, the 'Presses' should be scanned from the input device
// as they will be buffered and scanned by the CPU and flushed when instructed
type InputController interface {
	Press(button Button)
	// Reset prepares the controller for another read
	Reset()
}

var _ InputController = &controller{}

type controller struct {
	state   uint8
	buttons uint8
}

func newController() *controller {
	return &controller{
		buttons: 0x0,
		state:   0x0,
	}
}

func (c *controller) Write() {
	c.state = c.buttons
}

func (c *controller) Data() uint8 {
	val := uint8(0)
	if c.state&0x80 > 0 {
		val = 1
	}
	c.state <<= 1
	return val
}

func (c *controller) Press(button Button) {
	switch button {
	case ButtonA:
		c.buttons |= 0x80
	case ButtonB:
		c.buttons |= 0x40
	case ButtonSelect:
		c.buttons |= 0x20
	case ButtonStart:
		c.buttons |= 0x10
	case ButtonUP:
		c.buttons |= 0x08
	case ButtonDOWN:
		c.buttons |= 0x04
	case ButtonLEFT:
		c.buttons |= 0x02
	case ButtonRIGHT:
		c.buttons |= 0x01
	}
}

func (c *controller) Reset() {
	c.buttons = 0x0
}
