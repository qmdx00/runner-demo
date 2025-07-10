package internal

import (
	"runner-demo/internal/config"
	"runner-demo/internal/event"
	"runner-demo/internal/state"
	"runner-demo/internal/static"

	"github.com/hajimehoshi/ebiten/v2"
)

type Runner struct {
	ticker *Ticker
	pos    *Position

	stateM *StateMachine

	runningSpeed float64
	jumpingSpeed float64
}

func NewRunner() *Runner {
	return &Runner{
		ticker:       NewTicker(),
		pos:          NewPosition(0, 16),
		stateM:       NewStateMachine(state.RunnerStateIdle),
		runningSpeed: 0.1, // 8.0 grid cell per frame
		jumpingSpeed: 0.2, // 2.0 grid cell per frame
	}
}

func (r *Runner) HandleInput(e event.RunnerControlEvent) error {
	if err := r.stateM.HandleEvent(e); err != nil {
		return err
	}

	var dx, dy float64
	switch e {
	case event.EventRun:
		dx = r.runningSpeed
	case event.EventJump:
		dy = -r.jumpingSpeed
	case event.EventStop:
	default:
	}

	r.pos.MoveInWindow(dx, dy)

	return nil
}

func (r *Runner) Render(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	spriteW, spriteH := config.Global.Game.Sprite.Width, config.Global.Game.Sprite.Height
	columns, rows := config.Global.Game.Grid.Columns, config.Global.Game.Grid.Rows
	cellWidth := screen.Bounds().Dx() / columns
	cellHeight := screen.Bounds().Dy() / rows

	op.GeoM.Scale(float64(cellWidth)/float64(spriteW), float64(cellHeight)/float64(spriteH))
	op.GeoM.Translate(float64(r.pos.X*float64(cellWidth)), float64(r.pos.Y*float64(cellHeight)))

	screen.DrawImage(r.imageByState(), op)
}

func (r *Runner) imageByState() *ebiten.Image {
	switch r.stateM.currentState {
	case state.RunnerStateIdle:
		return static.RunnerIdleSprite.FrameByTicker(int(r.ticker.counter))
	case state.RunnerStateRunning:
		return static.RunnerRunSprite.FrameByTicker(int(r.ticker.counter))
	case state.RunnerStateJumping:
		return static.RunnerJumpSprite.FrameByTicker(int(r.ticker.counter))
	default:
		panic("unknown runner state")
	}
}
