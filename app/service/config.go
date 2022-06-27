package service

import (
	"database/sql"
	"scan/app/service/location"
)

type Config struct {
	db 		 *sql.DB
	location *locationService.Config
}

func New(db *sql.DB) *Config {
	return &Config{
		db: db,
	}
}

func (c *Config) Location() *locationService.Config {
	if c.location == nil {
		c.location = locationService.New(c.db)
	}
	
	return c.location
}
