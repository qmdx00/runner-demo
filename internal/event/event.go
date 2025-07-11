package event

type RunnerEvent uint8

const (
	Unknown RunnerEvent = iota

	InputMoveLeft
	InputMoveRight
	InputMoveRelease

	InputJumpPress
	InputJumpRelease

	RunnerVelocityYNegative
	RunnerGrounded
	RunnerLanded
)

func (e RunnerEvent) String() string {
	switch e {
	case InputMoveLeft:
		return "Move Left"
	case InputMoveRight:
		return "Move Right"
	case InputMoveRelease:
		return "Move Release"
	case InputJumpPress:
		return "Jump Press"
	case InputJumpRelease:
		return "Jump Release"
	case RunnerVelocityYNegative:
		return "Velocity Y Negative"
	case RunnerGrounded:
		return "Grounded"
	case RunnerLanded:
		return "Landed"
	default:
		return "Unknown"
	}
}

type EventBus struct {
	events chan RunnerEvent
}

func NewEventBus(cap ...int) *EventBus {
	var capacity = 10
	if len(cap) > 0 {
		capacity = cap[0]
	}
	return &EventBus{
		events: make(chan RunnerEvent, capacity),
	}
}

func (eb *EventBus) Push(event RunnerEvent) {
	eb.events <- event
}

func (eb *EventBus) Pop() (RunnerEvent, bool) {
	event, ok := <-eb.events
	return event, ok
}
