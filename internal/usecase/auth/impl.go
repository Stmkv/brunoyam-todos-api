package auth

import (
	"context"
	"errors"

	usersDomain "todos-api/internal/domain/users"
)

type Hasher interface {
	Compare(hash, password string) bool
}

type TokenManager interface {
	GenerateAccessToken(uid string) (string, error)
	GenerateRefreshToken(uid string) (string, error)
}

type useCase struct {
	userRepo     usersDomain.Repository
	hasher       Hasher
	tokenManager TokenManager
}

func New(
	userRepo usersDomain.Repository,
	hasher Hasher,
	tokenManager TokenManager,
) UseCase {
	return &useCase{
		userRepo:     userRepo,
		hasher:       hasher,
		tokenManager: tokenManager,
	}
}

func (uc *useCase) Login(ctx context.Context, email, password string) (string, string, error) {
	user, err := uc.userRepo.GetByEmail(ctx, email)
	if err != nil {
		return "", "", errors.New("invalid credentials")
	}

	if !uc.hasher.Compare(user.Password, password) {
		return "", "", errors.New("invalid credentials")
	}

	access, err := uc.tokenManager.GenerateAccessToken(user.UID)
	if err != nil {
		return "", "", err
	}

	refresh, err := uc.tokenManager.GenerateRefreshToken(user.UID)
	if err != nil {
		return "", "", err
	}

	return access, refresh, nil
}
