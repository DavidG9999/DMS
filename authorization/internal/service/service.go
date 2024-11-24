package service

import (
	"log/slog"
	"time"

	users_grpc "github.com/DavidG9999/DMS/authorization/internal/clients/users/grpc"
)

type Auth interface {
	Register
	Login
}

type Service struct {
	Auth
}

func NewService(logger *slog.Logger, userClient users_grpc.UserClient, tokenTTL time.Duration) *Service {
	return &Service{
		Auth: NewAuthService(logger, userClient, tokenTTL),
	}
}
