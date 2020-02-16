package controllers

import (
	"github.com/go-gl/glfw/v3.1/glfw"
	"github.com/i82orbom/nesgo/pkg/nes"
)

// KeyboardController represents a keyboard input controller for the emulator
type KeyboardController struct {
	controller nes.InputController
	mappings   map[glfw.Key]nes.Button
	keys       []glfw.Key
}

// NewKeyboardController creates a new keyboard input controller
func NewKeyboardController(controller nes.InputController) *KeyboardController {
	return &KeyboardController{
		controller: controller,
	}
}

// WithDefaultMapping creates a new keyboard controller with the default mapping
// W-A-S-D for keypad, N,M for A and B, Enter and Backspace for Start and Select
func (k *KeyboardController) WithDefaultMapping() *KeyboardController {
	k.SetMapping(defaultMapping())
	return k
}

// GetKeyMapping returns the list of keys in the mapping
func (k *KeyboardController) GetKeyMapping() []glfw.Key {
	return k.keys
}

// SetMapping sets the key mapping
func (k *KeyboardController) SetMapping(mappings map[glfw.Key]nes.Button) {
	k.mappings = mappings
	k.keys = []glfw.Key{}
	for key := range mappings {
		k.keys = append(k.keys, key)
	}
}

// HandleKey handles the key input and returns true if it was sucessful, false if the key was not mapped
func (k *KeyboardController) HandleKey(key glfw.Key) bool {
	action, ok := k.mappings[key]
	if !ok {
		return false
	}
	k.controller.Press(action)
	return true
}

// Reset calls the reset method for the controller to start a read
func (k *KeyboardController) Reset() {
	k.controller.Reset()
}

func defaultMapping() map[glfw.Key]nes.Button {
	return map[glfw.Key]nes.Button{
		glfw.KeyW:         nes.ButtonUP,
		glfw.KeyS:         nes.ButtonDOWN,
		glfw.KeyA:         nes.ButtonLEFT,
		glfw.KeyD:         nes.ButtonRIGHT,
		glfw.KeyN:         nes.ButtonA,
		glfw.KeyM:         nes.ButtonB,
		glfw.KeyEnter:     nes.ButtonStart,
		glfw.KeyBackspace: nes.ButtonSelect,
	}
}
