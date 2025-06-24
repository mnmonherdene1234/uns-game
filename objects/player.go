package objects

import (
	"github.com/mnmonherdene1234/uns-game/gameengine"
	"github.com/mnmonherdene1234/uns-game/gameengine/render"
)

type Player struct {
	GameEngine *gameengine.GameEngine
	ID         int
	Name       string
	Apple      uint32
	X          float32
}

func NewPlayer(gameEngine *gameengine.GameEngine, id int, name string) *Player {
	return &Player{
		GameEngine: gameEngine,
		ID:         id,
		Name:       name,
	}
}

func (p *Player) Start() {
	var err error

	p.Apple, err = render.LoadTexture("./assets/apple.png")
	println("Loading apple texture:", p.Apple)

	if err != nil {
		println("Failed to load apple texture: %v", err)
		return
	}
}

func (p *Player) Update() {
	p.X += 300 * float32(p.GameEngine.DeltaTime)
}

func (p *Player) Render() {
	if p.GameEngine == nil || p.GameEngine.Window == nil {
		return
	}
	w, h := p.GameEngine.Window.GetSize()
	render.DrawTexturedQuadWithWindow(p.Apple, p.X, 0, 100, 100, w, h)
}

func (p *Player) Destroy() {
}
