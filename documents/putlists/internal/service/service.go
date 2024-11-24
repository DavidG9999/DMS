package service

import (
	"errors"
	"log/slog"

	"github.com/DavidG9999/DMS/documents/putlists/internal/repository/postgres"
)

var (
	ErrInvalidCredentials    = errors.New("invalid credentials")
	ErrAutoExists            = errors.New("auto already exists")
	ErrBankAccExists         = errors.New("bank account already exists")
	ErrContragentExists      = errors.New("contragent already exists")
	ErrDriverExists          = errors.New("driver already exists")
	ErrOrganizationExists    = errors.New("organization already exists")
	ErrPutlistExists         = errors.New("putlist already exists")
	ErrInvalidDateTimeFormat = errors.New("invalid datetime format")
)

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

type Mehanic interface {
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

type Service struct {
	*slog.Logger
	Auto
	Contragent
	Dispetcher
	Driver
	Mehanic
	Organization
	BankAccount
	Putlist
	PutlistBody
}

func NewService(logger *slog.Logger, repos *postgres.Repository) *Service {
	return &Service{
		Auto:         NewAutoService(logger, repos.AutoRepository),
		Organization: NewOrganizationService(logger, repos.OrganizationRepository),
		BankAccount:  NewBankAccountService(logger, repos.BankAccountRepository),
		Contragent:   NewContragentService(logger, repos.ContragentRepository),
		Dispetcher:   NewDispetcherService(logger, repos.DispetcherRepository),
		Driver:       NewDriverService(logger, repos.DriverRepository),
		Mehanic:      NewMechanicService(logger, repos.MehanicRepository),
		Putlist:      NewPutlistService(logger, repos.PutlistRepository),
		PutlistBody:  NewPutlistBodyService(logger, repos.PutlistBodyRepository),
	}
}
