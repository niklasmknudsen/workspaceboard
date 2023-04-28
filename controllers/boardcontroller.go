package controllers

import (
	"example/WorkspaceBoard/mappers/dtos"
	"example/WorkspaceBoard/services"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BoardController interface {
	GetBoards() []dtos.Board
	CreateBoard(newBoard dtos.Board) bool
	DeleteBoard(ctx *gin.Context) (bool, error)
}

func GetBoards() []dtos.Board {
	fmt.Println("GetBoards:")
	boards := services.GetBoards()
	return boards
}

func CreateBoard(newBoard dtos.Board) bool {
	created := services.CreateBoard(newBoard)
	return created
}

func DeleteBoard(ctx *gin.Context) {
	id := ctx.Params.ByName("id")

	boardId, err := strconv.ParseInt(id, 0, 0)
	deletedboard, err := services.DeleteBoard(boardId)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusNotFound, deletedboard)
	}

	ctx.JSON(http.StatusOK, deletedboard)
}
