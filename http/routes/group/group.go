package group

import (
	"crud/http/handlers"
	"crud/http/middleware"
	"crud/infrastructure/database"
	"crud/usecase"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func RegisterGroupRoutes(router *gin.Engine, db *sql.DB) {
	protected := router.Group("/")
	protected.Use(middleware.AuthMiddleware())

	groupUsecase := usecase.NewGroupUsecase(
		database.NewGroupDb(db),
		usecase.NewUserUsecase(database.NewUserDb(db)),
	)

	groupHandler := handlers.NewGroupHandler(groupUsecase)

	protected.POST("/group/create", groupHandler.ActionCreateGroup)
	protected.POST("/group/:id/invite/send", groupHandler.ActionSendGroupInvite)
}
