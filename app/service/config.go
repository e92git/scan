package service

import (
	"database/sql"
	"fmt"
	"scan/app/service/location"
	"scan/app/service/scan"
)

type Config struct {
	db 		 *sql.DB
	location *locationService.Config
	scan 	 *scanService.Config
}

func New(db *sql.DB) *Config {
	return &Config{
		db: db,
	}
}

func (c *Config) Location() *locationService.Config {
	if c.location == nil {
		fmt.Println("Import locationService!")
		c.location = locationService.New(c.db)
	}
	
	return c.location
}

func (c *Config) Scan() *scanService.Config {
	if c.scan == nil {
		c.scan = scanService.New(c.db)
	}
	
	return c.scan
}
