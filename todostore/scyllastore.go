package todostore

import (
	"log"
	"sort"
	"time"

	"github.com/gocql/gocql"
)

// ScyllaStore implements the Store interface for ScyllaDB.
type ScyllaStore struct {
	session *gocql.Session
}

// NewScyllaStore creates a new ScyllaStore instance.
func NewScyllaStore(session *gocql.Session) *ScyllaStore {
	return &ScyllaStore{session: session}
}

// Create inserts a new TODO item into ScyllaDB.
func (s *ScyllaStore) Create(todo *CreateTodoRequest) error {
	query := `INSERT INTO todos (id, user_id, title, description, status, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?)`

	err := s.session.Query(query, todo.ID, todo.UserID, todo.Title, todo.Description, todo.Status, time.Now(), time.Now()).Exec()

	return err
}

// UpdateStatus updates the status of a TODO item in ScyllaDB.
func (s *ScyllaStore) UpdateStatus(req *UpdateTodoStatusRequest) error {
	query := `UPDATE todos SET status = ? WHERE id = ?`

	err := s.session.Query(query, req.Status, req.ID).Exec()
	return err
}

// GetUserTodos retrieves TODO items for a specific user from ScyllaDB.
func (s *ScyllaStore) GetUserTodos(req *GetUserTodosRequest) (*GetUserTodosResponse, error) {
	var todos []Todo
	baseQuery := `SELECT * FROM todos WHERE user_id = ?`
	var params []interface{}
	params = append(params, req.UserID)

	log.Printf("Status: %d\n", req.Status)

	if req.Status != -1 {
		baseQuery += " AND status = ? ALLOW FILTERING"
		params = append(params, req.Status)
	} else {
		baseQuery += "AND status != " + TodoStatusDeleted.String() + " ALLOW FILTERING"
	}

	log.Printf("Executing query: %s with UserID: %s\n", baseQuery, req.UserID)

	iter := s.session.Query(baseQuery, params...).Iter()

	for {
		var todo Todo
		if !iter.Scan(&todo.ID, &todo.CreatedAt, &todo.Description, &todo.Status, &todo.Title, &todo.UpdatedAt, &todo.UserID) {
			break
		}
		log.Printf("Fetched todo: %+v\n", todo)
		todos = append(todos, todo)
	}

	if len(todos) == 0 {
		log.Println("No todos found for the given user.")
	} else {
		// Sort todos
		sort.Slice(todos, func(i, j int) bool {
			if req.Order == "desc" {
				return todos[i].CreatedAt.After(todos[j].CreatedAt)
			}
			return todos[i].CreatedAt.Before(todos[j].CreatedAt)
		})
	}

	return &GetUserTodosResponse{Todos: todos}, nil
}

// GetUserTodo retrieves a specific TODO item for a user from ScyllaDB.
func (s *ScyllaStore) GetUserTodo(req *GetUserTodoRequest) (*GetUserTodoResponse, error) {
	return nil, gocql.ErrUnsupported
}

func (s *ScyllaStore) DeleteUserTodo(req *DeleteTodoRequest) error {
	query := `UPDATE todos SET status = ? WHERE id = ?`

	err := s.session.Query(query, TodoStatusDeleted, req.ID).Exec()
	return err
}
