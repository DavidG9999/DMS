package postgres

import (
	"github.com/jmoiron/sqlx"
)

type User interface {
	UserCreator
	UserProvider
	UserEditor
}

type Repository struct {
	User
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		User: NewUserPostgres(db),
	}
}
