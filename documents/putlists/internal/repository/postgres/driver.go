package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/DavidG9999/DMS/documents/putlists/internal/domain/entity"
	"github.com/DavidG9999/DMS/documents/putlists/internal/repository"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jmoiron/sqlx"
)

type DriverPostgres struct {
	db *sqlx.DB
}

func NewDriverPostgres(db *sqlx.DB) *DriverPostgres {
	return &DriverPostgres{
		db: db,
	}
}

type DriverCreator interface {
	CreateDriver(ctx context.Context, driver entity.Driver) (driverId int64, err error)
}

type DriverProvider interface {
	GetDrivers(ctx context.Context) ([]entity.Driver, error)
}

type DriverEditor interface {
	UpdateDriver(ctx context.Context, driverId int64, updateData entity.UpdateDriverInput) error
	DeleteDriver(ctx context.Context, driverId int64) error
}

func (r *DriverPostgres) CreateDriver(ctx context.Context, driver entity.Driver) (driverId int64, err error) {
	const op = "postgres.CreateDriver"

	query := fmt.Sprintf("INSERT INTO %s (fullname, license, class) VALUES ($1,$2,$3) RETURNING id", driversTable)

	row := r.db.QueryRowContext(ctx, query, driver.FullName, driver.License, driver.Class)

	if err := row.Scan(&driverId); err != nil {
		var postgresErr *pgconn.PgError
		if errors.As(err, &postgresErr) && postgresErr.Code == "23505" {
			return 0, fmt.Errorf("%s: %w", op, repository.ErrDriverExists)
		}
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	return driverId, nil
}

func (r *DriverPostgres) GetDrivers(ctx context.Context) ([]entity.Driver, error) {
	const op = "postgres.GetDrivers"
	var drivers []entity.Driver

	query := fmt.Sprintf("SELECT * FROM %s", driversTable)

	err := r.db.SelectContext(ctx, &drivers, query)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return []entity.Driver{}, fmt.Errorf("%s: %w", op, repository.ErrDriverNotFound)
		}
		return []entity.Driver{}, fmt.Errorf("%s: %w", op, err)
	}
	return drivers, err
}

func (r *DriverPostgres) UpdateDriver(ctx context.Context, driverId int64, updateData entity.UpdateDriverInput) error {
	const op = "postgres.UpdateDriver"
	var id int64

	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if updateData.FullName != nil {
		setValues = append(setValues, fmt.Sprintf("fullname=$%d", argId))
		args = append(args, updateData.FullName)
		argId++
	}

	if updateData.License != nil {
		setValues = append(setValues, fmt.Sprintf("license=$%d", argId))
		args = append(args, updateData.License)
		argId++
	}

	if updateData.Class != nil {
		setValues = append(setValues, fmt.Sprintf("class=$%d", argId))
		args = append(args, updateData.Class)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=$%d RETURNING id", driversTable, setQuery, argId)
	args = append(args, driverId)

	err := r.db.GetContext(ctx, &id, query, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("%s: %w", op, repository.ErrDriverNotFound)
		}
		return fmt.Errorf("%s %w", op, err)
	}
	return err
}

func (r *DriverPostgres) DeleteDriver(ctx context.Context, driverId int64) error {
	const op = "postgres.DeleteDriver"
	var id int64

	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1 RETURNING id", driversTable)

	err := r.db.GetContext(ctx, &id, query, driverId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("%s: %w", op, repository.ErrDriverNotFound)
		}
		return fmt.Errorf("%s %w", op, err)
	}
	return err
}
