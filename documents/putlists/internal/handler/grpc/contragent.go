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

type ContragentHandler struct {
	putlistv1.UnimplementedContragentServer
	ContragentService service.Contragent
}

func (h *ContragentHandler) CreateContragent(ctx context.Context, req *putlistv1.CreateContragentRequest) (*putlistv1.CreateContragentResponse, error) {
	if err := validateCreateContragent(req); err != nil {
		return nil, err
	}

	contragent := entity.Contragent{
		Name:    req.GetName(),
		Address: req.GetAddress(),
		InnKpp:  req.GetInnKpp(),
	}
	contragentId, err := h.ContragentService.CreateContragent(ctx, contragent)
	if err != nil {
		if errors.Is(err, service.ErrContragentExists) {
			return nil, status.Error(codes.AlreadyExists, "contragent already exists")
		}
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &putlistv1.CreateContragentResponse{
		ContragentId: contragentId,
	}, nil
}

func (h *ContragentHandler) GetContragents(ctx context.Context, req *putlistv1.GetContragentsRequest) (*putlistv1.GetContragentsResponse, error) {
	contragents, err := h.ContragentService.GetContragents(ctx)
	if err != nil {
		if errors.Is(err, service.ErrInvalidCredentials) {
			return nil, status.Error(codes.NotFound, "contragents not found")
		}
		return nil, status.Error(codes.Internal, "internal error")
	}

	contragentsList := make([]putlistv1.ContragentEntity, len(contragents))
	for idx, contragent := range contragents {
		contragentsList[idx].ContragentId = contragent.Id
		contragentsList[idx].Name = contragent.Name
		contragentsList[idx].Address = contragent.Address
		contragentsList[idx].InnKpp = contragent.InnKpp
	}

	contragentsResp := make([]*putlistv1.ContragentEntity, 0, len(contragents))
	for id := range contragentsList {
		contragentsResp = append(contragentsResp, &contragentsList[id])
	}

	return &putlistv1.GetContragentsResponse{
		Contragents: contragentsResp,
	}, nil
}

func (h *ContragentHandler) UpdateContragent(ctx context.Context, req *putlistv1.UpdateContragentRequest) (*putlistv1.UpdateContragentResponse, error) {
	if err := validateUpdateContragent(req); err != nil {
		return nil, err
	}

	updateData := entity.UpdateContragentInput{
		Name:    req.Name,
		Address: req.Address,
		InnKpp:  req.InnKpp,
	}
	err := h.ContragentService.UpdateContragent(ctx, req.GetContragentId(), updateData)
	if err != nil {
		if errors.Is(err, service.ErrInvalidCredentials) {
			return nil, status.Error(codes.NotFound, "contragent not found")
		}
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &putlistv1.UpdateContragentResponse{
		Message: "updated contragent",
	}, nil
}

func (h *ContragentHandler) DeleteContragent(ctx context.Context, req *putlistv1.DeleteContragentRequest) (*putlistv1.DeleteContragentResponse, error) {
	if err := validateDeleteContragent(req); err != nil {
		return nil, err
	}

	err := h.ContragentService.DeleteContragent(ctx, req.GetContragentId())
	if err != nil {
		if errors.Is(err, service.ErrInvalidCredentials) {
			return nil, status.Error(codes.NotFound, "contragent not found")
		}
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &putlistv1.DeleteContragentResponse{
		Message: "deleted",
	}, nil
}

func validateCreateContragent(req *putlistv1.CreateContragentRequest) error {
	if req.GetName() == "" {
		return status.Error(codes.InvalidArgument, "contragent name is required")
	}
	if req.GetAddress() == "" {
		return status.Error(codes.InvalidArgument, "address is required")
	}
	if req.GetInnKpp() == "" {
		return status.Error(codes.InvalidArgument, "inn/kpp is required")
	}
	if len(req.GetInnKpp()) != 20 {
		return status.Error(codes.InvalidArgument, "invalid field format: inn/kpp")
	}
	return nil
}

func validateUpdateContragent(req *putlistv1.UpdateContragentRequest) error {
	if req.GetContragentId() == 0 {
		return status.Error(codes.InvalidArgument, "contragent ID is required")
	}
	if req.Name == nil && req.Address == nil && req.InnKpp == nil {
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
	return nil
}

func validateDeleteContragent(req *putlistv1.DeleteContragentRequest) error {
	if req.GetContragentId() == 0 {
		return status.Error(codes.InvalidArgument, "contragent ID is required")
	}
	return nil
}
