package static

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type Sprite struct {
	frames []*ebiten.Image

	FrameWidth  int
	FrameHeight int
	FrameCount  int
}

func NewFrameSprite(img *ebiten.Image, frameWidth, frameHeight, frameCount int) *Sprite {
	frames := make([]*ebiten.Image, 0, frameCount)
	for index := range frameCount {
		frame, ok := img.SubImage(image.Rect(index*frameWidth, 0, (index+1)*frameWidth, frameHeight)).(*ebiten.Image)
		if !ok {
			panic("failed to create sub-image for sprite frame")
		}
		frames = append(frames, frame)
	}

	return &Sprite{
		frames:      frames,
		FrameWidth:  frameWidth,
		FrameHeight: frameHeight,
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
