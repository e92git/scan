package controller

import (
	"github.com/gin-gonic/gin"
)

func (c *Config) AddScan(g *gin.Context) {
	type request struct {
		Place    string `json:"place"`
		Plate    string `json:"plate"`
		Datetime string `json:"datetime"`
	}
	req := &request{}

	if err := g.BindJSON(req); err != nil {
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
