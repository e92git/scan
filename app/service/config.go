package service

import (
	"fmt"
	scanService "scan/app/service/scan"
	"scan/app/store"
)

type Config struct {
	store    *store.Store
	location *LocationService
	scan     *scanService.Config
}

func New(store *store.Store) *Config {
	return &Config{
		store: store,
	}
}

func (c *Config) Location() *LocationService {
	if c.location == nil {
		fmt.Println("Import locationService!")
		c.location = NewLocation(c.store)
	}

	return c.location
}

func (c *Config) Scan() *scanService.Config {
	if c.scan == nil {
		c.scan = scanService.New(c.store)
	}

	return c.scan
}
