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

type DispetcherService struct {
	logger *slog.Logger
	repo   postgres.DispetcherRepository
}

func NewDispetcherService(logger *slog.Logger,repo postgres.DispetcherRepository) *DispetcherService {
	return &DispetcherService{
		logger: logger,
		repo: repo,
	}
}

type DispetcherCreator interface {
	CreateDispetcher(ctx context.Context, dispetcher entity.Dispetcher) (dispetcherId int64, err error)
}

type DispetcherProvider interface {
	GetDispetchers(ctx context.Context) ([]entity.Dispetcher, error)
}

type DispetcherEditor interface {
	UpdateDispetcher(ctx context.Context, dispetcherId int64, updateData entity.UpdateDispetcherInput) error
	DeleteDispetcher(ctx context.Context, dispetcherId int64) error
}

func (s *DispetcherService) CreateDispetcher(ctx context.Context, dispetcher entity.Dispetcher) (dispetcherId int64, err error) {
	const op = "putlist_service.CreateDispetcher"

	logger := s.logger.With(
		slog.String("op", op),
	)
	logger.Info("creating dispetcher")

	dispetcherId, err = s.repo.CreateDispetcher(ctx, dispetcher)
	if err != nil {
		logger.Error("failed to create dispetcher")
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	logger.Info("dispetcher created")
	return dispetcherId, nil
}

func (s *DispetcherService) GetDispetchers(ctx context.Context) ([]entity.Dispetcher, error) {
	const op = "putlist_service.GetDispetchers"

	logger := s.logger.With(
		slog.String("op", op),
	)
	logger.Info("getting dispetchers")

	dispetchers, err := s.repo.GetDispetchers(ctx)
	if err != nil {
		if errors.Is(err, repository.ErrDispethcerNotFound) {
			logger.Error("dispetcher not found")
			return []entity.Dispetcher{}, fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
		}
		logger.Error("failed to get dispetchers")
		return []entity.Dispetcher{}, fmt.Errorf("%s: %w", op, err)
	}
	logger.Info("dispetchers received")
	return dispetchers, nil
}

func (s *DispetcherService) UpdateDispetcher(ctx context.Context, dispetcherId int64, updateData entity.UpdateDispetcherInput) error {
	const op = "putlist_service.UpdateDispetcher"

	logger := s.logger.With(
		slog.String("op", op),
	)
	logger.Info("updating dispetcher")

	err := s.repo.UpdateDispetcher(ctx, dispetcherId, updateData)
	if err != nil {
		if errors.Is(err, repository.ErrDispethcerNotFound) {
			logger.Error("dispethcer not found")
			return fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
		}
		logger.Error("failed to update dispetcher")
		return fmt.Errorf("%s: %w", op, err)
	}
	logger.Info("dispetcher updated")
	return err
}

func (s *DispetcherService) DeleteDispetcher(ctx context.Context, dispetcherId int64) error {
	const op = "putlist_service.DeleteDispetcher"

	logger := s.logger.With(
		slog.String("op", op),
	)
	logger.Info("deleting dispetcher")

	err := s.repo.DeleteDispetcher(ctx, dispetcherId)
	if err != nil {
		if errors.Is(err, repository.ErrDispethcerNotFound) {
			logger.Error("dispetcher not found")
			return fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
		}
		logger.Error("failed to delete dispetcher")
		return fmt.Errorf("%s: %w", op, err)
	}
	logger.Info("dispetcher deleted")
	return err
}
