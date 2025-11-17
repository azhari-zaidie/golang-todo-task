package main

import (
	"log"
	"net/http"
	route "todo-task/routes"

	"github.com/gorilla/mux"
)

// main
func main() {
	println("Hello, World!")
	// create route
	router := mux.NewRouter()

	route.SetupRoutes(router)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("could not start server: %s\n", err.Error())
	}

}
