package users

import "errors"

var (
	ErrUserNotFound      = errors.New("task not found")
	ErrUserIsNil         = errors.New("task is nil")
	ErrUserAlreadyExists = errors.New("task already exists")
	ErrEmptyUID          = errors.New("empty user uid")
	ErrEmptyName         = errors.New("empty user name")
	ErrEmptyEmail        = errors.New("empty user email")
	ErrEmptyPassword     = errors.New("empty user password")
)
