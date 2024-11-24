package service

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/DavidG9999/DMS/DMS_api_gateway/internal/entity"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type BankAccountCreator interface {
	CreateBankAccount(ctx context.Context, organizationId int64, bankAccount entity.BankAccount) (int64, error)
}

type BankAccountProvider interface {
	GetBankAccounts(ctx context.Context, organizationId int64) ([]entity.BankAccount, error)
}

type BankAccountEditor interface {
	UpdateBankAccount(ctx context.Context, bankAccountId int64, updateData entity.UpdateBankAccountInput) (string, error)
	DeleteBankAccount(ctx context.Context, bankAccountId int64) (string, error)
}

func (p *PutlistService) CreateBankAccount(ctx context.Context, organizationId int64, bankAccount entity.BankAccount) (int64, error) {
	const op = "service.CreateBankAccount"

	logger := p.logger.With(slog.String("op", op))
	logger.Info("creating bank account")

	bankAccountId, err := p.putlistClient.CreateBankAccount(ctx, organizationId, bankAccount.BankAccountNumber, bankAccount.BankName, bankAccount.BankIdNumber)
	if err != nil {
		if errors.Is(err, status.Error(codes.AlreadyExists, "bank account already exists")) {
			logger.Error("bank account already exists")
			return 0, fmt.Errorf("%s: %w", op, ErrOrganizationExists)
		}
		logger.Error("failed to create bank account")
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	logger.Info("bank account created")

	return bankAccountId, nil
}

func (p *PutlistService) GetBankAccounts(ctx context.Context, organizationId int64) ([]entity.BankAccount, error) {
	const op = "service.GetBankAccounts"

	logger := p.logger.With(slog.String("op", op))
	logger.Info("getting bank accounts")

	bankAccounts, err := p.putlistClient.GetBankAccounts(ctx, organizationId)
	if err != nil {
		logger.Error("failed to get bank accounts")
		return []entity.BankAccount{}, fmt.Errorf("%s: %w", op, err)
	}
	logger.Info("bank accounts received")

	return bankAccounts, nil
}

func (p *PutlistService) UpdateBankAccount(ctx context.Context, bankAccId int64, updateData entity.UpdateBankAccountInput) (string, error) {
	const op = "service.UpdateBankAccount"

	logger := p.logger.With(slog.String("op", op))
	logger.Info("updating bank account")

	err := updateData.Validate()
	if err != nil {
		return "", err
	}

	message, err := p.putlistClient.UpdateBankAccount(ctx, bankAccId, updateData.BankAccountNumber, updateData.BankName, updateData.BankIdNumber)
	if err != nil {
		if errors.Is(err, status.Error(codes.NotFound, "bank account not found")) {
			logger.Error("bank account not found")
			return "", fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
		}
		logger.Error("failed to update bank account")
		return "", fmt.Errorf("%s: %w", op, err)
	}
	logger.Info("bank account updated")

	return message, nil
}

func (p *PutlistService) DeleteBankAccount(ctx context.Context, bankAccId int64) (string, error) {
	const op = "service.DeleteBankAccount"

	logger := p.logger.With(slog.String("op", op))
	logger.Info("deleting bank account")

	message, err := p.putlistClient.DeleteBankAccount(ctx, bankAccId)
	if err != nil {
		if errors.Is(err, status.Error(codes.NotFound, "bank account not found")) {
			logger.Error("bank account not found")
			return "", fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
		}
		logger.Error("failed to delete bank account")
		return "", fmt.Errorf("%s: %w", op, err)
	}
	logger.Info("bank account deleted")

	return message, nil
}
