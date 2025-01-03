package main

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/DavidG9999/DMS/documents/putlists/internal/app"
	"github.com/DavidG9999/DMS/documents/putlists/internal/config"
	"github.com/DavidG9999/DMS/documents/putlists/internal/logger"
	"github.com/DavidG9999/DMS/documents/putlists/internal/repository/postgres"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	cfg := config.MustLoad()

	logger := logger.SetupLogger(cfg.Env, cfg.LogPath)

	logger.Info("starting application")

	psgdbCfg := postgres.Config{
		Host:     cfg.DB.Host,
		Port:     cfg.DB.Port,
		Username: cfg.DB.Username,
		Password: os.Getenv("PASSWORD_DB"),
		DBName:   cfg.DB.DBName,
		SSLMode:  cfg.DB.SSLMode,
	}

	application := app.NewApp(logger, cfg.GRPC.Port, psgdbCfg)

	go application.GRPCServer.MustRun()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	stopSignal := <-quit

	logger.Info("stopping application", slog.String("signal: ", stopSignal.String()))

	application.GRPCServer.Stop()

	logger.Info("application stopped")
}
