package main

import (
	"github.com/mnmonherdene1234/uns-game/gameengine"
	"github.com/mnmonherdene1234/uns-game/objects"
)

func main() {
	engine := gameengine.NewGameEngine("UNS Game")

	if err := engine.Start(); err != nil {
		panic(err)
	}

	engine.AssetsManager.AddImage("./assets/images/apple.png", "apple")
	engine.AssetsManager.AddImage("./assets/images/grass1.png", "grass1")
	engine.AssetsManager.LoadImages()

	engine.AddObject(objects.NewGrass(engine, 1920/2, 1080/2))

	engine.Logger.Info("Game engine started successfully")

	engine.Loop()
}
