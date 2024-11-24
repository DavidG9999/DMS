package service

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

	users_grpc "github.com/DavidG9999/DMS/authorization/internal/clients/users/grpc"
	genjwt "github.com/DavidG9999/DMS/authorization/internal/lib/jwt"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthService struct {
	logger     *slog.Logger
	userClient users_grpc.UserClient
	tokenTTL   time.Duration
}

func NewAuthService(logger *slog.Logger, userClient users_grpc.UserClient, tokenTTL time.Duration) *AuthService {
	return &AuthService{
		logger:     logger,
		userClient: userClient,
		tokenTTL:   tokenTTL,
	}
}

type Register interface {
	SignUp(ctx context.Context, name string, email string, password string) (userID int64, err error)
}

type Login interface {
	SignIn(ctx context.Context, email string, password string) (token string, err error)
}

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUserExist          = errors.New("user already exists")
)

func (a *AuthService) SignUp(ctx context.Context, name string, email string, password string) (userID int64, err error) {
	const op = "service.SignUp"

	logger := a.logger.With(slog.String("op", op))
	logger.Info("registering user")

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		logger.Error("failed to generate pasword hash")
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	userID, err = a.userClient.CreateUser(ctx, name, email, string(passwordHash))
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

	user, err := a.userClient.GetUser(ctx, email)
	if err != nil {
		if errors.Is(err, status.Error(codes.NotFound, "user not found")) {
			return "", fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
		}
		logger.Error("failed to get user")
		return "", fmt.Errorf("%s: %w", op, err)
	}
	if err := bcrypt.CompareHashAndPassword(user.PasswordHash, []byte(password)); err != nil {
		a.logger.Info("invalid credentials")
		return "", fmt.Errorf("%s: %w", op, err)
	}
	logger.Info("user logged in successfully")
	token, err = genjwt.GenerateJWTToken(user, a.tokenTTL)
	if err != nil {
		a.logger.Error("failed to generate token")
		return "", fmt.Errorf("%s: %w", op, err)
	}
	return token, nil
}
