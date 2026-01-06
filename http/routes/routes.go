package routes

import (
	"crud/http/routes/auth"
	"crud/http/routes/group"
	"crud/http/routes/movie"
	"crud/http/routes/rate"
	"crud/http/routes/task"
	"crud/http/routes/user"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(db *sql.DB) *gin.Engine {
	router := gin.Default()

	task.RegisterTaskRoutes(router, db)
	user.RegisterUserRoutes(router, db)
	auth.RegisterAuthRoutes(router, db)
	rate.RegisterRateRoutes(router, db)
	movie.RegisterMovieRoutes(router, db)
	group.RegisterGroupRoutes(router, db)

	return router
}
