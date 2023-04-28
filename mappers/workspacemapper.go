package mappers

import (
	"example/WorkspaceBoard/mappers/dtos"
	"example/WorkspaceBoard/models"
)

type WorkSpaceMapper interface {
	MapWorkSpace(workspace models.WorkSpaceBoard) dtos.WorkSpace
	MapWorkSpaces(workspaces []models.WorkSpaceBoard) []dtos.WorkSpace
}

func MapWorkSpace(workspace models.WorkSpaceBoard) dtos.WorkSpace {
	return dtos.WorkSpace{
		ID:     workspace.ID,
		Name:   workspace.Name,
		Boards: MapBoards(workspace.Boards),
	}
}

func MapWorkSpaces(workspaces []models.WorkSpaceBoard) []dtos.WorkSpace {
	var resultSet []dtos.WorkSpace

	for workspace := range workspaces {
		resultSet = append(resultSet, MapWorkSpace(workspaces[workspace]))
	}

	return resultSet
}

func MapWorkSpaceFromDTO(workspace dtos.WorkSpace) models.WorkSpaceBoard {
	return models.WorkSpaceBoard{
		ID:     workspace.ID,
		Name:   workspace.Name,
		Boards: MapBoardsFromDTO(workspace.Boards),
	}
}
