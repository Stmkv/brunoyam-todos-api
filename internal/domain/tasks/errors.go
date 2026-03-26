package tasks

import "errors"

var (
	ErrEmptyTID          = errors.New("tid cannot be empty")
	ErrEmptyTitle        = errors.New("title cannot be empty")
	ErrTaskNotFound      = errors.New("task not found")
	ErrTaskIsNil         = errors.New("task is nil")
	ErrTaskAlreadyExists = errors.New("task already exists")
	ErrStatusNotFound    = errors.New("status not found")
)
