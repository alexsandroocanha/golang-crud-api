package handlers

import (
	"MinhaApi/config"
	"MinhaApi/models"
	"database/sql"
	"encoding/json"
	"net/http"
)

type TaskHandler struct {
	DB *sql.DB
}

func NewTaskHandler(db *sql.DB) *TaskHandler {

	return &TaskHandler{DB: config.SetUpDB()}
}

func (taskHandler *TaskHandler) ReadTasks(w http.ResponseWriter, r *http.Request) {

	rows, err := taskHandler.DB.Query("SELECT * FROM tasks")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	var tasks []models.Task

	for rows.Next() {
		var task models.Task
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tasks = append(tasks, task)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}
