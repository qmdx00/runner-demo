package static

import (
	"bytes"
	"image"
	"io/fs"
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

var (
	TileImagesMap = make(map[string]*ebiten.Image)
)

func InitStatic() {
	InitBackground()
	InitSprites()
	InitStateFrames()
	InitTiles()
	InitTileMap()
}

func InitBackground() {
	bgImage, _, err := image.Decode(bytes.NewReader(assets.BackgroundImage))
	if err != nil {
		log.Fatal(err)
	}
	BackgroundImage_png = ebiten.NewImageFromImage(bgImage)
}

func InitSprites() {
	idleImage, _, err := image.Decode(bytes.NewReader(assets.RunnerIdleImage))
	if err != nil {
		log.Fatal(err)
	}
	RunnerIdleSprite = NewFrameSprite(ebiten.NewImageFromImage(idleImage), 8)

	runImage, _, err := image.Decode(bytes.NewReader(assets.RunnerRunImage))
	if err != nil {
		log.Fatal(err)
	}
	RunnerRunSprite = NewFrameSprite(ebiten.NewImageFromImage(runImage), 9)

	jumpImage, _, err := image.Decode(bytes.NewReader(assets.RunnerJumpImage))
	if err != nil {
		log.Fatal(err)
	}
	RunnerJumpSprite = NewFrameSprite(ebiten.NewImageFromImage(jumpImage), 12)
}

func InitTiles() {
	tileFS, err := fs.Sub(assets.TileFiles, "tiles")
	if err != nil {
		panic(err)
	}

	entries, err := fs.ReadDir(tileFS, ".")
	if err != nil {
		panic(err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		imgData, err := fs.ReadFile(tileFS, entry.Name())
		if err != nil {
			panic(err)
		}

		img, _, err := image.Decode(bytes.NewReader(imgData))
		if err != nil {
			panic(err)
		}

		TileImagesMap[entry.Name()] = ebiten.NewImageFromImage(img)
	}
}
