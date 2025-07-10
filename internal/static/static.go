package static

import (
	"bytes"
	"image"
	"log"
	"runner-demo/assets"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	RunnerRunSprite  *Sprite
	RunnerIdleSprite *Sprite
	RunnerJumpSprite *Sprite
)

var (
	BackgroundImage_png *ebiten.Image
)

func init() {
	idleImage, _, err := image.Decode(bytes.NewReader(assets.RunnerIdleImage))
	if err != nil {
		log.Fatal(err)
	}
	RunnerIdleSprite = NewFrameSprite(ebiten.NewImageFromImage(idleImage), 32, 32, 8)

	runImage, _, err := image.Decode(bytes.NewReader(assets.RunnerRunImage))
	if err != nil {
		log.Fatal(err)
	}
	RunnerRunSprite = NewFrameSprite(ebiten.NewImageFromImage(runImage), 32, 32, 9)

	jumpImage, _, err := image.Decode(bytes.NewReader(assets.RunnerJumpImage))
	if err != nil {
		log.Fatal(err)
	}
	RunnerJumpSprite = NewFrameSprite(ebiten.NewImageFromImage(jumpImage), 32, 32, 12)

	bgImage, _, err := image.Decode(bytes.NewReader(assets.BackgroundImage))
	if err != nil {
		log.Fatal(err)
	}
	BackgroundImage_png = ebiten.NewImageFromImage(bgImage)
}
