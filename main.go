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

	player := objects.NewPlayer(engine, 1, "Player 1")

	engine.AddObject(player)

	engine.Loop()
}
