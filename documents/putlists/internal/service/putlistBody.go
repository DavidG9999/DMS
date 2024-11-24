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

type PutlistBodyService struct {
	logger *slog.Logger
	repo   postgres.PutlistBodyRepository
}

func NewPutlistBodyService(logger *slog.Logger, repo postgres.PutlistBodyRepository) *PutlistBodyService {
	return &PutlistBodyService{
		logger: logger,
		repo:   repo,
	}
}

type PutlistBodyCreator interface {
	CreatePutlistBody(ctx context.Context, putlistBody entity.PutlistBody) (putlistBodyId int64, err error)
}

type PutlistBodyProvider interface {
	GetPutlistBodies(ctx context.Context, putlistNumber int64) ([]entity.PutlistBody, error)
}

type PutlistBodyEditor interface {
	UpdatePutlistBody(ctx context.Context, putlistBodyId int64, updateData entity.UpdatePutlistBodyInput) error
	DeletePutlistBody(ctx context.Context, putlistBodyId int64) error
}

func (s *PutlistBodyService) CreatePutlistBody(ctx context.Context, putlistBody entity.PutlistBody) (putlistBodyId int64, err error) {
	const op = "putlist_service.CreatePutlistBody"

	logger := s.logger.With(
		slog.String("op", op),
	)
	logger.Info("creating putlist body")

	putlistBodyId, err = s.repo.CreatePutlistBody(ctx, putlistBody)
	if err != nil {
		if errors.Is(err, repository.ErrInvalidDateTimeFormat) {
			logger.Error("invalid datetime format")
			return 0, fmt.Errorf("%s: %w", op, ErrInvalidDateTimeFormat)
		}
		logger.Error("failed to create putlist body")
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	logger.Info("putlist body created")
	return putlistBodyId, nil
}

func (s *PutlistBodyService) GetPutlistBodies(ctx context.Context, putlistNumber int64) ([]entity.PutlistBody, error) {
	const op = "putlist_service.GetPutlistBodies"

	logger := s.logger.With(
		slog.String("op", op),
	)
	logger.Info("getting putlist bodies")

	putlistBodies, err := s.repo.GetPutlistBodies(ctx, putlistNumber)
	if err != nil {
		logger.Error("failed to get putlist bodies")
		return []entity.PutlistBody{}, fmt.Errorf("%s: %w", op, err)
	}
	logger.Info("putlist bodies received")
	return putlistBodies, nil
}

func (s *PutlistBodyService) UpdatePutlistBody(ctx context.Context, putlistBodyId int64, updateData entity.UpdatePutlistBodyInput) error {
	const op = "putlist_service.UpdatePutlistBody"

	logger := s.logger.With(
		slog.String("op", op),
	)
	logger.Info("updating putlist body")

	err := s.repo.UpdatePutlistBody(ctx, putlistBodyId, updateData)
	if err != nil {
		if errors.Is(err, repository.ErrInvalidDateTimeFormat) {
			logger.Error("invalid datetime format")
			return fmt.Errorf("%s: %w", op, ErrInvalidDateTimeFormat)
		}
		if errors.Is(err, repository.ErrPutlistBodyNotFound) {
			logger.Error("putlist body not found")
			return fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
		}
		logger.Error("failed to update putlist body")
		return fmt.Errorf("%s: %w", op, err)
	}
	logger.Info("putlist body updated")
	return err
}

func (s *PutlistBodyService) DeletePutlistBody(ctx context.Context, putlistBodyId int64) error {
	const op = "putlist_service.DeletePutlistBody"

	logger := s.logger.With(
		slog.String("op", op),
	)
	logger.Info("deleting putlist body")

	err := s.repo.DeletePutlistBody(ctx, putlistBodyId)
	if err != nil {
		if errors.Is(err, repository.ErrPutlistBodyNotFound) {
			logger.Error("putlist body not found")
			return fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
		}
		logger.Error("failed to delete putlist body")
		return fmt.Errorf("%s: %w", op, err)
	}
	logger.Info("putlist body deleted")
	return err
}
