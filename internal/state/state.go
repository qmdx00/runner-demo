package state

type RunnerState uint8

const (
	RunnerStateIdle RunnerState = iota
	RunnerStateRunning
	RunnerStateJumping
	RunnerStatePaused
)

// String returns the string representation of the RunnerState.
func (s RunnerState) String() string {
	switch s {
	case RunnerStateIdle:
		return "Idle"
	case RunnerStateRunning:
		return "Running"
	case RunnerStateJumping:
		return "Jumping"
	case RunnerStatePaused:
		return "Paused"
	default:
		return "Unknown"
	}
}

// IsRunning checks if the RunnerState is Running.
func (s RunnerState) IsRunning() bool {
	return s == RunnerStateRunning
}

// IsPaused checks if the RunnerState is Paused.
func (s RunnerState) IsPaused() bool {
	return s == RunnerStatePaused
}
