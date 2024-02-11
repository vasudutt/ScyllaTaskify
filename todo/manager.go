package todo

import (
	"github.com/vasudutt/ScyllaTaskify/todostore"
	"github.com/vasudutt/ScyllaTaskify/uuid"
)

type TodoManager struct {
	uuidGenerator uuid.Generator
	store         todostore.Store
}

func NewTodoManager(uuidGenerator uuid.Generator, store todostore.Store) *TodoManager {
	return &TodoManager{
		uuidGenerator: uuidGenerator,
		store:         store,
	}
}

func (m *TodoManager) Create(req *CreateTodoRequest) (*CreateTodoResponse, error) {
	todoId := m.uuidGenerator.Generate()
	defaultStatus := todostore.TodoStatusPending

	r := &todostore.CreateTodoRequest{
		ID:          todoId,
		UserID:      req.UserID,
		Title:       req.Title,
		Description: req.Description,
		Status:      defaultStatus,
	}

	err := m.store.Create(r)

	if err != nil {
		return nil, err
	}

	return &CreateTodoResponse{
		ID: todoId,
	}, nil
}

func (m *TodoManager) UpdateStatus(req *UpdateTodoStatusRequest) error {
	r := &todostore.UpdateTodoStatusRequest{
		ID:     req.ID,
		Status: todostore.ParseTodoStatus(req.Status),
	}

	return m.store.UpdateStatus(r)
}

func (m *TodoManager) GetUserTodos(req *GetUserTodosRequest) (*GetUserTodosResponse, error) {
	r := &todostore.GetUserTodosRequest{
		UserID: req.UserID,
		Status: req.Status,
		Order:  req.Order,
	}

	resp, err := m.store.GetUserTodos(r)

	if err != nil {
		return nil, err
	}

	todos := make([]Todo, len(resp.Todos))
	for i, todo := range resp.Todos {
		todos[i] = Todo{
			ID:          todo.ID,
			UserID:      todo.UserID,
			Title:       todo.Title,
			Description: todo.Description,
			Status:      todo.Status.String(),
			CreatedAt:   todo.CreatedAt,
			UpdatedAt:   todo.UpdatedAt,
		}
	}

	return &GetUserTodosResponse{
		Todos: todos,
	}, nil
}

func (m *TodoManager) GetUserTodo(req *GetUserTodoRequest) (*GetUserTodoResponse, error) {
	r := &todostore.GetUserTodoRequest{
		UserID: req.UserID,
		TodoID: req.TodoID,
	}

	resp, err := m.store.GetUserTodo(r)

	if err != nil {
		return nil, err
	}

	todo := Todo{
		ID:          resp.Todo.ID,
		UserID:      resp.Todo.UserID,
		Title:       resp.Todo.Title,
		Description: resp.Todo.Description,
		Status:      resp.Todo.Status.String(),
		CreatedAt:   resp.Todo.CreatedAt,
		UpdatedAt:   resp.Todo.UpdatedAt,
	}

	return &GetUserTodoResponse{
		Todo: todo,
	}, nil
}

func (m *TodoManager) Delete(req *DeleteTodoRequest) error {
	r := &todostore.DeleteTodoRequest{
		ID: req.ID,
	}

	return m.store.DeleteUserTodo(r)
}
