package main

import (
	handlers "MinhaApi/Handlers"
	"MinhaApi/config"
	"MinhaApi/models"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	dbConnection := config.SetUpDB()

	_, err := dbConnection.Exec(models.CreateTableSQL)

	if err != nil {
		log.Fatal(err)
	}

	defer dbConnection.Close()

	taskHandler := handlers.NewTaskHandler(dbConnection)

	router := mux.NewRouter()

	router.HandleFunc("/tasks", taskHandler.ReadTasks).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}
