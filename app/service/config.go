package service

import (
	"fmt"
	"scan/app/apiserver"
	"scan/app/store"
)

type Config struct {
	config      *apiserver.Config
	store       *store.Store
	location    *LocationService
	scan        *ScanService
	user        *UserService
	vinAutocode *VinAutocodeService
	vinCloud    *VinCloudService
	vin         *VinService
	car         *CarService
	tire        *TireService
}

func New(config *apiserver.Config, store *store.Store) *Config {
	return &Config{
		config: config,
		store:  store,
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

func (c *Config) VinAutocode() *VinAutocodeService {
	if c.vinAutocode == nil {
		c.vinAutocode = NewVinAutocode(c.config ,c.store, c.Car())
	}
	return c.vinAutocode
}

func (c *Config) VinCloud() *VinCloudService {
	if c.vinCloud == nil {
		c.vinCloud = NewVinCloud(c.config, c.store, c.Car())
	}
	return c.vinCloud
}

func (c *Config) Vin() *VinService {
	if c.vin == nil {
		c.vin = NewVin(c.store, c.VinAutocode(), c.VinCloud())
	}
	return c.vin
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
