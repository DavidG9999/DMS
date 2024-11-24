package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/DavidG9999/DMS/users/internal/domain/entity"
	"github.com/DavidG9999/DMS/users/internal/repository"
	"github.com/jmoiron/sqlx"

	"github.com/jackc/pgx/v5/pgconn"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

type UserCreator interface {
	CreateUser(ctx context.Context, name string, email string, passwordHash string) (userID int64, err error)
}

type UserProvider interface {
	GetUser(ctx context.Context, email string) (entity.User, error)
	GetUserById(ctx context.Context, userID int64) (entity.User, error)
}

type UserEditor interface {
	UpdateName(ctx context.Context, userID int64, updateName string) error
	UpdatePassword(ctx context.Context, userID int64, updatePasswordHash string) error
	DeleteUser(ctx context.Context, userID int64) error
}

func (r *UserPostgres) CreateUser(ctx context.Context, name string, email string, passwordHash string) (userID int64, err error) {
	const op = "postgres.CreateUser"

	query := fmt.Sprintf("INSERT INTO %s (name, email, passwordhash) VALUES ($1, $2, $3) returning id", usersTable)

	row := r.db.QueryRowContext(ctx, query, name, email, passwordHash)

	if err := row.Scan(&userID); err != nil {
		var postgresErr *pgconn.PgError
		if errors.As(err, &postgresErr) && postgresErr.Code == "23505" {

			return 0, fmt.Errorf("%s: %w", op, repository.ErrUserExists)
		}
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return userID, err
}

func (r *UserPostgres) GetUser(ctx context.Context, email string) (entity.User, error) {
	const op = "postgres.GetUser"
	var user entity.User

	query := fmt.Sprintf("SELECT id, name, email, passwordhash FROM %s WHERE email=$1", usersTable)

	err := r.db.GetContext(ctx, &user, query, email)
	if errors.Is(err, sql.ErrNoRows) {
		return entity.User{}, fmt.Errorf("%s: %w", op, repository.ErrUserNotFound)
	}
	return user, err
}

func (r *UserPostgres) GetUserById(ctx context.Context, userID int64) (entity.User, error) {
	const op = "postgres.GetUserById"
	var user entity.User

	query := fmt.Sprintf("SELECT name, email FROM %s WHERE id=$1", usersTable)

	err := r.db.GetContext(ctx, &user, query, userID)
	if errors.Is(err, sql.ErrNoRows) {
		return entity.User{}, fmt.Errorf("%s: %w", op, repository.ErrUserNotFound)
	}
	return user, err
}

func (r *UserPostgres) UpdateName(ctx context.Context, userID int64, updateName string) error {
	const op = "postgres.UpdateName"
	var id int64

	query := fmt.Sprintf("UPDATE %s SET name=$1 WHERE id=$2 RETURNING id", usersTable)

	err := r.db.GetContext(ctx, &id, query, updateName, userID)
	if errors.Is(err, sql.ErrNoRows) {
		return fmt.Errorf("%s: %w", op, repository.ErrUserNotFound)
	}
	return err
}

func (r *UserPostgres) UpdatePassword(ctx context.Context, userID int64, updatePasswordHash string) error {
	const op = "postgres.UpdatePassword"
	var id int64

	query := fmt.Sprintf("UPDATE %s SET passwordhash=$1 WHERE id=$2 RETURNING id", usersTable)

	err := r.db.GetContext(ctx, &id, query, updatePasswordHash, userID)
	if errors.Is(err, sql.ErrNoRows) {
		return fmt.Errorf("%s: %w", op, repository.ErrUserNotFound)
	}
	return err
}

func (r *UserPostgres) DeleteUser(ctx context.Context, userID int64) error {
	const op = "postgres.DeleteUser"
	var id int64

	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1 RETURNING id", usersTable)

	err := r.db.GetContext(ctx, &id, query, userID)
	if errors.Is(err, sql.ErrNoRows) {
		return fmt.Errorf("%s: %w", op, repository.ErrUserNotFound)
	}
	return err
}
