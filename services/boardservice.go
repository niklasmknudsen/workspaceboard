package services

import (
	"errors"
	"example/WorkspaceBoard/mappers"
	"example/WorkspaceBoard/mappers/dtos"
	"example/WorkspaceBoard/repositories"
)

type BoardService interface {
	GetBoards() []dtos.Board
	CreateBoard(board dtos.Board) bool
	DeleteBoard(id int64) (bool, error)
}

func GetBoards() []dtos.Board {
	boards := mappers.MapBoards(repositories.GetBoards())
	return boards
}

func CreateBoard(board dtos.Board) bool {
	newBoard := mappers.MapBoardFromDTO(board)
	createdBoard := repositories.CreateBoard(newBoard)
	if createdBoard != 0 {
		return true
	}

	return false
}

func DeleteBoard(id int64) (bool, error) {
	deletedElement, err := repositories.DeleteBoard(id)
	if err != nil {
		return deletedElement, errors.New(err.Error())
	}

	return deletedElement, nil
}
