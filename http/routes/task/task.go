package task

import (
	"crud/http/handlers"
	"crud/infrastructure/database"
	"crud/usecase"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func RegisterTaskRoutes(router *gin.Engine, db *sql.DB) {
	taskRepo := database.NewTaskDb(db)
	taskUsecase := usecase.NewTaskUsecase(taskRepo)
	taskHandler := handlers.NewTaskHandler(taskUsecase)

	router.GET("/tasks", taskHandler.List)
	router.POST("/tasks", taskHandler.Create)
}
