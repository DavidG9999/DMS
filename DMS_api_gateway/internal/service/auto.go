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


type AutoCreator interface {
	CreateAuto(ctx context.Context, auto entity.Auto) (int64, error)
}

type AutoProvider interface {
	GetAutos(ctx context.Context) ([]entity.Auto, error)
}

type AutoEditor interface {
	UpdateAuto(ctx context.Context, autoId int64, updateData entity.UpdateAutoInput) (string, error)
	DeleteAuto(ctx context.Context, autoId int64) (string, error)
}

func (p *PutlistService) CreateAuto(ctx context.Context, auto entity.Auto) (int64, error) {
	const op = "service.CreateAuto"

	logger := p.logger.With(slog.String("op", op))
	logger.Info("creating auto")

	autoId, err := p.putlistClient.CreateAuto(ctx, auto.Brand, auto.Model, auto.StateNumber)
	if err != nil {
		if errors.Is(err, status.Error(codes.AlreadyExists, "auto already exists")) {
			logger.Error("auto already exists")
			return 0, fmt.Errorf("%s: %w", op, ErrAutoExists)
		}
		logger.Error("failed to create auto")
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	logger.Info("auto created")

	return autoId, nil
}

func (p *PutlistService) GetAutos(ctx context.Context) ([]entity.Auto, error) {
	const op = "service.GetAutos"

	logger := p.logger.With(slog.String("op", op))
	logger.Info("getting auto")

	autos, err := p.putlistClient.GetAutos(ctx)
	if err != nil {
		logger.Error("failed to get autos")
		return []entity.Auto{}, fmt.Errorf("%s: %w", op, err)
	}
	logger.Info("autos received")

	return autos, nil
}

func (p *PutlistService) UpdateAuto(ctx context.Context, autoId int64, updateData entity.UpdateAutoInput) (string, error) {
	const op = "service.UpdateAuto"

	err := updateData.Validate()
	if err != nil {
		return "", err
	}

	logger := p.logger.With(slog.String("op", op))
	logger.Info("updating auto")

	message, err := p.putlistClient.UpdateAuto(ctx, autoId, updateData.Brand, updateData.Model, updateData.StateNumber)
	if err != nil {
		if errors.Is(err, status.Error(codes.NotFound, "auto not found")) {
			logger.Error("auto not found")
			return "", fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
		}
		logger.Error("failed to update auto")
		return "", fmt.Errorf("%s: %w", op, err)
	}
	logger.Info("auto updated")

	return message, nil
}

func (p *PutlistService) DeleteAuto(ctx context.Context, autoId int64) (string, error) {
	const op = "service.DeleteAuto"

	logger := p.logger.With(slog.String("op", op))
	logger.Info("deleting auto")

	message, err := p.putlistClient.DeleteAuto(ctx, autoId)
	if err != nil {
		if errors.Is(err, status.Error(codes.NotFound, "auto not found")) {
			logger.Error("auto not found")
			return "", fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
		}
		logger.Error("failed to delete auto")
		return "", fmt.Errorf("%s: %w", op, err)
	}
	logger.Info("auto deleted")

	return message, nil
}
