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

type OrganizationPostgres struct {
	db *sqlx.DB
}

func NewOrganizationPostgres(db *sqlx.DB) *OrganizationPostgres {
	return &OrganizationPostgres{
		db: db,
	}
}

type OrganizationCreator interface {
	CreateOrganization(ctx context.Context, organization entity.Organization) (organizationId int64, err error)
}

type OrganizationProvider interface {
	GetOrganizations(ctx context.Context) ([]entity.Organization, error)
}

type OrganizationEditor interface {
	UpdateOrganization(ctx context.Context, organizationId int64, updateData entity.UpdateOrganizationInput) error
	DeleteOrganization(ctx context.Context, organizationId int64) error
}

func (r *OrganizationPostgres) CreateOrganization(ctx context.Context, organization entity.Organization) (organizationId int64, err error) {
	const op = "postgres.CreateOrganization"

	query := fmt.Sprintf("INSERT INTO %s (name, address, chief, financialchief, innkpp) VALUES ($1, $2, $3, $4, $5) RETURNING id", organizationsTable)

	row := r.db.QueryRowContext(ctx, query, organization.Name, organization.Address, organization.Chief, organization.FinancialChief, organization.InnKpp)

	if err := row.Scan(&organizationId); err != nil {
		var postgressErr *pgconn.PgError
		if errors.As(err, &postgressErr) && postgressErr.Code == "23505" {
			return 0, fmt.Errorf("%s: %w", op, repository.ErrOrganizationExists)
		}
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	return organizationId, nil
}

func (r *OrganizationPostgres) GetOrganizations(ctx context.Context) ([]entity.Organization, error) {
	const op = "postgres.GetOrganizations"
	var organizations []entity.Organization

	query := fmt.Sprintf("SELECT * FROM %s", organizationsTable)

	err := r.db.Select(&organizations, query)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return []entity.Organization{}, fmt.Errorf("%s: %w", op, repository.ErrOrganizationNotFound)
		}
		return []entity.Organization{}, fmt.Errorf("%s: %w", op, err)
	}
	return organizations, err
}

func (r *OrganizationPostgres) UpdateOrganization(ctx context.Context, organizationId int64, updateData entity.UpdateOrganizationInput) error {
	const op = "postgres.UpdateOrganization"
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
	if updateData.Chief != nil {
		setValues = append(setValues, fmt.Sprintf("chief=$%d", argId))
		args = append(args, updateData.Chief)
		argId++
	}
	if updateData.FinancialChief != nil {
		setValues = append(setValues, fmt.Sprintf("financialchief=$%d", argId))
		args = append(args, updateData.FinancialChief)
		argId++
	}
	if updateData.InnKpp != nil {
		setValues = append(setValues, fmt.Sprintf("innkpp=$%d", argId))
		args = append(args, updateData.InnKpp)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=$%d RETURNING id", organizationsTable, setQuery, argId)
	args = append(args, organizationId)

	err := r.db.GetContext(ctx, &id, query, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("%s: %w", op, repository.ErrOrganizationNotFound)
		}
		return fmt.Errorf("%s: %w", op, err)
	}
	return err
}

func (r *OrganizationPostgres) DeleteOrganization(ctx context.Context, organizationId int64) error {
	const op = "postgres.DeleteOrganization"
	var id int64

	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1 RETURNING id", organizationsTable)

	err := r.db.GetContext(ctx, &id, query, organizationId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("%s: %w", op, repository.ErrOrganizationNotFound)
		}
		return fmt.Errorf("%s: %w", op, err)
	}
	return err
}
