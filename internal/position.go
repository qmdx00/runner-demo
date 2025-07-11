package internal

import (
	"fmt"
	"runner-demo/internal/config"
)

// Position represents a position in the game grid.

type Position struct {
	X, Y          float64
	columns, rows int
}

// NewPosition creates a new Position instance.
func NewPosition(x, y float64) *Position {
	return &Position{
		X:       x,
		Y:       y,
		columns: config.Global.Game.Grid.Columns,
		rows:    config.Global.Game.Grid.Rows,
	}
}

func (p *Position) String() string {
	return fmt.Sprintf("Position{X: %f, Y: %f}", p.X, p.Y)
}

func (p *Position) MoveInWindow(dx, dy float64) {
	p.X += dx
	p.Y += dy

	// Ensure the position does not out of game grid bounds
	if p.X < 0 {
		p.X = 0
	}
	if p.Y < 0 {
		p.Y = 0
	}
	if p.X > float64(p.columns-1) {
		p.X = float64(p.columns - 1)
	}
	if p.Y > float64(p.rows-1) {
		p.Y = float64(p.rows - 1)
	}
}

// IsOnGround checks if the position is on the ground or on a platform
func (p *Position) IsOnGround(groundLevel float64) bool {
	// Check if on ground
	if p.Y >= groundLevel {
		return true
	}
	return false
}
