package handler

import (
	"context"
	"errors"

putlistv1 "github.com/DavidG9999/DMS/api/grpc/document_api/putlist/gen/go"
	"github.com/DavidG9999/DMS/documents/putlists/internal/domain/entity"
	"github.com/DavidG9999/DMS/documents/putlists/internal/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type OrganizationHandler struct {
	putlistv1.UnimplementedOrganizationServer
	OrganizationService service.Organization
}

func (h *OrganizationHandler) CreateOrganization(ctx context.Context, req *putlistv1.CreateOrganizationRequest) (*putlistv1.CreateOrganizationResponse, error) {
	if err := validateCreateOrganization(req); err != nil {
		return nil, err
	}

	organization := entity.Organization{
		Name:           req.GetName(),
		Chief:          req.GetChief(),
		FinancialChief: req.GetFinChief(),
		Address:        req.GetAddress(),
		InnKpp:         req.GetInnKpp(),
	}
	organizationId, err := h.OrganizationService.CreateOrganization(ctx, organization)
	if err != nil {
		if errors.Is(err, service.ErrOrganizationExists) {
			return nil, status.Error(codes.AlreadyExists, "organization already exists")
		}
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &putlistv1.CreateOrganizationResponse{
		OrganizationId: organizationId,
	}, nil
}

func (h *OrganizationHandler) GetOrganizations(ctx context.Context, req *putlistv1.GetOrganizationsRequest) (*putlistv1.GetOrganizationsResponse, error) {
	organizations, err := h.OrganizationService.GetOrganizations(ctx)
	if err != nil {
		if errors.Is(err, service.ErrInvalidCredentials) {
			return nil, status.Error(codes.NotFound, "organizations not found")
		}
		return nil, status.Error(codes.Internal, "internal error")
	}

	organizationsList := make([]putlistv1.OrganizationEntity, len(organizations))
	for idx, organization := range organizations {
		organizationsList[idx].OrganizationId = organization.Id
		organizationsList[idx].Name = organization.Name
		organizationsList[idx].Chief = organization.Chief
		organizationsList[idx].FinChief = organization.FinancialChief
		organizationsList[idx].Address = organization.Address
		organizationsList[idx].InnKpp = organization.InnKpp
	}

	organizationsResp := make([]*putlistv1.OrganizationEntity, 0, len(organizations))
	for id := range organizationsList {
		organizationsResp = append(organizationsResp, &organizationsList[id])
	}

	return &putlistv1.GetOrganizationsResponse{
		Organizations: organizationsResp,
	}, nil
}

func (h *OrganizationHandler) UpdateOrganization(ctx context.Context, req *putlistv1.UpdateOrganizationRequest) (*putlistv1.UpdateOrganizationResponse, error) {
	if err := validateUpdateOrganization(req); err != nil {
		return nil, err
	}

	updateData := entity.UpdateOrganizationInput{
		Name:           req.Name,
		Address:        req.Address,
		Chief:          req.Chief,
		FinancialChief: req.FinChief,
		InnKpp:         req.InnKpp,
	}
	err := h.OrganizationService.UpdateOrganization(ctx, req.GetOrganizationId(), updateData)
	if err != nil {
		if errors.Is(err, service.ErrInvalidCredentials) {
			return nil, status.Error(codes.NotFound, "organization not found")
		}
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &putlistv1.UpdateOrganizationResponse{
		Message: "updated organization",
	}, nil
}

func (h *OrganizationHandler) DeleteOrganization(ctx context.Context, req *putlistv1.DeleteOrganizationRequest) (*putlistv1.DeleteOrganizationResponse, error) {
	if err := validateDeleteOrganization(req); err != nil {
		return nil, err
	}

	err := h.OrganizationService.DeleteOrganization(ctx, req.GetOrganizationId())
	if err != nil {
		if errors.Is(err, service.ErrInvalidCredentials) {
			return nil, status.Error(codes.NotFound, "organization not found")
		}
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &putlistv1.DeleteOrganizationResponse{
		Message: "deleted",
	}, nil
}

func validateCreateOrganization(req *putlistv1.CreateOrganizationRequest) error {
	if req.GetName() == "" {
		return status.Error(codes.InvalidArgument, "organization name is required")
	}
	if req.GetAddress() == "" {
		return status.Error(codes.InvalidArgument, "address is required")
	}
	if req.GetChief() == "" {
		return status.Error(codes.InvalidArgument, "chief is required")
	}
	if req.GetFinChief() == "" {
		return status.Error(codes.InvalidArgument, "financial chief is required")
	}
	if req.GetInnKpp() == "" {
		return status.Error(codes.InvalidArgument, "inn/kpp is required")
	}
	if len(req.GetInnKpp()) != 20 {
		return status.Error(codes.InvalidArgument, "invalid field format: inn/kpp")
	}
	return nil
}

func validateUpdateOrganization(req *putlistv1.UpdateOrganizationRequest) error {
	if req.GetOrganizationId() == 0 {
		return status.Error(codes.InvalidArgument, "organization ID is required")
	}
	if req.Name == nil && req.Address == nil && req.Chief == nil && req.FinChief == nil && req.InnKpp == nil {
		return status.Error(codes.InvalidArgument, "updating structure has no values")
	}
	if req.InnKpp != nil {
		if *req.InnKpp == "" {
			return status.Error(codes.InvalidArgument, "updating structure has empty values")
		}
		if len(*req.InnKpp) != 20 {
			return status.Error(codes.InvalidArgument, "invalid field format: inn/kpp")
		}
	}
	if req.Name != nil {
		if *req.Name == "" {
			return status.Error(codes.InvalidArgument, "updating structure has empty values")
		}
	}
	if req.Address != nil {
		if *req.Address == "" {
			return status.Error(codes.InvalidArgument, "updating structure has empty values")
		}
	}
	if req.Chief != nil {
		if *req.Chief == "" {
			return status.Error(codes.InvalidArgument, "updating structure has empty values")
		}
	}
	if req.FinChief != nil {
		if *req.FinChief == "" {
			return status.Error(codes.InvalidArgument, "updating structure has empty values")
		}
	}
	return nil
}

func validateDeleteOrganization(req *putlistv1.DeleteOrganizationRequest) error {
	if req.GetOrganizationId() == 0 {
		return status.Error(codes.InvalidArgument, "organization ID is required")
	}
	return nil
}
