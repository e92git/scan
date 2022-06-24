package service

import (
	"database/sql"
	"scan/app/service/location"
)

type Config struct {
	Db 		 *sql.DB
	location *locationService.Config
}

func (c *Config) Location() *locationService.Config {
	if c.location != nil {
		return c.location
	}
	c.location = locationService.New(c.Db)
	return c.location
}
