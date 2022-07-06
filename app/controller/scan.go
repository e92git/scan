package controller

import (
	// "scan/app/model"

	"github.com/gin-gonic/gin"
)

func (c *Config) AddScans(g *gin.Context) {
	newScan, err := c.service.Scan().FirstOrCreate(
		g.Query("place"),
		g.Query("plate"),
		g.Query("datetime"),
	)
	if err != nil {
		c.error(g, err)
		return
	}

	c.respond(g, newScan)
}

func (c *Config) AddScan(g *gin.Context) {
	newScan, err := c.service.Scan().FirstOrCreate(
		g.Query("place"),
		g.Query("plate"),
		g.Query("datetime"),
	)
	if err != nil {
		c.error(g, err)
		return
	}

	c.respond(g, newScan)
}
