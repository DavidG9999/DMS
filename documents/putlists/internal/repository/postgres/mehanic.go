package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/DavidG9999/DMS/documents/putlists/internal/domain/entity"
	"github.com/DavidG9999/DMS/documents/putlists/internal/repository"
	"github.com/jmoiron/sqlx"
)

type MehanicPostgres struct {
	db *sqlx.DB
}

func NewMehanicPostgres(db *sqlx.DB) *MehanicPostgres {
	return &MehanicPostgres{
		db: db,
	}
}

type MehanicCreator interface {
	CreateMehanic(ctx context.Context, mehanic entity.Mehanic) (mehanicId int64, err error)
}

type MehanicProvider interface {
	GetMehanics(ctx context.Context) ([]entity.Mehanic, error)
}

type MehanicEditor interface {
	UpdateMehanic(ctx context.Context, mehanicId int64, updateData entity.UpdateMehanicInput) error
	DeleteMehanic(ctx context.Context, mehanicId int64) error
}

func (r *MehanicPostgres) CreateMehanic(ctx context.Context, mehanic entity.Mehanic) (mehanicId int64, err error) {
	const op = "postgres.CreateMehanic"

	query := fmt.Sprintf("INSERT INTO %s (fullname) VALUES ($1) RETURNING id", mehanicsTable)

	row := r.db.QueryRowContext(ctx, query, mehanic.FullName)

	if err := row.Scan(&mehanicId); err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	return mehanicId, nil
}

func (r *MehanicPostgres) GetMehanics(ctx context.Context) ([]entity.Mehanic, error) {
	const op = "postgres.GetMehanics"
	var meсhanics []entity.Mehanic

	query := fmt.Sprintf("SELECT * FROM %s", mehanicsTable)

	err := r.db.SelectContext(ctx, &meсhanics, query)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return []entity.Mehanic{}, fmt.Errorf("%s: %w", op, repository.ErrMehanicNotFound)
		}
		return []entity.Mehanic{}, fmt.Errorf("%s: %w", op, err)
	}
	return meсhanics, err
}

func (r *MehanicPostgres) UpdateMehanic(ctx context.Context, mehanicId int64, updateData entity.UpdateMehanicInput) error {
	const op = "postgres.UpdateMehanic"
	var id int64

	query := fmt.Sprintf("UPDATE %s SET fullname=$1 WHERE id=$2 RETURNING id", mehanicsTable)

	err := r.db.GetContext(ctx, &id, query, updateData.FullName, mehanicId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("%s: %w", op, repository.ErrMehanicNotFound)
		}
		return fmt.Errorf("%s: %w", op, err)
	}
	return err
}

func (r *MehanicPostgres) DeleteMehanic(ctx context.Context, mehanicId int64) error {
	const op = "postgres.DeleteMehanic"
	var id int64

	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1 RETURNING id", mehanicsTable)

	err := r.db.GetContext(ctx, &id, query, mehanicId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("%s: %w", op, repository.ErrMehanicNotFound)
		}
		return fmt.Errorf("%s: %w", op, err)
	}
	return err
}
