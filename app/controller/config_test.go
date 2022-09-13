package controller

import (
	"scan/app/apiserver"
	"scan/app/service"
	"scan/app/store"
	"testing"
)

func TestNew(t *testing.T) {
	_, err := GetTestController()
	if err != nil {
		t.Error(err)
		return
	}
}

var TestController *Config

func GetTestController() (*Config, error) {
	if TestController != nil {
		return TestController, nil
	}

	// load config
	config, err := apiserver.LoadConfig()
	if err != nil {
		return nil, err
	}
	// подмена бд на тестовую!
	config.Dsn = config.DsnTest
	// load store
	store, err := store.New(config)
	if err != nil {
		return nil, err
	}
	// load service
	service := service.New(store)

	// load controller
	TestController = New(config, store, service)

	return TestController, nil
}
