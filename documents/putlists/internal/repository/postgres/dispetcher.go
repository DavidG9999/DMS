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

type DispetcherPostgres struct {
	db *sqlx.DB
}

func NewDispetcherPostgres(db *sqlx.DB) *DispetcherPostgres {
	return &DispetcherPostgres{
		db: db,
	}
}

type DispetcherCreator interface {
	CreateDispetcher(ctx context.Context, dispetcher entity.Dispetcher) (dispetcherId int64, err error)
}

type DispetcherProvider interface {
	GetDispetchers(ctx context.Context) ([]entity.Dispetcher, error)
}

type DispetcherEditor interface {
	UpdateDispetcher(ctx context.Context, dispetcherId int64, updateData entity.UpdateDispetcherInput) error
	DeleteDispetcher(ctx context.Context, dispetcherId int64) error
}

func (r *DispetcherPostgres) CreateDispetcher(ctx context.Context, dispetcher entity.Dispetcher) (dispetcherId int64, err error) {
	const op = "postgres.CreateDispetcher"

	query := fmt.Sprintf("INSERT INTO %s (fullname) VALUES ($1) RETURNING id", dispetchersTable)

	row := r.db.QueryRowContext(ctx, query, dispetcher.FullName)

	if err := row.Scan(&dispetcherId); err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	return dispetcherId, nil
}

func (r *DispetcherPostgres) GetDispetchers(ctx context.Context) ([]entity.Dispetcher, error) {
	const op = "postgres.GetDispetchers"

	var dispetchers []entity.Dispetcher

	query := fmt.Sprintf("SELECT * FROM %s", dispetchersTable)

	err := r.db.SelectContext(ctx, &dispetchers, query)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return []entity.Dispetcher{}, fmt.Errorf("%s: %w", op, repository.ErrDispethcerNotFound)
		}
		return []entity.Dispetcher{}, fmt.Errorf("%s: %w", op, err)
	}
	return dispetchers, err
}

func (r *DispetcherPostgres) UpdateDispetcher(ctx context.Context, dispetcherId int64, updateData entity.UpdateDispetcherInput) error {
	const op = "postgres.UpdateDispetcher"
	var id int64

	query := fmt.Sprintf("UPDATE %s SET fullname=$1 WHERE id=$2 RETURNING id", dispetchersTable)

	err := r.db.GetContext(ctx, &id, query, updateData.FullName, dispetcherId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("%s: %w", op, repository.ErrDispethcerNotFound)
		}
		return fmt.Errorf("%s: %w", op, err)
	}
	return err
}

func (r *DispetcherPostgres) DeleteDispetcher(ctx context.Context, dispetcherId int64) error {
	const op = "postgres.DeleteDispetcher"
	var id int64

	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1 RETURNING id", dispetchersTable)

	err := r.db.GetContext(ctx, &id, query, dispetcherId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("%s: %w", op, repository.ErrDispethcerNotFound)
		}
		return fmt.Errorf("%s: %w", op, err)
	}
	return err
}
