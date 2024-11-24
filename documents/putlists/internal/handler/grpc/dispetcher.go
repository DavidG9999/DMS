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

type DispetcherHandler struct {
	putlistv1.UnimplementedDispetcherServer
	DispetcherService service.Dispetcher
}

func (h *DispetcherHandler) CreateDispetcher(ctx context.Context, req *putlistv1.CreateDispetcherRequest) (*putlistv1.CreateDispetcherResponse, error) {
	if err := validateCreateDispetcher(req); err != nil {
		return nil, err
	}

	dispetcher := entity.Dispetcher{
		FullName: req.GetFullName(),
	}
	
	dispetcherId, err := h.DispetcherService.CreateDispetcher(ctx, dispetcher)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &putlistv1.CreateDispetcherResponse{
		DispetcherId: dispetcherId,
	}, nil
}

func (h *DispetcherHandler) GetDispetchers(ctx context.Context, req *putlistv1.GetDispetchersRequest) (*putlistv1.GetDispetchersResponse, error) {
	dispetchers, err := h.DispetcherService.GetDispetchers(ctx)
	if err != nil {
		if errors.Is(err, service.ErrInvalidCredentials) {
			return nil, status.Error(codes.NotFound, "dispetchers not found")
		}
		return nil, status.Error(codes.Internal, "internal error")
	}

	dispetchersList := make([]putlistv1.DispetcherEntity, len(dispetchers))
	for idx, dispetcher := range dispetchers {
		dispetchersList[idx].DispetcherId = dispetcher.Id
		dispetchersList[idx].FullName = dispetcher.FullName
	}

	dispetchersResp := make([]*putlistv1.DispetcherEntity, 0, len(dispetchers))
	for id := range dispetchersList {
		dispetchersResp = append(dispetchersResp, &dispetchersList[id])
	}

	return &putlistv1.GetDispetchersResponse{
		Dispetchers: dispetchersResp,
	}, nil
}

func (h *DispetcherHandler) UpdateDispetcher(ctx context.Context, req *putlistv1.UpdateDispetcherRequest) (*putlistv1.UpdateDispetcherResponse, error) {
	if err := validateUpdateDispetcher(req); err != nil {
		return nil, err
	}

	updateData := entity.UpdateDispetcherInput{
		FullName: req.FullName,
	}
	err := h.DispetcherService.UpdateDispetcher(ctx, req.GetDispetcherId(), updateData)
	if err != nil {
		if errors.Is(err, service.ErrInvalidCredentials) {
			return nil, status.Error(codes.NotFound, "dispetcher not found")
		}
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &putlistv1.UpdateDispetcherResponse{
		Message: "updated dispetcher",
	}, nil
}

func (h *DispetcherHandler) DeleteDispetcher(ctx context.Context, req *putlistv1.DeleteDispetcherRequest) (*putlistv1.DeleteDispetcherResponse, error) {
	if err := validateDeleteDispetcher(req); err != nil {
		return nil, err
	}

	err := h.DispetcherService.DeleteDispetcher(ctx, req.GetDispetcherId())
	if err != nil {
		if errors.Is(err, service.ErrInvalidCredentials) {
			return nil, status.Error(codes.NotFound, "dispetcher not found")
		}
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &putlistv1.DeleteDispetcherResponse{
		Message: "deleted",
	}, nil
}

func validateCreateDispetcher(req *putlistv1.CreateDispetcherRequest) error {
	if req.GetFullName() == "" {
		return status.Error(codes.InvalidArgument, "fullname is required")
	}
	return nil
}

func validateUpdateDispetcher(req *putlistv1.UpdateDispetcherRequest) error {
	if req.GetDispetcherId() == 0 {
		return status.Error(codes.InvalidArgument, "dispetcher ID is required")
	}
	if req.FullName == nil {
		return status.Error(codes.InvalidArgument, "updating structure has no values")
	}
	if req.FullName != nil {
		if *req.FullName == "" {
			return status.Error(codes.InvalidArgument, "updating structure has empty values")
		}
	}
	return nil
}

func validateDeleteDispetcher(req *putlistv1.DeleteDispetcherRequest) error {
	if req.GetDispetcherId() == 0 {
		return status.Error(codes.InvalidArgument, "dispetcher ID is required")
	}
	return nil
}
