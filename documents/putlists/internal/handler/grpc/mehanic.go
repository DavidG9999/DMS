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

type MehanicHandler struct {
	putlistv1.UnimplementedMehanicServer
	MehanicService service.Mehanic
}

func (h *MehanicHandler) CreateMehanic(ctx context.Context, req *putlistv1.CreateMehanicRequest) (*putlistv1.CreateMehanicResponse, error) {
	if err := validateCreateMehanic(req); err != nil {
		return nil, err
	}

	mehanic := entity.Mehanic{
		FullName: req.GetFullName(),
	}
	mehanicId, err := h.MehanicService.CreateMehanic(ctx, mehanic)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &putlistv1.CreateMehanicResponse{
		MehanicId: mehanicId,
	}, nil
}

func (h *MehanicHandler) GetMehanics(ctx context.Context, req *putlistv1.GetMehanicsRequest) (*putlistv1.GetMehanicsResponse, error) {
	mehanics, err := h.MehanicService.GetMehanics(ctx)
	if err != nil {
		if errors.Is(err, service.ErrInvalidCredentials) {
			return nil, status.Error(codes.NotFound, "mehanics not found")
		}
		return nil, status.Error(codes.Internal, "internal error")
	}

	mehanicsList := make([]putlistv1.MehanicEntity, len(mehanics))
	for idx, mehanic := range mehanics {
		mehanicsList[idx].MehanicId = mehanic.Id
		mehanicsList[idx].FullName = mehanic.FullName
	}

	mehanicsResp := make([]*putlistv1.MehanicEntity, 0, len(mehanics))
	for id := range mehanicsList {
		mehanicsResp = append(mehanicsResp, &mehanicsList[id])
	}

	return &putlistv1.GetMehanicsResponse{
		Mehanics: mehanicsResp,
	}, nil
}

func (h *MehanicHandler) UpdateMehanic(ctx context.Context, req *putlistv1.UpdateMehanicRequest) (*putlistv1.UpdateMehanicResponse, error) {
	if err := validateUpdateMehanic(req); err != nil {
		return nil, err
	}

	updateData := entity.UpdateMehanicInput{
		FullName: req.FullName,
	}
	err := h.MehanicService.UpdateMehanic(ctx, req.GetMehanicId(), updateData)
	if err != nil {
		if errors.Is(err, service.ErrInvalidCredentials) {
			return nil, status.Error(codes.NotFound, "mehanic not found")
		}
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &putlistv1.UpdateMehanicResponse{
		Message: "updated mehanic",
	}, nil
}

func (h *MehanicHandler) DeleteMehanic(ctx context.Context, req *putlistv1.DeleteMehanicRequest) (*putlistv1.DeleteMehanicResponse, error) {
	if err := validateDeleteMehanic(req); err != nil {
		return nil, err
	}

	err := h.MehanicService.DeleteMehanic(ctx, req.GetMehanicId())
	if err != nil {
		if errors.Is(err, service.ErrInvalidCredentials) {
			return nil, status.Error(codes.NotFound, "mehanic not found")
		}
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &putlistv1.DeleteMehanicResponse{
		Message: "deleted",
	}, nil
}

func validateCreateMehanic(req *putlistv1.CreateMehanicRequest) error {
	if req.GetFullName() == "" {
		return status.Error(codes.InvalidArgument, "fullname is required")
	}
	return nil
}

func validateUpdateMehanic(req *putlistv1.UpdateMehanicRequest) error {
	if req.GetMehanicId() == 0 {
		return status.Error(codes.InvalidArgument, "mehanic ID is required")
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

func validateDeleteMehanic(req *putlistv1.DeleteMehanicRequest) error {
	if req.GetMehanicId() == 0 {
		return status.Error(codes.InvalidArgument, "mehanic ID is required")
	}
	return nil
}
