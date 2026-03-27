package users

import (
	"context"
	"errors"
	"fmt"
	domain "todos-api/internal/domain/users"
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

func (r *Repository) GetAll(ctx context.Context) ([]*domain.User, error) {
	rows, err := r.db.Query(ctx,
		`SELECT uid, name, email FROM users`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*domain.User
	for rows.Next() {
		u := new(domain.User)

		err := rows.Scan(
			&u.UID,
			&u.Name,
			&u.Email,
		)
		if err != nil {
			return nil, err
		}
		result = append(result, u)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return result, nil
}

func (r *Repository) GetByID(ctx context.Context, id string) (*domain.User, error) {
	row := r.db.QueryRow(ctx, `SELECT uid, name, email FROM users WHERE uid = $1`, id)

	result := new(domain.User)
	err := row.Scan(
		&result.UID,
		&result.Name,
		&result.Email,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.ErrUserNotFound
		}
		return nil, err
	}
	return result, nil
}

func (r *Repository) Create(ctx context.Context, user *domain.User) error {
	if user == nil {
		return domain.ErrUserIsNil
	}
	_, err := r.db.Exec(ctx,
		`INSERT INTO users (uid, name, email, password) VALUES ($1, $2, $3, $4)`,
		user.UID,
		user.Name,
		user.Email,
		user.Password,
	)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == postgres.ErrCodeUniqueViolation {
			return domain.ErrUserAlreadyExists
		}
		return fmt.Errorf("create user: %w", err)
	}
	return nil
}

func (r *Repository) Update(ctx context.Context, user *domain.User) error {
	if user == nil {
		return domain.ErrUserIsNil
	}

	commandTag, err := r.db.Exec(ctx,
		`UPDATE users SET name = $1, email = $2 WHERE uid = $4`,
		user.Name,
		user.Email,
		user.UID,
	)

	if commandTag.RowsAffected() == 0 {
		return domain.ErrUserNotFound
	}

	if err != nil {
		return fmt.Errorf("update user: %s %w", user.UID, err)
	}
	return nil
}

func (r *Repository) Delete(ctx context.Context, id string) error {
	commandTag, err := r.db.Exec(ctx,
		"DELETE FROM users WHERE uid = $1",
		id,
	)

	if err != nil {
		return fmt.Errorf("delete task %s: %w", id, err)
	}
	if commandTag.RowsAffected() == 0 {
		return domain.ErrUserNotFound
	}

	return nil
}

func (r *Repository) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	row := r.db.QueryRow(ctx, `SELECT uid, name, email, password FROM users WHERE email = $1`, email)

	result := new(domain.User)
	err := row.Scan(
		&result.UID,
		&result.Name,
		&result.Email,
		&result.Password,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.ErrUserNotFound
		}
		return nil, err
	}
	return result, nil
}
