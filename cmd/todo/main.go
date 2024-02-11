// main.go
package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/vasudutt/ScyllaTaskify/config"
	"github.com/vasudutt/ScyllaTaskify/handler"
	"github.com/vasudutt/ScyllaTaskify/scylla"
	"github.com/vasudutt/ScyllaTaskify/todo"
	"github.com/vasudutt/ScyllaTaskify/todostore"
	"github.com/vasudutt/ScyllaTaskify/uuid"
)

func main() {
	// Initialize your configuration and ScyllaDB manager
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	scyllaManager := scylla.NewManager(cfg)
	err = scyllaManager.CreateKeyspace(cfg.ScyllaKeyspace)
	if err != nil {
		log.Fatal(err)
	}

	session, err := scyllaManager.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	// Initialize your store and manager
	store := todostore.NewScyllaStore(session)
	todoManager := todo.NewTodoManager(uuid.New(), store)

	// Initialize API handlers
	todoHandler := handler.NewTodoAPIHandler(todoManager)

	// Setup routes using chi
	r := chi.NewRouter()

	// Define your routes
	r.Route("/api/todos", func(r chi.Router) {
		r.Post("/", todoHandler.CreateTodoHandler)
		r.Get("/", todoHandler.GetUserTodosHandler)
		r.Put("/update/status", todoHandler.UpdateTodoStatusHandler)
		r.Delete("/{todoID}", todoHandler.DeleteTodoHandler)
	})

	// Start the server
	port := ":8080"
	log.Printf("Server listening on port %s", port)
	log.Fatal(http.ListenAndServe(port, r))
}
