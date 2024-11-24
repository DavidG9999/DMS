package grpcserver

import (
	"fmt"
	"log/slog"
	"net"

	userv1 "github.com/DavidG9999/DMS/api/grpc/user_api/gen/go"
	handler "github.com/DavidG9999/DMS/users/internal/handler/grpc"
	"github.com/DavidG9999/DMS/users/internal/service"

	"google.golang.org/grpc"
)

type GRPCServer struct {
	logger     *slog.Logger
	gRPCServer *grpc.Server
	port       int
}

func NewGRPCServer(logger *slog.Logger, port int, service *service.Service) *GRPCServer {
	gRPCServer := grpc.NewServer()
	userv1.RegisterUserServer(gRPCServer, &handler.Handler{Service: *service})

	return &GRPCServer{
		logger:     logger,
		gRPCServer: gRPCServer,
		port:       port,
	}
}

func (srv *GRPCServer) MustRun() {
	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}

func (srv *GRPCServer) ListenAndServe() error {
	const op = "grpcserver.Run"

	logger := srv.logger.With(slog.String("op", op), slog.Int("port", srv.port))

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", srv.port))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	logger.Info("grpc server is running", slog.String("addr", listener.Addr().String()))

	if err := srv.gRPCServer.Serve(listener); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (srv *GRPCServer) Stop() {
	const op = "grpcserver.Stop"

	logger := srv.logger.With(slog.String("op", op))

	srv.gRPCServer.GracefulStop()

	logger.Info("stopping grpc server")
}
