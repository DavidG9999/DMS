package app

import (
	"log/slog"

	grpcserver "github.com/DavidG9999/DMS/users/internal/app/grpc"
	"github.com/DavidG9999/DMS/users/internal/repository/postgres"
	"github.com/DavidG9999/DMS/users/internal/service"
)

type App struct {
	GRPCServer *grpcserver.GRPCServer
}

func NewApp(logger *slog.Logger, gRPCPort int, dbCfg postgres.Config) *App {

	db, err := postgres.NewPostgresDB(dbCfg)
	if err != nil {
		panic(err)
	}

	repositories := postgres.NewRepository(db)

	services := service.NewService(logger, *repositories)

	gRPCServer := grpcserver.NewGRPCServer(logger, gRPCPort, services)

	return &App{
		GRPCServer: gRPCServer,
	}
}
