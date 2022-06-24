package controller

import (
	"database/sql"
	"fmt"
	"scan/app/apiserver"
	"scan/app/service"

	"github.com/gin-gonic/gin"
)

type Config struct {
	Router  *gin.Engine
	Db      *sql.DB
	Config  *apiserver.Config
	service *service.Config
}

func (c *Config) Service() *service.Config {
	if c.service != nil {
		return c.service
	}

	fmt.Println("Import Service!")

	c.service = &service.Config{
		Db: c.Db,
	}

	return c.service
}
