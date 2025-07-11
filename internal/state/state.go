package state

type RunnerState uint8

const (
	RunnerStateIdle RunnerState = iota

	RunnerStateRunAccelerating
	RunnerStateRunCruising
	RunnerStateRunDecelerating
	RunnerStateRunStopping

	RunnerStateJumpCharging
	RunnerStateJumpRising
	RunnerStateJumpFalling
	RunnerStateJumpLanding
)

// String returns the string representation of the RunnerState.
func (s RunnerState) String() string {
	switch s {
	case RunnerStateIdle:
		return "Idle"
	case RunnerStateRunAccelerating:
		return "Run Accelerating"
	case RunnerStateRunCruising:
		return "Run Cruising"
	case RunnerStateRunDecelerating:
		return "Run Decelerating"
	case RunnerStateRunStopping:
		return "Run Stopping"
	case RunnerStateJumpCharging:
		return "Jump Charging"
	case RunnerStateJumpRising:
		return "Jump Rising"
	case RunnerStateJumpFalling:
		return "Jump Falling"
	case RunnerStateJumpLanding:
		return "Jump Landing"
	default:
		return "Unknown State"
	}
}

// IsRunning checks if the RunnerState is Running.
func (s RunnerState) IsRunning() bool {
	return s == RunnerStateRunAccelerating ||
		s == RunnerStateRunCruising ||
		s == RunnerStateRunDecelerating ||
		s == RunnerStateRunStopping
}

// IsJumping checks if the RunnerState is Jumping.
func (s RunnerState) IsJumping() bool {
	return s == RunnerStateJumpCharging ||
		s == RunnerStateJumpRising ||
		s == RunnerStateJumpFalling ||
		s == RunnerStateJumpLanding
}
