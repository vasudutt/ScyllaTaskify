package handler

import (
	"encoding/json"
	"net/http"

	"github.com/vasudutt/ScyllaTaskify/todo"
	"github.com/vasudutt/ScyllaTaskify/todostore"
)

type TodoAPIHandler struct {
	manager *todo.TodoManager
}

func NewTodoAPIHandler(manager *todo.TodoManager) *TodoAPIHandler {
	return &TodoAPIHandler{manager: manager}
}

func (h *TodoAPIHandler) CreateTodoHandler(w http.ResponseWriter, r *http.Request) {
	var createReq todo.CreateTodoRequest
	if err := json.NewDecoder(r.Body).Decode(&createReq); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	response, err := h.manager.Create(&createReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *TodoAPIHandler) GetUserTodosHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("userID")
	status := r.URL.Query().Get("status")
	order := r.URL.Query().Get("order")

	req := todo.GetUserTodosRequest{UserID: userID, Status: todostore.ParseTodoStatus(status), Order: order}

	resp, err := h.manager.GetUserTodos(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *TodoAPIHandler) UpdateTodoStatusHandler(w http.ResponseWriter, r *http.Request) {
	var updateReq todo.UpdateTodoStatusRequest
	if err := json.NewDecoder(r.Body).Decode(&updateReq); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err := h.manager.UpdateStatus(&updateReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *TodoAPIHandler) DeleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	todoID := r.URL.Query().Get("todoID")

	req := todo.DeleteTodoRequest{ID: todoID}

	err := h.manager.Delete(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
