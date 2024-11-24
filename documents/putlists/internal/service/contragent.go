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

type ContragentService struct {
	logger *slog.Logger
	repo   postgres.ContragentRepository
}

func NewContragentService(logger *slog.Logger,repo postgres.ContragentRepository) *ContragentService {
	return &ContragentService{
		logger: logger,
		repo: repo,
	}
}

type ContragentCreator interface {
	CreateContragent(ctx context.Context, contragent entity.Contragent) (contragentId int64, err error)
}

type ContragentProvider interface {
	GetContragents(ctx context.Context) ([]entity.Contragent, error)
}

type ContragentEditor interface {
	UpdateContragent(ctx context.Context, contragentId int64, updateData entity.UpdateContragentInput) error
	DeleteContragent(ctx context.Context, contragentId int64) error
}

func (s *ContragentService) CreateContragent(ctx context.Context, contragent entity.Contragent) (contragentId int64, err error) {
	const op = "putlist_service.CreateContragent"

	logger := s.logger.With(
		slog.String("op", op),
	)
	logger.Info("creating contragent")

	contragentId, err = s.repo.CreateContragent(ctx, contragent)
	if err != nil {
		if errors.Is(err, repository.ErrContragentExists) {
			logger.Error("contragent already exists")
			return 0, fmt.Errorf("%s: %w", op, ErrContragentExists)
		}
		logger.Error("failed to create contragent")
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	logger.Info("contragent created")
	return contragentId, nil
}

func (s *ContragentService) GetContragents(ctx context.Context) ([]entity.Contragent, error) {
	const op = "putlist_service.GetContragents"

	logger := s.logger.With(
		slog.String("op", op),
	)
	logger.Info("getting contragents")

	contragents, err := s.repo.GetContragents(ctx)
	if err != nil {
		if errors.Is(err, repository.ErrContragentNotFound) {
			logger.Error("contragent not found")
			return []entity.Contragent{}, fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
		}
		logger.Error("failed to get contragents")
		return []entity.Contragent{}, fmt.Errorf("%s: %w", op, err)
	}
	logger.Info("contragents received")
	return contragents, nil
}

func (s *ContragentService) UpdateContragent(ctx context.Context, contragentId int64, updateData entity.UpdateContragentInput) error {
	const op = "putlist_service.UpdateContragent"

	logger := s.logger.With(
		slog.String("op", op),
	)
	logger.Info("updating contragent")

	err := s.repo.UpdateContragent(ctx, contragentId, updateData)
	if err != nil {
		if errors.Is(err, repository.ErrContragentNotFound) {
			logger.Error("contragent not found")
			return fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
		}
		logger.Error("failed to update contragent")
		return fmt.Errorf("%s: %w", op, err)
	}
	logger.Info("contragent updated")
	return err
}

func (s *ContragentService) DeleteContragent(ctx context.Context, contragentId int64) error {
	const op = "putlist_service.DeleteContragent"

	logger := s.logger.With(
		slog.String("op", op),
	)
	logger.Info("deleting contragent")

	err := s.repo.DeleteContragent(ctx, contragentId)
	if err != nil {
		if errors.Is(err, repository.ErrContragentNotFound) {
			logger.Error("contragent not found")
			return fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
		}
		logger.Error("failed to delete contragent")
		return fmt.Errorf("%s: %w", op, err)
	}
	logger.Info("contragent deleted")
	return err

}
