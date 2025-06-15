package task

import "time"

type CreateTaskRequest struct {
	Title string `json:"title" validate:"required"`
}

type UpdateTaskRequest struct {
	Title  string `json:"title" validate:"required"`
	IsDone bool   `json:"is_done" validate:"boolean"`
}

type Task struct {
	ID        uint      `json:"id"`
	UserID    uint      `json:"user_id"`
	Title     string    `json:"title"`
	IsDone    bool      `json:"is_done"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}
