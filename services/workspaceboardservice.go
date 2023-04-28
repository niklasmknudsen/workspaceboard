package services

import (
	"errors"
	"example/WorkspaceBoard/mappers"
	"example/WorkspaceBoard/mappers/dtos"
	"example/WorkspaceBoard/repositories"
	"fmt"
)

type WorkspaceBoardServices interface {
	GetWorkSpaces() ([]dtos.WorkSpace, error)
	CreateWorkSpace(workspace dtos.WorkSpace) (bool, error)
	DeleteWorkSpace(id int64) (bool, error)
}

func CreateWorkSpace(workspace dtos.WorkSpace) (bool, error) {
	if workspace.Name == "" {
		return false, errors.New("no workspace was passed as argument")
	}
	fmt.Println("kommer ind i service CreateWorkSpace")
	createdWorkspaceBoard := repositories.CreateWorkSpace(mappers.MapWorkSpaceFromDTO(workspace))

	if createdWorkspaceBoard != 0 {
		return true, nil
	}

	return false, errors.New("could not create new workspace")
}

func GetWorkSpaces() ([]dtos.WorkSpace, error) {
	workspaces := repositories.GetWorkSpaces()

	if workspaces != nil {
		return mappers.MapWorkSpaces(workspaces), nil
	}

	return []dtos.WorkSpace{}, errors.New("Could not find any workspaces in repository")
}

func DeleteWorkSpace(id int64) (bool, error) {
	result, error := repositories.DeleteWorkSpace(id)
	if error != nil {
		return false, errors.New(error.Error())
	}

	return result, nil
}
