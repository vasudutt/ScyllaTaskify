package todostore

type Store interface {
	Create(todo *CreateTodoRequest) error
	UpdateStatus(req *UpdateTodoStatusRequest) error
	GetUserTodos(req *GetUserTodosRequest) (*GetUserTodosResponse, error)
	GetUserTodo(req *GetUserTodoRequest) (*GetUserTodoResponse, error)
	DeleteUserTodo(req *DeleteTodoRequest) error
}

type CreateTodoRequest struct {
	ID          string     `json:"id"`
	UserID      string     `json:"user_id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      TodoStatus `json:"status"`
}

type UpdateTodoStatusRequest struct {
	ID     string
	Status TodoStatus
}

type GetUserTodosRequest struct {
	UserID string
	Status TodoStatus
	Order  string
}

type GetUserTodosResponse struct {
	Todos []Todo
}

type GetUserTodoRequest struct {
	UserID string
	TodoID string
}

type GetUserTodoResponse struct {
	Todo Todo
}

type DeleteTodoRequest struct {
	ID string
}
