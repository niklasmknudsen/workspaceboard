package mappers

import (
	"example/WorkspaceBoard/mappers/dtos"
	"example/WorkspaceBoard/models"
)

type TaskMapper interface {
	MapTask(task models.Task) dtos.Task
	MapTaskFromDTO(task dtos.Task) models.Task
	MapTasks(task []models.Task) []dtos.Task
}

func MapTask(task models.Task) dtos.Task {
	return dtos.Task{
		ID:          task.ID,
		Description: task.Description,
		Board:       task.Board,
	}
}

func MapTasks(tasks []models.Task) []dtos.Task {
	var resultSet []dtos.Task

	for task := range tasks {
		resultSet = append(resultSet, MapTask(tasks[task]))
	}

	return resultSet
}

func MapTaskFromDTO(task dtos.Task) models.Task {
	return models.Task{
		ID:          task.ID,
		Description: task.Description,
		Board:       task.Board,
	}
}

func MapTasksFromDTO(tasks []dtos.Task) []models.Task {
	var resultSet []models.Task

	for task := range tasks {
		resultSet = append(resultSet, MapTaskFromDTO(tasks[task]))
	}

	return resultSet
}
