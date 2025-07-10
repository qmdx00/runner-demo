package internal

import (
	"fmt"
	"runner-demo/internal/config"
)

type Position struct {
	X, Y   float64
	width  int
	height int
}

// NewPosition creates a new Position instance.
func NewPosition(x, y float64, w, h int) *Position {
	return &Position{
		X:      x,
		Y:      y,
		width:  w,
		height: h,
	}
}

func (p *Position) String() string {
	return fmt.Sprintf("Position{X: %f, Y: %f}", p.X, p.Y)
}

func (p *Position) MoveInWindow(dx, dy float64) {
	p.X += dx
	p.Y += dy

	// Ensure the position does not exceed the game boundaries
	if p.X < 0 {
		p.X = 0
	}
	if p.Y < 0 {
		p.Y = 0
	}
	if p.X+float64(p.width) > float64(config.Global.Game.Window.Width) {
		p.X = float64(config.Global.Game.Window.Width) - float64(p.width)
	}
	if p.Y+float64(p.height) > float64(config.Global.Game.Window.Height) {
		p.Y = float64(config.Global.Game.Window.Height) - float64(p.height)
	}
}
