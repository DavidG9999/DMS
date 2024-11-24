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

type DriverService struct {
	logger *slog.Logger
	repo   postgres.DriverRepository
}

func NewDriverService(logger *slog.Logger,repo postgres.DriverRepository) *DriverService {
	return &DriverService{
		logger: logger,
		repo: repo,
	}
}

type DriverCreator interface {
	CreateDriver(ctx context.Context, driver entity.Driver) (driverId int64, err error)
}

type DriverProvider interface {
	GetDrivers(ctx context.Context) ([]entity.Driver, error)
}

type DriverEditor interface {
	UpdateDriver(ctx context.Context, driverId int64, updateData entity.UpdateDriverInput) error
	DeleteDriver(ctx context.Context, driverId int64) error
}

func (s *DriverService) CreateDriver(ctx context.Context, driver entity.Driver) (driverId int64, err error) {
	const op = "putlist_service.CreateDriver"

	logger := s.logger.With(
		slog.String("op", op),
	)
	logger.Info("creating driver")

	driverId, err = s.repo.CreateDriver(ctx, driver)
	if err != nil {
		if errors.Is(err, repository.ErrDriverExists) {
			logger.Error("driver already exists")
			return 0, fmt.Errorf("%s: %w", op, ErrDriverExists)
		}
		logger.Error("failed to create driver")
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	logger.Info("driver created")
	return driverId, nil
}

func (s *DriverService) GetDrivers(ctx context.Context) ([]entity.Driver, error) {
	const op = "putlist_service.GetDrivers"

	logger := s.logger.With(
		slog.String("op", op),
	)
	logger.Info("getting drivers")

	drivers, err := s.repo.GetDrivers(ctx)
	if err != nil {
		if errors.Is(err, repository.ErrDriverNotFound) {
			logger.Error("driver not found")
			return []entity.Driver{}, fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
		}
		logger.Error("failed to get drivers")
		return []entity.Driver{}, fmt.Errorf("%s: %w", op, err)
	}
	logger.Info("drivers received")
	return drivers, nil
}

func (s *DriverService) UpdateDriver(ctx context.Context, driverId int64, updateData entity.UpdateDriverInput) error {
	const op = "putlist_service.UpdateDriver"

	logger := s.logger.With(
		slog.String("op", op),
	)
	logger.Info("updating driver")

	err := s.repo.UpdateDriver(ctx, driverId, updateData)
	if err != nil {
		if errors.Is(err, repository.ErrDriverNotFound) {
			logger.Error("driver not found")
			return fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
		}
		logger.Error("failed to update driver")
		return fmt.Errorf("%s: %w", op, err)
	}
	logger.Info("driver updated")
	return err
}

func (s *DriverService) DeleteDriver(ctx context.Context, driverId int64) error {
	const op = "putlist_service.DeleteDriver"

	logger := s.logger.With(
		slog.String("op", op),
	)
	logger.Info("deleting driver")

	err := s.repo.DeleteDriver(ctx, driverId)
	if err != nil {
		if errors.Is(err, repository.ErrDriverNotFound) {
			logger.Error("driver not found")
			return fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
		}
		logger.Error("failed to delete driver")
		return fmt.Errorf("%s: %w", op, err)
	}
	logger.Info("driver deleted")
	return err
}
