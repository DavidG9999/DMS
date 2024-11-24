package app

import (
	httpserver "github.com/DavidG9999/DMS/DMS_api_gateway/internal/app/http"
	authgrpc "github.com/DavidG9999/DMS/DMS_api_gateway/internal/clients/grpc/auth"
	putlistgrpc "github.com/DavidG9999/DMS/DMS_api_gateway/internal/clients/grpc/documents/putlist"
	usergrpc "github.com/DavidG9999/DMS/DMS_api_gateway/internal/clients/grpc/user"
	"github.com/DavidG9999/DMS/DMS_api_gateway/internal/handler"
	"github.com/DavidG9999/DMS/DMS_api_gateway/internal/service"
	"context"
	"log/slog"
)

type App struct {
	HTTPServer *httpserver.Server
	Handler    handler.Handler
}

func NewApp(logger *slog.Logger, authCfg authgrpc.Config, userCfg usergrpc.Config, putlistCfg putlistgrpc.Config) *App {

	authClient, err := authgrpc.NewClient(context.Background(), logger, authCfg)
	if err != nil {
		panic(err)
	}
	userClient, err := usergrpc.NewClient(context.Background(), logger, userCfg)
	if err != nil {
		panic(err)
	}
	putlistClient, err := putlistgrpc.NewClient(context.Background(), logger, putlistCfg)
	if err != nil {
		panic(err)
	}
	services := service.NewService(logger, *authClient, *userClient, *putlistClient)

	handlers := handler.NewHandler(services)
	httpsrv := new(httpserver.Server)

	return &App{
		HTTPServer: httpsrv,
		Handler:    *handlers,
	}
}
