package todo

import (
	"time"

	"github.com/vasudutt/ScyllaTaskify/todostore"
)

type Todo struct {
	ID          string    `json:"id"`
	UserID      string    `json:"user_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type CreateTodoRequest struct {
	UserID      string `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type CreateTodoResponse struct {
	ID string `json:"id"`
}

type UpdateTodoStatusRequest struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

type GetUserTodosRequest struct {
	UserID string               `json:"user_id"`
	Status todostore.TodoStatus `json:"status"`
	Order  string               `json:"order"`
}

type GetUserTodosResponse struct {
	Todos []Todo `json:"todos"`
}

type GetUserTodoRequest struct {
	UserID string `json:"user_id"`
	TodoID string `json:"todo_id"`
}

type GetUserTodoResponse struct {
	Todo Todo `json:"todo"`
}

type DeleteTodoRequest struct {
	ID string `json:"id"`
}
