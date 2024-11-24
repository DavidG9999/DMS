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

type DriverCreator interface {
	CreateDriver(ctx context.Context, driver entity.Driver) (int64, error)
}

type DriverProvider interface {
	GetDrivers(ctx context.Context) ([]entity.Driver, error)
}

type DriverEditor interface {
	UpdateDriver(ctx context.Context, driverId int64, updateData entity.UpdateDriverInput) (string, error)
	DeleteDriver(ctx context.Context, driverId int64) (string, error)
}

func (p *PutlistService) CreateDriver(ctx context.Context, driver entity.Driver) (int64, error) {
	const op = "service.CreateDriver"

	logger := p.logger.With(slog.String("op", op))
	logger.Info("creating driver")

	driverId, err := p.putlistClient.CreateDriver(ctx, driver.FullName, driver.License, driver.Class)
	if err != nil {
		if errors.Is(err, status.Error(codes.AlreadyExists, "driver already exists")) {
			logger.Error("driver already exists")
			return 0, fmt.Errorf("%s: %w", op, ErrDriverExists)
		}
		logger.Error("failed to create driver")
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	logger.Info("driver created")

	return driverId, nil
}

func (p *PutlistService) GetDrivers(ctx context.Context) ([]entity.Driver, error) {
	const op = "service.GetDrivers"

	logger := p.logger.With(slog.String("op", op))
	logger.Info("getting drivers")

	drivers, err := p.putlistClient.GetDrivers(ctx)
	if err != nil {
		logger.Error("failed to get drivers")
		return []entity.Driver{}, fmt.Errorf("%s: %w", op, err)
	}
	logger.Info("drivers received")

	return drivers, nil
}

func (p *PutlistService) UpdateDriver(ctx context.Context, driverId int64, updateData entity.UpdateDriverInput) (string, error) {
	const op = "service.UpdateDriver"

	logger := p.logger.With(slog.String("op", op))
	logger.Info("updating driver")

	err := updateData.Validate()
	if err != nil {
		return "", err
	}

	message, err := p.putlistClient.UpdateDriver(ctx, driverId, updateData.FullName, updateData.License, updateData.Class)
	if err != nil {
		if errors.Is(err, status.Error(codes.NotFound, "driver not found")) {
			logger.Error("driver not found")
			return "", fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
		}
		logger.Error("failed to update driver")
		return "", fmt.Errorf("%s: %w", op, err)
	}
	logger.Info("driver updated")

	return message, nil
}

func (p *PutlistService) DeleteDriver(ctx context.Context, driverId int64) (string, error) {
	const op = "service.DeleteDriver"

	logger := p.logger.With(slog.String("op", op))
	logger.Info("deleting driver")

	message, err := p.putlistClient.DeleteDriver(ctx, driverId)
	if err != nil {
		if errors.Is(err, status.Error(codes.NotFound, "driver not found")) {
			logger.Error("driver not found")
			return "", fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
		}
		logger.Error("failed to delete driver")
		return "", fmt.Errorf("%s: %w", op, err)
	}
	logger.Info("driver deleted")

	return message, nil
}
