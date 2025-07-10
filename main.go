package main

import (
	"runner-demo/internal"
	"runner-demo/internal/config"
	"runner-demo/internal/static"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	config.InitConfig()
	static.InitStatic()
	game := internal.NewGame()
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
