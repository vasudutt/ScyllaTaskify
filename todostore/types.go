package todostore

import "time"

type TodoStatus int

const (
	TodoStatusPending TodoStatus = iota
	TodoStatusCompleted
	TodoStatusDeleted
)

func (s TodoStatus) String() string {
	switch s {
	case TodoStatusPending:
		return "pending"
	case TodoStatusCompleted:
		return "completed"
	case TodoStatusDeleted:
		return "deleted"
	default:
		return ""
	}
}

func ParseTodoStatus(s string) TodoStatus {
	switch s {
	case "pending":
		return TodoStatusPending
	case "completed":
		return TodoStatusCompleted
	case "deleted":
		return TodoStatusDeleted
	default:
		return -1
	}
}

type Todo struct {
	ID          string     `json:"id"`
	UserID      string     `json:"user_id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      TodoStatus `json:"status"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
}
