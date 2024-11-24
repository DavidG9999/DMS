package service

import (
	putlistgrpc "github.com/DavidG9999/DMS/DMS_api_gateway/internal/clients/grpc/documents/putlist"
	"log/slog"
)

type PutlistService struct {
	logger        *slog.Logger
	putlistClient putlistgrpc.PutlistClient
}

func NewPutlistService(logger *slog.Logger, putlistClient putlistgrpc.PutlistClient) *PutlistService {
	return &PutlistService{
		logger:        logger,
		putlistClient: putlistClient,
	}
}
