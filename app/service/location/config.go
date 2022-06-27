package locationService

import (
	"database/sql"
	"scan/app/service/location/repository"
)

type Config struct {
	db                 *sql.DB
	locationRepository *repository.Config
}

func New(db *sql.DB) *Config {
	return &Config{
		db:                 db,
		locationRepository: repository.New(db),
	}
}