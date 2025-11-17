package models

import "time"

// create model
type Todo struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"created_at"`
}

var Todos []Todo
var nextID int = 1

// initialize todos
func InitiliazeTodos() {
	Todos = []Todo{
		{
			ID:          nextID,
			Title:       "Samples First Todo",
			Description: "My First project on learning GO Language",
			Completed:   false,
			CreatedAt:   time.Now(),
		},
	}
	nextID++
}

func GetNextID() int {
	id := nextID
	nextID++
	return id
}
