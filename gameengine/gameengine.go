package gameengine

import (
	"runtime"
	"time"

	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

type GameEngine struct {
	window    *glfw.Window
	deltaTime float64
}

func init() {
	runtime.LockOSThread() // Required for OpenGL
}

func NewGameEngine() *GameEngine {
	return &GameEngine{}
}

func (ge *GameEngine) InitWindow(width, height int, title string) error {
	if err := glfw.Init(); err != nil {
		return err
	}
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 6)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	window, err := glfw.CreateWindow(width, height, title, nil, nil)
	if err != nil {
		return err
	}
	window.MakeContextCurrent()
	if err := gl.Init(); err != nil {
		return err
	}
	ge.window = window
	return nil
}

func (ge *GameEngine) Loop() {
	var lastTime = time.Now()
	for !ge.window.ShouldClose() {
		now := time.Now()
		ge.deltaTime = now.Sub(lastTime).Seconds()
		lastTime = now

		ge.input()
		ge.update()
		ge.render()
	}
	ge.window.Destroy()
	glfw.Terminate()
}

func (ge *GameEngine) input() {
	if ge.window.GetKey(glfw.KeyEscape) == glfw.Press {
		ge.window.SetShouldClose(true)
	}
	// Add more input handling as needed
}

func (ge *GameEngine) update() {
	// Update game state logic here
	// This is where you would update positions, handle collisions, etc.
	println("Updating game state, deltaTime:", ge.deltaTime)
}

func (ge *GameEngine) render() {
	// OpenGL rendering
	gl.ClearColor(0.9, 0.3, 0.9, 1.0)
	gl.Clear(gl.COLOR_BUFFER_BIT)

	ge.window.SwapBuffers()
	glfw.PollEvents()
}
