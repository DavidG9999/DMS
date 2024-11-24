package usergrpc

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	userv1 "github.com/DavidG9999/DMS/api/grpc/user_api/gen/go"
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

func InterceptorLogger(logger *slog.Logger) logging.Logger {
	return logging.LoggerFunc(func(ctx context.Context, level logging.Level, msg string, fields ...any) {
		logger.Log(ctx, slog.Level(level), msg, fields...)
	})
}

func (uc *UserClient) GetUserById(ctx context.Context, userId int64) (name string, email string, err error) {
	const op = "grpc.ClientGetUserById"

	resp, err := uc.api.GetUserById(ctx, &userv1.GetUserByIdRequest{UserId: userId})
	if err != nil {
		return "", "", fmt.Errorf("%s: %w", op, err)
	}

	return resp.GetName(), resp.GetEmail(), nil
}

func (uc *UserClient) UpdateName(ctx context.Context, userId int64, updateName string) (string, error) {
	const op = "grpc.ClientUpdateName"

	resp, err := uc.api.UpdateName(ctx, &userv1.UpdateNameRequest{UserId: userId, UpdateName: updateName})
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return resp.GetMessage(), nil
}

func (uc *UserClient) UpdatePassword(ctx context.Context, userId int64, updatePassword string) (string, error) {
	const op = "grpc.ClientUpdatePassword"

	resp, err := uc.api.UpdatePassword(ctx, &userv1.UpdatePasswordRequest{UserId: userId, UpdatePassword: updatePassword})
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return resp.GetMessage(), nil
}

func (uc *UserClient) DeleteUser(ctx context.Context, userId int64) (string, error) {
	const op = "grpc.ClientDeleteUser"

	resp, err := uc.api.DeleteUser(ctx, &userv1.DeleteUserRequest{UserId: userId})
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return resp.GetMessage(), nil
}
