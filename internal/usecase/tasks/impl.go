package tasks

import (
	"context"
	domain "todos-api/internal/domain/tasks"

	"github.com/google/uuid"
)

type useCase struct {
	repo domain.Repository
}

func New(repo domain.Repository) UseCase {
	return &useCase{
		repo: repo,
	}
}

func (uc *useCase) GetAll(ctx context.Context) ([]*domain.Task, error) {
	return uc.repo.GetAll(ctx)
}

func (uc *useCase) GetByID(ctx context.Context, id string) (*domain.Task, error) {
	if id == "" {
		return nil, domain.ErrEmptyTID
	}

	return uc.repo.GetByID(ctx, id)
}

func (uc *useCase) Create(ctx context.Context, title, description string) (*domain.Task, error) {
	id := uuid.New().String()

	task, err := domain.NewTask(id, title, description)
	if err != nil {
		return nil, err
	}

	if err := uc.repo.Create(ctx, task); err != nil {
		return nil, err
	}

	return task, nil
}

func (uc *useCase) Update(ctx context.Context, id, title, description string, status domain.Status) (*domain.Task, error) {
	if id == "" {
		return nil, domain.ErrEmptyTID
	}

	task, err := uc.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	task.Title = title
	task.Description = description
	task.Status = status

	if err := uc.repo.Update(ctx, task); err != nil {
		return nil, err
	}

	return task, nil
}

func (uc *useCase) Delete(ctx context.Context, id string) error {
	if id == "" {
		return domain.ErrEmptyTID
	}

	return uc.repo.Delete(ctx, id)
}
