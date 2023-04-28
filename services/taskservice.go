package services

import (
	"errors"
	"example/WorkspaceBoard/mappers"
	"example/WorkspaceBoard/mappers/dtos"
	"example/WorkspaceBoard/repositories"
)

type TaskService interface {
	GetTasks() []dtos.Task
	GetAllTaskByBoard(id int64) []dtos.Task
	CreateTask(task dtos.Task) (dtos.Task, error)
	DeleteTask(id int64) bool
}

func GetTasks() []dtos.Task {
	tasks := repositories.GetTasks()
	return mappers.MapTasks(tasks)
}

func GetAllTaskByBoard(id int64) []dtos.Task {
	tasks := repositories.GetAllByBoard(id)
	return mappers.MapTasks(tasks)
}

func CreateTask(task dtos.Task) (dtos.Task, error) {
	if task.Description == "" {
		return dtos.Task{}, errors.New("Cannot create a new task without a description provided")
	}

	createTask := repositories.CreateTask(mappers.MapTaskFromDTO(task))
	newTask := mappers.MapTask(repositories.GetById(createTask))

	return newTask, nil
}

func DeleteTask(id int64) bool {
	deleted, err := repositories.DeleteTask(id)
	if err != nil {
		return deleted
	}

	return deleted
}
