package app

import (
	"log/slog"

	grpcserver "github.com/DavidG9999/DMS/documents/putlists/internal/app/grpc"
	"github.com/DavidG9999/DMS/documents/putlists/internal/repository/postgres"
	"github.com/DavidG9999/DMS/documents/putlists/internal/service"
)

type App struct {
	GRPCServer *grpcserver.GRPCServer
}

func NewApp(logger *slog.Logger, gPRCPort int, dbCfg postgres.Config) *App {

	db, err := postgres.NewPostgresDB(dbCfg)
	if err != nil {
		panic(err)
	}

	repositories := postgres.NewRepository(db)

	services := service.NewService(logger, repositories)

	gRPCServer := grpcserver.NewGRPCServer(logger, gPRCPort, services)

	return &App{
		GRPCServer: gRPCServer,
	}
}
