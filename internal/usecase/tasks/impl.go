package tasks

import (
	"context"
	domain "todos-api/internal/domain/tasks"

	"github.com/google/uuid"
)

type Repository interface {
	GetAll(ctx context.Context, userID string) ([]*domain.Task, error)
	GetByID(ctx context.Context, id, userID string) (*domain.Task, error)
	Create(ctx context.Context, task *domain.Task) error
	Update(ctx context.Context, task *domain.Task) error
	Delete(ctx context.Context, id, userID string) error
}

type useCase struct {
	repo Repository
}

func New(repo Repository) UseCase {
	return &useCase{
		repo: repo,
	}
}

func (uc *useCase) GetAll(ctx context.Context, userID string) ([]*domain.Task, error) {
	return uc.repo.GetAll(ctx, userID)
}

func (uc *useCase) GetByID(ctx context.Context, id, userID string) (*domain.Task, error) {
	if id == "" {
		return nil, domain.ErrEmptyTID
	}

	return uc.repo.GetByID(ctx, id, userID)
}

func (uc *useCase) Create(ctx context.Context, userID, title, description string) (*domain.Task, error) {
	id := uuid.New().String()

	task, err := domain.NewTask(id, userID, title, description)
	if err != nil {
		return nil, err
	}

	if err := uc.repo.Create(ctx, task); err != nil {
		return nil, err
	}

	return task, nil
}

func (uc *useCase) Update(ctx context.Context, id, userID, title, description string, status domain.Status) (*domain.Task, error) {
	if id == "" {
		return nil, domain.ErrEmptyTID
	}

	task, err := uc.repo.GetByID(ctx, id, userID)
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

func (uc *useCase) Delete(ctx context.Context, id, userID string) error {
	if id == "" {
		return domain.ErrEmptyTID
	}

	return uc.repo.Delete(ctx, id, userID)
}
