package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

const (
	bankAccountsTable   = "bankaccounts"
	autosTable          = "autos"
	contragentsTable    = "contragents"
	dispetchersTable    = "dispetchers"
	driversTable        = "drivers"
	itemsTable          = "items"
	mehanicsTable      = "mehanics"
	organizationsTable  = "organizations"
	putlistHeadersTable = "putlistheaders"
	putlistBodiesTable  = "putlistbodies"
)

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	const op = "postgres.NewPostgresDB"

	db, err := sqlx.Open("pgx", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)

	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return db, nil
}
