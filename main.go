package main

import (
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

	router := mux.NewRouter()

	router.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, World"))
	}).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}
