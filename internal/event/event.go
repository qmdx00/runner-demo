package event

type RunnerControlEvent uint8

const (
	EventUnknown RunnerControlEvent = iota
	EventRun
	EventJump
	EventStop
	EventPause
	EventResume
)

func (e RunnerControlEvent) String() string {
	switch e {
	case EventRun:
		return "Run"
	case EventJump:
		return "Jump"
	case EventStop:
		return "Stop"
	case EventPause:
		return "Pause"
	case EventResume:
		return "Resume"
	default:
		return "Unknown"
	}
}
