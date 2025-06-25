package gameengine

import (
	"runtime"
	"time"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/mnmonherdene1234/uns-game/gameengine/render"
)

type GameEngine struct {
	Name      string
	Objects   []GameObject
	Window    *glfw.Window
	DeltaTime float64
}

func init() {
	runtime.LockOSThread() // Required for OpenGL
}

func NewGameEngine(name string) *GameEngine {
	return &GameEngine{Name: name, Objects: make([]GameObject, 0)}
}

func (ge *GameEngine) Start() error {
	if err := ge.InitWindow(1920, 1080, ge.Name); err != nil {
		return err
	}

	render.InitQuadShader()

	return nil
}

func (ge *GameEngine) InitWindow(width, height int, title string) error {
	if err := glfw.Init(); err != nil {
		return err
	}
	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)

	// Get the primary monitor for fullscreen
	monitor := glfw.GetPrimaryMonitor()
	mode := monitor.GetVideoMode()
	window, err := glfw.CreateWindow(mode.Width, mode.Height, title, monitor, nil)
	if err != nil {
		return err
	}
	window.MakeContextCurrent()
	if err := gl.Init(); err != nil {
		return err
	}
	ge.Window = window
	return nil
}

func (ge *GameEngine) Loop() {
	var lastTime = time.Now()
	for !ge.Window.ShouldClose() {
		now := time.Now()
		ge.DeltaTime = now.Sub(lastTime).Seconds()
		lastTime = now

		ge.input()
		ge.update()
		ge.render()
	}
	ge.Window.Destroy()
	glfw.Terminate()
}

func (ge *GameEngine) input() {
	if ge.Window.GetKey(glfw.KeyEscape) == glfw.Press {
		ge.Window.SetShouldClose(true)
	}
	// Add more input handling as needed
}

func (ge *GameEngine) update() {
	// Update game state logic here
	// This is where you would update positions, handle collisions, etc.
	for _, obj := range ge.Objects {
		obj.Update()
	}
}

func (ge *GameEngine) render() {
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
	gl.ClearColor(0, 0, 1, 0)
	gl.Clear(gl.COLOR_BUFFER_BIT)

	for _, obj := range ge.Objects {
		obj.Render()
	}

	ge.Window.SwapBuffers()
	glfw.PollEvents()
}

func (ge *GameEngine) AddObject(obj GameObject) {
	ge.Objects = append(ge.Objects, obj)
	obj.Start()
}

func (ge *GameEngine) RemoveObject(obj GameObject) {
	for i, o := range ge.Objects {
		if o == obj {
			o.Destroy()
			ge.Objects = append(ge.Objects[:i], ge.Objects[i+1:]...)
			return
		}
	}
}
