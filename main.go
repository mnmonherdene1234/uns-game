package main

import gameengine "uns/gameengine"

func main() {
	engine := gameengine.NewGameEngine()

	if err := engine.InitWindow(800, 600, "Game Engine Example"); err != nil {
		panic(err)
	}
	engine.Loop()
}
