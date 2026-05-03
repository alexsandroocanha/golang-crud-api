package main

import (
	handlers "MinhaApi/Handlers"
	"MinhaApi/config"
	"MinhaApi/models"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {

	go func() {
		reg := prometheus.NewRegistry()
		reg.MustRegister(
			collectors.NewGoCollector(),
			collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}),
		)

		http.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{}))

		log.Println("Prometheus rodando na porta 2112")
		log.Fatal(http.ListenAndServe(":2112", nil))
	}()

	dbConnection := config.SetUpDB()

	defer dbConnection.Close()

	_, err := dbConnection.Exec(models.CreateTableSQL)
	if err != nil {
		log.Fatal(err)
	}

	taskHandler := handlers.NewTaskHandler(dbConnection)

	router := mux.NewRouter()

	router.HandleFunc("/tasks", taskHandler.ReadTasks).Methods("GET")
	router.HandleFunc("/tasks", taskHandler.CreateTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", taskHandler.UpdateTask).Methods("PUT")
	router.HandleFunc("/tasks/{id}", taskHandler.DeleteTask).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}
