package service

import (
	"log/slog"

	"github.com/DavidG9999/DMS/users/internal/repository/postgres"
)

type User interface {
	UserCreator
	UserProvider
	UserEditor
}

type Service struct {
	*slog.Logger
	User
}

func NewService(logger *slog.Logger, repo postgres.Repository) *Service {
	return &Service{
		User: NewUserService(logger, repo.User),
	}
}
