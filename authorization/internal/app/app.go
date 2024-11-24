package app

import (
	"context"
	"log/slog"
	"time"

	grpcserver "github.com/DavidG9999/DMS/authorization/internal/app/grpc"
	users_grpc "github.com/DavidG9999/DMS/authorization/internal/clients/users/grpc"
	"github.com/DavidG9999/DMS/authorization/internal/service"
)

type App struct {
	GRPCServer *grpcserver.GRPCServer
}

func NewApp(logger *slog.Logger, gRPCPort int, userClientCfg users_grpc.Config, tokenTTL time.Duration) *App {

	userClient, err := users_grpc.NewClient(context.Background(), logger, userClientCfg)
	if err != nil {
		panic(err)
	}

	services := service.NewService(logger, *userClient, tokenTTL)
	gRPCserver := grpcserver.NewGRPCServer(logger, gRPCPort, services)

	return &App{
		GRPCServer: gRPCserver,
	}
}
