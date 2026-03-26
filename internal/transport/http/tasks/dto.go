package tasks

import "todos-api/internal/domain/tasks"

type createTaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type updateTaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type taskResponse struct {
	ID          string       `json:"id"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	Status      tasks.Status `json:"status"`
}
