package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"context"

	"github.com/DavidG9999/DMS/documents/putlists/internal/domain/entity"
	"github.com/DavidG9999/DMS/documents/putlists/internal/repository"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jmoiron/sqlx"
)

type PutlistPostgres struct {
	db *sqlx.DB
}

func NewPutlistPostgres(db *sqlx.DB) *PutlistPostgres {
	return &PutlistPostgres{
		db: db,
	}
}

type PutlistCreator interface {
	CreatePutlist(ctx context.Context, putlist entity.PutlistHeader) (putlistId int64, err error)
}

type PutlistProvider interface {
	GetPutlists(ctx context.Context, userId int64) ([]entity.PutlistHeader, error)
	GetPutlistByNumber(ctx context.Context, userId int64, putlistNumber int64) (entity.PutlistHeader, error)
}

type PutlistEditor interface {
	UpdatePutlist(ctx context.Context, userId int64, putlistNumber int64, updateData entity.UpdatePutlistHeaderInput) error
	DeletePutlist(ctx context.Context, userId int64, putlistNumber int64) error
}

func (r *PutlistPostgres) CreatePutlist(ctx context.Context, putlist entity.PutlistHeader) (putlistId int64, err error) {
	const op = "postgres.CreatePutlist"

	query := fmt.Sprintf("INSERT INTO %s (userid, number, bankaccountid, datewith, datefor, autoid, driverid, dispetcherid, mehanicid) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id", putlistHeadersTable)

	row := r.db.QueryRowContext(ctx, query, putlist.UserId, putlist.Number, putlist.BankAccountId, putlist.DateWith, putlist.DateFor, putlist.AutoId, putlist.DriverId, putlist.DispetcherId, putlist.MehanicId)

	if err := row.Scan(&putlistId); err != nil {
		var postgresErr *pgconn.PgError
		if errors.As(err, &postgresErr) && postgresErr.Code == "22007" {
			return 0, fmt.Errorf("%s: %w", op, repository.ErrInvalidDateTimeFormat)
		}
		if errors.As(err, &postgresErr) && postgresErr.Code == "22008" {
			return 0, fmt.Errorf("%s: %w", op, repository.ErrInvalidDateTimeFormat)
		}
		if errors.As(err, &postgresErr) && postgresErr.Code == "23505" {
			return 0, fmt.Errorf("%s: %w", op, repository.ErrPutlistExists)
		}
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	return putlistId, nil
}

func (r *PutlistPostgres) GetPutlists(ctx context.Context, userId int64) ([]entity.PutlistHeader, error) {
	const op = "postgres.GetPutlistHeaders"
	var putlists []entity.PutlistHeader

	query := fmt.Sprintf("SELECT * FROM %s WHERE userid=$1", putlistHeadersTable)

	err := r.db.SelectContext(ctx, &putlists, query, userId)
	if err != nil {
		return []entity.PutlistHeader{}, fmt.Errorf("%s: %w", op, err)
	}

	return putlists, err
}

func (r *PutlistPostgres) GetPutlistByNumber(ctx context.Context, userId int64, putlistNumber int64) (entity.PutlistHeader, error) {
	const op = "postgres.GetPutlistByNumber"
	var putlist entity.PutlistHeader

	query := fmt.Sprintf("SELECT * FROM %s WHERE userid = $1 AND number = $2", putlistHeadersTable)

	err := r.db.GetContext(ctx, &putlist, query, userId, putlistNumber)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.PutlistHeader{}, fmt.Errorf("%s: %w", op, repository.ErrPutlistNotFound)
		}
		return entity.PutlistHeader{}, fmt.Errorf("%s: %w", op, err)
	}

	return putlist, err
}

func (r *PutlistPostgres) UpdatePutlist(ctx context.Context, userId int64, putlistNumber int64, updateData entity.UpdatePutlistHeaderInput) error {
	const op = "postgres.UpdatePutlist"
	var id int64

	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if updateData.BankAccountId != nil {
		setValues = append(setValues, fmt.Sprintf("bankaccountid=$%d", argId))
		args = append(args, updateData.BankAccountId)
		argId++
	}
	if updateData.DateWith != nil {
		setValues = append(setValues, fmt.Sprintf("datewith=$%d", argId))
		args = append(args, updateData.DateWith)
		argId++
	}
	if updateData.DateFor != nil {
		setValues = append(setValues, fmt.Sprintf("datefor=$%d", argId))
		args = append(args, updateData.DateFor)
		argId++
	}
	if updateData.AutoId != nil {
		setValues = append(setValues, fmt.Sprintf("autoid=$%d", argId))
		args = append(args, updateData.AutoId)
		argId++
	}
	if updateData.DriverId != nil {
		setValues = append(setValues, fmt.Sprintf("driverid=$%d", argId))
		args = append(args, updateData.DriverId)
		argId++
	}
	if updateData.DispetcherId != nil {
		setValues = append(setValues, fmt.Sprintf("dispetcherid=$%d", argId))
		args = append(args, updateData.DispetcherId)
		argId++
	}
	if updateData.MehanicId != nil {
		setValues = append(setValues, fmt.Sprintf("mehanicid=$%d", argId))
		args = append(args, updateData.MehanicId)
		argId++
	}
	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE userid=$%d AND number=$%d  RETURNING id", putlistHeadersTable, setQuery, argId, argId+1)
	args = append(args, userId, putlistNumber)

	err := r.db.GetContext(ctx, &id, query, args...)
	if err != nil {
		var postgresErr *pgconn.PgError
		if errors.As(err, &postgresErr) && postgresErr.Code == "22007" {
			return fmt.Errorf("%s: %w", op, repository.ErrInvalidDateTimeFormat)
		}
		if errors.As(err, &postgresErr) && postgresErr.Code == "22008" {
			return fmt.Errorf("%s: %w", op, repository.ErrInvalidDateTimeFormat)
		}
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("%s: %w", op, repository.ErrPutlistNotFound)
		}
		return fmt.Errorf("%s: %w", op, err)
	}

	return err
}

func (r *PutlistPostgres) DeletePutlist(ctx context.Context, userId int64, putlistNumber int64) error {
	const op = "postgres.DeletePutlist"
	var id int64

	query := fmt.Sprintf("DELETE FROM %s WHERE userid =$1 AND number = $2 RETURNING id", putlistHeadersTable)

	err := r.db.GetContext(ctx, &id, query, userId, putlistNumber)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("%s: %w", op, repository.ErrPutlistNotFound)
		}
		return fmt.Errorf("%s: %w", op, err)
	}

	return err
}
