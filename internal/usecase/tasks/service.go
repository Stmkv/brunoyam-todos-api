package tasks

import (
	"context"

	domain "todos-api/internal/domain/tasks"
)

type UseCase interface {
	GetAll(ctx context.Context, userID string) ([]*domain.Task, error)
	GetByID(ctx context.Context, id, userID string) (*domain.Task, error)
	Create(ctx context.Context, userID, title, description string) (*domain.Task, error)
	Update(ctx context.Context, id, userID, title, description string, status domain.Status) (*domain.Task, error)
	Delete(ctx context.Context, id, userID string) error
}
