package static

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type TileID byte

const (
	TileIDGrass_1 TileID = iota + 1
	TileIDGrass_2
	TileIDGrass_3
	TileIDGrass_4
	TileIDGrass_5
	TileIDGrass_6
	TileIDGrass_7
	TileIDGrass_8

	TileIDSoil_1
	TileIDSoil_2
	TileIDSoil_3
	TileIDSoil_4
	TileIDSoil_5
	TileIDSoil_6
	TileIDSoil_7
	TileIDSoil_8

	TileIDWater_1
	TileIDWater_2
)

var TileMap = make(map[TileID]*Tile)

type TileAlign string

const (
	TileAlignCenter TileAlign = "center"
	TileAlignLeft   TileAlign = "left"
	TileAlignRight  TileAlign = "right"
	TileAlignTop    TileAlign = "top"
	TileAlignBottom TileAlign = "bottom"
)

type Tile struct {
	ID            TileID
	Image         *ebiten.Image
	Align         TileAlign
	Width, Height int
}

func NewTile(id TileID, filename string, align TileAlign) *Tile {
	image, ok := TileImagesMap[filename]
	if !ok {
		panic("tile image not found: " + filename)
	}
	return &Tile{ID: id, Image: image, Width: image.Bounds().Dx(), Height: image.Bounds().Dy(), Align: align}
}

func InitTileMap() {
	TileMap = map[TileID]*Tile{
		TileIDGrass_1: NewTile(TileIDGrass_1, "1.png", TileAlignCenter),
		TileIDGrass_2: NewTile(TileIDGrass_2, "2.png", TileAlignCenter),
		TileIDGrass_3: NewTile(TileIDGrass_3, "3.png", TileAlignCenter),
		TileIDGrass_4: NewTile(TileIDGrass_4, "7.png", TileAlignCenter),
		TileIDGrass_5: NewTile(TileIDGrass_5, "11.png", TileAlignCenter),
		TileIDGrass_6: NewTile(TileIDGrass_6, "13.png", TileAlignTop),
		TileIDGrass_7: NewTile(TileIDGrass_7, "14.png", TileAlignTop),
		TileIDGrass_8: NewTile(TileIDGrass_8, "15.png", TileAlignTop),

		TileIDSoil_1: NewTile(TileIDSoil_1, "4.png", TileAlignCenter),
		TileIDSoil_2: NewTile(TileIDSoil_2, "5.png", TileAlignCenter),
		TileIDSoil_3: NewTile(TileIDSoil_3, "6.png", TileAlignCenter),
		TileIDSoil_4: NewTile(TileIDSoil_4, "8.png", TileAlignCenter),
		TileIDSoil_5: NewTile(TileIDSoil_5, "9.png", TileAlignCenter),
		TileIDSoil_6: NewTile(TileIDSoil_6, "10.png", TileAlignCenter),
		TileIDSoil_7: NewTile(TileIDSoil_7, "12.png", TileAlignCenter),
		TileIDSoil_8: NewTile(TileIDSoil_8, "16.png", TileAlignCenter),

		TileIDWater_1: NewTile(TileIDWater_1, "17.png", TileAlignBottom),
		TileIDWater_2: NewTile(TileIDWater_2, "18.png", TileAlignCenter),
	}
}
