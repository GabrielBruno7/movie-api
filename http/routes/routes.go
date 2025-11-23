package routes

import (
	"crud/http/handlers"
	"database/sql"

	"github.com/gorilla/mux"
)

func SetupRoutes(databaseConnection *sql.DB) *mux.Router {
	router := mux.NewRouter()

	taskHandler := handlers.NewTaskHandler(databaseConnection)

	router.HandleFunc("/tasks", taskHandler.List).Methods("GET")
	router.HandleFunc("/tasks", taskHandler.Create).Methods("POST")

	return router
}
