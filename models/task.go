package models

type Task struct {
	ID          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
	Status      bool   `json:"status" db:"status"`
}

const (
	TableName      = "tasks"
	CreateTableSQL = `CREATE TABLE IF NOT EXISTS tasks(
	id SERIAL PRIMARY KEY,
	title VARCHAR(100) NOT NULL,
	description TEXT,
	status BOOLEAN NOT NULL DEFAULT false
	);`
)
