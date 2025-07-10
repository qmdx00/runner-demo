package internal

import (
	"fmt"
	"runner-demo/internal/config"
	"runner-demo/internal/event"
	"runner-demo/internal/input"
	"runner-demo/internal/static"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var _ ebiten.Game = (*Game)(nil)

type Game struct {
	ticker *Ticker
	runner *Runner
}

// NewGame creates a new instance of Game.
func NewGame() *Game {
	ebiten.SetWindowTitle(config.Global.Game.Title)
	ebiten.SetWindowSize(config.Global.Game.Window.Width, config.Global.Game.Window.Height)

	return &Game{
		ticker: NewTicker(),
		runner: NewRunner(),
	}
}

// Layout ...
func (g *Game) Layout(outsideWidth int, outsideHeight int) (screenWidth int, screenHeight int) {
	scale := ebiten.Monitor().DeviceScaleFactor()
	return int(float64(outsideWidth) * scale), int(float64(outsideHeight) * scale)
}

// Draw ...
func (g *Game) Draw(screen *ebiten.Image) {
	// // render the background grid for debugging purposes
	// for y := 0; y < config.Global.Game.Window.Height; y += config.Global.Game.Cell.Height {
	// 	for x := 0; x < config.Global.Game.Window.Width; x += config.Global.Game.Cell.Width {
	// 		opt := &ebiten.DrawImageOptions{}
	// 		opt.GeoM.Translate(float64(x), float64(y))
	// 		cell := ebiten.NewImage(config.Global.Game.Cell.Width, config.Global.Game.Cell.Height)
	// 		random := uint8(rand.IntN(255)) // Random value for color
	// 		cell.Fill(color.RGBA{R: random, G: random, B: random, A: 255})
	// 		screen.DrawImage(cell, opt)
	// 	}
	// }

	// for index := range static.RunnerIdleSprite.Frames() {
	// 	ops := &ebiten.DrawImageOptions{}
	// 	ops.GeoM.Translate(float64(index*static.RunnerIdleSprite.FrameWidth), 100)
	// 	screen.DrawImage(static.RunnerIdleSprite.FrameByTicker(index), ops)
	// }

	// for index := range static.RunnerRunSprite.Frames() {
	// 	ops := &ebiten.DrawImageOptions{}
	// 	ops.GeoM.Translate(float64(index*static.RunnerRunSprite.FrameWidth), 200)
	// 	screen.DrawImage(static.RunnerRunSprite.FrameByTicker(index), ops)
	// }

	g.drawBackground(screen, static.BackgroundImage_png)
	g.runner.Render(screen)

	ebitenutil.DebugPrint(screen, fmt.Sprintf("Welcome to the Runner Game!\nUse arrow keys to move the player.\nPosition: (%.2f, %.2f)\nState: %s", g.runner.pos.X, g.runner.pos.Y, g.runner.stateM.CurrentState()))
}

// Update ...
func (g *Game) Update() error {
	g.ticker.Tick()

	var controlEvent event.RunnerControlEvent = event.EventStop
	for key, e := range input.KeyboardEventMap {
		if ebiten.IsKeyPressed(key) {
			controlEvent = e
		}
	}
	g.runner.HandleInput(controlEvent)

	return nil
}

func (g *Game) drawBackground(screen *ebiten.Image, backgroundImage *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	bgBounds := backgroundImage.Bounds()
	scaleX := float64(screen.Bounds().Dx()) / float64(bgBounds.Dx())
	scaleY := float64(screen.Bounds().Dy()) / float64(bgBounds.Dy())
	op.GeoM.Scale(scaleX, scaleY)
	screen.DrawImage(backgroundImage, op)
}
