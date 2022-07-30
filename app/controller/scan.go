package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *Config) AddScan(g *gin.Context) {
	err := http.ErrServerClosed
	if err != nil {
		c.error(g, err)
		return
	}

	type request struct {
		Place    string `json:"place"`
		Plate    string `json:"plate"`
		Datetime string `json:"datetime"`
	}
	req := &request{}

	if err := g.BindJSON(req); err != nil {
		c.error(g, err)
		return
	}

	newScan, err := c.service.Scan().FirstOrCreate(
		req.Place,
		req.Plate,
		req.Datetime,
	)
	if err != nil {
		c.error(g, err)
		return
	}

	c.respond(g, newScan)
}

func (c *Config) AddScanGet(g *gin.Context) {
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
