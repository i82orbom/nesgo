package gui

import (
	"image"
)

// GameWindow represents a window with a canvas and provides a KeyCallback
type GameWindow interface {
	Draw()
	SetKeyCallback(fn KeyCallback)
	Destroy()
	ShouldClose() bool
	SetTextureID(id int)
}

// KeyCallback represents a keycallback function
type KeyCallback func(key int, isPress bool)

// TextureProvider represents the interface to the outer gui world to set a texture in the gl context
type TextureProvider interface {
	Texture(idx int) *image.RGBA
}
