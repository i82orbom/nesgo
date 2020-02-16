package nesgo

import (
	"fmt"

	"github.com/go-gl/glfw/v3.1/glfw"
	"github.com/i82orbom/nesgo/pkg/gui"
	"github.com/i82orbom/nesgo/pkg/nes"
)

// Emulator represents a NES emulator handles the user input
type Emulator struct {
	console *nes.Console
	window  gui.GameWindow

	// Status
	emulationEnabled       bool
	disassemble            bool
	currentTexture         int
	currentSelectedPalette int
}

// NewEmulator creates a new instance of the emulator
func NewEmulator(window gui.GameWindow, console *nes.Console) *Emulator {
	return &Emulator{
		window:                 window,
		console:                console,
		emulationEnabled:       false,
		currentTexture:         0,
		currentSelectedPalette: 0,
	}
}

// KeyCallback provides the input handler to interact with the emulator
func (e *Emulator) KeyCallback(key int, isPress bool) {
	if !isPress {
		return
	}
	// Currently the keys are harcoded
	switch glfw.Key(key) {
	case glfw.KeySpace:
		e.stepFrame()
		if e.disassemble {
			e.console.Disassemble()
		}
	case glfw.KeyE:
		e.emulationEnabled = !e.emulationEnabled
		fmt.Printf("Emulation enabled: %v\n", e.emulationEnabled)
	case glfw.KeyL:
		e.disassemble = !e.disassemble
		fmt.Printf("Dissasemble enabled: %v\n", e.disassemble)
	case glfw.KeyO:
		e.currentSelectedPalette++
		e.currentSelectedPalette %= 7
		e.window.SetTextureID(e.currentTexture, e.currentSelectedPalette)
		fmt.Printf("Current palette: %v\n", e.currentSelectedPalette)
	case glfw.KeyP:
		e.currentTexture++
		e.currentTexture %= 3
		e.window.SetTextureID(e.currentTexture, e.currentSelectedPalette)
		fmt.Printf("Current texture: %v\n", e.currentTexture)
	}
}

// Step makes the console step forward a frame
func (e *Emulator) Step() {
	if !e.emulationEnabled {
		return
	}
	e.stepFrame()
	if e.disassemble {
		e.console.Disassemble()
	}
}

func (e *Emulator) stepFrame() {
	e.console.StepFrame()
}
