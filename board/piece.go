package board

import (
	"image/color"
	"main/utils"
)

type Piece struct {
	Position []float64
	Image    [][]color.RGBA
}

type SendPiecesBackSturct struct {
	Pieces []Piece
}

type PieceSentToServer struct {
	Position utils.Vec2
	Image    [][]color.RGBA
}
