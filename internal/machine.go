package internal

import (
	"fmt"
	"runner-demo/internal/event"
	"runner-demo/internal/state"
)

type StateMachine struct {
	currentState state.RunnerState
	eventBus     *event.EventBus
}

func NewStateMachine(currentState state.RunnerState) *StateMachine {
	return &StateMachine{
		currentState: currentState,
		eventBus:     event.NewEventBus(),
	}
}

func (sm *StateMachine) PushEvent(e event.RunnerEvent) {
	sm.eventBus.Push(e)
}

func (sm *StateMachine) CurrentState() state.RunnerState {
	return sm.currentState
}

func (sm *StateMachine) HandleEvent() {
	for {
		e, ok := sm.eventBus.Pop()
		if !ok {
			break
		}
		sm.transition(e)
	}
}

func (sm *StateMachine) transition(e event.RunnerEvent) {
	fromState := sm.currentState
	toState := fromState // Default to current state (no transition)

	switch fromState {
	case state.RunnerStateIdle:
		switch e {
		case event.InputMoveRight:
			toState = state.RunnerStateRunAccelerating
		case event.InputJumpPress:
			toState = state.RunnerStateJumpCharging
		}
	case state.RunnerStateRunAccelerating:
		switch e {
		case event.RunnerReachedMaxHorizontalSpeed:
			toState = state.RunnerStateRunCruising
		case event.InputMoveRelease:
			toState = state.RunnerStateRunDecelerating
		case event.InputJumpPress:
			toState = state.RunnerStateJumpCharging
		}
	case state.RunnerStateRunCruising:
		switch e {
		case event.InputMoveLeft, event.InputMoveRelease:
			toState = state.RunnerStateRunDecelerating
		case event.InputJumpPress:
			toState = state.RunnerStateJumpCharging
		}
	case state.RunnerStateRunDecelerating:
		switch e {
		case event.InputMoveRight:
			toState = state.RunnerStateRunAccelerating
		case event.InputJumpPress:
			toState = state.RunnerStateJumpCharging
		case event.RunnerHorizontalStopped:
			toState = state.RunnerStateRunStopped
		}
	case state.RunnerStateRunStopped:
		switch e {
		case event.InputMoveRight:
			toState = state.RunnerStateRunAccelerating
		case event.InputMoveRelease:
			toState = state.RunnerStateIdle
		case event.InputJumpPress:
			toState = state.RunnerStateJumpCharging
		}
	case state.RunnerStateJumpCharging:
		switch e {
		case event.InputJumpRelease:
			toState = state.RunnerStateJumpRising
		}
	case state.RunnerStateJumpRising:
		switch e {
		case event.RunnerReachedMaxVerticalHeight:
			toState = state.RunnerStateJumpFalling
		}
	case state.RunnerStateJumpFalling:
		switch e {
		case event.RunnerVerticalLanded:
			toState = state.RunnerStateJumpLanded
		}
	case state.RunnerStateJumpLanded:
		switch e {
		case event.InputMoveRight:
			toState = state.RunnerStateRunAccelerating
		case event.InputJumpPress:
			toState = state.RunnerStateJumpCharging
		case event.InputMoveRelease:
			toState = state.RunnerStateIdle
		}
	}

	sm.currentState = toState

	// DEBUG: Log the state transition
	if fromState != toState {
		fmt.Printf("Transitioning from [%s] to [%s] due to event [%s]\n", fromState, toState, e)
	}
}
