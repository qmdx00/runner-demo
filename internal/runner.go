package internal

import (
	"runner-demo/internal/config"
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

	maxVX        float64 // maximum horizontal velocity for running
	maxVY        float64 // maximum vertical velocity for jumping
	vX           float64 // horizontal velocity for running
	vY           float64 // vertical velocity for jumping
	acceleration float64 // acceleration rate for running
	latestUpdate int64   // latest update time for the runner
}

func NewRunner() *Runner {
	stateM := NewStateMachine(state.RunnerStateIdle)
	go stateM.HandleEvent()

	return &Runner{
		ticker: NewTicker(),
		pos:    NewPosition(0, 16),
		stateM: stateM,

		maxVX:        0.5,  // 0.1 grid cell per 100ms
		maxVY:        1.0,  // 0.2 grid cell per 100ms
		acceleration: 0.05, // acceleration rate for running
	}
}

func (r *Runner) HandleStateTransitions() error {
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

	// if r.vY > 0 {
	// 	r.stateM.PushEvent(event.RunnerVelocityYNegative)
	// } else if r.vY < 0 {
	// 	r.stateM.PushEvent(event.RunnerGrounded)
	// } else {
	// 	r.stateM.PushEvent(event.RunnerLanded)
	// }

	if r.latestUpdate == 0 {
		r.latestUpdate = time.Now().UnixMilli()
	}

	switch r.stateM.CurrentState() {
	case state.RunnerStateRunAccelerating:
		if r.vX < r.maxVX {
			r.vX += r.acceleration
		} else {
			// Velocity has reached maximum, transition to cruising
			r.vX = r.maxVX
			r.stateM.PushEvent(event.InputMoveRight) // This will trigger transition to cruising
		}
	case state.RunnerStateRunCruising:
		// Maintain maximum velocity
		r.vX = r.maxVX
	case state.RunnerStateRunDecelerating:
		if r.vX > 0 {
			r.vX -= r.acceleration
			if r.vX <= 0 {
				r.vX = 0
				// Transition to idle when velocity reaches zero
				r.stateM.currentState = state.RunnerStateIdle
			}
		}
	case state.RunnerStateJumpCharging:
		if r.vY < r.maxVY {
			r.vY += r.acceleration
		} else {
			// Velocity has reached maximum, transition to rising
			r.vY = r.maxVY
			r.stateM.PushEvent(event.InputJumpRelease)
		}
	case state.RunnerStateJumpRising:
		if r.vY > 0 {
			r.vY -= r.acceleration // simulate gravity
		} else {
			// Transition to falling when velocity becomes zero or negative
			r.stateM.currentState = state.RunnerStateJumpFalling
		}
	case state.RunnerStateJumpFalling:
		if r.vY < 0 {
			r.vY += r.acceleration // simulate gravity
		} else {
			r.stateM.PushEvent(event.RunnerLanded)
			r.vY = 0 // reset vertical velocity
		}
	case state.RunnerStateIdle:
		r.vX = 0 // reset horizontal velocity
	default:
	}

	// calculate the movement distance based on the velocity
	var dx, dy float64
	currentTime := time.Now().UnixMilli()
	dt := float64(currentTime-r.latestUpdate) / 100.0
	dx = r.vX * dt
	dy = -r.vY * dt // negative because positive vY should move up (negative screen coordinates)
	r.pos.MoveInWindow(dx, dy)

	// update the latest update time
	r.latestUpdate = currentTime

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
	case state.RunnerStateRunAccelerating, state.RunnerStateRunCruising, state.RunnerStateRunDecelerating, state.RunnerStateRunStopping:
		return static.RunnerRunSprite.FrameByTicker(int(r.ticker.counter))
	case state.RunnerStateJumpCharging, state.RunnerStateJumpRising, state.RunnerStateJumpFalling, state.RunnerStateJumpLanding:
		return static.RunnerJumpSprite.FrameByTicker(int(r.ticker.counter))
	default:
		panic("unknown runner state")
	}
}
