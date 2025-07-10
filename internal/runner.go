package internal

import (
	"runner-demo/internal/event"
	"runner-demo/internal/state"
	"runner-demo/internal/static"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Runner struct {
	ticker *Ticker
	pos    *Position

	stateM *StateMachine

	runningSpeed float64
	jumpingSpeed float64
	lastUpdate   int64
}

func NewRunner() *Runner {
	return &Runner{
		ticker:       NewTicker(),
		pos:          NewPosition(0, 320, static.RunnerIdleSprite.FrameWidth, static.RunnerIdleSprite.FrameHeight),
		stateM:       NewStateMachine(state.RunnerStateIdle),
		runningSpeed: 8.0,  // pixels per 100 millisecond
		jumpingSpeed: 10.0, // pixels per 100 millisecond
	}
}

func (r *Runner) HandleInput(e event.RunnerControlEvent) error {
	if err := r.stateM.HandleEvent(e); err != nil {
		return err
	}

	if r.lastUpdate == 0 {
		r.lastUpdate = time.Now().UnixMilli()
	}

	now := time.Now().UnixMilli()
	dt := now - r.lastUpdate
	r.lastUpdate = now

	moveDistance := float64(dt) * r.runningSpeed / 100.0
	jumpDistance := float64(dt) * r.jumpingSpeed / 100.0

	var dx, dy float64
	switch e {
	case event.EventRun:
		dx = moveDistance
	case event.EventJump:
		dy = -jumpDistance
	case event.EventStop:
	default:
	}

	r.pos.MoveInWindow(dx, dy)

	return nil
}

func (r *Runner) Render(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(r.pos.X), float64(r.pos.Y))
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
