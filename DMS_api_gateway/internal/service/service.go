package service

import (
	"errors"
	"log/slog"

	authgrpc "github.com/DavidG9999/DMS/DMS_api_gateway/internal/clients/grpc/auth"
	putlistgrpc "github.com/DavidG9999/DMS/DMS_api_gateway/internal/clients/grpc/documents/putlist"
	usergrpc "github.com/DavidG9999/DMS/DMS_api_gateway/internal/clients/grpc/user"
)

var (
	ErrInvalidCredentials    = errors.New("invalid credentials")
	ErrUserExist             = errors.New("user already exists")
	ErrAutoExists            = errors.New("auto already exists")
	ErrBankAccExists         = errors.New("bank account already exists")
	ErrContragentExists      = errors.New("contragent already exists")
	ErrDriverExists          = errors.New("driver already exists")
	ErrOrganizationExists    = errors.New("organization already exists")
	ErrPutlistExists         = errors.New("putlist already exists")
	ErrInvalidDateTimeFormat = errors.New("invalid datetime format")
)

type Authorization interface {
	Register
	Login
}

type User interface {
	UserProvider
	UserEditor
}

type Auto interface {
	AutoCreator
	AutoProvider
	AutoEditor
}

type Contragent interface {
	ContragentCreator
	ContragentProvider
	ContragentEditor
}

type Dispetcher interface {
	DispetcherCreator
	DispetcherProvider
	DispetcherEditor
}

type Driver interface {
	DriverCreator
	DriverProvider
	DriverEditor
}

type Mechanic interface {
	MehanicCreator
	MehanicProvider
	MehanicEditor
}

type Organization interface {
	OrganizationCreator
	OrganizationProvider
	OrganizationEditor
}

type BankAccount interface {
	BankAccountCreator
	BankAccountProvider
	BankAccountEditor
}
type Putlist interface {
	PutlistCreator
	PutlistProvider
	PutlistEditor
}

type PutlistBody interface {
	PutlistBodyCreator
	PutlistBodyProvider
	PutlistBodyEditor
}

type PutlistAPI interface {
	Auto
	Contragent
	Dispetcher
	Driver
	Mechanic
	Organization
	BankAccount
	Putlist
	PutlistBody
}

type Service struct {
	Authorization
	User
	PutlistAPI
}

func NewService(logger *slog.Logger, authClient authgrpc.AuthClient, userClient usergrpc.UserClient, putlistClient putlistgrpc.PutlistClient) *Service {
	return &Service{
		Authorization: NewAuthService(logger, authClient),
		User:          NewUserService(logger, userClient),
		PutlistAPI:    NewPutlistService(logger, putlistClient),
	}
}
