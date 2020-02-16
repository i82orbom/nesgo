package nesgo

import (
	"fmt"

	"github.com/go-gl/glfw/v3.1/glfw"
	"github.com/i82orbom/nesgo/pkg/gui"
	"github.com/i82orbom/nesgo/pkg/nes"
	"github.com/i82orbom/nesgo/pkg/nesgo/controllers"
)

// Emulator represents a NES emulator handles the user input
type Emulator struct {
	console *nes.Console
	// Abstract to support joystick for example
	controller1 *controllers.KeyboardController

	window gui.GameWindow

	// Status
	emulationEnabled       bool
	disassemble            bool
	currentTexture         int
	currentSelectedPalette int
}

// NewEmulator creates a new instance of the emulator
func NewEmulator(window gui.GameWindow, console *nes.Console) *Emulator {
	return &Emulator{
		controller1:            controllers.NewKeyboardController(console.Controller1()).WithDefaultMapping(),
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
	case glfw.KeyR:
		fmt.Printf("Reset emulation\n")
		e.console.Reset()
	}
}

// Step makes the console step forward a frame
func (e *Emulator) Step() {
	if !e.emulationEnabled {
		return
	}
	// Read input from controllers
	e.sampleKeys()
	e.stepFrame()
	if e.disassemble {
		e.console.Disassemble()
	}
}

func (e *Emulator) sampleKeys() {
	e.controller1.Reset()
	for _, key := range e.controller1.GetKeyMapping() {
		if e.window.IsKeyPress(int(key)) {
			e.controller1.HandleKey(glfw.Key(key))
		}
	}
}

func (e *Emulator) stepFrame() {
	e.console.StepFrame()
}
