package scanService

import (
	"database/sql"
	"scan/app/service/scan/repository"
)

type Config struct {
	db                 *sql.DB
	repository *repository.Config
}

func New(db *sql.DB) *Config {
	return &Config{
		db:         db,
		repository: repository.New(db),
	}
}