package service

import (
	"fmt"
	"scan/app/store"
)

type Config struct {
	store    *store.Store
	location *LocationService
	scan     *ScanService
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

func (c *Config) Scan() *ScanService {
	if c.scan == nil {
		c.scan = NewScan(c.store)
	}
	return c.scan
}
