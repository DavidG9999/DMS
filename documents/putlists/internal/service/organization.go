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

type OrganizationService struct {
	logger *slog.Logger
	repo   postgres.OrganizationRepository
}

func NewOrganizationService(logger *slog.Logger, repo postgres.OrganizationRepository) *OrganizationService {
	return &OrganizationService{
		logger: logger,
		repo:   repo,
	}
}

type OrganizationCreator interface {
	CreateOrganization(ctx context.Context, organization entity.Organization) (organizationId int64, err error)
}

type OrganizationProvider interface {
	GetOrganizations(ctx context.Context) ([]entity.Organization, error)
}

type OrganizationEditor interface {
	UpdateOrganization(ctx context.Context, organizationId int64, updateData entity.UpdateOrganizationInput) error
	DeleteOrganization(ctx context.Context, organizationId int64) error
}

func (s *OrganizationService) CreateOrganization(ctx context.Context, organization entity.Organization) (organizationId int64, err error) {
	const op = "putlist_service.CreateOrganization"

	logger := s.logger.With(
		slog.String("op", op),
	)
	logger.Info("creating organization")

	organizationId, err = s.repo.CreateOrganization(ctx, organization)
	if err != nil {
		if errors.Is(err, repository.ErrOrganizationExists) {
			logger.Error("organization already exists")
			return 0, fmt.Errorf("%s: %w", op, ErrOrganizationExists)
		}
		logger.Error("failed to create organization")
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	logger.Info("organization created")
	return organizationId, nil
}

func (s *OrganizationService) GetOrganizations(ctx context.Context) ([]entity.Organization, error) {
	const op = "putlist_service.GetOrganizations"

	logger := s.logger.With(
		slog.String("op", op),
	)
	logger.Info("getting organizations")

	organizations, err := s.repo.GetOrganizations(ctx)
	if err != nil {
		if errors.Is(err, repository.ErrOrganizationNotFound) {
			logger.Error("organization not found")
			return []entity.Organization{}, fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
		}
		logger.Error("failed to get organizations")
		return []entity.Organization{}, fmt.Errorf("%s: %w", op, err)
	}
	logger.Info("organizations received")
	return organizations, nil
}

func (s *OrganizationService) UpdateOrganization(ctx context.Context, organizationId int64, updateData entity.UpdateOrganizationInput) error {
	const op = "putlist_service.UpdateOrganization"

	logger := s.logger.With(
		slog.String("op", op),
	)
	logger.Info("updating organization")

	err := s.repo.UpdateOrganization(ctx, organizationId, updateData)
	if err != nil {
		if errors.Is(err, repository.ErrOrganizationNotFound) {
			logger.Error("organization not found")
			return fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
		}
		logger.Error("failed to update organization")
		return fmt.Errorf("%s: %w", op, err)
	}
	logger.Info("organization updated")
	return err
}

func (s *OrganizationService) DeleteOrganization(ctx context.Context, organizationId int64) error {
	const op = "putlist_service.DeleteOrganization"

	logger := s.logger.With(
		slog.String("op", op),
	)
	logger.Info("deleting organization")

	err := s.repo.DeleteOrganization(ctx, organizationId)
	if err != nil {
		if errors.Is(err, repository.ErrOrganizationNotFound) {
			logger.Error("organization not found")
			return fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
		}
		logger.Error("failed to delete organization")
		return fmt.Errorf("%s: %w", op, err)
	}
	logger.Info("organization deleted")
	return err
}
