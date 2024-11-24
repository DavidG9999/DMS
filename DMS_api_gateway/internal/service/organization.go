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

type OrganizationCreator interface {
	CreateOrganization(ctx context.Context, organization entity.Organization) (int64, error)
}

type OrganizationProvider interface {
	GetOrganizations(ctx context.Context) ([]entity.Organization, error)
}

type OrganizationEditor interface {
	UpdateOrganization(ctx context.Context, organizationId int64, updateData entity.UpdateOrganizationInput) (string, error)
	DeleteOrganization(ctx context.Context, organizationId int64) (string, error)
}

func (p *PutlistService) CreateOrganization(ctx context.Context, organization entity.Organization) (int64, error) {
	const op = "service.CreateOrganization"

	logger := p.logger.With(slog.String("op", op))
	logger.Info("creating organization")

	organizationId, err := p.putlistClient.CreateOrganization(ctx, organization.Name, organization.Address, organization.Chief, organization.FinancialChief, organization.InnKpp)
	if err != nil {
		if errors.Is(err, status.Error(codes.AlreadyExists, "organization already exists")) {
			logger.Error("organization already exists")
			return 0, fmt.Errorf("%s: %w", op, ErrOrganizationExists)
		}
		logger.Error("failed to create organization")
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	logger.Info("organization created")

	return organizationId, nil
}

func (p *PutlistService) GetOrganizations(ctx context.Context) ([]entity.Organization, error) {
	const op = "service.GetOrganizations"

	logger := p.logger.With(slog.String("op", op))
	logger.Info("getting organizations")

	organizations, err := p.putlistClient.GetOrganizations(ctx)
	if err != nil {
		logger.Error("failed to get organizations")
		return []entity.Organization{}, fmt.Errorf("%s: %w", op, err)
	}
	logger.Info("organizations received")

	return organizations, nil
}

func (p *PutlistService) UpdateOrganization(ctx context.Context, organizationId int64, updateData entity.UpdateOrganizationInput) (string, error) {
	const op = "service.UpdateOrganization"

	logger := p.logger.With(slog.String("op", op))
	logger.Info("updating organization")

	err := updateData.Validate()
	if err != nil {
		return "", err
	}

	message, err := p.putlistClient.UpdateOrganization(ctx, organizationId, updateData.Name, updateData.Address, updateData.Chief, updateData.FinancialChief, updateData.InnKpp)
	if err != nil {
		if errors.Is(err, status.Error(codes.NotFound, "organization not found")) {
			logger.Error("organization not found")
			return "", fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
		}
		logger.Error("failed to update organization")
		return "", fmt.Errorf("%s: %w", op, err)
	}
	logger.Info("organization updated")

	return message, nil
}

func (p *PutlistService) DeleteOrganization(ctx context.Context, organizationId int64) (string, error) {
	const op = "service.DeleteOrganization"

	logger := p.logger.With(slog.String("op", op))
	logger.Info("deleting organization")

	message, err := p.putlistClient.DeleteOrganization(ctx, organizationId)
	if err != nil {
		if errors.Is(err, status.Error(codes.NotFound, "organization not found")) {
			logger.Error("organization not found")
			return "", fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
		}
		logger.Error("failed to delete organization")
		return "", fmt.Errorf("%s: %w", op, err)
	}
	logger.Info("organization deleted")

	return message, nil
}
