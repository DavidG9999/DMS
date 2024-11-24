package putlistgrpc

import (
	"context"
	"fmt"

	"github.com/DavidG9999/DMS/DMS_api_gateway/internal/entity"
	putlistv1 "github.com/DavidG9999/DMS/api/grpc/document_api/putlist/gen/go"
)

func (pc *PutlistClient) CreateDriver(ctx context.Context, fullname, license, class string) (int64, error) {
	const op = "grpc.ClientCreateDriver"

	resp, err := pc.apiDriver.CreateDriver(ctx, &putlistv1.CreateDriverRequest{FullName: fullname, License: license, Class: class})
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return resp.GetDriverId(), nil
}

func (pc *PutlistClient) GetDrivers(ctx context.Context) ([]entity.Driver, error) {
	const op = "grpc.ClientGetDrivers"

	resp, err := pc.apiDriver.GetDrivers(ctx, &putlistv1.GetDriversRequest{})
	if err != nil {
		return []entity.Driver{}, fmt.Errorf("%s: %w", op, err)
	}

	drivers := make([]entity.Driver, len(resp.Drivers))
	for id, driver := range resp.GetDrivers() {
		drivers[id].Id = driver.DriverId
		drivers[id].FullName = driver.FullName
		drivers[id].License = driver.License
		drivers[id].Class = driver.Class
	}
	return drivers, nil
}

func (pc *PutlistClient) UpdateDriver(ctx context.Context, driverId int64, fullname, license, class *string) (string, error) {
	const op = "grpc.ClientUpdateDriver"

	resp, err := pc.apiDriver.UpdateDriver(ctx, &putlistv1.UpdateDriverRequest{DriverId: driverId, FullName: fullname, License: license, Class: class})
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return resp.GetMessage(), nil
}

func (pc *PutlistClient) DeleteDriver(ctx context.Context, driverId int64) (string, error) {
	const op = "grpc.ClientDeleteDriver"

	resp, err := pc.apiDriver.DeleteDriver(ctx, &putlistv1.DeleteDriverRequest{DriverId: driverId})
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return resp.GetMessage(), nil
}
