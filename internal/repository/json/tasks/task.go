package tasks

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	domain "todos-api/internal/domain/tasks"
)

type Repository struct {
	filePath string
	mu       sync.Mutex
}

func NewRepository(path string) *Repository {
	dir := filepath.Dir(path)

	if err := os.MkdirAll(dir, 0755); err != nil {
		fmt.Println(err)
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		file, err := os.Create(path)
		if err != nil {
			panic(err)
		} else {
			defer func() {
				if err := file.Close(); err != nil {
					panic(err)
				}
			}()
		}
	}
	return &Repository{
		filePath: path,
	}
}

func (r *Repository) read() ([]*domain.Task, error) {
	data, err := os.ReadFile(r.filePath)

	if err != nil {
		if os.IsNotExist(err) {
			return []*domain.Task{}, nil
		}

		return nil, err
	}

	var tasks []*domain.Task

	if len(data) == 0 {
		return []*domain.Task{}, nil
	}

	err = json.Unmarshal(data, &tasks)

	return tasks, err
}

func (r *Repository) write(tasks []*domain.Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(r.filePath, data, 0644)
}

func (r *Repository) GetAll(_ context.Context, userID string) ([]*domain.Task, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	all, err := r.read()
	if err != nil {
		return nil, err
	}

	var result []*domain.Task
	for _, t := range all {
		if t.UserID == userID {
			result = append(result, t)
		}
	}
	return result, nil
}

func (r *Repository) GetByID(_ context.Context, id, userID string) (*domain.Task, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	tasks, err := r.read()
	if err != nil {
		return nil, err
	}
	for _, task := range tasks {
		if task.TID == id && task.UserID == userID {
			return task, nil
		}
	}
	return nil, domain.ErrTaskNotFound
}

func (r *Repository) Create(_ context.Context, task *domain.Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	tasksList, err := r.read()
	if err != nil {
		return err
	}

	tasks := append(tasksList, task)
	return r.write(tasks)
}

func (r *Repository) Update(_ context.Context, task *domain.Task) error {
	tasksList, err := r.read()
	if err != nil {
		return err
	}

	for i := range tasksList {
		if tasksList[i].TID == task.TID && tasksList[i].UserID == task.UserID {
			tasksList[i] = task

			return r.write(tasksList)
		}
	}

	return errors.New("task not found")
}

func (r *Repository) Delete(_ context.Context, id, userID string) error {
	tasksList, err := r.read()
	if err != nil {
		return err
	}

	for i := range tasksList {
		if tasksList[i].TID == id && tasksList[i].UserID == userID {
			tasksList = append(tasksList[:i], tasksList[i+1:]...)

			return r.write(tasksList)
		}
	}

	return domain.ErrTaskNotFound
}
