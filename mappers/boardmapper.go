package mappers

import (
	"example/WorkspaceBoard/mappers/dtos"
	"example/WorkspaceBoard/models"
)

type BoardMapper interface {
	MapBoard(board models.Board) dtos.Board
	MapBoards(task []models.Board) []dtos.Board
	MapBoardFromDTO(board dtos.Board) models.Board
	MapBoardsFromDTO(boards []dtos.Board) []models.Board
}

func MapBoard(board models.Board) dtos.Board {
	return dtos.Board{
		ID:    board.ID,
		Name:  board.Name,
		Tasks: MapTasks(board.Tasks),
	}
}

func MapBoards(boards []models.Board) []dtos.Board {
	var resultSet []dtos.Board

	for board := range boards {
		resultSet = append(resultSet, MapBoard(boards[board]))
	}

	return resultSet
}

func MapBoardFromDTO(board dtos.Board) models.Board {
	return models.Board{
		ID:    board.ID,
		Name:  board.Name,
		Tasks: MapTasksFromDTO(board.Tasks),
	}
}

func MapBoardsFromDTO(boards []dtos.Board) []models.Board {
	var resultSet []models.Board

	for board := range boards {
		resultSet = append(resultSet, MapBoardFromDTO(boards[board]))
	}

	return resultSet
}
