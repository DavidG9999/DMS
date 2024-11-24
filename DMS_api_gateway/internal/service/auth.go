package service

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	authgrpc "github.com/DavidG9999/DMS/DMS_api_gateway/internal/clients/grpc/auth"
	parsejwt "github.com/DavidG9999/DMS/DMS_api_gateway/internal/lib/jwt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthService struct {
	logger     *slog.Logger
	authClient authgrpc.AuthClient
}

func NewAuthService(logger *slog.Logger, authClient authgrpc.AuthClient) *AuthService {
	return &AuthService{
		logger:     logger,
		authClient: authClient,
	}
}

type Register interface {
	SignUp(ctx context.Context, name string, email string, password string) (userID int64, err error)
}

type Login interface {
	SignIn(ctx context.Context, email string, password string) (token string, err error)
	ParseToken(ctx context.Context, accessToken string) (int64, error)
}

func (a *AuthService) SignUp(ctx context.Context, name string, email string, password string) (userID int64, err error) {
	const op = "service.SignUp"

	logger := a.logger.With(slog.String("op", op))
	logger.Info("registering user")

	userID, err = a.authClient.SignUp(ctx, name, email, password)
	if err != nil {
		if errors.Is(err, status.Error(codes.AlreadyExists, "user already exists")) {
			logger.Error("user already exists")
			return 0, fmt.Errorf("%s: %w", op, ErrUserExist)
		}
		logger.Error("failed to save user")
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	logger.Info("user registered")
	return userID, nil
}

func (a *AuthService) SignIn(ctx context.Context, email string, password string) (token string, err error) {
	const op = "service.SignIn"

	logger := a.logger.With(slog.String("op", op))
	logger.Info("logining user")

	token, err = a.authClient.SignIn(ctx, email, password)
	if err != nil {
		if errors.Is(err, status.Error(codes.NotFound, "user not found")) {
			logger.Error("user not found")
			return "", fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
		}
		logger.Error("failed to get user")
		return "", fmt.Errorf("%s: %w", op, err)
	}
	return token, nil
}

func (a *AuthService) ParseToken(ctx context.Context, accessToken string) (int64, error) {
	const op = "service.ParseToken"

	logger := a.logger.With(slog.String("op", op))
	logger.Info("parsing token")

	userId, err := parsejwt.ParseToken(ctx, accessToken)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return userId, nil
}
