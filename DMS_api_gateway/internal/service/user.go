package service

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	usergrpc "github.com/DavidG9999/DMS/DMS_api_gateway/internal/clients/grpc/user"
	"github.com/DavidG9999/DMS/DMS_api_gateway/internal/entity"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserService struct {
	logger     *slog.Logger
	userClient usergrpc.UserClient
}

func NewUserService(logger *slog.Logger, userClient usergrpc.UserClient) *UserService {
	return &UserService{
		logger:     logger,
		userClient: userClient,
	}
}

type UserProvider interface {
	GetUserById(ctx context.Context, userId int64) (name string, email string, err error)
}

type UserEditor interface {
	UpdateName(ctx context.Context, userId int64, updateData entity.UpdateNameUserInput) (string, error)
	UpdatePassword(ctx context.Context, userId int64, updateData entity.UpdatePasswordUserInput) (string, error)
	DeleteUser(ctx context.Context, userId int64) (string, error)
}

func (u *UserService) GetUserById(ctx context.Context, userId int64) (name string, email string, err error) {
	const op = "service.GetUserById"

	logger := u.logger.With(slog.String("op", op))
	logger.Info("getting user")

	userName, userEmail, err := u.userClient.GetUserById(ctx, userId)
	if err != nil {
		if errors.Is(err, status.Error(codes.NotFound, "user not found")) {
			logger.Error("user not found")
			return "", "", fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
		}
		logger.Error("failed to get user")
		return "", "", fmt.Errorf("%s: %w", op, err)
	}
	logger.Info("user received")

	return userName, userEmail, nil
}

func (u *UserService) UpdateName(ctx context.Context, userId int64, updateData entity.UpdateNameUserInput) (string, error) {
	const op = "service.UpdateName"

	logger := u.logger.With(slog.String("op", op))
	logger.Info("updating user name")

	err := updateData.Validate()
	if err != nil {
		return "", err
	}

	message, err := u.userClient.UpdateName(ctx, userId, *updateData.Name)
	if err != nil {
		if errors.Is(err, status.Error(codes.NotFound, "user not found")) {
			logger.Error("user not found")
			return "", fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
		}
		logger.Error("failed to update user name")
		return "", fmt.Errorf("%s: %w", op, err)
	}
	logger.Info("updated name")

	return message, nil
}

func (u *UserService) UpdatePassword(ctx context.Context, userId int64, updateData entity.UpdatePasswordUserInput) (string, error) {
	const op = "service.UpdatePassword"

	logger := u.logger.With(slog.String("op", op))
	logger.Info("updating user password")

	err := updateData.Validate()
	if err != nil {
		return "", err
	}

	message, err := u.userClient.UpdatePassword(ctx, userId, *updateData.Password)
	if err != nil {
		if errors.Is(err, status.Error(codes.NotFound, "user not found")) {
			logger.Error("user not found")
			return "", fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
		}
		logger.Error("failed to update user password")
		return "", fmt.Errorf("%s: %w", op, err)
	}
	logger.Info("updated password")

	return message, nil
}

func (u *UserService) DeleteUser(ctx context.Context, userId int64) (string, error) {
	const op = "service.DeleteUser"

	logger := u.logger.With(slog.String("op", op))
	logger.Info("deleting user")

	message, err := u.userClient.DeleteUser(ctx, userId)
	if err != nil {
		if errors.Is(err, status.Error(codes.NotFound, "user not found")) {
			logger.Error("user not found")
			return "", fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
		}
		logger.Error("failed to delete user")
		return "", fmt.Errorf("%s: %w", op, err)
	}
	logger.Info("user deleted")

	return message, nil
}
