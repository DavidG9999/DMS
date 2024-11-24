package postgres

import (
	"github.com/jmoiron/sqlx"
)

type AutoRepository interface {
	AutoCreator
	AutoProvider
	AutoEditor
}

type ContragentRepository interface {
	ContragentCreator
	ContragentProvider
	ContragentEditor
}

type DispetcherRepository interface {
	DispetcherCreator
	DispetcherProvider
	DispetcherEditor
}

type DriverRepository interface {
	DriverCreator
	DriverProvider
	DriverEditor
}

type MehanicRepository interface {
	MehanicCreator
	MehanicProvider
	MehanicEditor
}

type OrganizationRepository interface {
	OrganizationCreator
	OrganizationProvider
	OrganizationEditor
}

type BankAccountRepository interface {
	BankAccountCreator
	BankAccountProvider
	BankAccountEditor
}

type PutlistRepository interface {
	PutlistCreator
	PutlistProvider
	PutlistEditor
}

type PutlistBodyRepository interface {
	PutlistBodyCreator
	PutlistBodyProvider
	PutlistBodyEditor
}

type Repository struct {
	BankAccountRepository
	AutoRepository
	ContragentRepository
	DispetcherRepository
	DriverRepository
	MehanicRepository
	OrganizationRepository
	PutlistRepository
	PutlistBodyRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		AutoRepository:         NewAutoPostgres(db),
		OrganizationRepository: NewOrganizationPostgres(db),
		BankAccountRepository:  NewBankAccountPostgres(db),
		ContragentRepository:   NewContragentPostgres(db),
		DispetcherRepository:   NewDispetcherPostgres(db),
		DriverRepository:       NewDriverPostgres(db),
		MehanicRepository:      NewMehanicPostgres(db),
		PutlistRepository:      NewPutlistPostgres(db),
		PutlistBodyRepository:  NewPutlistBodyPostgres(db),
	}
}
