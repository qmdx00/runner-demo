package static

import (
	"image"
	"runner-demo/internal/config"
	"runner-demo/internal/state"

	"github.com/hajimehoshi/ebiten/v2"
)

var StateFrameMap = make(map[state.RunnerState][]*ebiten.Image)

type Sprite struct {
	frames []*ebiten.Image

	FrameWidth  int
	FrameHeight int
	FrameCount  int
}

func NewFrameSprite(img *ebiten.Image, frameCount int) *Sprite {
	spriteW, spriteH := config.Global.Game.Sprite.Width, config.Global.Game.Sprite.Height

	frames := make([]*ebiten.Image, 0, frameCount)
	for index := range frameCount {
		frame, ok := img.SubImage(image.Rect(index*spriteW, 0, (index+1)*spriteW, spriteH)).(*ebiten.Image)
		if !ok {
			panic("failed to create sub-image for sprite frame")
		}
		frames = append(frames, frame)
	}

	return &Sprite{
		frames:      frames,
		FrameWidth:  spriteW,
		FrameHeight: spriteH,
		FrameCount:  len(frames),
	}
}

func (s *Sprite) FrameByStateAndTicker(runnerState state.RunnerState, tick int) *ebiten.Image {
	if frames := StateFrameMap[runnerState]; len(frames) > 0 {
		index := (tick / 10) % len(frames)
		return frames[index]
	}
	return RunnerIdleSprite.Frames()[0] // Fallback to idle frame if no frames found for the state
}

func (s *Sprite) Frames() []*ebiten.Image {
	return s.frames
}

func InitStateFrames() {
	StateFrameMap = map[state.RunnerState][]*ebiten.Image{
		state.RunnerStateIdle:            RunnerIdleSprite.Frames(),
		state.RunnerStateRunAccelerating: RunnerRunSprite.Frames()[:4],
		state.RunnerStateRunCruising:     RunnerRunSprite.Frames(),
		state.RunnerStateRunDecelerating: RunnerRunSprite.Frames()[4:],
		state.RunnerStateRunStopped:      RunnerIdleSprite.Frames(),
		state.RunnerStateJumpCharging:    RunnerJumpSprite.Frames()[:4],
		state.RunnerStateJumpRising:      RunnerJumpSprite.Frames()[4:6],
		state.RunnerStateJumpFalling:     RunnerJumpSprite.Frames()[6:10],
		state.RunnerStateJumpLanded:      RunnerJumpSprite.Frames()[11:],
	}
}
