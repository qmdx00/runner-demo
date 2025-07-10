package static

import (
	"image"
	"runner-demo/internal/config"

	"github.com/hajimehoshi/ebiten/v2"
)

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

func (s *Sprite) FrameByTicker(tick int) *ebiten.Image {
	index := (tick / 10) % s.FrameCount
	return s.frames[index]
}

func (s *Sprite) Frames() []*ebiten.Image {
	return s.frames
}
