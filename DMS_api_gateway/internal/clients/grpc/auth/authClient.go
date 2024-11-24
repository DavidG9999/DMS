package authgrpc

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	authv1 "github.com/DavidG9999/DMS/api/grpc/auth_api/gen/go"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/retry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
)

type AuthClient struct {
	api authv1.AuthClient
}

type Config struct {
	Addr         string
	Timeout      time.Duration
	RetriesCount int
}

func NewClient(ctx context.Context, logger *slog.Logger, cfg Config) (*AuthClient, error) {
	const op = "grpc.NewAuthClient"

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
	return &AuthClient{
		api: authv1.NewAuthClient(cc),
	}, nil
}

func InterceptorLogger(logger *slog.Logger) logging.Logger {
	return logging.LoggerFunc(func(ctx context.Context, level logging.Level, msg string, fields ...any) {
		logger.Log(ctx, slog.Level(level), msg, fields...)
	})
}

func (ac *AuthClient) SignUp(ctx context.Context, name, email, password string) (int64, error) {
	const op = "grpc.ClientSignUp"

	resp, err := ac.api.SignUp(ctx, &authv1.SignUpRequest{Name: name, Email: email, Password: password})
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return resp.GetUserId(), nil
}

func (ac *AuthClient) SignIn(ctx context.Context, email, password string) (string, error) {
	const op = "grpc.ClientSignIn"

	resp, err := ac.api.SignIn(ctx, &authv1.SignInRequest{Email: email, Password: password})
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return resp.GetToken(), nil
}
