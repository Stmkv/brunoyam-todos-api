package users

import "errors"

var (
	ErrUserNotFound      = errors.New("user not found")
	ErrUserIsNil         = errors.New("user is nil")
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrEmptyUID          = errors.New("empty user uid")
	ErrEmptyName         = errors.New("empty user name")
	ErrEmptyEmail        = errors.New("empty user email")
	ErrEmptyPassword     = errors.New("empty user password")
)
