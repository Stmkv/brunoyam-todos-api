package users

import "context"

type Repository interface {
	GetAll(ctx context.Context) ([]*User, error)
	GetByID(ctx context.Context, id string) (*User, error)
	Create(ctx context.Context, task *User) error
	Update(ctx context.Context, task *User) error
	Delete(ctx context.Context, id string) error
}
