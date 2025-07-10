package scenes

import (
	"runner-demo/internal/config"
	"runner-demo/internal/static"

	"github.com/hajimehoshi/ebiten/v2"
)

type Scene struct {
	layers [][][]static.TileID
}

func NewDefaultScene(mps ...[][]byte) *Scene {
	layers := make([][][]static.TileID, 0, len(mps))

	for _, mp := range mps {
		layer := make([][]static.TileID, len(mp))

		for row := range mp {
			layer[row] = make([]static.TileID, len(mp[row]))

			for column := range mp[row] {
				layer[row][column] = static.TileID(mp[row][column])
			}
		}
		layers = append(layers, layer)
	}

	return &Scene{layers: layers}
}

func (s *Scene) Render(screen *ebiten.Image) {
	tileW, tileH := config.Global.Game.Tile.Width, config.Global.Game.Tile.Height
	columns, rows := config.Global.Game.Grid.Columns, config.Global.Game.Grid.Rows
	cellWidth := screen.Bounds().Dx() / columns
	cellHeight := screen.Bounds().Dy() / rows

	for layerIndex := range s.layers {
		for row := range s.layers[layerIndex] {
			for column := range s.layers[layerIndex][row] {
				tileID := s.layers[layerIndex][row][column]
				if tile, ok := static.TileMap[tileID]; ok && tile.Image != nil {
					op := &ebiten.DrawImageOptions{}
					sx, sy := float64(cellWidth)/float64(tileW), float64(cellHeight)/float64(tileH)
					tx, ty := float64(column*cellWidth), float64(row*cellHeight)

					if tile.Width != tile.Height {
						switch tile.Align {
						case static.TileAlignLeft:
							tx = float64(column*cellWidth) - (float64(cellWidth) - float64(tile.Width)*sx)
						case static.TileAlignRight:
							tx = float64(column*cellWidth) + (float64(cellWidth) - float64(tile.Width)*sx)
						case static.TileAlignTop:
							ty = float64(row*cellHeight) - (float64(cellHeight) - float64(tile.Height)*sy)
						case static.TileAlignBottom:
							ty = float64(row*cellHeight) + (float64(cellHeight) - float64(tile.Height)*sy)
						}
					}

					op.GeoM.Scale(sx, sy)
					op.GeoM.Translate(tx, ty)
					screen.DrawImage(tile.Image, op)
				}
			}
		}
	}
}
