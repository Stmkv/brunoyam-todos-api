package tasks

type Status string

const (
	StatusNew       Status = "new"
	StatusInProcess Status = "in_process"
	StatusCompleted Status = "completed"
)

type Task struct {
	TID         string
	UserID      string
	Title       string
	Description string
	Status      Status
}

func NewTask(tid string, userID string, title string, description string) (*Task, error) {
	if tid == "" {
		return nil, ErrEmptyTID
	}
	if title == "" {
		return nil, ErrEmptyTitle
	}

	return &Task{
		TID:         tid,
		UserID:      userID,
		Title:       title,
		Description: description,
		Status:      StatusNew,
	}, nil
}

func ValidateStatus(s string) error {
	switch Status(s) {
	case StatusNew,
		StatusInProcess,
		StatusCompleted:
		return nil
	default:
		return ErrStatusNotFound
	}
}
