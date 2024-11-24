package grpcserver

import (
	"fmt"
	"log/slog"
	"net"

	putlistv1 "github.com/DavidG9999/DMS/api/grpc/document_api/putlist/gen/go"
	handler "github.com/DavidG9999/DMS/documents/putlists/internal/handler/grpc"
	"github.com/DavidG9999/DMS/documents/putlists/internal/service"
	"google.golang.org/grpc"
)

type GRPCServer struct {
	logger     *slog.Logger
	gRPCServer *grpc.Server
	port       int
}

func NewGRPCServer(logger *slog.Logger, port int, service *service.Service) *GRPCServer {
	gRPCServer := grpc.NewServer()
	putlistv1.RegisterAutoServer(gRPCServer, &handler.AutoHandler{AutoService: service.Auto})
	putlistv1.RegisterBankAccountServer(gRPCServer, &handler.BankAccountHandler{BankAccountService: service.BankAccount})
	putlistv1.RegisterContragentServer(gRPCServer, &handler.ContragentHandler{ContragentService: service.Contragent})
	putlistv1.RegisterDispetcherServer(gRPCServer, &handler.DispetcherHandler{DispetcherService: service.Dispetcher})
	putlistv1.RegisterDriverServer(gRPCServer, &handler.DriverHandler{DriverService: service.Driver})
	putlistv1.RegisterMehanicServer(gRPCServer, &handler.MehanicHandler{MehanicService: service.Mehanic})
	putlistv1.RegisterOrganizationServer(gRPCServer, &handler.OrganizationHandler{OrganizationService: service.Organization})
	putlistv1.RegisterPutlistServer(gRPCServer, &handler.PutlistHandler{PutlistService: service.Putlist})
	putlistv1.RegisterPutlistBodyServer(gRPCServer, &handler.PutlistBodyHandler{PutlistBodyService: service.PutlistBody})
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

	logger := srv.logger.With(slog.String("op ", op), slog.Int("port ", srv.port))

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

	logger := srv.logger.With(slog.String("op ", op))

	srv.gRPCServer.GracefulStop()

	logger.Info("stopping grpc server")
}
