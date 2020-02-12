package glfw

import (
	"image"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/i82orbom/nesgo/pkg/gui"
)

// gameTexture represents a GL texture
type gameTexture struct {
	texture uint32
	window  gui.GameWindow
}

func newGameTexture(window gui.GameWindow) *gameTexture {
	texture := createTexture()

	// Initialize
	gl.ClearColor(0, 0, 0, 1)
	gl.Enable(gl.TEXTURE_2D)
	gl.Clear(gl.COLOR_BUFFER_BIT)

	return &gameTexture{
		texture: texture,
	}
}

// SetTexture sets the current texture
func (tex *gameTexture) SetTexture(texture *image.RGBA) {
	gl.BindTexture(gl.TEXTURE_2D, tex.texture)
	size := texture.Rect.Size()
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA,
		int32(size.X),
		int32(size.Y),
		0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(texture.Pix))
}

func createTexture() uint32 {
	texture := uint32(0)
	gl.GenTextures(1, &texture)
	gl.BindTexture(gl.TEXTURE_2D, texture)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	gl.BindTexture(gl.TEXTURE_2D, 0)
	return texture
}
