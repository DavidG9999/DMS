package putlistgrpc

import (
	"context"
	"fmt"

	"github.com/DavidG9999/DMS/DMS_api_gateway/internal/entity"
	putlistv1 "github.com/DavidG9999/DMS/api/grpc/document_api/putlist/gen/go"
)

func (pc *PutlistClient) CreateMehanic(ctx context.Context, fullname string) (int64, error) {
	const op = "grpc.ClientCreateMehanic"

	resp, err := pc.apiMehanic.CreateMehanic(ctx, &putlistv1.CreateMehanicRequest{FullName: fullname})
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	return resp.GetMehanicId(), nil
}

func (pc *PutlistClient) GetMehanics(ctx context.Context) ([]entity.Mehanic, error) {
	const op = "grpc.ClientGetMehanics"

	resp, err := pc.apiMehanic.GetMehanics(ctx, &putlistv1.GetMehanicsRequest{})
	if err != nil {
		return []entity.Mehanic{}, fmt.Errorf("%s: %w", op, err)
	}

	mehanics := make([]entity.Mehanic, len(resp.Mehanics))
	for id, mehanic := range resp.GetMehanics() {
		mehanics[id].Id = mehanic.MehanicId
		mehanics[id].FullName = mehanic.FullName
	}
	return mehanics, nil
}

func (pc *PutlistClient) UpdateMehanic(ctx context.Context, mehanicId int64, fullname *string) (string, error) {
	const op = "grpc.ClientUpdateMehanic"

	resp, err := pc.apiMehanic.UpdateMehanic(ctx, &putlistv1.UpdateMehanicRequest{MehanicId: mehanicId, FullName: fullname})
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return resp.GetMessage(), nil
}

func (pc *PutlistClient) DeleteMehanic(ctx context.Context, mehanicId int64) (string, error) {
	const op = "grpc.ClientDeleteMehanic"

	resp, err := pc.apiMehanic.DeleteMehanic(ctx, &putlistv1.DeleteMehanicRequest{MehanicId: mehanicId})
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return resp.GetMessage(), nil
}
