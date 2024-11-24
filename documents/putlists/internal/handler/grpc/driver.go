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

type DriverHandler struct {
	putlistv1.UnimplementedDriverServer
	DriverService service.Driver
}

func (h *DriverHandler) CreateDriver(ctx context.Context, req *putlistv1.CreateDriverRequest) (*putlistv1.CreateDriverResponse, error) {
	if err := validateCreateDriver(req); err != nil {
		return nil, err
	}

	driver := entity.Driver{
		FullName: req.GetFullName(),
		License:  req.GetLicense(),
		Class:    req.GetClass(),
	}
	driverId, err := h.DriverService.CreateDriver(ctx, driver)
	if err != nil {
		if errors.Is(err, service.ErrDriverExists) {
			return nil, status.Error(codes.AlreadyExists, "driver already exists")
		}
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &putlistv1.CreateDriverResponse{
		DriverId: driverId,
	}, nil
}

func (h *DriverHandler) GetDrivers(ctx context.Context, req *putlistv1.GetDriversRequest) (*putlistv1.GetDriversResponse, error) {
	drivers, err := h.DriverService.GetDrivers(ctx)
	if err != nil {
		if errors.Is(err, service.ErrInvalidCredentials) {
			return nil, status.Error(codes.NotFound, "drivers not found")
		}
		return nil, status.Error(codes.Internal, "internal error")
	}

	driversList := make([]putlistv1.DriverEntity, len(drivers))
	for idx, driver := range drivers {
		driversList[idx].DriverId = driver.Id
		driversList[idx].FullName = driver.FullName
		driversList[idx].License = driver.License
		driversList[idx].Class = driver.Class
	}

	driversResp := make([]*putlistv1.DriverEntity, 0, len(drivers))
	for id := range driversList {
		driversResp = append(driversResp, &driversList[id])
	}

	return &putlistv1.GetDriversResponse{
		Drivers: driversResp,
	}, nil
}

func (h *DriverHandler) UpdateDriver(ctx context.Context, req *putlistv1.UpdateDriverRequest) (*putlistv1.UpdateDriverResponse, error) {
	if err := validateUpdateDriver(req); err != nil {
		return nil, err
	}

	updateData := entity.UpdateDriverInput{
		FullName: req.FullName,
		License:  req.License,
		Class:    req.Class,
	}
	err := h.DriverService.UpdateDriver(ctx, req.GetDriverId(), updateData)
	if err != nil {
		if errors.Is(err, service.ErrInvalidCredentials) {
			return nil, status.Error(codes.NotFound, "driver not found")
		}
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &putlistv1.UpdateDriverResponse{
		Message: "updated driver",
	}, nil
}

func (h *DriverHandler) DeleteDriver(ctx context.Context, req *putlistv1.DeleteDriverRequest) (*putlistv1.DeleteDriverResponse, error) {
	if err := validateDeleteDriver(req); err != nil {
		return nil, err
	}

	err := h.DriverService.DeleteDriver(ctx, req.GetDriverId())
	if err != nil {
		if errors.Is(err, service.ErrInvalidCredentials) {
			return nil, status.Error(codes.NotFound, "driver not found")
		}
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &putlistv1.DeleteDriverResponse{
		Message: "deleted",
	}, nil
}

func validateCreateDriver(req *putlistv1.CreateDriverRequest) error {
	if req.GetFullName() == "" {
		return status.Error(codes.InvalidArgument, "fullname is required")
	}
	if req.GetLicense() == "" {
		return status.Error(codes.InvalidArgument, "license is required")
	}
	if req.GetClass() == "" {
		return status.Error(codes.InvalidArgument, "class is required")
	}
	if len(req.GetLicense()) != 10 {
		return status.Error(codes.InvalidArgument, "invalid field format: license")
	}
	return nil
}

func validateUpdateDriver(req *putlistv1.UpdateDriverRequest) error {
	if req.GetDriverId() == 0 {
		return status.Error(codes.InvalidArgument, "driver ID is required")
	}
	if req.FullName == nil && req.License == nil && req.Class == nil {
		return status.Error(codes.InvalidArgument, "updating structure has no values")
	}
	if req.License != nil {
		if *req.License == "" {
			return status.Error(codes.InvalidArgument, "updating structure has empty values")
		}
		if len(*req.License) != 10 {
			return status.Error(codes.InvalidArgument, "invalid field format: license")
		}
	}
	if req.FullName != nil {
		if *req.FullName == "" {
			return status.Error(codes.InvalidArgument, "updating structure has empty values")
		}
	}
	if req.Class != nil {
		if *req.Class == "" {
			return status.Error(codes.InvalidArgument, "updating structure has empty values")
		}
	}
	return nil
}

func validateDeleteDriver(req *putlistv1.DeleteDriverRequest) error {
	if req.GetDriverId() == 0 {
		return status.Error(codes.InvalidArgument, "driver ID is required")
	}
	return nil
}
