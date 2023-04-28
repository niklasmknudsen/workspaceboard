package main

import (
	"net/http"
	"os"

	"example/WorkspaceBoard/controllers"
	"example/WorkspaceBoard/mappers/dtos"

	"github.com/gin-gonic/gin"
)

func main() {
	setupEnviroment() // sets up env variables
	// Setup router & api endpoints
	router := gin.Default()

	// Workspace endpoints
	router.GET("/api/workspaces", func(ctx *gin.Context) {
		controllers.GetWorkSpaces(ctx)
	})
	router.POST("/api/workspaces", func(ctx *gin.Context) {
		controllers.CreateWorkSpace(ctx)
	})
	router.DELETE("/api/workspaces/:id", func(ctx *gin.Context) { controllers.DeleteWorkSpace(ctx) })

	// Board endpoints
	router.GET("/api/boards", func(ctx *gin.Context) {
		boards := controllers.GetBoards()
		ctx.JSON(http.StatusOK, gin.H{
			"status":  "200",
			"data":    boards,
			"message": "Success"})
	})

	router.POST("/api/boards", func(ctx *gin.Context) {
		var newBoard dtos.Board
		err := ctx.BindJSON(newBoard)
		if err != nil {
			ctx.Error(err)
			ctx.JSON(http.StatusBadRequest, newBoard)
		}
		created := controllers.CreateBoard(newBoard)
		ctx.JSON(http.StatusOK, created)
	})
	router.DELETE("/api/boards/:id", func(ctx *gin.Context) { controllers.DeleteBoard(ctx) })

	// Task endpoints
	router.GET("/api/tasks", func(ctx *gin.Context) { controllers.GetTasks(ctx) })
	router.GET("/api/tasks/board/:id", func(ctx *gin.Context) { controllers.GetAllTasksByBoard(ctx) })
	router.POST("/api/tasks", func(ctx *gin.Context) { controllers.CreateTask(ctx) })
	router.DELETE("/api/tasks/:id", func(ctx *gin.Context) { controllers.DeleteTask(ctx) })

	// run webserver
	router.Run(("localhost:9000"))
}

func setupEnviroment() {
	os.Setenv("MYSQL_ConnectionString", "nmk@localhost:passpass@tcp(127.0.0.1:3306)/test")
	os.Setenv("APP_ENV", "development")
}
