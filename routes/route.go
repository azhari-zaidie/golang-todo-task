package route

import (
	"net/http"
	"todo-task/handlers"
	"todo-task/models"

	"github.com/gorilla/mux"
)

func SetupRoutes(router *mux.Router) {
	models.InitiliazeTodos()

	// todo route
	router.HandleFunc("/api/todos", handlers.GetAllTodos).Methods("GET")
	router.HandleFunc("/api/todos", handlers.CreateTodo).Methods("POST")
	router.HandleFunc("/api/todos/{id}", handlers.GetTodoDetails).Methods("GET")
	router.HandleFunc("/api/todos/{id}", handlers.DeleteTodo).Methods("DELETE")
	router.HandleFunc("/api/todos/{id}", handlers.UpdateTodo).Methods("PUT")

	//check health
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK TEST NEW"))
	}).Methods("GET")
}
