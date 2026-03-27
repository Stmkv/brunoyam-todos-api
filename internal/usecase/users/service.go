package users

import (
	"context"

	domain "todos-api/internal/domain/users"
)

type UseCase interface {
	GetAll(ctx context.Context) ([]*domain.User, error)
	GetByID(ctx context.Context, id string) (*domain.User, error)
	Create(ctx context.Context, name, email, password string) (*domain.User, error)
	Update(ctx context.Context, uid, name, email string) (*domain.User, error)
	Delete(ctx context.Context, uid string) error
}
