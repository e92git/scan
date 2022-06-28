package service

import (
	"fmt"
	"scan/app/service/location"
	"scan/app/service/scan"
	"scan/app/store"
)

type Config struct {
	store 	 *store.Store
	location *locationService.Config
	scan 	 *scanService.Config
}

func New(store *store.Store) *Config {
	return &Config{
		store: store,
	}
}

func (c *Config) Location() *locationService.Config {
	if c.location == nil {
		fmt.Println("Import locationService!")
		c.location = locationService.New(c.store)
	}
	
	return c.location
}

func (c *Config) Scan() *scanService.Config {
	if c.scan == nil {
		c.scan = scanService.New(c.store)
	}
	
	return c.scan
}
