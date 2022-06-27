package locationService

import (
	"database/sql"
	"fmt"
	"scan/app/service/location/repository"
)

type Config struct {
	db                 *sql.DB
	locationRepository *repository.Location
}

func New(db *sql.DB) *Config {
	return &Config{
		db: db,
	}
}

func (c *Config) LocationRepository() (*repository.Location) {
	if c.locationRepository == nil {
		fmt.Println("Import LocationRepository!")
		c.locationRepository = repository.New(c.db)
	}
	
	return c.locationRepository
} 