package gui

import (
	"image"

	"github.com/go-gl/glfw/v3.1/glfw"
)

// GameWindow represents a window with a canvas and provides a KeyCallback
type GameWindow interface {
	Draw()
	SetKeyCallback(fn KeyCallback)
	Destroy()
	ShouldClose() bool
}

// KeyCallback represents a keycallback function
type KeyCallback func(key glfw.Key, action glfw.Action)

// TextureProvider represents the interface to the outer gui world to set a texture in the gl context
type TextureProvider interface {
	Texture() *image.RGBA
}
