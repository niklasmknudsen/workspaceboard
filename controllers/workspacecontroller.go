package controllers

import (
	"errors"
	"example/WorkspaceBoard/mappers/dtos"
	"example/WorkspaceBoard/services"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type WorkSpaceController interface {
	GetWorkSpaces(c *gin.Context)
	CreateWorkSpace(c *gin.Context)
	DeleteWorkSpace(c *gin.Context)
}

func GetWorkSpaces(ctx *gin.Context) {
	fmt.Println("GetWorkSpaces:")
	workspaces, err := services.GetWorkSpaces()
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusNotFound, workspaces)
		return
	}

	ctx.JSON(http.StatusOK, workspaces)
}

func CreateWorkSpace(ctx *gin.Context) {
	var workspace dtos.WorkSpace
	err := ctx.BindJSON(workspace)
	if err != nil {
		fmt.Println("error i at bind json")
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, errors.New(err.Error()))
		return
	}
	createWorkspace, err := services.CreateWorkSpace(workspace)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, createWorkspace)
		return
	}

	ctx.JSON(http.StatusCreated, createWorkspace)
}

func DeleteWorkSpace(ctx *gin.Context) {
	id := ctx.Params.ByName("id")

	workspaceId, err := strconv.ParseInt(id, 0, 0)
	deletedWorkSpace, err := services.DeleteWorkSpace(workspaceId)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusNotFound, deletedWorkSpace)
		return
	}

	ctx.JSON(http.StatusOK, deletedWorkSpace)
}
