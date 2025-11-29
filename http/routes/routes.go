package routes

import (
	"crud/http/routes/task"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(db *sql.DB) *gin.Engine {
	router := gin.Default()

	task.RegisterTaskRoutes(router, db)

	return router
}
