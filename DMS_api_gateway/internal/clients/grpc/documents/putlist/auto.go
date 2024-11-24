package putlistgrpc

import (
	"context"
	"fmt"

	"github.com/DavidG9999/DMS/DMS_api_gateway/internal/entity"
	putlistv1 "github.com/DavidG9999/DMS/api/grpc/document_api/putlist/gen/go"
)

func (pc *PutlistClient) CreateAuto(ctx context.Context, brand, model, stateNumber string) (autoId int64, err error) {
	const op = "grpc.ClientCreateAuto"

	resp, err := pc.apiAuto.CreateAuto(ctx, &putlistv1.CreateAutoRequest{Brand: brand, Model: model, StateNumber: stateNumber})
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return resp.GetAutoId(), nil
}

func (pc *PutlistClient) GetAutos(ctx context.Context) ([]entity.Auto, error) {
	const op = "grpc.ClientGetAutos"

	resp, err := pc.apiAuto.GetAutos(ctx, &putlistv1.GetAutosRequest{})
	if err != nil {
		return []entity.Auto{}, fmt.Errorf("%s: %w", op, err)
	}

	autos := make([]entity.Auto, len(resp.Autos))
	for id, auto := range resp.GetAutos() {
		autos[id].Id = auto.AutoId
		autos[id].Brand = auto.Brand
		autos[id].Model = auto.Model
		autos[id].StateNumber = auto.StateNumber
	}
	return autos, nil
}

func (pc *PutlistClient) UpdateAuto(ctx context.Context, autoId int64, brand, model, stateNumber *string) (string, error) {
	const op = "grpc.ClientUpdateAuto"


	resp, err := pc.apiAuto.UpdateAuto(ctx, &putlistv1.UpdateAutoRequest{AutoId: autoId, Brand: brand, Model: model, StateNumber: stateNumber})
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return resp.GetMessage(), nil
}

func (pc *PutlistClient) DeleteAuto(ctx context.Context, autoId int64) (string, error) {
	const op = "grpc.ClientDeleteAuto"

	resp, err := pc.apiAuto.DeleteAuto(ctx, &putlistv1.DeleteAutoRequest{AutoId: autoId})
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return resp.GetMessage(), nil
}
