package controllers

import (
	"errors"
	"example/WorkspaceBoard/mappers/dtos"
	"example/WorkspaceBoard/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TaskController interface {
	GetTasks(ctx *gin.Context)
	GetAllTasksByBoard(ctx *gin.Context)
	CreateTask(ctx *gin.Context)
	DeleteTask(ctx *gin.Context)
}

func GetTasks(ctx *gin.Context) {
	tasks := services.GetTasks()
	ctx.JSON(http.StatusOK, tasks)
}

func CreateTask(ctx *gin.Context) {
	var task dtos.Task

	err := ctx.BindJSON(task)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, errors.New(err.Error()))
	}

	created, err := services.CreateTask(task)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, created)
	}

	ctx.JSON(http.StatusCreated, created)
}

func GetAllTasksByBoard(ctx *gin.Context) {
	id := ctx.Params.ByName("id")

	boardId, err := strconv.ParseInt(id, 0, 0)
	tasks := services.GetAllTaskByBoard(boardId)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusNotFound, tasks)
	}

	ctx.JSON(http.StatusOK, tasks)
}

func DeleteTask(ctx *gin.Context) {
	id := ctx.Params.ByName("id")

	taskId, err := strconv.ParseInt(id, 0, 0)
	taskToDelete := services.DeleteTask(taskId)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, taskToDelete)
	}

	ctx.JSON(http.StatusOK, taskToDelete)
}
