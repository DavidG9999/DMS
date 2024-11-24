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

type PutlistHandler struct {
	putlistv1.UnimplementedPutlistServer
	PutlistService service.Putlist
}

func (h *PutlistHandler) CreatePutlist(ctx context.Context, req *putlistv1.CreatePutlistRequest) (*putlistv1.CreatePutlistResponse, error) {
	if err := validateCreatePutlist(req); err != nil {
		return nil, err
	}

	putlist := entity.PutlistHeader{
		UserId:        req.GetUserId(),
		Number:        req.GetNumber(),
		BankAccountId: req.GetBankAccountId(),
		DateWith:      req.GetDateWith(),
		DateFor:       req.GetDateFor(),
		AutoId:        req.GetAutoId(),
		DriverId:      req.GetDriverId(),
		DispetcherId:  req.GetDispetcherId(),
		MehanicId:     req.GetMehanicId(),
	}
	putlistId, err := h.PutlistService.CreatePutlist(ctx, putlist)
	if err != nil {
		if errors.Is(err, service.ErrInvalidDateTimeFormat) {
			return nil, status.Error(codes.InvalidArgument, "invalid datetime format")
		}
		if errors.Is(err, service.ErrPutlistExists) {
			return nil, status.Error(codes.AlreadyExists, "putlist already exists")
		}
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &putlistv1.CreatePutlistResponse{
		PutlistId: putlistId,
	}, nil
}

func (h *PutlistHandler) GetPutlists(ctx context.Context, req *putlistv1.GetPutlistsRequest) (*putlistv1.GetPutlistsResponse, error) {
	if err := validateGetPutlists(req); err != nil {
		return nil, err
	}

	putlists, err := h.PutlistService.GetPutlists(ctx, req.GetUserId())
	if err != nil {
		if errors.Is(err, service.ErrInvalidCredentials) {
			return nil, status.Error(codes.NotFound, "putlists not found")
		}
		return nil, status.Error(codes.Internal, "internal error")
	}

	putlistsList := make([]putlistv1.PutlistEntity, len(putlists))
	for idx, putlist := range putlists {
		putlistsList[idx].PutlistId = putlist.Id
		putlistsList[idx].UserId = putlist.UserId
		putlistsList[idx].Number = putlist.Number
		putlistsList[idx].BankAccountId = putlist.BankAccountId
		putlistsList[idx].DateWith = putlist.DateWith
		putlistsList[idx].DateFor = putlist.DateFor
		putlistsList[idx].AutoId = putlist.AutoId
		putlistsList[idx].DriverId = putlist.DriverId
		putlistsList[idx].DispetcherId = putlist.DispetcherId
		putlistsList[idx].MehanicId = putlist.MehanicId
	}

	putlistsResp := make([]*putlistv1.PutlistEntity, 0, len(putlists))
	for id := range putlistsList {
		putlistsResp = append(putlistsResp, &putlistsList[id])
	}

	return &putlistv1.GetPutlistsResponse{
		Putlists: putlistsResp,
	}, nil
}

func (h *PutlistHandler) GetPutlistByNumber(ctx context.Context, req *putlistv1.GetPutlistByNumberRequest) (*putlistv1.GetPutlistByNumberResponse, error) {
	if err := validateGetPutlistByNumber(req); err != nil {
		return nil, err
	}

	putlist, err := h.PutlistService.GetPutlistByNumber(ctx, req.GetUserId(), req.GetNumber())
	if err != nil {
		if errors.Is(err, service.ErrInvalidCredentials) {
			return nil, status.Error(codes.NotFound, "putlist by this number not found")
		}
		return nil, status.Error(codes.Internal, "internal error")
	}
	putlistResp := putlistv1.PutlistEntity{
		PutlistId:     putlist.Id,
		UserId:        putlist.UserId,
		Number:        putlist.Number,
		BankAccountId: putlist.BankAccountId,
		DateWith:      putlist.DateWith,
		DateFor:       putlist.DateFor,
		AutoId:        putlist.AutoId,
		DriverId:      putlist.DriverId,
		DispetcherId:  putlist.DispetcherId,
		MehanicId:     putlist.MehanicId,
	}

	return &putlistv1.GetPutlistByNumberResponse{
		Putlist: &putlistResp,
	}, nil
}

func (h *PutlistHandler) UpdatePutlist(ctx context.Context, req *putlistv1.UpdatePutlistRequest) (*putlistv1.UpdatePutlistResponse, error) {
	if err := validateUpdatePutlist(req); err != nil {
		return nil, err
	}

	updateData := entity.UpdatePutlistHeaderInput{
		BankAccountId: req.BankAccountId,
		DateWith:      req.DateWith,
		DateFor:       req.DateFor,
		AutoId:        req.AutoId,
		DriverId:      req.DriverId,
		DispetcherId:  req.DispetcherId,
		MehanicId:     req.MehanicId,
	}
	err := h.PutlistService.UpdatePutlist(ctx, req.GetUserId(), req.GetNumber(), updateData)
	if err != nil {
		if errors.Is(err, service.ErrInvalidDateTimeFormat) {
			return nil, status.Error(codes.InvalidArgument, "invalid datetime format")
		}
		if errors.Is(err, service.ErrInvalidCredentials) {
			return nil, status.Error(codes.NotFound, "putlist not found")
		}
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &putlistv1.UpdatePutlistResponse{
		Message: "updated putlist",
	}, nil
}

func (h *PutlistHandler) DeletePutlist(ctx context.Context, req *putlistv1.DeletePutlistRequest) (*putlistv1.DeletePutlistResponse, error) {
	if err := validateDeletePutlist(req); err != nil {
		return nil, err
	}

	err := h.PutlistService.DeletePutlist(ctx, req.GetUserId(), req.GetNumber())
	if err != nil {
		if errors.Is(err, service.ErrInvalidCredentials) {
			return nil, status.Error(codes.NotFound, "putlist not found")
		}
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &putlistv1.DeletePutlistResponse{
		Message: "deleted",
	}, nil
}

func validateCreatePutlist(req *putlistv1.CreatePutlistRequest) error {
	if req.GetUserId() == 0 {
		return status.Error(codes.InvalidArgument, "user ID is required")
	}
	if req.GetNumber() == 0 {
		return status.Error(codes.InvalidArgument, "putlist number is required")
	}
	if req.GetBankAccountId() == 0 {
		return status.Error(codes.InvalidArgument, "bank account ID is required")
	}
	if req.GetDateWith() == "" {
		return status.Error(codes.InvalidArgument, "date with is required")
	}
	if req.GetDateFor() == "" {
		return status.Error(codes.InvalidArgument, "date for is required")
	}
	if req.GetAutoId() == 0 {
		return status.Error(codes.InvalidArgument, "auto ID is required")
	}
	if req.GetDriverId() == 0 {
		return status.Error(codes.InvalidArgument, "driver ID is required")
	}
	if req.GetDispetcherId() == 0 {
		return status.Error(codes.InvalidArgument, "dispetcher ID is required")
	}
	if req.GetMehanicId() == 0 {
		return status.Error(codes.InvalidArgument, "mehanic ID is required")
	}
	return nil
}

func validateGetPutlists(req *putlistv1.GetPutlistsRequest) error {
	if req.GetUserId() == 0 {
		return status.Error(codes.InvalidArgument, "user ID is required")
	}
	return nil
}

func validateGetPutlistByNumber(req *putlistv1.GetPutlistByNumberRequest) error {
	if req.GetUserId() == 0 {
		return status.Error(codes.InvalidArgument, "user ID is required")
	}
	if req.GetNumber() == 0 {
		return status.Error(codes.InvalidArgument, "number is required")
	}
	return nil
}

func validateUpdatePutlist(req *putlistv1.UpdatePutlistRequest) error {
	if req.GetUserId() == 0 {
		return status.Error(codes.InvalidArgument, "user ID is required")
	}
	if req.GetNumber() == 0 {
		return status.Error(codes.InvalidArgument, "number is required")
	}
	if req.BankAccountId == nil && req.DateWith == nil && req.DateFor == nil && req.AutoId == nil && req.DriverId == nil && req.DispetcherId == nil && req.MehanicId == nil {
		return status.Error(codes.InvalidArgument, "updating structure has no values")
	}
	if req.BankAccountId != nil {
		if *req.BankAccountId == 0 {
			return status.Error(codes.InvalidArgument, "updating structure has empty values")
		}
	}
	if req.DateWith != nil {
		if *req.DateWith == "" {
			return status.Error(codes.InvalidArgument, "updating structure has empty values")
		}
	}
	if req.DateFor != nil {
		if *req.DateFor == "" {
			return status.Error(codes.InvalidArgument, "updating structure has empty values")
		}
	}
	if req.AutoId != nil {
		if *req.AutoId == 0 {
			return status.Error(codes.InvalidArgument, "updating structure has empty values")
		}
	}
	if req.DriverId != nil {
		if *req.DriverId == 0 {
			return status.Error(codes.InvalidArgument, "updating structure has empty values")
		}
	}
	if req.DispetcherId != nil {
		if *req.DispetcherId == 0 {
			return status.Error(codes.InvalidArgument, "updating structure has empty values")
		}
	}
	if req.MehanicId != nil {
		if *req.MehanicId == 0 {
			return status.Error(codes.InvalidArgument, "updating structure has empty values")
		}
	}
	return nil
}

func validateDeletePutlist(req *putlistv1.DeletePutlistRequest) error {
	if req.GetUserId() == 0 {
		return status.Error(codes.InvalidArgument, "user ID is required")
	}
	if req.GetNumber() == 0 {
		return status.Error(codes.InvalidArgument, "number is required")
	}
	return nil
}
