package main

import (
	"github.com/DavidG9999/DMS/DMS_api_gateway/internal/app"
	authgrpc "github.com/DavidG9999/DMS/DMS_api_gateway/internal/clients/grpc/auth"
	putlistgrpc "github.com/DavidG9999/DMS/DMS_api_gateway/internal/clients/grpc/documents/putlist"
	usergrpc "github.com/DavidG9999/DMS/DMS_api_gateway/internal/clients/grpc/user"
	"github.com/DavidG9999/DMS/DMS_api_gateway/internal/config"
	"github.com/DavidG9999/DMS/DMS_api_gateway/internal/logger"
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/viper"
)

// @title DMS API
// @version 1.0
// @description API Server for Document Management Service

// @host localhost:8000
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {

	config.InitConfig()

	logger := logger.SetupLogger(viper.GetString("env"))

	logger.Info("starting application")

	authCfg := authgrpc.Config{
		Addr:         viper.GetString("auth_client.address"),
		Timeout:      viper.GetDuration("auth_client.timeout"),
		RetriesCount: viper.GetInt("auth_client.retries_count"),
	}

	userCfg := usergrpc.Config{
		Addr:         viper.GetString("user_client.address"),
		Timeout:      viper.GetDuration("user_client.timeout"),
		RetriesCount: viper.GetInt("user_client.retries_count"),
	}

	putlistCfg := putlistgrpc.Config{
		Addr:         viper.GetString("putlist_client.address"),
		Timeout:      viper.GetDuration("putlist_client.timeout"),
		RetriesCount: viper.GetInt("putlist_client.retries_count"),
	}

	application := app.NewApp(logger, authCfg, userCfg, putlistCfg)

	go application.HTTPServer.Run(viper.GetString("port"), application.Handler.InitRoutes())

	logger.Info("HTTP Server is running", slog.String("port", viper.GetString("port")))

	logger.Info("App started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	stopSignal := <-quit

	logger.Info("App shutting down", slog.String("signal: ", stopSignal.String()))

	application.HTTPServer.Shutdown(context.Background())

	logger.Info("stopping HTTP server")

	logger.Info("App stopped")

}
