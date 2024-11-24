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

type PutlistService struct {
	logger *slog.Logger
	repo   postgres.PutlistRepository
}

func NewPutlistService(logger *slog.Logger, repo postgres.PutlistRepository) *PutlistService {
	return &PutlistService{
		logger: logger,
		repo:   repo,
	}
}

type PutlistCreator interface {
	CreatePutlist(ctx context.Context, putlist entity.PutlistHeader) (putlistId int64, err error)
}

type PutlistProvider interface {
	GetPutlists(ctx context.Context, userId int64) ([]entity.PutlistHeader, error)
	GetPutlistByNumber(ctx context.Context, userId int64, putlistNumber int64) (entity.PutlistHeader, error)
}

type PutlistEditor interface {
	UpdatePutlist(ctx context.Context, userId int64, putlistNumber int64, updateData entity.UpdatePutlistHeaderInput) error
	DeletePutlist(ctx context.Context, userId int64, putlistNumber int64) error
}

func (s *PutlistService) CreatePutlist(ctx context.Context, putlist entity.PutlistHeader) (putlistId int64, err error) {
	const op = "putlist_service.CreatePutlist"

	logger := s.logger.With(
		slog.String("op", op),
	)
	logger.Info("creating putlist")

	putlistId, err = s.repo.CreatePutlist(ctx, putlist)
	if err != nil {
		if errors.Is(err, repository.ErrInvalidDateTimeFormat) {
			logger.Error("invalid datetime format")
			return 0, fmt.Errorf("%s: %w", op, ErrInvalidDateTimeFormat)
		}
		if errors.Is(err, repository.ErrPutlistExists) {
			logger.Error("putlist already exists")
			return 0, fmt.Errorf("%s: %w", op, ErrPutlistExists)
		}
		logger.Error("failed to create putlist")
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	logger.Info("putlist created")
	return putlistId, nil
}

func (s *PutlistService) GetPutlists(ctx context.Context, userId int64) ([]entity.PutlistHeader, error) {
	const op = "putlist_service.GetPutlists"

	logger := s.logger.With(
		slog.String("op", op),
	)
	logger.Info("getting putlists")

	putlists, err := s.repo.GetPutlists(ctx, userId)
	if err != nil {
		if errors.Is(err, repository.ErrPutlistNotFound) {
			logger.Error("putlist not found")
			return []entity.PutlistHeader{}, fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
		}
		logger.Error("failed to get putlists")
		return []entity.PutlistHeader{}, fmt.Errorf("%s: %w", op, err)
	}
	logger.Info("putlists received")
	return putlists, nil
}

func (s *PutlistService) GetPutlistByNumber(ctx context.Context, userId int64, putlistNumber int64) (entity.PutlistHeader, error) {
	const op = "putlist_service.GetPutlistByNumber"

	logger := s.logger.With(
		slog.String("op", op),
	)
	logger.Info("getting putlist by number")

	putlist, err := s.repo.GetPutlistByNumber(ctx, userId, putlistNumber)
	if err != nil {
		if errors.Is(err, repository.ErrPutlistNotFound) {
			logger.Error("putlist not found")
			return entity.PutlistHeader{}, fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
		}
		logger.Error("failed to get putlist")
		return entity.PutlistHeader{}, fmt.Errorf("%s: %w", op, err)
	}
	logger.Info("putlist received")
	return putlist, nil
}

func (s *PutlistService) UpdatePutlist(ctx context.Context, userId int64, putlistNumber int64, updateData entity.UpdatePutlistHeaderInput) error {
	const op = "putlist_service.UpdatePutlist"

	logger := s.logger.With(
		slog.String("op", op),
	)
	logger.Info("updating putlist")

	err := s.repo.UpdatePutlist(ctx, userId, putlistNumber, updateData)
	if err != nil {
		if errors.Is(err, repository.ErrInvalidDateTimeFormat) {
			logger.Error("invalid datetime format")
			return fmt.Errorf("%s: %w", op, ErrInvalidDateTimeFormat)
		}
		if errors.Is(err, repository.ErrPutlistNotFound) {
			logger.Error("putlist not found")
			return fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
		}
		logger.Error("failed to update putlist")
		return fmt.Errorf("%s: %w", op, err)
	}
	logger.Info("putlist updated")
	return err
}

func (s *PutlistService) DeletePutlist(ctx context.Context, userId int64, putlistNumber int64) error {
	const op = "putlist_service.DeletePutlist"

	logger := s.logger.With(
		slog.String("op", op),
	)
	logger.Info("deleting putlist")

	err := s.repo.DeletePutlist(ctx, userId, putlistNumber)
	if err != nil {
		if errors.Is(err, repository.ErrPutlistNotFound) {
			logger.Error("putlist not found")
			return fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
		}
		logger.Error("failed to delete putlist")
		return fmt.Errorf("%s: %w", op, err)
	}
	logger.Info("putlist deleted")
	return err
}
