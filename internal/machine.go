package internal

import (
	"fmt"
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
			event.EventRun: {
				{event: event.EventRun, From: state.RunnerStateIdle, To: state.RunnerStateRunning},
				{event: event.EventRun, From: state.RunnerStateRunning, To: state.RunnerStateRunning},
			},
			event.EventStop: {
				{event: event.EventStop, From: state.RunnerStateRunning, To: state.RunnerStateIdle},
				{event: event.EventStop, From: state.RunnerStateJumping, To: state.RunnerStateIdle},
			},
			event.EventJump: {
				{event: event.EventJump, From: state.RunnerStateRunning, To: state.RunnerStateJumping},
				{event: event.EventJump, From: state.RunnerStateIdle, To: state.RunnerStateJumping},
			},
			event.EventPause: {
				{event: event.EventPause, From: state.RunnerStateRunning, To: state.RunnerStatePaused},
			},
			event.EventResume: {
				{event: event.EventResume, From: state.RunnerStatePaused, To: state.RunnerStateRunning},
			},
		},
	}
}

func (sm *StateMachine) CurrentState() state.RunnerState {
	return sm.currentState
}

func (sm *StateMachine) HandleEvent(event event.RunnerControlEvent) error {
	transitions, exists := sm.transitions[event]
	if !exists {
		return fmt.Errorf("event %v not supported", event)
	}

	for _, transition := range transitions {
		if transition.From == sm.currentState {
			sm.currentState = transition.To
			return nil
		}
	}
	return fmt.Errorf("invalid transition from %v on event %v", sm.currentState, event)
}
