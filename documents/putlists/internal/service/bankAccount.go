package service

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/DavidG9999/DMS/documents/putlists/internal/domain/entity"
	"github.com/DavidG9999/DMS/documents/putlists/internal/repository"
	"github.com/DavidG9999/DMS/documents/putlists/internal/repository/postgres"
)

type BankAccountService struct {
	repo   postgres.BankAccountRepository
	logger *slog.Logger
}

func NewBankAccountService(logger *slog.Logger,repo postgres.BankAccountRepository) *BankAccountService {
	return &BankAccountService{
		logger: logger,
		repo: repo,
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

func (s *BankAccountService) CreateBankAccount(ctx context.Context, bankAccount entity.BankAccount) (bankAccountId int64, err error) {
	const op = "putlist_service.CreateBankAccount"

	logger := s.logger.With(
		slog.String("op", op),
	)
	logger.Info("creating bank account")

	bankAccountId, err = s.repo.CreateBankAccount(ctx, bankAccount)
	if err != nil {
		if errors.Is(err, repository.ErrBankAccExists) {
			logger.Error("bank account already exists")
			return 0, fmt.Errorf("%s: %w", op, ErrBankAccExists)
		}
		logger.Error("failed to create bank account")
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	logger.Info("bank account created")
	return bankAccountId, nil
}

func (s *BankAccountService) GetBankAccounts(ctx context.Context, organizationId int64) ([]entity.BankAccount, error) {
	const op = "putlist_service.GetBankAccounts"

	logger := s.logger.With(
		slog.String("op", op),
	)
	logger.Info("getting bank account")

	bankAccounts, err := s.repo.GetBankAccounts(ctx, organizationId)
	if err != nil {
		if errors.Is(err, repository.ErrBankAccNotFound) {
			logger.Error("bank account not found")
			return []entity.BankAccount{}, fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
		}
		logger.Error("failed to get bank accounts")
		return []entity.BankAccount{}, fmt.Errorf("%s: %w", op, err)
	}
	logger.Info("bank accounts received")
	return bankAccounts, nil
}

func (s *BankAccountService) UpdateBankAccount(ctx context.Context, bankAccountId int64, updateData entity.UpdateBankAccountInput) error {
	const op = "putlist_service.UpdateBankAccount"

	logger := s.logger.With(
		slog.String("op", op),
	)
	logger.Info("updating bank account")

	err := s.repo.UpdateBankAccount(ctx, bankAccountId, updateData)
	if err != nil {
		if errors.Is(err, repository.ErrBankAccNotFound) {
			logger.Error("bank account not found")
			return fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
		}
		logger.Error("failed to update bank account")
		return fmt.Errorf("%s: %w", op, err)
	}
	logger.Info("bank account updated")
	return err
}

func (s *BankAccountService) DeleteBankAccount(ctx context.Context, bankAccountId int64) error {
	const op = "putlist_service.DeleteBankAccount"

	logger := s.logger.With(
		slog.String("op", op),
	)
	logger.Info("deleting bank account")

	err := s.repo.DeleteBankAccount(ctx, bankAccountId)
	if err != nil {
		if errors.Is(err, repository.ErrBankAccNotFound) {
			logger.Error("bank account not found")
			return fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
		}
		logger.Error("failed to delete bank account")
		return fmt.Errorf("%s: %w", op, err)
	}
	logger.Info("bank account deleted")
	return err
}
