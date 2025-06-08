package board

import (
	"image/color"
	"main/utils"
)

type Tile struct {
	Position []uint16
	Color    color.RGBA
}

type TileSentToServer struct {
	Pos   utils.Vec2
	Color color.RGBA
}
