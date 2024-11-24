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

type DispetcherCreator interface {
	CreateDispetcher(ctx context.Context, dispetcher entity.Dispetcher) (int64, error)
}

type DispetcherProvider interface {
	GetDispetchers(ctx context.Context) ([]entity.Dispetcher, error)
}

type DispetcherEditor interface {
	UpdateDispetcher(ctx context.Context, dispetcherId int64, updateData entity.UpdateDispetcherInput) (string, error)
	DeleteDispetcher(ctx context.Context, dispetcherId int64) (string, error)
}

func (p *PutlistService) CreateDispetcher(ctx context.Context, dispetcher entity.Dispetcher) (int64, error) {
	const op = "service.CreateDispetcher"

	logger := p.logger.With(slog.String("op", op))
	logger.Info("creating dispetcher")

	dispetcherId, err := p.putlistClient.CreateDispetcher(ctx, dispetcher.FullName)
	if err != nil {
		logger.Error("failed to create dispetcher")
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	logger.Info("dispetcher created")

	return dispetcherId, nil
}

func (p *PutlistService) GetDispetchers(ctx context.Context) ([]entity.Dispetcher, error) {
	const op = "service.GetDispetchers"

	logger := p.logger.With(slog.String("op", op))
	logger.Info("getting dispetchers")

	dispetchers, err := p.putlistClient.GetDispetchers(ctx)
	if err != nil {
		logger.Error("failed to get dispetchers")
		return []entity.Dispetcher{}, fmt.Errorf("%s: %w", op, err)
	}
	logger.Info("dispetchers received")

	return dispetchers, nil
}

func (p *PutlistService) UpdateDispetcher(ctx context.Context, dispetcherId int64, updateData entity.UpdateDispetcherInput) (string, error) {
	const op = "service.UpdateDispetcher"

	logger := p.logger.With(slog.String("op", op))
	logger.Info("updating dispetcher")

	err := updateData.Validate()
	if err != nil {
		return "", err
	}

	message, err := p.putlistClient.UpdateDispetcher(ctx, dispetcherId, updateData.FullName)
	if err != nil {
		if errors.Is(err, status.Error(codes.NotFound, "dispetcher not found")) {
			logger.Error("dispetcher not found")
			return "", fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
		}
		logger.Error("failed to update dispetcher")
		return "", fmt.Errorf("%s: %w", op, err)
	}
	logger.Info("dispetcher updated")

	return message, nil
}

func (p *PutlistService) DeleteDispetcher(ctx context.Context, dispetcherId int64) (string, error) {
	const op = "service.DeleteDispetcher"

	logger := p.logger.With(slog.String("op", op))
	logger.Info("deleting dispetcher")

	message, err := p.putlistClient.DeleteDispetcher(ctx, dispetcherId)
	if err != nil {
		if errors.Is(err, status.Error(codes.NotFound, "dispetcher not found")) {
			logger.Error("dispetcher not found")
			return "", fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
		}
		logger.Error("failed to delete dispetcher")
		return "", fmt.Errorf("%s: %w", op, err)
	}
	logger.Info("dispetcher deleted")

	return message, nil
}
