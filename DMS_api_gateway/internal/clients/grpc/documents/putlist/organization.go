package putlistgrpc

import (
		"github.com/DavidG9999/DMS/DMS_api_gateway/internal/entity"
	"context"
	"fmt"

	putlistv1 "github.com/DavidG9999/DMS/api/grpc/document_api/putlist/gen/go"
)

func (pc *PutlistClient) CreateOrganization(ctx context.Context, name, address, chief, filChief, innKpp string) (int64, error) {
	const op = "grpc.ClientCreateOrganization"

	resp, err := pc.apiOrganization.CreateOrganization(ctx, &putlistv1.CreateOrganizationRequest{Name: name, Address: address, Chief: chief, FinChief: filChief, InnKpp: innKpp})
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return resp.GetOrganizationId(), nil
}

func (pc *PutlistClient) GetOrganizations(ctx context.Context) ([]entity.Organization, error) {
	const op = "grpc.ClientGetOrganizations"

	resp, err := pc.apiOrganization.GetOrganizations(ctx, &putlistv1.GetOrganizationsRequest{})
	if err != nil {
		return []entity.Organization{}, fmt.Errorf("%s: %w", op, err)
	}

	organizations := make([]entity.Organization, len(resp.Organizations))
	for id, organization := range resp.GetOrganizations() {
		organizations[id].Id = organization.OrganizationId
		organizations[id].Name = organization.Name
		organizations[id].Address = organization.Address
		organizations[id].Chief = organization.Chief
		organizations[id].FinancialChief = organization.FinChief
		organizations[id].InnKpp = organization.InnKpp
	}
	return organizations, nil
}

func (pc *PutlistClient) UpdateOrganization(ctx context.Context, organizationId int64, name, address, chief, finChief, innKpp *string) (string, error) {
	const op = "grpc.ClientUpdateOrganization"

	resp, err := pc.apiOrganization.UpdateOrganization(ctx, &putlistv1.UpdateOrganizationRequest{OrganizationId: organizationId, Name: name, Address: address, Chief: chief, FinChief: finChief, InnKpp: innKpp})
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return resp.GetMessage(), nil
}

func (pc *PutlistClient) DeleteOrganization(ctx context.Context, organizationId int64) (string, error) {
	const op = "grpc.ClientDeleteOrganization"

	resp, err := pc.apiOrganization.DeleteOrganization(ctx, &putlistv1.DeleteOrganizationRequest{OrganizationId: organizationId})
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return resp.GetMessage(), nil
}
