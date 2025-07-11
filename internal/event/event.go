package event

type RunnerControlEvent uint8

const (
	EventUnknown RunnerControlEvent = iota
	EventAccelerate
	EventDecelerate
	EventJump
)

func (e RunnerControlEvent) String() string {
	switch e {
	case EventAccelerate:
		return "Accelerate"
	case EventDecelerate:
		return "Decelerate"
	case EventJump:
		return "Jump"
	default:
		return "Unknown"
	}
}
