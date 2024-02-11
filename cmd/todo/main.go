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

	store := todostore.NewScyllaStore(session)
	todoManager := todo.NewTodoManager(uuid.New(), store)
	todoHandler := handler.NewTodoAPIHandler(todoManager)

	r := chi.NewRouter()

	// Define routes
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
