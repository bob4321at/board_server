package board

import (
	"encoding/json"
	"image/color"
	"io"
	"main/users"
	"main/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BoardSentToServer struct {
	Size  utils.Vec2
	Tiles [][]TileSentToServer
}

func (board *BoardSentToServer) GetBoard() (new_board Board) {
	new_board.Size = []uint16{uint16(board.Size.X), uint16(board.Size.Y)}

	for y := range board.Tiles {
		new_board.Tiles = append(new_board.Tiles, []Tile{})
		for _, tile := range board.Tiles[y] {
			new_board.Tiles[y] = append(new_board.Tiles[y], Tile{[]uint16{uint16(tile.Pos.X), uint16(tile.Pos.Y)}, tile.Color})
		}
	}

	return new_board
}

type Board struct {
	Size  []uint16
	Tiles [][]Tile
}

type ChangeNetworkedPieceStruct struct {
	ID    uint8
	Piece Piece
}

type ChangedPiece struct {
	ID       uint8
	Position [2]float64
	Image    [][]color.RGBA
}

type ListOfChangedPiece struct {
	Pieces []ChangedPiece
}

var Gameboard = Board{}
var Gameboard_Set = false

var Pieces = []PieceSentToServer{}
var Origonal_Pieces = []PieceSentToServer{}
var Users = []users.User{}

func CheckForChangeForUser(c *gin.Context) {
	data_in_bytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		panic(err)
	}

	data := users.User{}
	if err := json.Unmarshal(data_in_bytes, &data); err != nil {
		panic(err)
	}

	send_back_user := Users[data.ID-1]
	Users[data.ID-1].Got_Changes = true

	c.JSON(http.StatusOK, send_back_user)

}

func AddUser(c *gin.Context) {
	Users = append(Users, users.User{ID: uint8(len(Users)), Got_Changes: true})

	c.JSON(http.StatusOK, users.User{ID: uint8(len(Users)), Got_Changes: true})
}

func ChangePiece(c *gin.Context) {
	data_in_bytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		panic(err)
	}

	data := ChangedPiece{}
	if err := json.Unmarshal(data_in_bytes, &data); err != nil {
		panic(err)
	}

	Pieces[data.ID] = PieceSentToServer{utils.Vec2{X: data.Position[0], Y: data.Position[1]}, data.Image}

	for i := range Users {
		Users[i].Got_Changes = false
	}
}

func GetPieceChanges(c *gin.Context) {
	pieces_to_send := []ChangedPiece{}

	for i, piece := range Pieces {
		pieces_to_send = append(pieces_to_send, ChangedPiece{uint8(i), [2]float64{piece.Position.X, piece.Position.Y}, piece.Image})
	}

	data_to_send := ListOfChangedPiece{pieces_to_send}

	c.JSON(http.StatusOK, data_to_send)
}

func MakeBoardToServer(c *gin.Context) {
	data_in_bytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		panic(err)
	}

	data := BoardSentToServer{}
	if err := json.Unmarshal(data_in_bytes, &data); err != nil {
		panic(err)
	}

	Gameboard = data.GetBoard()
	Gameboard_Set = true

	c.Status(http.StatusAccepted)
}

func GivePiecesToServer(c *gin.Context) {
	data_in_bytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		panic(err)
	}

	data := []PieceSentToServer{}
	if err := json.Unmarshal(data_in_bytes, &data); err != nil {
		panic(err)
	}

	Pieces = data
	Origonal_Pieces = data
	c.Status(http.StatusAccepted)
}

func GetBoardFromServer(c *gin.Context) {
	c.JSON(http.StatusOK, Gameboard)
}

func GetPiecesFromServer(c *gin.Context) {
	New_Pieces := SendPiecesBackSturct{}

	for _, piece := range Pieces {
		New_Pieces.Pieces = append(New_Pieces.Pieces, Piece{[]float64{piece.Position.X, piece.Position.Y}, piece.Image})
	}

	c.JSON(http.StatusOK, New_Pieces)
}
