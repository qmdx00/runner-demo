package input

import (
	"runner-demo/internal/event"

	"github.com/hajimehoshi/ebiten/v2"
)

var KeyboardEventMap = map[ebiten.Key]event.RunnerControlEvent{
	ebiten.KeyArrowRight: event.EventRun,
	ebiten.KeySpace:      event.EventJump,
}
