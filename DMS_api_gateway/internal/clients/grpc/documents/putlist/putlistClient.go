package putlistgrpc

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	putlistv1 "github.com/DavidG9999/DMS/api/grpc/document_api/putlist/gen/go"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/retry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
)

type PutlistClient struct {
	apiAuto         putlistv1.AutoClient
	apiBankAccount  putlistv1.BankAccountClient
	apiContragent   putlistv1.ContragentClient
	apiDriver       putlistv1.DriverClient
	apiDispetcher   putlistv1.DispetcherClient
	apiMehanic      putlistv1.MehanicClient
	apiOrganization putlistv1.OrganizationClient
	apiPutlist      putlistv1.PutlistClient
	apiPutlistBody  putlistv1.PutlistBodyClient
}

type Config struct {
	Addr         string
	Timeout      time.Duration
	RetriesCount int
}

func NewClient(ctx context.Context, logger *slog.Logger, cfg Config) (*PutlistClient, error) {
	const op = "grpc.NewPutlistClient"

	retryOpts := []retry.CallOption{
		retry.WithCodes(codes.NotFound, codes.Aborted, codes.DeadlineExceeded),
		retry.WithMax(uint(cfg.RetriesCount)),
		retry.WithPerRetryTimeout(cfg.Timeout),
	}

	logOpts := []logging.Option{
		logging.WithLogOnEvents(logging.PayloadReceived, logging.PayloadSent),
	}

	cc, err := grpc.NewClient(cfg.Addr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithChainUnaryInterceptor(logging.UnaryClientInterceptor(InterceptorLogger(logger), logOpts...), retry.UnaryClientInterceptor(retryOpts...)))
	if err != nil {
		return nil, fmt.Errorf("%s: %v", op, err)
	}
	return &PutlistClient{
		apiAuto:         putlistv1.NewAutoClient(cc),
		apiBankAccount:  putlistv1.NewBankAccountClient(cc),
		apiContragent:   putlistv1.NewContragentClient(cc),
		apiDriver:       putlistv1.NewDriverClient(cc),
		apiDispetcher:   putlistv1.NewDispetcherClient(cc),
		apiMehanic:      putlistv1.NewMehanicClient(cc),
		apiOrganization: putlistv1.NewOrganizationClient(cc),
		apiPutlist:      putlistv1.NewPutlistClient(cc),
		apiPutlistBody:  putlistv1.NewPutlistBodyClient(cc),
	}, nil
}

func InterceptorLogger(logger *slog.Logger) logging.Logger {
	return logging.LoggerFunc(func(ctx context.Context, level logging.Level, msg string, fields ...any) {
		logger.Log(ctx, slog.Level(level), msg, fields...)
	})
}
