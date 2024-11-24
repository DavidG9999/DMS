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

type PutlistBodyHandler struct {
	putlistv1.UnimplementedPutlistBodyServer
	PutlistBodyService service.PutlistBody
}

func (h *PutlistBodyHandler) CreatePutlistBody(ctx context.Context, req *putlistv1.CreatePutlistBodyRequest) (*putlistv1.CreatePutlistBodyResponse, error) {
	if err := validateCreatePutlistBody(req); err != nil {
		return nil, err
	}

	putlistBody := entity.PutlistBody{
		PutlistNumber: req.GetPutlistNumber(),
		Number:        req.GetNumber(),
		ContragentId:  req.GetContragentId(),
		Item:          req.GetItem(),
		TimeWith:      req.GetTimeWith(),
		TimeFor:       req.GetTimeFor(),
	}
	putlistBodyId, err := h.PutlistBodyService.CreatePutlistBody(ctx, putlistBody)
	if err != nil {
		if errors.Is(err, service.ErrInvalidDateTimeFormat) {
			return nil, status.Error(codes.InvalidArgument, "invalid datetime format")
		}
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &putlistv1.CreatePutlistBodyResponse{
		PutlistBodyId: putlistBodyId,
	}, nil
}

func (h *PutlistBodyHandler) GetPutlistBodies(ctx context.Context, req *putlistv1.GetPutlistBodiesRequest) (*putlistv1.GetPutlistBodiesResponse, error) {
	if err := validateGetPutlistBodies(req); err != nil {
		return nil, err
	}

	putlistBodies, err := h.PutlistBodyService.GetPutlistBodies(ctx, req.GetPutlistNumber())
	if err != nil {
		if errors.Is(err, service.ErrInvalidCredentials) {
			return nil, status.Error(codes.NotFound, "putlist bodies not found")
		}
		return nil, status.Error(codes.Internal, "internal error")
	}

	putlistBodiesList := make([]putlistv1.PutlistBodyEntity, len(putlistBodies))
	for idx, putlistBody := range putlistBodies {
		putlistBodiesList[idx].PutlistBodyId = putlistBody.Id
		putlistBodiesList[idx].PutlistNumber = putlistBody.PutlistNumber
		putlistBodiesList[idx].Number = putlistBody.Number
		putlistBodiesList[idx].ContragentId = putlistBody.ContragentId
		putlistBodiesList[idx].Item = putlistBody.Item
		putlistBodiesList[idx].TimeWith = putlistBody.TimeWith
		putlistBodiesList[idx].TimeFor = putlistBody.TimeFor
	}

	putlistBodiesResp := make([]*putlistv1.PutlistBodyEntity, 0, len(putlistBodies))
	for id := range putlistBodiesList {
		putlistBodiesResp = append(putlistBodiesResp, &putlistBodiesList[id])
	}

	return &putlistv1.GetPutlistBodiesResponse{
		PutlistBodies: putlistBodiesResp,
	}, nil
}

func (h *PutlistBodyHandler) UpdatePutlistBody(ctx context.Context, req *putlistv1.UpdatePutlistBodyRequest) (*putlistv1.UpdatePutlistBodyResponse, error) {
	if err := validateUpdatePutlistBody(req); err != nil {
		return nil, err
	}

	updateData := entity.UpdatePutlistBodyInput{
		Number:       req.Number,
		ContragentId: req.ContragentId,
		Item:         req.Item,
		TimeWith:     req.TimeWith,
		TimeFor:      req.TimeFor,
	}
	err := h.PutlistBodyService.UpdatePutlistBody(ctx, req.GetPutlistBodyId(), updateData)
	if err != nil {
		if errors.Is(err, service.ErrInvalidDateTimeFormat) {
			return nil, status.Error(codes.InvalidArgument, "invalid datetime format")
		}
		if errors.Is(err, service.ErrInvalidCredentials) {
			return nil, status.Error(codes.NotFound, "putlist body not found")
		}
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &putlistv1.UpdatePutlistBodyResponse{
		Message: "updated putlist body",
	}, nil
}

func (h *PutlistBodyHandler) DeletePutlistBody(ctx context.Context, req *putlistv1.DeletePutlistBodyRequest) (*putlistv1.DeletePutlistBodyResponse, error) {
	if err := validateDeletePutlistBody(req); err != nil {
		return nil, err
	}

	err := h.PutlistBodyService.DeletePutlistBody(ctx, req.GetPutlistBodyId())
	if err != nil {
		if errors.Is(err, service.ErrInvalidCredentials) {
			return nil, status.Error(codes.NotFound, "putlist body not found")
		}
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &putlistv1.DeletePutlistBodyResponse{
		Message: "deleted",
	}, nil
}

func validateCreatePutlistBody(req *putlistv1.CreatePutlistBodyRequest) error {
	if req.GetPutlistNumber() == 0 {
		return status.Error(codes.InvalidArgument, "putlist number is required")
	}
	if req.GetNumber() == 0 {
		return status.Error(codes.InvalidArgument, "putlist body number is required")
	}
	if req.GetContragentId() == 0 {
		return status.Error(codes.InvalidArgument, "contragent ID is required")
	}
	if req.GetItem() == "" {
		return status.Error(codes.InvalidArgument, "item is required")
	}
	if req.GetTimeWith() == "" {
		return status.Error(codes.InvalidArgument, "time with is required")
	}
	if req.GetTimeFor() == "" {
		return status.Error(codes.InvalidArgument, "time for is required")
	}
	return nil
}

func validateGetPutlistBodies(req *putlistv1.GetPutlistBodiesRequest) error {
	if req.GetPutlistNumber() == 0 {
		return status.Error(codes.InvalidArgument, "putlist number is required")
	}
	return nil
}

func validateUpdatePutlistBody(req *putlistv1.UpdatePutlistBodyRequest) error {
	if req.GetPutlistBodyId() == 0 {
		return status.Error(codes.InvalidArgument, "putlist body ID is required")
	}
	if req.Number == nil && req.ContragentId == nil && req.Item == nil && req.TimeWith == nil && req.TimeFor == nil {
		return status.Error(codes.InvalidArgument, "updating structure has no values")
	}
	if req.Number != nil {
		if *req.Number == 0 {
			return status.Error(codes.InvalidArgument, "updating structure has empty values")
		}
	}
	if req.ContragentId != nil {
		if *req.ContragentId == 0 {
			return status.Error(codes.InvalidArgument, "updating structure has empty values")
		}
	}
	if req.Item != nil {
		if *req.Item == "" {
			return status.Error(codes.InvalidArgument, "updating structure has empty values")
		}
	}
	if req.TimeWith != nil {
		if *req.TimeWith == "" {
			return status.Error(codes.InvalidArgument, "updating structure has empty values")
		}
	}
	if req.TimeFor != nil {
		if *req.TimeFor == "" {
			return status.Error(codes.InvalidArgument, "updating structure has empty values")
		}
	}
	return nil
}

func validateDeletePutlistBody(req *putlistv1.DeletePutlistBodyRequest) error {
	if req.GetPutlistBodyId() == 0 {
		return status.Error(codes.InvalidArgument, "putlist body ID is required")
	}
	return nil
}
