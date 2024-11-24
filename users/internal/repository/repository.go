package repository

import (
	"errors"
)

var (
	ErrUserExists   = errors.New("user already exist")
	ErrUserNotFound = errors.New("user not found")
)
