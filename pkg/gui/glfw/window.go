package glfw

import (
	"runtime"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.1/glfw"
	"github.com/i82orbom/nesgo/pkg/gui"
)

const (
	windowTitle = "NES GO"
	width       = 256
	height      = 240
)

type window struct {
	*glfw.Window
	texture          *gameTexture
	textureProvider  gui.TextureProvider
	currentTextureID int
}

// NewGameWindow creates a new gamewindow
func NewGameWindow(textureProvider gui.TextureProvider) (gui.GameWindow, error) {
	runtime.LockOSThread()

	if err := glfw.Init(); err != nil {
		return nil, err
	}
	if err := gl.Init(); err != nil {
		return nil, err
	}
	w := &window{}

	glfwWindow, err := glfw.CreateWindow(width, height, windowTitle, nil, nil)
	if err != nil {
		return nil, err
	}
	glfwWindow.MakeContextCurrent()
	w.Window = glfwWindow
	w.texture = newGameTexture(w)
	w.textureProvider = textureProvider
	w.currentTextureID = 0
	return w, nil
}

// SetTextureID allows to cycle the textures if the texture provider allows it
func (w *window) SetTextureID(id int) {
	w.currentTextureID = id
}

func (w *window) Draw() {
	glfw.PollEvents()

	currentTexture := w.textureProvider.Texture(w.currentTextureID)
	w.texture.SetTexture(currentTexture)

	width, height := w.getFrameBufferSizeF32()
	width /= 256
	height /= 240
	// Begin draw

	gl.Begin(gl.QUADS)

	gl.TexCoord2f(0, 1)
	gl.Vertex2f(-width, -height)

	gl.TexCoord2f(1, 1)
	gl.Vertex2f(width, -height)

	gl.TexCoord2f(1, 0)
	gl.Vertex2f(width, height)

	gl.TexCoord2f(0, 0)
	gl.Vertex2f(-width, height)

	gl.End()

	// End draw
	w.Window.SwapBuffers()
}

func (w *window) getFrameBufferSizeF32() (float32, float32) {
	width, height := w.Window.GetFramebufferSize()
	return float32(width), float32(height)
}

func (w *window) SetKeyCallback(fnCallback gui.KeyCallback) {
	w.Window.SetKeyCallback(func(window *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
		fnCallback(int(key), action == glfw.Press)
	})
}

func (w *window) Destroy() {
	glfw.Terminate()
}

func (w *window) ShouldClose() bool {
	return w.Window.ShouldClose()
}
