package tasks

import (
	"context"
	"errors"
	"fmt"
	domain "todos-api/internal/domain/tasks"
	"todos-api/internal/repository/postgres"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) GetAll(ctx context.Context) ([]*domain.Task, error) {
	rows, err := r.db.Query(ctx,
		`SELECT tid, title, description, status FROM tasks`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*domain.Task
	for rows.Next() {
		t := new(domain.Task)

		err := rows.Scan(
			&t.TID,
			&t.Title,
			&t.Description,
			&t.Status,
		)
		if err != nil {
			return nil, err
		}
		result = append(result, t)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return result, nil
}

func (r *Repository) GetByID(ctx context.Context, id string) (*domain.Task, error) {
	row := r.db.QueryRow(ctx, `SELECT tid, title, description, status  FROM tasks WHERE tid = $1`, id)

	result := new(domain.Task)
	err := row.Scan(
		&result.TID,
		&result.Title,
		&result.Description,
		&result.Status,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.ErrTaskNotFound
		}
		return nil, err
	}
	return result, nil
}

func (r *Repository) Create(ctx context.Context, task *domain.Task) error {
	if task == nil {
		return domain.ErrTaskIsNil
	}
	_, err := r.db.Exec(ctx,
		`INSERT INTO tasks (tid, title, description, status) VALUES ($1, $2, $3, $4)`,
		task.TID,
		task.Title,
		task.Description,
		task.Status,
	)

	// Не знаю насколько правильно. https://www.postgresql.org/docs/current/errcodes-appendix.html, действительно postgres кидает эти коды и их так обрабатывают?
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == postgres.ErrCodeUniqueViolation {
			return domain.ErrTaskAlreadyExists
		}
		return fmt.Errorf("create task: %w", err)
	}
	return nil
}

func (r *Repository) Update(ctx context.Context, task *domain.Task) error {
	if task == nil {
		return domain.ErrTaskIsNil
	}

	commandTag, err := r.db.Exec(ctx,
		`UPDATE tasks SET title = $1, description = $2, status = $3 WHERE tid = $4`,
		task.Title,
		task.Description,
		task.Status,
		task.TID,
	)

	if commandTag.RowsAffected() == 0 {
		return domain.ErrTaskNotFound
	}

	if err != nil {
		return fmt.Errorf("update task: %s %w", task.TID, err)
	}
	return nil
}

func (r *Repository) Delete(ctx context.Context, id string) error {
	commandTag, err := r.db.Exec(ctx,
		"DELETE FROM tasks WHERE tid = $1",
		id,
	)

	if err != nil {
		return fmt.Errorf("delete task %s: %w", id, err)
	}
	if commandTag.RowsAffected() == 0 {
		return domain.ErrTaskNotFound
	}

	return nil
}
