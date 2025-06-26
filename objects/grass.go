package objects

import (
	"time"

	"github.com/mnmonherdene1234/uns-game/gameengine"
	"github.com/mnmonherdene1234/uns-game/gameengine/render"
	"github.com/mnmonherdene1234/uns-game/utils"
)

var GrassLimit = 10

type Grass struct {
	GameEngine                  *gameengine.GameEngine
	X                           float32
	Y                           float32
	W                           float32
	H                           float32
	Growth                      float32
	GrowthRate                  float32
	MaxWidth                    float32
	MaxHeight                   float32
	ReproductionCooldownSeconds float32
	ReproductedDate             time.Time
	ImageID                     uint32
}

func NewGrass(gameEngine *gameengine.GameEngine, x, y float32) *Grass {

	randomReproductionCooldownSeconds := utils.RandomFloat32(3.0, 15.0) // Random cooldown between 5 and 15 seconds

	return &Grass{
		GameEngine:                  gameEngine,
		X:                           x,
		Y:                           y,
		W:                           0,
		H:                           0,
		Growth:                      0.0,
		GrowthRate:                  0.001,
		MaxWidth:                    256 * 0.3,
		MaxHeight:                   247 * 0.3,
		ReproductionCooldownSeconds: randomReproductionCooldownSeconds, // Cooldown in seconds for reproduction
	}
}

func (g *Grass) Start() {
	if image, found := g.GameEngine.AssetsManager.GetImageByName("grass1"); found {
		g.ImageID = image.ID
	} else {
		g.GameEngine.Logger.Error("Failed to load grass texture: grass1 not found")
		return
	}
}

func totalGrassCount(gameEngine *gameengine.GameEngine) int {
	count := 0
	for _, obj := range gameEngine.Objects {
		if _, ok := obj.(*Grass); ok {
			count++
		}
	}
	return count
}

func (g *Grass) Update() {
	g.Growth += g.GrowthRate
	if g.Growth > 1.0 {
		g.Growth = 1.0
	}

	g.Reproduction()

	previousWidth := g.W
	g.W = g.MaxWidth * g.Growth
	widthDifference := g.W - previousWidth
	g.X -= widthDifference / 2

	previousHeight := g.H
	g.H = g.MaxHeight * g.Growth
	heightDifference := g.H - previousHeight
	g.Y -= heightDifference / 2

	if g.W < 0 || g.H < 0 {
		g.W = 0
		g.H = 0
	}
}

func (g *Grass) Render() {
	if g.GameEngine == nil || g.GameEngine.Window == nil {
		return
	}
	w, h := g.GameEngine.Window.GetSize()
	render.DrawTexturedQuadWithWindow(g.ImageID, g.X, g.Y, g.W, g.H, w, h)
}

func (g *Grass) Destroy() {
}

func (g *Grass) Reproduction() {
	if g.Growth < 1.0 {
		return // Grass is not fully grown, do not reproduce
	}

	if totalGrassCount(g.GameEngine) >= GrassLimit {
		return // Limit reached, do not reproduce
	}

	now := time.Now()

	if g.ReproductedDate.IsZero() || now.Sub(g.ReproductedDate).Seconds() >= float64(g.ReproductionCooldownSeconds) {
		g.ReproductedDate = now

		var maxDistance float32 = 150.0 // Maximum distance for reproduction

		randomX := utils.RandomFloat32(g.X-maxDistance, g.X+maxDistance)

		if randomX < 0 {
			randomX = 0
		}

		randomY := utils.RandomFloat32(g.Y-maxDistance, g.Y+maxDistance)

		if randomY < 0 {
			randomY = 0
		}

		g.GameEngine.AddObject(NewGrass(g.GameEngine, randomX, randomY))
	}
}
