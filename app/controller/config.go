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
	store   *store.Store
	service *service.Config
}

func New() (*Config, error) {
	config, err := apiserver.LoadConfig()
	if err != nil {
		return nil, err
	}

	db, err := apiserver.ConnectGorm(config.Dsn, config.LogLevel)
	if err != nil {
		return nil, err
	}

	store := store.New(db)

	c := &Config{
		config:  config,
		store:   store,
		service: service.New(store),
	}

	return c, nil
}

func (c *Config) Addr() string {
	return c.config.BindAddr
}

func (c *Config) error(g *gin.Context, err error) {
	log.Print(err)
	g.IndentedJSON(http.StatusBadRequest, gin.H{
		"url": g.Request.URL,
		"body": g.Request.Body,
		"error": err.Error(),
	})
}

func (c *Config) respond(g *gin.Context, obj any) {
	g.IndentedJSON(http.StatusOK, obj)
}
