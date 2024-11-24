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

type PutlistCreator interface {
	CreatePutlist(ctx context.Context, userId int64, putlist entity.Putlist) (int64, error)
}

type PutlistProvider interface {
	GetPutlists(ctx context.Context, userId int64) ([]entity.Putlist, error)
	GetPutlistByNumber(ctx context.Context, userId int64, putlistNumber int64) (entity.Putlist, error)
}

type PutlistEditor interface {
	UpdatePutlist(ctx context.Context, userId int64, putlistNumber int64, updateData entity.UpdatePutlistHeaderInput) (string, error)
	DeletePutlist(ctx context.Context, userId int64, putlistNumber int64) (string, error)
}

func (p *PutlistService) CreatePutlist(ctx context.Context, userId int64, putlist entity.Putlist) (int64, error) {
	const op = "service.CreatePutlist"

	logger := p.logger.With(slog.String("op", op))
	logger.Info("creating putlist")

	putlistId, err := p.putlistClient.CreatePutlist(ctx, userId, putlist.Number, putlist.BankAccountId, putlist.DateWith, putlist.DateFor, putlist.AutoId, putlist.DriverId, putlist.DispetcherId, putlist.MehanicId)
	if err != nil {
		if errors.Is(err, status.Error(codes.AlreadyExists, "putlist already exists")) {
			logger.Error("putlist already exists")
			return 0, fmt.Errorf("%s: %w", op, ErrPutlistExists)
		}
		if errors.Is(err, status.Error(codes.InvalidArgument, "invalid datetime format")) {
			logger.Error("invalid datetime format")
			return 0, fmt.Errorf("%s: %w", op, ErrInvalidDateTimeFormat)
		}
		logger.Error("failed to create putlist")
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	logger.Info("putlist created")

	return putlistId, nil
}

func (p *PutlistService) GetPutlists(ctx context.Context, userId int64) ([]entity.Putlist, error) {
	const op = "service.GetPutlists"

	logger := p.logger.With(slog.String("op", op))
	logger.Info("getting putlists")

	putlists, err := p.putlistClient.GetPutlists(ctx, userId)
	if err != nil {
		logger.Error("failed to get putlists")
		return []entity.Putlist{}, fmt.Errorf("%s: %w", op, err)
	}
	logger.Info("putlists received")

	return putlists, nil
}

func (p *PutlistService) GetPutlistByNumber(ctx context.Context, userId int64, putlistNumber int64) (entity.Putlist, error) {
	const op = "service.GetPutlistByNumber"

	logger := p.logger.With(slog.String("op", op))
	logger.Info("getting putlist")

	putlist, err := p.putlistClient.GetPutlistByNumber(ctx, userId, putlistNumber)
	if err != nil {
		if errors.Is(err, status.Error(codes.NotFound, "putlist by this number not found")) {
			logger.Error("putlist by this number not found")
			return entity.Putlist{}, fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
		}
		logger.Error("failed to get putlist")
		return entity.Putlist{}, fmt.Errorf("%s: %w", op, err)
	}
	logger.Info("putlist received")

	return putlist, nil
}

func (p *PutlistService) UpdatePutlist(ctx context.Context, userId int64, putlistNumber int64, updateData entity.UpdatePutlistHeaderInput) (string, error) {
	const op = "service.UpdatePutlist"

	logger := p.logger.With(slog.String("op", op))
	logger.Info("updating putlist")

	err := updateData.Validate()
	if err != nil {
		return "", err
	}

	message, err := p.putlistClient.UpdatePutlist(ctx, userId, putlistNumber, updateData.BankAccountId, updateData.DateWith, updateData.DateFor, updateData.AutoId, updateData.DriverId, updateData.DispetcherId, updateData.MehanicId)
	if err != nil {
		if errors.Is(err, status.Error(codes.NotFound, "putlist not found")) {
			logger.Error("putlist not found")
			return "", fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
		}
		if errors.Is(err, status.Error(codes.InvalidArgument, "invalid datetime format")) {
			logger.Error("invalid datetime format")
			return "", fmt.Errorf("%s: %w", op, ErrInvalidDateTimeFormat)
		}
		logger.Error("failed to update putlist")
		return "", fmt.Errorf("%s: %w", op, err)
	}
	logger.Info("putlist updated")

	return message, nil
}

func (p *PutlistService) DeletePutlist(ctx context.Context, userId int64, putlistNumber int64) (string, error) {
	const op = "service.DeletePutlist"

	logger := p.logger.With(slog.String("op", op))
	logger.Info("deleting putlist")

	message, err := p.putlistClient.DeletePutlist(ctx, userId, putlistNumber)
	if err != nil {
		if errors.Is(err, status.Error(codes.NotFound, "putlist not found")) {
			logger.Error("putlist not found")
			return "", fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
		}
		logger.Error("failed to delete putlist")
		return "", fmt.Errorf("%s: %w", op, err)
	}
	logger.Info("putlist deleted")

	return message, nil
}
