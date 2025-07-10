package main

import (
	"runner-demo/internal"
	"runner-demo/internal/config"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	config.InitConfig()
	game := internal.NewGame()
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
