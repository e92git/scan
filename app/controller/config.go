package controller

import (
	"log"
	"net/http"
	"scan/app/apiserver"
	"scan/app/service"
	"scan/app/store"

	"github.com/gin-gonic/gin"
)

type Config struct {
	config  *apiserver.Config
	Router  *gin.Engine
	store   *store.Store
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
		store:   store.New(db),
		service: service.New(db),
	}

	return c, nil
}

func (c *Config) RunServer() error {
	return c.Router.Run(c.config.BindAddr)
}

func (c *Config) error(g *gin.Context, err error) {
	log.Print(err)
	g.JSON(http.StatusBadRequest, gin.H{
		"error": err.Error(),
	})
}

func (c *Config) respond(g *gin.Context, obj any) {
	g.IndentedJSON(http.StatusOK, obj)
}
