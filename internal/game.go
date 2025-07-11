package internal

import (
	"fmt"
	"runner-demo/assets/maps"
	"runner-demo/internal/config"
	"runner-demo/internal/scenes"
	"runner-demo/internal/static"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var _ ebiten.Game = (*Game)(nil)

type Game struct {
	ticker *Ticker
	runner *Runner
	scene  *scenes.Scene
}

// NewGame creates a new instance of Game.
func NewGame() *Game {
	ebiten.SetWindowTitle(config.Global.Game.Title)
	ebiten.SetWindowSize(config.Global.Game.Window.Width, config.Global.Game.Window.Height)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	return &Game{
		ticker: NewTicker(),
		runner: NewRunner(),
		// scene:  scenes.NewDefaultScene(maps.MAP_1_Water, maps.MAP_1_Soil),
		scene: scenes.NewDefaultScene(maps.MAP_1_Test),
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
	// columns, rows := config.Global.Game.Grid.Columns, config.Global.Game.Grid.Rows
	// cellWidth := screen.Bounds().Dx() / columns
	// cellHeight := screen.Bounds().Dy() / rows
	// for y := range rows {
	// 	for x := range columns {
	// 		opt := &ebiten.DrawImageOptions{}
	// 		opt.GeoM.Translate(float64(x*cellWidth), float64(y*cellHeight))
	// 		cell := ebiten.NewImage(cellWidth, cellHeight)
	// 		random := uint8(rand.IntN(255)) // Random value for color
	// 		cell.Fill(color.RGBA{R: random, G: random, B: random, A: 255})
	// 		screen.DrawImage(cell, opt)
	// 	}
	// }

	g.drawBackground(screen, static.BackgroundImage_png)
	g.scene.Render(screen)
	g.runner.Render(screen)

	ebitenutil.DebugPrint(screen, fmt.Sprintf(
		`Welcome to the Runner Game!
Window Size: (%d, %d)
Position: (%.2f, %.2f)
Velocity: (%.2f, %.2f)
Max Velocity: (%.2f, %.2f)
State: %s`,
		screen.Bounds().Dx(), screen.Bounds().Dy(),
		g.runner.pos.X, g.runner.pos.Y,
		g.runner.vX, g.runner.vY,
		g.runner.maxVX, g.runner.maxVY,
		g.runner.stateM.CurrentState()))
}

// Update ...
func (g *Game) Update() error {
	g.ticker.Tick()

	g.runner.HandleStateTransitions()

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
