package putlistgrpc

import (
	"github.com/DavidG9999/DMS/DMS_api_gateway/internal/entity"
	"context"
	"fmt"

	putlistv1 "github.com/DavidG9999/DMS/api/grpc/document_api/putlist/gen/go"
)

func (pc *PutlistClient) CreateContragent(ctx context.Context, name, address, innKpp string) (int64, error) {
	const op = "grpc.ClientCreateContragent"

	resp, err := pc.apiContragent.CreateContragent(ctx, &putlistv1.CreateContragentRequest{Name: name, Address: address, InnKpp: innKpp})
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return resp.GetContragentId(), nil
}

func (pc *PutlistClient) GetContragents(ctx context.Context) ([]entity.Contragent, error) {
	const op = "grpc.ClientGetContragents"

	resp, err := pc.apiContragent.GetContragents(ctx, &putlistv1.GetContragentsRequest{})
	if err != nil {
		return []entity.Contragent{}, fmt.Errorf("%s: %w", op, err)
	}

	contragents := make([]entity.Contragent, len(resp.Contragents))
	for id, contragent := range resp.GetContragents() {
		contragents[id].Id = contragent.ContragentId
		contragents[id].Name = contragent.Name
		contragents[id].Address = contragent.Address
		contragents[id].InnKpp = contragent.InnKpp
	}
	return contragents, nil
}

func (pc *PutlistClient) UpdateContragent(ctx context.Context, contragentId int64, name, address, innKpp *string) (string, error) {
	const op = "grpc.ClientUpdateContragent"

	resp, err := pc.apiContragent.UpdateContragent(ctx, &putlistv1.UpdateContragentRequest{ContragentId: contragentId, Name: name, Address: address, InnKpp: innKpp})
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return resp.GetMessage(), nil
}

func (pc *PutlistClient) DeleteContragent(ctx context.Context, contragentId int64) (string, error) {
	const op = "grpc.ClientDeleteContragent"

	resp, err := pc.apiContragent.DeleteContragent(ctx, &putlistv1.DeleteContragentRequest{ContragentId: contragentId})
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return resp.GetMessage(), nil
}
