package main

import (
	"fmt"
	"log"
	"os"

	"github.com/i82orbom/nesgo/pkg/gui/glfw"
	"github.com/i82orbom/nesgo/pkg/nes"
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

	window, err := glfw.NewGameWindow(console)
	if err != nil {
		log.Fatalf("Could not create gamewindow")
	}
	defer window.Destroy()
	for !window.ShouldClose() {
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
