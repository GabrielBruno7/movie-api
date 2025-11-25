package main

import (
	"crud/config"
	"crud/http/routes"
	"log"
	"net/http"
	"os"
)

func main() {
	databaseConnection := config.SetupDatabase()
	defer databaseConnection.Close()

	router := routes.SetupRoutes(databaseConnection)

	log.Println("Servidor rodando na porta " + os.Getenv("SERVER_PORT"))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("SERVER_PORT"), router))
}
