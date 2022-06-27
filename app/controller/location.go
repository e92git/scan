package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *Config) GetLocations(g *gin.Context) {
	locations, err := c.Service.Location().All()
	if err != nil {
		fmt.Println(err)
	}
	g.IndentedJSON(http.StatusOK, locations)
}
