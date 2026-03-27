package users

import (
	"context"
	domain "todos-api/internal/domain/users"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type useCase struct {
	repo domain.Repository
}

func New(repo domain.Repository) UseCase {
	return &useCase{
		repo: repo,
	}
}

func (uc *useCase) GetAll(ctx context.Context) ([]*domain.User, error) {
	return uc.repo.GetAll(ctx)
}

func (uc *useCase) GetByID(ctx context.Context, id string) (*domain.User, error) {
	if id == "" {
		return nil, domain.ErrEmptyUID
	}

	return uc.repo.GetByID(ctx, id)
}

func (uc *useCase) Create(ctx context.Context, name, email, password string) (*domain.User, error) {
	uid := uuid.New().String()
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user, err := domain.NewUser(uid, name, email, string(hash))
	if err != nil {
		return nil, err
	}

	if err := uc.repo.Create(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (uc *useCase) Update(ctx context.Context, uid, name, email string) (*domain.User, error) {
	if uid == "" {
		return nil, domain.ErrEmptyUID
	}

	user, err := uc.repo.GetByID(ctx, uid)
	if err != nil {
		return nil, err
	}

	user.Name = name
	user.Email = email

	if err := uc.repo.Update(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (uc *useCase) Delete(ctx context.Context, uid string) error {
	if uid == "" {
		return domain.ErrEmptyUID
	}

	return uc.repo.Delete(ctx, uid)
}
