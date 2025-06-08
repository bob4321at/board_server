package main

import (
	"main/board"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	go func() {
		refresh := true
		for _, user := range board.Users {
			if !user.Got_Changes {
				refresh = false
			}
		}

		if refresh {
			board.Origonal_Pieces = board.Pieces
		}
	}()

	r.GET("/GameMadeYet", func(c *gin.Context) {
		if board.Gameboard_Set {
			c.JSON(http.StatusOK, "made")
		} else {
			c.JSON(http.StatusOK, "not")
		}
	})
	r.GET("/GetBoardFromServer", board.GetBoardFromServer)
	r.GET("/GetPiecesFromServer", board.GetPiecesFromServer)
	r.GET("/AddUser", board.AddUser)

	r.POST("/SendBoardToServer", board.MakeBoardToServer)
	r.POST("/GivePiecesToServer", board.GivePiecesToServer)
	r.POST("/ChangePiece", board.ChangePiece)
	r.POST("/CheckForChangeForUser", board.CheckForChangeForUser)
	r.GET("/GetPieceChanges", board.GetPieceChanges)

	r.Run(":8080")
}
