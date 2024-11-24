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

type AutoService struct {
	logger *slog.Logger
	repo   postgres.AutoRepository
}

func NewAutoService(logger *slog.Logger, repo postgres.AutoRepository) *AutoService {
	return &AutoService{
		logger: logger,
		repo:   repo,
	}
}

type AutoCreator interface {
	CreateAuto(ctx context.Context, auto entity.Auto) (autoId int64, err error)
}

type AutoProvider interface {
	GetAutos(ctx context.Context) ([]entity.Auto, error)
}

type AutoEditor interface {
	UpdateAuto(ctx context.Context, autoId int64, updateData entity.UpdateAutoInput) error
	DeleteAuto(ctx context.Context, autoId int64) error
}

func (s *AutoService) CreateAuto(ctx context.Context, auto entity.Auto) (autoId int64, err error) {
	const op = "putlist_service.CreateAuto"

	logger := s.logger.With(
		slog.String("op", op),
	)
	logger.Info("creating auto")

	autoId, err = s.repo.CreateAuto(ctx, auto)
	if err != nil {
		if errors.Is(err, repository.ErrAutoExists) {
			logger.Error("auto already exists")
			return 0, fmt.Errorf("%s: %w", op, ErrAutoExists)
		}
		logger.Error("failed to create auto")
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	logger.Info("auto created")
	return autoId, nil
}

func (s *AutoService) GetAutos(ctx context.Context) ([]entity.Auto, error) {
	const op = "putlist_service.GetAutos"

	logger := s.logger.With(
		slog.String("op", op),
	)
	logger.Info("getting auto")

	autos, err := s.repo.GetAutos(ctx)
	if err != nil {
		if errors.Is(err, repository.ErrAutoNotFound) {
			logger.Error("auto not found")
			return []entity.Auto{}, fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
		}
		logger.Error("failed to get autos")
		return []entity.Auto{}, fmt.Errorf("%s: %w", op, err)
	}
	logger.Info("auto received")
	return autos, nil
}

func (s *AutoService) UpdateAuto(ctx context.Context, autoId int64, updateData entity.UpdateAutoInput) error {
	const op = "putlist_service.UpdateAuto"

	logger := s.logger.With(
		slog.String("op", op),
	)
	logger.Info("updating auto")

	err := s.repo.UpdateAuto(ctx, autoId, updateData)
	if err != nil {
		if errors.Is(err, repository.ErrAutoNotFound) {
			logger.Error("auto not found")
			return fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
		}
		logger.Error("failed to update auto")
		return fmt.Errorf("%s: %w", op, err)
	}
	logger.Info("auto updated")
	return err
}

func (s *AutoService) DeleteAuto(ctx context.Context, autoId int64) error {
	const op = "putlist_service.DeleteAuto"

	logger := s.logger.With(
		slog.String("op", op),
	)
	logger.Info("deleting auto")

	err := s.repo.DeleteAuto(ctx, autoId)
	if err != nil {
		if errors.Is(err, repository.ErrAutoNotFound) {
			logger.Error("auto not found")
			return fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
		}
		logger.Error("failed to delete auto")
		return fmt.Errorf("%s: %w", op, err)
	}
	logger.Info("auto deleted")
	return err
}
