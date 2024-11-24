package service

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/DavidG9999/DMS/users/internal/domain/entity"
	"github.com/DavidG9999/DMS/users/internal/repository"
	"github.com/DavidG9999/DMS/users/internal/repository/postgres"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	logger   *slog.Logger
	userRepo postgres.User
}

func NewUserService(logger *slog.Logger, userRepo postgres.User) *UserService {
	return &UserService{
		logger:   logger,
		userRepo: userRepo,
	}
}

type UserCreator interface {
	CreateUser(ctx context.Context, name string, email string, passwordHash string) (userID int64, err error)
}

type UserProvider interface {
	GetUser(ctx context.Context, email string) (entity.User, error)
	GetUserById(ctx context.Context, userID int64) (entity.User, error)
}

type UserEditor interface {
	UpdateName(ctx context.Context, userID int64, updateName string) error
	UpdatePassword(ctx context.Context, userID int64, updatePassword string) error
	DeleteUser(ctx context.Context, userID int64) error
}

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUserExist          = errors.New("user already exist")
)

func (s *UserService) CreateUser(ctx context.Context, name string, email string, passwordHash string) (userID int64, err error) {
	const op = "user_service.CreateUser"

	logger := s.logger.With(
		slog.String("op", op),
	)
	logger.Info("creating user")

	userID, err = s.userRepo.CreateUser(ctx, name, email, passwordHash)
	if err != nil {
		if errors.Is(err, repository.ErrUserExists) {
			logger.Error("user already exists")
			return 0, fmt.Errorf("%s: %w", op, ErrUserExist)
		}
		logger.Error("failed to create user")
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	logger.Info("user created")
	return userID, nil
}

func (s *UserService) GetUser(ctx context.Context, email string) (entity.User, error) {
	const op = "user_service.GetUser"

	logger := s.logger.With(
		slog.String("op", op),
	)
	logger.Info("getting user")

	user, err := s.userRepo.GetUser(ctx, email)
	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			logger.Error("user not found")
			return entity.User{}, fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
		}
		logger.Error("failed to get user")
		return entity.User{}, fmt.Errorf("%s: %w", op, err)
	}
	logger.Info("user received")
	return user, nil
}

func (s *UserService) GetUserById(ctx context.Context, userID int64) (entity.User, error) {
	const op = "user_service.GetUserById"

	logger := s.logger.With(
		slog.String("op", op),
	)
	logger.Info("getting user by user id")

	user, err := s.userRepo.GetUserById(ctx, userID)
	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			logger.Error("user not found")
			return entity.User{}, fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
		}
		logger.Error("failed to get user")
		return entity.User{}, fmt.Errorf("%s: %w", op, err)
	}
	logger.Info("user received")
	return user, nil
}

func (s *UserService) UpdateName(ctx context.Context, userID int64, updateName string) error {
	const op = "user_service.UpdateName"

	logger := s.logger.With(
		slog.String("op", op),
	)
	logger.Info("updating username")

	err := s.userRepo.UpdateName(ctx, userID, updateName)
	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			logger.Error("user not found")
			return fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
		}
		logger.Error("failed to update user")
		return fmt.Errorf("%s: %w", op, err)
	}
	logger.Info("username updated")
	return nil
}

func (s *UserService) UpdatePassword(ctx context.Context, userID int64, updatePassword string) error {
	const op = "user_service.UpdatePassword"

	logger := s.logger.With(
		slog.String("op", op),
	)
	logger.Info("updating user password")

	updatePasswordHash, err := bcrypt.GenerateFromPassword([]byte(updatePassword), bcrypt.DefaultCost)
	if err != nil {
		logger.Error("failed to generate password hash")
		return fmt.Errorf("%s: %w", op, err)
	}

	err = s.userRepo.UpdatePassword(ctx, userID, string(updatePasswordHash))
	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			logger.Error("user not found")
			return fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
		}
		logger.Error("failed to update user")
		return fmt.Errorf("%s: %w", op, err)
	}
	logger.Info("user password updated")
	return nil
}

func (s *UserService) DeleteUser(ctx context.Context, userID int64) error {
	const op = "user_service.DeleteUser"

	logger := s.logger.With(
		slog.String("op", op),
	)
	logger.Info("deleting user")

	err := s.userRepo.DeleteUser(ctx, userID)
	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			logger.Error("user not found")
			return fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
		}
		logger.Error("failed to delete user")
		return fmt.Errorf("%s: %w", op, err)
	}
	logger.Info("user deleted")
	return nil
}
