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

type ContragentPostgres struct {
	db *sqlx.DB
}

func NewContragentPostgres(db *sqlx.DB) *ContragentPostgres {
	return &ContragentPostgres{
		db: db,
	}
}

type ContragentCreator interface {
	CreateContragent(ctx context.Context, contragent entity.Contragent) (contragentId int64, err error)
}

type ContragentProvider interface {
	GetContragents(ctx context.Context) ([]entity.Contragent, error)
}

type ContragentEditor interface {
	UpdateContragent(ctx context.Context, contragentId int64, updateData entity.UpdateContragentInput) error
	DeleteContragent(ctx context.Context, contragentId int64) error
}

func (r *ContragentPostgres) CreateContragent(ctx context.Context, contragent entity.Contragent) (contragentId int64, err error) {
	const op = "postgres.CreateContragent"

	query := fmt.Sprintf("INSERT INTO %s (name, address, innkpp) VALUES ($1, $2, $3) RETURNING id", contragentsTable)

	row := r.db.QueryRowContext(ctx, query, contragent.Name, contragent.Address, contragent.InnKpp)

	if err := row.Scan(&contragentId); err != nil {
		var postgresErr *pgconn.PgError
		if errors.As(err, &postgresErr) && postgresErr.Code == "23505" {
			return 0, fmt.Errorf("%s: %w", op, repository.ErrContragentExists)
		}
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	return contragentId, nil
}

func (r *ContragentPostgres) GetContragents(ctx context.Context) ([]entity.Contragent, error) {
	const op = "postgres.GetContragents"

	var contragents []entity.Contragent
	query := fmt.Sprintf("SELECT * FROM %s", contragentsTable)

	err := r.db.SelectContext(ctx, &contragents, query)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return []entity.Contragent{}, fmt.Errorf("%s: %w", op, repository.ErrContragentNotFound)
		}
		return []entity.Contragent{}, fmt.Errorf("%s: %w", op, err)
	}

	return contragents, err
}

func (r *ContragentPostgres) UpdateContragent(ctx context.Context, contragentId int64, updateData entity.UpdateContragentInput) error {
	const op = "postgres.UpgrateContragent"
	var id int64

	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if updateData.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, updateData.Name)
		argId++
	}

	if updateData.Address != nil {
		setValues = append(setValues, fmt.Sprintf("address=$%d", argId))
		args = append(args, updateData.Address)
		argId++
	}

	if updateData.InnKpp != nil {
		setValues = append(setValues, fmt.Sprintf("innkpp=$%d", argId))
		args = append(args, updateData.InnKpp)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=$%d RETURNING id", contragentsTable, setQuery, argId)
	args = append(args, contragentId)

	err := r.db.GetContext(ctx, &id, query, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("%s: %w", op, repository.ErrContragentNotFound)
		}
		return fmt.Errorf("%s: %w", op, err)
	}
	return err
}

func (r *ContragentPostgres) DeleteContragent(ctx context.Context, contragentId int64) error {
	const op = "postgres.DeleteContragent"
	var id int64

	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1 RETURNING id", contragentsTable)

	err := r.db.GetContext(ctx, &id, query, contragentId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("%s: %w", op, repository.ErrContragentNotFound)
		}
		return fmt.Errorf("%s: %w", op, err)
	}
	return err
}
