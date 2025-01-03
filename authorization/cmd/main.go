package main

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/DavidG9999/DMS/authorization/internal/app"
	users_grpc "github.com/DavidG9999/DMS/authorization/internal/clients/users/grpc"
	"github.com/DavidG9999/DMS/authorization/internal/config"
	"github.com/DavidG9999/DMS/authorization/internal/logger"
)

func main() {
	cfg := config.MustLoad()

	logger := logger.SetupLogger(cfg.Env, cfg.LogPath)

	logger.Info("starting application")

	userClCgf := users_grpc.Config{
		Addr:         cfg.UserClient.Address,
		Timeout:      cfg.UserClient.Timeout,
		RetriesCount: cfg.UserClient.RetriesCount,
	}

	application := app.NewApp(logger, cfg.GRPC.Port, userClCgf, cfg.TokenTTL) //что то тут не так, попробуй переписать по друглмк

	go application.GRPCServer.MustRun()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	stopSignal := <-quit

	logger.Info("stopping application", slog.String("signal: ", stopSignal.String()))

	application.GRPCServer.Stop()

	logger.Info("application stopped")
}
