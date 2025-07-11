package internal

import (
	"runner-demo/internal/event"
	"runner-demo/internal/state"
)

type Transition struct {
	event    event.RunnerControlEvent
	From, To state.RunnerState
}

type StateMachine struct {
	currentState state.RunnerState
	transitions  map[event.RunnerControlEvent][]Transition
}

func NewStateMachine(currentState state.RunnerState) *StateMachine {
	return &StateMachine{
		currentState: currentState,
		transitions: map[event.RunnerControlEvent][]Transition{
			event.EventAccelerate: {
				{event: event.EventAccelerate, From: state.RunnerStateIdle, To: state.RunnerStateRunAccelerating},
				{event: event.EventAccelerate, From: state.RunnerStateRunDecelerating, To: state.RunnerStateRunAccelerating},
				{event: event.EventAccelerate, From: state.RunnerStateRunCruising, To: state.RunnerStateRunAccelerating},
				{event: event.EventAccelerate, From: state.RunnerStateRunStopping, To: state.RunnerStateRunAccelerating},
				{event: event.EventAccelerate, From: state.RunnerStateJumpLanding, To: state.RunnerStateRunAccelerating},
			},
			event.EventDecelerate: {
				{event: event.EventDecelerate, From: state.RunnerStateRunAccelerating, To: state.RunnerStateRunDecelerating},
				{event: event.EventDecelerate, From: state.RunnerStateRunCruising, To: state.RunnerStateRunDecelerating},
			},
			event.EventJump: {
				{event: event.EventJump, From: state.RunnerStateIdle, To: state.RunnerStateJumpCharging},
				{event: event.EventJump, From: state.RunnerStateRunAccelerating, To: state.RunnerStateJumpCharging},
				{event: event.EventJump, From: state.RunnerStateRunCruising, To: state.RunnerStateJumpCharging},
				{event: event.EventJump, From: state.RunnerStateRunDecelerating, To: state.RunnerStateJumpCharging},
				{event: event.EventJump, From: state.RunnerStateRunStopping, To: state.RunnerStateJumpCharging},
			},
		},
	}
}

func (sm *StateMachine) CurrentState() state.RunnerState {
	return sm.currentState
}

// func (sm *StateMachine) HandleEvent(event event.RunnerControlEvent) error {
// 	transitions, exists := sm.transitions[event]
// 	if !exists {
// 		return fmt.Errorf("event %v not supported", event)
// 	}

// 	for _, transition := range transitions {
// 		if transition.From == sm.currentState {
// 			sm.currentState = transition.To
// 			return nil
// 		}
// 	}
// 	return fmt.Errorf("invalid transition from %v on event %v", sm.currentState, event)
// }
