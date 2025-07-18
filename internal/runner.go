package internal

import (
	"runner-demo/internal/config"
	"runner-demo/internal/event"
	"runner-demo/internal/state"
	"runner-demo/internal/static"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	// HACK: hardcoded ground level for the runner test
	defaultGroundLevel = 16.0
)

type Runner struct {
	ticker *Ticker
	pos    *Position

	stateM *StateMachine

	maxVX        float64 // maximum horizontal velocity for running
	maxVY        float64 // maximum vertical velocity for jumping
	vX           float64 // horizontal velocity for running
	vY           float64 // vertical velocity for jumping
	acceleration float64 // acceleration rate for running

	latestUpdateAt    int64 // latest update time for the runner
	jumpChargeStartAt int64 // jump charge start time
	maxJumpChargeTime int64 // maximum time for jump charge
}

func NewRunner() *Runner {
	stateM := NewStateMachine(state.RunnerStateIdle)
	go stateM.HandleEvent()

	return &Runner{
		ticker: NewTicker(),
		pos:    NewPosition(0, defaultGroundLevel),
		stateM: stateM,

		maxVX:        0.5,  // 0.5 grid cell per 100ms
		maxVY:        1.0,  // 1.0 grid cell per 100ms
		acceleration: 0.03, // slightly slower acceleration for better control

		maxJumpChargeTime: 500,
	}
}

func (r *Runner) handleKeyBoardInput() {
	// move input handling
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) || ebiten.IsKeyPressed(ebiten.KeyD) {
		r.stateM.PushEvent(event.InputMoveRight)
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) || ebiten.IsKeyPressed(ebiten.KeyA) {
		r.stateM.PushEvent(event.InputMoveLeft)
	} else {
		r.stateM.PushEvent(event.InputMoveRelease)
	}

	// jump input handling
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		r.stateM.PushEvent(event.InputJumpPress)
	} else {
		r.stateM.PushEvent(event.InputJumpRelease)
	}
}

func (r *Runner) HandleStateTransitions() error {
	r.handleKeyBoardInput()

	if r.latestUpdateAt == 0 {
		r.latestUpdateAt = time.Now().UnixMilli()
	}

	switch r.stateM.CurrentState() {
	case state.RunnerStateIdle:
		r.vX = 0 // reset horizontal velocity
	case state.RunnerStateRunAccelerating:
		if r.vX < r.maxVX {
			r.vX += r.acceleration
		} else {
			r.vX = r.maxVX
			r.stateM.PushEvent(event.RunnerReachedMaxHorizontalSpeed)
		}
	case state.RunnerStateRunCruising:
		// Maintain maximum velocity
		r.vX = r.maxVX
	case state.RunnerStateRunDecelerating:
		if r.vX > 0 {
			r.vX -= r.acceleration
			if r.vX <= 0 {
				r.vX = 0
				r.stateM.PushEvent(event.RunnerHorizontalStopped)
			}
		}
	case state.RunnerStateRunStopped:
		r.vX = 0 // Reset horizontal velocity on stop
	case state.RunnerStateJumpCharging:
		// initialize jump charge start time if not set
		if r.jumpChargeStartAt == 0 {
			r.jumpChargeStartAt = time.Now().UnixMilli()
		}

		// check if the jump charge timeout
		chargeTime := time.Now().UnixMilli() - r.jumpChargeStartAt
		if chargeTime >= r.maxJumpChargeTime {
			r.stateM.PushEvent(event.RunnerJumpChargeTimeout)
		} else {
			if r.vY < r.maxVY {
				r.vY += r.acceleration
			} else {
				r.vY = r.maxVY
			}
		}
	case state.RunnerStateJumpRising:
		// reset jump charge start time
		r.jumpChargeStartAt = 0

		// Apply gravity
		r.vY -= r.acceleration * 1.5
		if r.vY <= 0 {
			r.vY = 0
			r.stateM.PushEvent(event.RunnerReachedMaxVerticalHeight)
		}
	case state.RunnerStateJumpFalling:
		// Apply gravity (falling)
		r.vY -= r.acceleration * 1.5

		// Check if we hit the ground
		groundLevel := defaultGroundLevel
		if r.pos.Y >= groundLevel {
			// r.pos.Y = groundLevel
			r.vY = 0
			r.stateM.PushEvent(event.RunnerVerticalLanded)
		}
	case state.RunnerStateJumpLanded:
		r.vY = 0
		// Reduce horizontal velocity on landing
		if r.vX > 0 {
			r.vX -= r.acceleration * 0.5
			if r.vX < 0 {
				r.vX = 0
				r.stateM.PushEvent(event.RunnerHorizontalStopped)
			}
		} else {
			r.vX = 0 // Reset horizontal velocity on landing
			r.stateM.PushEvent(event.RunnerHorizontalStopped)
		}
	default:
		// do nothing
	}

	// calculate the movement distance based on the velocity
	var dx, dy float64
	currentTime := time.Now().UnixMilli()
	dt := float64(currentTime-r.latestUpdateAt) / 100.0
	dx = r.vX * dt
	dy = -r.vY * dt // negative because positive vY should move up (negative screen coordinates)
	r.pos.MoveInWindow(dx, dy)

	// update the latest update time
	r.latestUpdateAt = currentTime

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
	return static.RunnerIdleSprite.FrameByStateAndTicker(r.stateM.CurrentState(), int(r.ticker.counter))
}
