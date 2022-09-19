package service

import (
	"fmt"
	"scan/app/store"
)

type Config struct {
	store       *store.Store
	location    *LocationService
	scan        *ScanService
	user        *UserService
	vin         *VinService
	vinAutocode *VinAutocodeService
	car         *CarService
	tire        *TireService
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
		c.scan = NewScan(c.store, c.Location())
	}
	return c.scan
}

func (c *Config) User() *UserService {
	if c.user == nil {
		c.user = NewUser(c.store)
	}
	return c.user
}

func (c *Config) Vin() *VinService {
	if c.vin == nil {
		c.vin = NewVin(c.store, c.VinAutocode())
	}
	return c.vin
}

func (c *Config) VinAutocode() *VinAutocodeService {
	if c.vinAutocode == nil {
		c.vinAutocode = NewVinAutocode(c.store, c.Car())
	}
	return c.vinAutocode
}

func (c *Config) Car() *CarService {
	if c.car == nil {
		c.car = NewCar(c.store)
	}
	return c.car
}

func (c *Config) Tire() *TireService {
	if c.tire == nil {
		c.tire = NewTire(c.store)
	}
	return c.tire
}
