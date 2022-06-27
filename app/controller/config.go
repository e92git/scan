package controller

import (
	"scan/app/apiserver"
	"scan/app/service"

	"github.com/gin-gonic/gin"
)

type Config struct {
	config  *apiserver.Config
	Router  *gin.Engine
	service *service.Config
}

func New() (*Config, error) {
	config, err := apiserver.LoadConfig()
	if err != nil {
		return nil, err
	}

	db, err := apiserver.ConnectDb(config.DatabaseURL)
	if err != nil {
		return nil, err
	}

	c := &Config{
		config:  config,
		Router:  gin.Default(),
		service: service.New(db),
	}

	return c, nil
}

func (c *Config) RunServer() error {
	return c.Router.Run(c.config.BindAddr)
}
