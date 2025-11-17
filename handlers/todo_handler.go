// logic CRUD

package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
	"todo-task/models"

	"github.com/gorilla/mux"
)

func GetAllTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.Todos)
}

// create
func CreateTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var todo models.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		// w.Write([]byte("Invalid request payload"))
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request payload"})
		return
	}

	todo.ID = models.GetNextID()
	todo.CreatedAt = time.Now()
	models.Todos = append(models.Todos, todo)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todo)

}

// delete
func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid todo ID"})
		return
	}

	for i, todo := range models.Todos {
		if todo.ID == id {
			models.Todos = append(models.Todos[:i], models.Todos[i+1:]...)
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]string{
				"message": "Todo deleted successfully",
			})
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{
		"error": "Todo not found",
	})
}

// read details
func GetTodoDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid todo ID"})
		return
	}

	for _, todo := range models.Todos {
		if todo.ID == id {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(todo)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "Todo not found"})
}

// update
func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Invalid todo ID",
		})
		return
	}

	var updatedTodo models.Todo
	for _, todo := range models.Todos {
		if todo.ID == id {
			if err := json.NewDecoder(r.Body).Decode(&updatedTodo); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request payload"})
				return
			}
		}
	}

	for i, todo := range models.Todos {
		if todo.ID == id {
			models.Todos[i].Title = updatedTodo.Title
			models.Todos[i].Description = updatedTodo.Description
			models.Todos[i].Completed = updatedTodo.Completed
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(models.Todos[i])
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "Todo not found"})
}
