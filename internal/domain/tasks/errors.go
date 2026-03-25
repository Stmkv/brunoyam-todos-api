package tasks

import "errors"

var (
	ErrEmptyTID          = errors.New("tid cannot be empty")
	ErrEmptyTitle        = errors.New("title cannot be empty")
	ErrCannotStart       = errors.New("task cannot be started from current status")
	ErrCannotComplete    = errors.New("task cannot be completed from current status")
	ErrTaskNotFound      = errors.New("task not found")
	ErrTaskIsNil         = errors.New("task is nil")
	ErrTaskAlreadyExists = errors.New("task already exists")
)
