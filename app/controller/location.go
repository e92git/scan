package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func (c *Config) GetLocations(g *gin.Context) {
	locations, err := c.service.Location().All()
	if err != nil {
		c.error(g, err)
		return
	}
	g.IndentedJSON(http.StatusOK, locations)
}
