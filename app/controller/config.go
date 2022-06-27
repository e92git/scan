package controller

import (
	"database/sql"
	"scan/app/apiserver"
	"scan/app/service"

	"github.com/gin-gonic/gin"
)

type Config struct {
	Config  *apiserver.Config
	Db      *sql.DB
	Router  *gin.Engine
	Service *service.Config
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
		Config:  config,
		Db:      db,
		Router:  gin.Default(),
		Service: service.New(db),
	}

	return c, nil
}

func (c *Config) RunServer() error {
	return c.Router.Run(c.Config.BindAddr)
}
