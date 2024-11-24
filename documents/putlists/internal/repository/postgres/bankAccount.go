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

type BankAccountPostgres struct {
	db *sqlx.DB
}

func NewBankAccountPostgres(db *sqlx.DB) *BankAccountPostgres {
	return &BankAccountPostgres{
		db: db,
	}
}

type BankAccountCreator interface {
	CreateBankAccount(ctx context.Context, bankAccount entity.BankAccount) (bankAccountId int64, err error)
}

type BankAccountProvider interface {
	GetBankAccounts(ctx context.Context, organizationId int64) ([]entity.BankAccount, error)
}

type BankAccountEditor interface {
	UpdateBankAccount(ctx context.Context, bankAccountId int64, updateData entity.UpdateBankAccountInput) error
	DeleteBankAccount(ctx context.Context, bankAccountId int64) error
}

func (r *BankAccountPostgres) CreateBankAccount(ctx context.Context, bankAccount entity.BankAccount) (bankAccountId int64, err error) {
	const op = "postgres.CreateBankAccount"

	query := fmt.Sprintf("INSERT INTO %s (bankaccountnumber, bankname, bankidnumber, organizationid) VALUES ($1, $2, $3, $4) RETURNING id", bankAccountsTable)

	row := r.db.QueryRowContext(ctx, query, bankAccount.BankAccountNumber, bankAccount.BankName, bankAccount.BankIdNumber, bankAccount.OrganizationId)

	if err := row.Scan(&bankAccountId); err != nil {
		var postgresErr *pgconn.PgError
		if errors.As(err, &postgresErr) && postgresErr.Code == "23505" {
			return 0, fmt.Errorf("%s: %w", op, repository.ErrBankAccExists)
		}
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	return bankAccountId, nil
}

func (r *BankAccountPostgres) GetBankAccounts(ctx context.Context, organizationId int64) ([]entity.BankAccount, error) {
	const op = "postgres.GetBankAccounts"

	var bankAccounts []entity.BankAccount

	query := fmt.Sprintf("SELECT * FROM %s WHERE organizationid=$1", bankAccountsTable)

	err := r.db.SelectContext(ctx, &bankAccounts, query, organizationId)

	if errors.Is(err, sql.ErrNoRows) {
		return []entity.BankAccount{}, fmt.Errorf("%s: %w", op, repository.ErrBankAccNotFound)
	}
	return bankAccounts, err
}

func (r *BankAccountPostgres) UpdateBankAccount(ctx context.Context, bankAccountId int64, updateData entity.UpdateBankAccountInput) error {
	const op = "postgres.UpdateBankAccount"
	var id int64

	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if updateData.BankAccountNumber != nil {
		setValues = append(setValues, fmt.Sprintf("bankaccountnumber=$%d", argId))
		args = append(args, updateData.BankAccountNumber)
		argId++
	}
	if updateData.BankName != nil {
		setValues = append(setValues, fmt.Sprintf("bankname=$%d", argId))
		args = append(args, updateData.BankName)
		argId++
	}
	if updateData.BankIdNumber != nil {
		setValues = append(setValues, fmt.Sprintf("bankidnumber=$%d", argId))
		args = append(args, updateData.BankIdNumber)
		argId++
	}
	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=$%d RETURNING id", bankAccountsTable, setQuery, argId)
	args = append(args, bankAccountId)

	err := r.db.GetContext(ctx, &id, query, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("%s: %w", op, repository.ErrBankAccNotFound)
		}
		return fmt.Errorf("%s: %w", op, err)
	}

	return err
}

func (r *BankAccountPostgres) DeleteBankAccount(ctx context.Context, bankAccountId int64) error {
	const op = "postgres.DeleteBankAccount"
	var id int64

	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1 RETURNING id", bankAccountsTable)

	err := r.db.GetContext(ctx, &id, query, bankAccountId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("%s: %w", op, repository.ErrBankAccNotFound)
		}
		return fmt.Errorf("%s: %w", op, err)
	}

	return err
}
