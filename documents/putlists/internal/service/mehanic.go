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

type MehanicService struct {
	logger *slog.Logger
	repo   postgres.MehanicRepository
}

func NewMechanicService(logger *slog.Logger,repo postgres.MehanicRepository) *MehanicService {
	return &MehanicService{
		logger: logger,
		repo: repo,
	}
}

type MehanicCreator interface {
	CreateMehanic(ctx context.Context, mehanic entity.Mehanic) (mehanicId int64, err error)
}

type MehanicProvider interface {
	GetMehanics(ctx context.Context) ([]entity.Mehanic, error)
}

type MehanicEditor interface {
	UpdateMehanic(ctx context.Context, mehanicId int64, updateData entity.UpdateMehanicInput) error
	DeleteMehanic(ctx context.Context, mehanicId int64) error
}

func (s *MehanicService) CreateMehanic(ctx context.Context, mehanic entity.Mehanic) (mehanicId int64, err error) {
	const op = "putlist_service.CreateMehanic"

	logger := s.logger.With(
		slog.String("op", op),
	)
	logger.Info("creating mehanic")

	mehanicId, err = s.repo.CreateMehanic(ctx, mehanic)
	if err != nil {
		logger.Error("failed to create mehanic")
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	logger.Info("mehanic created")
	return mehanicId, nil
}

func (s *MehanicService) GetMehanics(ctx context.Context) ([]entity.Mehanic, error) {
	const op = "putlist_service.GetMehanics"

	logger := s.logger.With(
		slog.String("op", op),
	)
	logger.Info("getting mehanics")

	mehanics, err := s.repo.GetMehanics(ctx)
	if err != nil {
		if errors.Is(err, repository.ErrMehanicNotFound) {
			logger.Error("mehanic not found")
			return []entity.Mehanic{}, fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
		}
		logger.Error("failed to get mehanics")
		return []entity.Mehanic{}, fmt.Errorf("%s: %w", op, err)
	}
	logger.Info("mehanics received")
	return mehanics, nil
}

func (s *MehanicService) UpdateMehanic(ctx context.Context, mehanicId int64, updateData entity.UpdateMehanicInput) error {
	const op = "putlist_service.UpdateMehanic"

	logger := s.logger.With(
		slog.String("op", op),
	)
	logger.Info("updating mehanics")

	err := s.repo.UpdateMehanic(ctx, mehanicId, updateData)
	if err != nil {
		if errors.Is(err, repository.ErrMehanicNotFound) {
			logger.Error("mehanic not found")
			return fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
		}
		logger.Error("failed to update mehanics")
		return fmt.Errorf("%s: %w", op, err)
	}
	logger.Info("mehanics updated")
	return err
}

func (s *MehanicService) DeleteMehanic(ctx context.Context, mehanicId int64) error {
	const op = "putlist_service.DeleteMehanic"

	logger := s.logger.With(
		slog.String("op", op),
	)
	logger.Info("deleting mehanic")

	err := s.repo.DeleteMehanic(ctx, mehanicId)
	if err != nil {
		if errors.Is(err, repository.ErrMehanicNotFound) {
			logger.Error("mehanic not found")
			return fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
		}
		logger.Error("failed to deletemehanicr")
		return fmt.Errorf("%s: %w", op, err)
	}
	logger.Info("mehanic deleted")
	return err
}
