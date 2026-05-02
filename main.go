package main

import (
	handlers "MinhaApi/Handlers"
	"MinhaApi/config"
	"MinhaApi/models"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {

	time.Sleep(10 * time.Second)
	time.Sleep(10 * time.Second)

	dbConnection := config.SetUpDB()

	defer dbConnection.Close()

	_, err := dbConnection.Exec(models.CreateTableSQL)
	if err != nil {
		log.Println(err)
	}

	taskHandler := handlers.NewTaskHandler(dbConnection)

	router := mux.NewRouter()

	router.HandleFunc("/tasks", taskHandler.ReadTasks).Methods("GET")
	router.HandleFunc("/tasks", taskHandler.CreateTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", taskHandler.UpdateTask).Methods("PUT")
	router.HandleFunc("/tasks/{id}", taskHandler.DeleteTask).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}
