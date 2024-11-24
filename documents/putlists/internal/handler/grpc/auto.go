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

type AutoHandler struct {
	putlistv1.UnimplementedAutoServer
	AutoService service.Auto
}

func (h *AutoHandler) CreateAuto(ctx context.Context, req *putlistv1.CreateAutoRequest) (*putlistv1.CreateAutoResponse, error) {
	if err := validateCreateAuto(req); err != nil {
		return nil, err
	}

	auto := entity.Auto{
		Brand:       req.GetBrand(),
		Model:       req.GetModel(),
		StateNumber: req.GetStateNumber(),
	}

	autoId, err := h.AutoService.CreateAuto(ctx, auto)
	if err != nil {
		if errors.Is(err, service.ErrAutoExists) {
			return nil, status.Error(codes.AlreadyExists, "auto already exists")
		}
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &putlistv1.CreateAutoResponse{
		AutoId: autoId,
	}, nil
}

func (h *AutoHandler) GetAutos(ctx context.Context, req *putlistv1.GetAutosRequest) (*putlistv1.GetAutosResponse, error) {
	autos, err := h.AutoService.GetAutos(ctx)
	if err != nil {
		if errors.Is(err, service.ErrInvalidCredentials) {
			return nil, status.Error(codes.NotFound, "autos not found")
		}
		return nil, status.Error(codes.Internal, "internal error")
	}

	autosList := make([]putlistv1.AutoEntity, len(autos))
	for idx, auto := range autos {
		autosList[idx].AutoId = auto.Id
		autosList[idx].Brand = auto.Brand
		autosList[idx].Model = auto.Model
		autosList[idx].StateNumber = auto.StateNumber
	}

	autosResp := make([]*putlistv1.AutoEntity, 0, len(autos))

	for id := range autosList {
		autosResp = append(autosResp, &autosList[id])
	}

	return &putlistv1.GetAutosResponse{
		Autos: autosResp,
	}, nil
}

func (h *AutoHandler) UpdateAuto(ctx context.Context, req *putlistv1.UpdateAutoRequest) (*putlistv1.UpdateAutoResponse, error) {
	if err := validateUpdateAuto(req); err != nil {
		return nil, err
	}

	updateData := entity.UpdateAutoInput{
		Brand:       req.Brand,
		Model:       req.Model,
		StateNumber: req.StateNumber,
	}

	err := h.AutoService.UpdateAuto(ctx, req.GetAutoId(), updateData)
	if err != nil {
		if errors.Is(err, service.ErrInvalidCredentials) {
			return nil, status.Error(codes.NotFound, "auto not found")
		}
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &putlistv1.UpdateAutoResponse{
		Message: "updated auto",
	}, nil
}

func (h *AutoHandler) DeleteAuto(ctx context.Context, req *putlistv1.DeleteAutoRequest) (*putlistv1.DeleteAutoResponse, error) {
	if err := validateDeleteAuto(req); err != nil {
		return nil, err
	}

	err := h.AutoService.DeleteAuto(ctx, req.GetAutoId())
	if err != nil {
		if errors.Is(err, service.ErrInvalidCredentials) {
			return nil, status.Error(codes.NotFound, "auto not found")
		}
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &putlistv1.DeleteAutoResponse{
		Message: "deleted",
	}, nil
}

func validateCreateAuto(req *putlistv1.CreateAutoRequest) error {
	if req.GetBrand() == "" {
		return status.Error(codes.InvalidArgument, "brand is required")
	}
	if req.GetModel() == "" {
		return status.Error(codes.InvalidArgument, "model is required")
	}
	if req.GetStateNumber() == "" {
		return status.Error(codes.InvalidArgument, "state number is required")
	}
	if len(req.GetStateNumber()) > 12 || len(req.GetStateNumber()) < 11 {
		return status.Error(codes.InvalidArgument, "invalid field format: state number")
	}
	return nil
}

func validateUpdateAuto(req *putlistv1.UpdateAutoRequest) error {
	if req.GetAutoId() == 0 {
		return status.Error(codes.InvalidArgument, "auto ID is required")
	}
	if req.Brand == nil && req.Model == nil && req.StateNumber == nil {
		return status.Error(codes.InvalidArgument, "updating structure has no values")
	}
	if req.StateNumber != nil {
		if *req.StateNumber == "" {
			return status.Error(codes.InvalidArgument, "updating structure has empty values")
		}
		if len(*req.StateNumber) > 12 || len(*req.StateNumber) < 11 {
			return status.Error(codes.InvalidArgument, "invalid field format: state number")
		}
	}
	if req.Brand != nil {
		if *req.Brand == "" {
			return status.Error(codes.InvalidArgument, "updating structure has empty values")
		}
	}
	if req.Model != nil {
		if *req.Model == "" {
			return status.Error(codes.InvalidArgument, "updating structure has empty values")
		}
	}
	return nil
}

func validateDeleteAuto(req *putlistv1.DeleteAutoRequest) error {
	if req.GetAutoId() == 0 {
		return status.Error(codes.InvalidArgument, "auto ID is required")
	}
	return nil
}
