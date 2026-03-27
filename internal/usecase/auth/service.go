package auth

import "context"

type UseCase interface {
	Login(ctx context.Context, email, password string) (string, string, error)
}
