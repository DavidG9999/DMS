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

type MehanicCreator interface {
	CreateMehanic(ctx context.Context, mehanic entity.Mehanic) (int64, error)
}

type MehanicProvider interface {
	GetMehanics(ctx context.Context) ([]entity.Mehanic, error)
}

type MehanicEditor interface {
	UpdateMehanic(ctx context.Context, mehanicId int64, updateData entity.UpdateMehanicInput) (string, error)
	DeleteMehanic(ctx context.Context, mehanicId int64) (string, error)
}

func (p *PutlistService) CreateMehanic(ctx context.Context, mehanic entity.Mehanic) (int64, error) {
	const op = "service.CreateMehanic"

	logger := p.logger.With(slog.String("op", op))
	logger.Info("creating mehanic")

	mehanicId, err := p.putlistClient.CreateMehanic(ctx, mehanic.FullName)
	if err != nil {
		logger.Error("failed to create mehanic")
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	logger.Info("mehanic created")

	return mehanicId, nil
}

func (p *PutlistService) GetMehanics(ctx context.Context) ([]entity.Mehanic, error) {
	const op = "service.GetMehanics"

	logger := p.logger.With(slog.String("op", op))
	logger.Info("getting mehanics")

	mehanics, err := p.putlistClient.GetMehanics(ctx)
	if err != nil {
		logger.Error("failed to get mehanics")
		return []entity.Mehanic{}, fmt.Errorf("%s: %w", op, err)
	}
	logger.Info("mehanics received")

	return mehanics, nil
}

func (p *PutlistService) UpdateMehanic(ctx context.Context, mehanicId int64, updateData entity.UpdateMehanicInput) (string, error) {
	const op = "service.UpdateMehanic"

	logger := p.logger.With(slog.String("op", op))
	logger.Info("updating mehanic")

	err := updateData.Validate()
	if err != nil {
		return "", err
	}

	message, err := p.putlistClient.UpdateMehanic(ctx, mehanicId, updateData.FullName)
	if err != nil {
		if errors.Is(err, status.Error(codes.NotFound, "mehanic not found")) {
			logger.Error("mehanic not found")
			return "", fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
		}
		logger.Error("failed to update mehanic")
		return "", fmt.Errorf("%s: %w", op, err)
	}
	logger.Info("mehanic updated")

	return message, nil
}

func (p *PutlistService) DeleteMehanic(ctx context.Context, mehanicId int64) (string, error) {
	const op = "service.DeleteMehanic"

	logger := p.logger.With(slog.String("op", op))
	logger.Info("deleting mehanic")

	message, err := p.putlistClient.DeleteMehanic(ctx, mehanicId)
	if err != nil {
		if errors.Is(err, status.Error(codes.NotFound, "mehanic not found")) {
			logger.Error("mehanic not found")
			return "", fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
		}
		logger.Error("failed to delete mehanic")
		return "", fmt.Errorf("%s: %w", op, err)
	}
	logger.Info("mehanic deleted")

	return message, nil
}
