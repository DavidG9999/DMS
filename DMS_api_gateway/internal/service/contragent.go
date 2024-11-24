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

type ContragentCreator interface {
	CreateContragent(ctx context.Context, contragent entity.Contragent) (int64, error)
}

type ContragentProvider interface {
	GetContragents(ctx context.Context) ([]entity.Contragent, error)
}

type ContragentEditor interface {
	UpdateContragent(ctx context.Context, contragentId int64, updateData entity.UpdateContragentInput) (string, error)
	DeleteContragent(ctx context.Context, contragentId int64) (string, error)
}

func (p *PutlistService) CreateContragent(ctx context.Context, contragent entity.Contragent) (int64, error) {
	const op = "service.CreateContragent"

	logger := p.logger.With(slog.String("op", op))
	logger.Info("creating contragent")

	contragentId, err := p.putlistClient.CreateContragent(ctx, contragent.Name, contragent.Address, contragent.InnKpp)
	if err != nil {
		if errors.Is(err, status.Error(codes.AlreadyExists, "contragent already exists")) {
			logger.Error("contragent already exists")
			return 0, fmt.Errorf("%s: %w", op, ErrContragentExists)
		}
		logger.Error("failed to create contragent")
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	logger.Info("contragent created")

	return contragentId, nil
}

func (p *PutlistService) GetContragents(ctx context.Context) ([]entity.Contragent, error) {
	const op = "service.GetContragents"

	logger := p.logger.With(slog.String("op", op))
	logger.Info("getting contragents")

	contragents, err := p.putlistClient.GetContragents(ctx)
	if err != nil {
		logger.Error("failed to get contragents")
		return []entity.Contragent{}, fmt.Errorf("%s: %w", op, err)
	}
	logger.Info("contragents received")

	return contragents, nil
}

func (p *PutlistService) UpdateContragent(ctx context.Context, contragentId int64, updateData entity.UpdateContragentInput) (string, error) {
	const op = "service.UpdateContragent"

	logger := p.logger.With(slog.String("op", op))
	logger.Info("updating contragent")

	err := updateData.Validate()
	if err != nil {
		return "", err
	}

	message, err := p.putlistClient.UpdateContragent(ctx, contragentId, updateData.Name, updateData.Address, updateData.InnKpp)
	if err != nil {
		if errors.Is(err, status.Error(codes.NotFound, "contragent not found")) {
			logger.Error("contragent not found")
			return "", fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
		}
		logger.Error("failed to update contragent")
		return "", fmt.Errorf("%s: %w", op, err)
	}
	logger.Info("contragent updated")

	return message, nil
}

func (p *PutlistService) DeleteContragent(ctx context.Context, contragentId int64) (string, error) {
	const op = "service.DeleteContragent"

	logger := p.logger.With(slog.String("op", op))
	logger.Info("deleting contragent")

	message, err := p.putlistClient.DeleteContragent(ctx, contragentId)
	if err != nil {
		if errors.Is(err, status.Error(codes.NotFound, "contragent not found")) {
			logger.Error("contragent not found")
			return "", fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
		}
		logger.Error("failed to delete contragent")
		return "", fmt.Errorf("%s: %w", op, err)
	}
	logger.Info("contragent deleted")

	return message, nil
}
