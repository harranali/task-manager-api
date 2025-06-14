package task

import (
	"net/http"

	"github.com/harranali/task-manager-api/internal/user"
)

func RegisterRoutes(mux *http.ServeMux) {
	repo := NewRepository()
	srv := NewService(repo)
	h := NewHandler(srv, user.UserService)

	mux.HandleFunc("POST /tasks", h.SaveTask)
	mux.HandleFunc("GET /tasks/{id}", h.GetByID)
	mux.HandleFunc("GET /tasks", h.GetUserTasks)
	mux.HandleFunc("PUT /tasks/{id}", h.Update)
}
