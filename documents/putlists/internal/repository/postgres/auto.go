package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/DavidG9999/DMS/documents/putlists/internal/domain/entity"
	"github.com/DavidG9999/DMS/documents/putlists/internal/repository"
	"github.com/jmoiron/sqlx"

	"github.com/jackc/pgx/v5/pgconn"
)

type AutoPostgres struct {
	db *sqlx.DB
}

func NewAutoPostgres(db *sqlx.DB) *AutoPostgres {
	return &AutoPostgres{
		db: db,
	}
}

type AutoCreator interface {
	CreateAuto(ctx context.Context, auto entity.Auto) (autoId int64, err error)
}

type AutoProvider interface {
	GetAutos(ctx context.Context) ([]entity.Auto, error)
}

type AutoEditor interface {
	UpdateAuto(ctx context.Context, autoId int64, updateData entity.UpdateAutoInput) error
	DeleteAuto(ctx context.Context, autoId int64) error
}

func (r *AutoPostgres) CreateAuto(ctx context.Context, auto entity.Auto) (autoId int64, err error) {
	const op = "postgres.CreateAuto"

	query := fmt.Sprintf("INSERT INTO %s (brand, model, statenumber) VALUES ($1, $2, $3) RETURNING id", autosTable)

	row := r.db.QueryRowContext(ctx, query, auto.Brand, auto.Model, auto.StateNumber)

	if err := row.Scan(&autoId); err != nil {
		var postgresErr *pgconn.PgError
		if errors.As(err, &postgresErr) && postgresErr.Code == "23505" {
			return 0, fmt.Errorf("%s: %w", op, repository.ErrAutoExists)
		}
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	return autoId, nil
}

func (r *AutoPostgres) GetAutos(ctx context.Context) ([]entity.Auto, error) {
	const op = "postgres.GetAutos"

	var autos []entity.Auto

	query := fmt.Sprintf("SELECT * FROM %s", autosTable)

	err := r.db.SelectContext(ctx, &autos, query)
	if errors.Is(err, sql.ErrNoRows) {
		return []entity.Auto{}, fmt.Errorf("%s: %w", op, repository.ErrAutoNotFound)
	}
	return autos, err
}

func (r *AutoPostgres) UpdateAuto(ctx context.Context, autoId int64, updateData entity.UpdateAutoInput) error {
	const op = "postgres.UpdateAuto"
	var id int64

	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if updateData.Brand != nil {
		setValues = append(setValues, fmt.Sprintf("brand=$%d", argId))
		args = append(args, updateData.Brand)
		argId++
	}

	if updateData.Model != nil {
		setValues = append(setValues, fmt.Sprintf("model=$%d", argId))
		args = append(args, updateData.Model)
		argId++
	}

	if updateData.StateNumber != nil {
		setValues = append(setValues, fmt.Sprintf("statenumber=$%d", argId))
		args = append(args, updateData.StateNumber)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=$%d RETURNING id", autosTable, setQuery, argId)
	args = append(args, autoId)

	err := r.db.GetContext(ctx, &id, query, args...)
	if errors.Is(err, sql.ErrNoRows) {
		return fmt.Errorf("%s: %w", op, repository.ErrAutoNotFound)
	}
	return err
}

func (r *AutoPostgres) DeleteAuto(ctx context.Context, autoId int64) error {
	const op = "postgres.DeleteUser"
	var id int64

	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1 RETURNING id", autosTable)

	err := r.db.GetContext(ctx, &id, query, autoId)
	if errors.Is(err, sql.ErrNoRows) {
		return fmt.Errorf("%s: %w", op, repository.ErrAutoNotFound)
	}
	return err
}
