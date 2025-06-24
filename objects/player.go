package objects

import (
	"github.com/mnmonherdene1234/uns-game/gameengine"
	"github.com/mnmonherdene1234/uns-game/gameengine/render"
)

type Player struct {
	GameEngine *gameengine.GameEngine
	ID         int
	Name       string
}

func NewPlayer(gameEngine *gameengine.GameEngine, id int, name string) *Player {
	return &Player{
		GameEngine: gameEngine,
		ID:         id,
		Name:       name,
	}
}

func (p Player) Start() {
	println("Player started:", p.Name, "with ID:", p.ID)
}

func (p Player) Update() {
}

func (p Player) Render() {
	triangle := render.NewTriangle2D([]float32{
		-0.5, -0.5,
		0.5, -0.5,
		0.0, 0.5,
	})
	triangle.Draw()
}

func (p Player) Destroy() {
}
