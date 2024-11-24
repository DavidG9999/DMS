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

type PutlistBodyPostgres struct {
	db *sqlx.DB
}

func NewPutlistBodyPostgres(db *sqlx.DB) *PutlistBodyPostgres {
	return &PutlistBodyPostgres{
		db: db,
	}
}

type PutlistBodyCreator interface {
	CreatePutlistBody(ctx context.Context, putlistBody entity.PutlistBody) (putlistBodyId int64, err error)
}

type PutlistBodyProvider interface {
	GetPutlistBodies(ctx context.Context, putlistNumber int64) ([]entity.PutlistBody, error)
}

type PutlistBodyEditor interface {
	UpdatePutlistBody(ctx context.Context, putlistBodyId int64, updateData entity.UpdatePutlistBodyInput) error
	DeletePutlistBody(ctx context.Context, putlistBodyId int64) error
}

func (r *PutlistBodyPostgres) CreatePutlistBody(ctx context.Context, putlistBody entity.PutlistBody) (putlistBodyId int64, err error) {
	const op = "postgres.CreatePutlistBody"

	query := fmt.Sprintf("INSERT INTO %s (putlistheadernumber, number, contragentid, item, timewith, timefor) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id", putlistBodiesTable)

	row := r.db.QueryRowContext(ctx, query, putlistBody.PutlistNumber, putlistBody.Number, putlistBody.ContragentId, putlistBody.Item, putlistBody.TimeWith, putlistBody.TimeFor)

	if err := row.Scan(&putlistBodyId); err != nil {
		var postgresErr *pgconn.PgError
		if errors.As(err, &postgresErr) && postgresErr.Code == "22007" {
			return 0, fmt.Errorf("%s: %w", op, repository.ErrInvalidDateTimeFormat)
		}
		if errors.As(err, &postgresErr) && postgresErr.Code == "22008" {
			return 0, fmt.Errorf("%s: %w", op, repository.ErrInvalidDateTimeFormat)
		}
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	return putlistBodyId, nil
}

func (r *PutlistBodyPostgres) GetPutlistBodies(ctx context.Context, putlistNumber int64) ([]entity.PutlistBody, error) {
	const op = "postgres.GetPutlistBodies"
	var putlistBodies []entity.PutlistBody

	query := fmt.Sprintf("SELECT * FROM %s WHERE putlistheadernumber=$1", putlistBodiesTable)

	err := r.db.SelectContext(ctx, &putlistBodies, query, putlistNumber)
	if err != nil {
		return []entity.PutlistBody{}, fmt.Errorf("%s: %w", op, err)
	}
	return putlistBodies, err
}

func (r *PutlistBodyPostgres) UpdatePutlistBody(ctx context.Context, putlistBodyId int64, updateData entity.UpdatePutlistBodyInput) error {
	const op = "postgres.UpdatePutlistBody"
	var id int64

	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1
	if updateData.Number != nil {
		setValues = append(setValues, fmt.Sprintf("number=$%d", argId))
		args = append(args, updateData.Number)
		argId++
	}
	if updateData.ContragentId != nil {
		setValues = append(setValues, fmt.Sprintf("contragentid=$%d", argId))
		args = append(args, updateData.ContragentId)
		argId++
	}
	if updateData.Item != nil {
		setValues = append(setValues, fmt.Sprintf("item=$%d", argId))
		args = append(args, updateData.Item)
		argId++
	}
	if updateData.TimeWith != nil {
		setValues = append(setValues, fmt.Sprintf("timewith=$%d", argId))
		args = append(args, updateData.TimeWith)
		argId++
	}
	if updateData.TimeFor != nil {
		setValues = append(setValues, fmt.Sprintf("timefor=$%d", argId))
		args = append(args, updateData.TimeFor)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=$%d RETURNING id", putlistBodiesTable, setQuery, argId)
	args = append(args, putlistBodyId)

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
			return fmt.Errorf("%s: %w", op, repository.ErrPutlistBodyNotFound)
		}
		return fmt.Errorf("%s: %w", op, err)
	}
	return err
}

func (r *PutlistBodyPostgres) DeletePutlistBody(ctx context.Context, putlistBodyId int64) error {
	const op = "postgres.DeletePutlistBody"
	var id int64

	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1 RETURNING id", putlistBodiesTable)

	err := r.db.GetContext(ctx, &id, query, putlistBodyId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("%s: %w", op, repository.ErrPutlistBodyNotFound)
		}
		return fmt.Errorf("%s: %w", op, err)
	}
	return err
}
