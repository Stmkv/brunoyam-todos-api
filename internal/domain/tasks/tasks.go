package tasks

type Status int

const (
	StatusNew Status = iota
	StatusInProcess
	StatusCompleted
)

type Task struct {
	TID         string
	Title       string
	Description string
	Status      Status
}

func NewTask(tid string, title string, description string) (*Task, error) {
	if tid == "" {
		return nil, ErrEmptyTID
	}
	if title == "" {
		return nil, ErrEmptyTitle
	}

	return &Task{
		TID:         tid,
		Title:       title,
		Description: description,
		Status:      StatusNew,
	}, nil
}

func (t *Task) Start() error {
	if t.Status != StatusNew {
		return ErrCannotStart
	}
	t.Status = StatusInProcess
	return nil
}

func (t *Task) Complete() error {
	if t.Status != StatusInProcess {
		return ErrCannotComplete
	}
	t.Status = StatusCompleted
	return nil
}
