package main

import (
	"fmt"
	"log"
	"os"

	"github.com/i82orbom/nesgo/pkg/gui/glfw"
	"github.com/i82orbom/nesgo/pkg/nes"
	"github.com/i82orbom/nesgo/pkg/nesgo"
)

func main() {
	var (
		romPath string
		err     error
		cart    *nes.Cartridge
	)

	if romPath, err = romFromFlags(); err != nil {
		log.Fatalf(err.Error())
	}
	if cart, err = nes.NewCartridge(romPath); err != nil {
		log.Fatalf(err.Error())
	}

	console := nes.NewConsole()
	console.InsertCartridge(cart)
	console.Reset() // Reset signal to start emulating

	window, err := glfw.NewGameWindow(console.TextureProvider())
	if err != nil {
		log.Fatalf("Could not create gamewindow")
	}
	emulator := nesgo.NewEmulator(window, console)
	window.SetKeyCallback(emulator.KeyCallback)

	defer window.Destroy()
	for !window.ShouldClose() {
		emulator.Step()
		window.Draw()
	}
}

// Returns the path to the rom
func romFromFlags() (string, error) {
	if len(os.Args) == 1 {
		return "", fmt.Errorf("a path to a rom has to be specified")
	}
	return os.Args[1], nil
}
