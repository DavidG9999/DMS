package main

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/DavidG9999/DMS/authorization/internal/app"
	users_grpc "github.com/DavidG9999/DMS/authorization/internal/clients/users/grpc"
	"github.com/DavidG9999/DMS/authorization/internal/config"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()

	logger := setupLogger(cfg.Env)

	logger.Info("starting application")

	userClCgf := users_grpc.Config{
		Addr:         cfg.UserClient.Address,
		Timeout:      cfg.UserClient.Timeout,
		RetriesCount: cfg.UserClient.RetriesCount,
	}

	application := app.NewApp(logger, cfg.GRPC.Port, userClCgf, cfg.TokenTTL)//что то тут не так, попробуй переписать по друглмк 

	go application.GRPCServer.MustRun()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	stopSignal := <-quit

	logger.Info("stopping application", slog.String("signal: ", stopSignal.String()))

	application.GRPCServer.Stop()

	logger.Info("application stopped")
}

func setupLogger(env string) *slog.Logger {
	var logger *slog.Logger

	switch env {
	case envLocal:
		logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envDev:
		logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}
	return logger
}
