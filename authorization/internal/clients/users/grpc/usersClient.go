package users_grpc

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	userv1 "github.com/DavidG9999/DMS/api/grpc/user_api/gen/go"
	"github.com/DavidG9999/DMS/authorization/internal/domain/entity"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/retry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
)

type UserClient struct {
	api userv1.UserClient
}

type Config struct {
	Addr         string
	Timeout      time.Duration
	RetriesCount int
}

func NewClient(ctx context.Context, logger *slog.Logger, cfg Config) (*UserClient, error) {
	const op = "grpc.NewUserClient"

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
	return &UserClient{
		api: userv1.NewUserClient(cc),
	}, nil
}

func (us *UserClient) CreateUser(ctx context.Context, name, email string, passwordHash string) (int64, error) {

	response, err := us.api.CreateUser(ctx, &userv1.CreateUserRequest{
		Name:         name,
		Email:        email,
		PasswordHash: passwordHash,
	})
	if err != nil {
		return 0, err
	}

	return response.GetUserId(), nil
}

func (us *UserClient) GetUser(ctx context.Context, email string) (entity.User, error) {

	response, err := us.api.GetUser(ctx, &userv1.GetUserRequest{
		Email: email,
	})
	if err != nil {
		return entity.User{}, err
	}

	user := entity.User{
		ID:           response.GetUserId(),
		Name:         response.GetName(),
		Email:        response.GetEmail(),
		PasswordHash: []byte(response.GetPasswordHash()),
	}

	return user, nil
}

func InterceptorLogger(logger *slog.Logger) logging.Logger {
	return logging.LoggerFunc(func(ctx context.Context, level logging.Level, msg string, fields ...any) {
		logger.Log(ctx, slog.Level(level), msg, fields...)
	})
}
