package putlistgrpc

import (
	"context"
	"fmt"

	"github.com/DavidG9999/DMS/DMS_api_gateway/internal/entity"
	putlistv1 "github.com/DavidG9999/DMS/api/grpc/document_api/putlist/gen/go"
)

func (pc *PutlistClient) CreateDispetcher(ctx context.Context, fullname string) (int64, error) {
	const op = "grpc.ClientCreateDispetcher"

	resp, err := pc.apiDispetcher.CreateDispetcher(ctx, &putlistv1.CreateDispetcherRequest{FullName: fullname})
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return resp.GetDispetcherId(), nil
}

func (pc *PutlistClient) GetDispetchers(ctx context.Context) ([]entity.Dispetcher, error) {
	const op = "grpc.ClientGetDispetchers"

	resp, err := pc.apiDispetcher.GetDispetchers(ctx, &putlistv1.GetDispetchersRequest{})
	if err != nil {
		return []entity.Dispetcher{}, fmt.Errorf("%s: %w", op, err)
	}

	dispetchers := make([]entity.Dispetcher, len(resp.Dispetchers))
	for id, dispetcher := range resp.GetDispetchers() {
		dispetchers[id].Id = dispetcher.DispetcherId
		dispetchers[id].FullName = dispetcher.FullName
	}
	return dispetchers, nil
}

func (pc *PutlistClient) UpdateDispetcher(ctx context.Context, dispetcherId int64, fullname *string) (string, error) {
	const op = "grpc.ClientUpdateDispetcher"

	resp, err := pc.apiDispetcher.UpdateDispetcher(ctx, &putlistv1.UpdateDispetcherRequest{DispetcherId: dispetcherId, FullName: fullname})
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return resp.GetMessage(), nil
}

func (pc *PutlistClient) DeleteDispetcher(ctx context.Context, dispetcherId int64) (string, error) {
	const op = "grpc.ClientDeleteDispetcher"

	resp, err := pc.apiDispetcher.DeleteDispetcher(ctx, &putlistv1.DeleteDispetcherRequest{DispetcherId: dispetcherId})
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return resp.GetMessage(), nil
}
