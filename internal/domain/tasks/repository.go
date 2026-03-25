package tasks

import "context"

type Repository interface {
	GetAll(ctx context.Context) ([]*Task, error)
	GetByID(ctx context.Context, id string) (*Task, error)
	Create(ctx context.Context, task *Task) error
	Update(ctx context.Context, task *Task) error
	Delete(ctx context.Context, id string) error
}
