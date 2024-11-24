package putlistgrpc

import (
		"github.com/DavidG9999/DMS/DMS_api_gateway/internal/entity"
	"context"
	"fmt"

	putlistv1 "github.com/DavidG9999/DMS/api/grpc/document_api/putlist/gen/go"
)

func (pc *PutlistClient) CreatePutlistBody(ctx context.Context, putlistNumber, number, contragentId int64, item, timeWith, timeFor string) (int64, error) {
	const op = "grpc.ClientCreatePutlistBody"

	resp, err := pc.apiPutlistBody.CreatePutlistBody(ctx, &putlistv1.CreatePutlistBodyRequest{PutlistNumber: putlistNumber, Number: number, ContragentId: contragentId, Item: item, TimeWith: timeWith, TimeFor: timeFor})
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return resp.GetPutlistBodyId(), nil
}

func (pc *PutlistClient) GetPutlistBodies(ctx context.Context, putlistNumber int64) ([]entity.PutlistBody, error) {
	const op = "grpc.ClientGetOrganizations"

	resp, err := pc.apiPutlistBody.GetPutlistBodies(ctx, &putlistv1.GetPutlistBodiesRequest{PutlistNumber: putlistNumber})
	if err != nil {
		return []entity.PutlistBody{}, fmt.Errorf("%s: %w", op, err)
	}

	putlistBodies := make([]entity.PutlistBody, len(resp.PutlistBodies))
	for id, putlistBody := range resp.GetPutlistBodies() {
		putlistBodies[id].Id = putlistBody.PutlistBodyId
		putlistBodies[id].PutlistNumber = putlistBody.PutlistNumber
		putlistBodies[id].Number = putlistBody.Number
		putlistBodies[id].ContragentId = putlistBody.ContragentId
		putlistBodies[id].Item = putlistBody.Item
		putlistBodies[id].TimeWith = putlistBody.TimeWith
		putlistBodies[id].TimeFor = putlistBody.TimeFor
	}
	return putlistBodies, nil
}

func (pc *PutlistClient) UpdatePutlistBody(ctx context.Context, putlistBodyId int64,  number, contragentId *int64, item, timeWith, timeFor *string) (string, error) {
	const op = "grpc.ClientUpdatePutlistBody"

	resp, err := pc.apiPutlistBody.UpdatePutlistBody(ctx, &putlistv1.UpdatePutlistBodyRequest{PutlistBodyId: putlistBodyId, Number: number, ContragentId: contragentId, Item: item, TimeWith: timeWith, TimeFor: timeFor})
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return resp.GetMessage(), nil
}

func (pc *PutlistClient) DeletePutlistBody(ctx context.Context, putlistBodyId int64) (string, error) {
	const op = "grpc.ClientDeletePutlistBody"

	resp, err := pc.apiPutlistBody.DeletePutlistBody(ctx, &putlistv1.DeletePutlistBodyRequest{PutlistBodyId: putlistBodyId})
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return resp.GetMessage(), nil
}
