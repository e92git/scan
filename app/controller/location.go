package controller

import (
	"net/http"
	"scan/app/model"

	"github.com/gin-gonic/gin"
)

func (s *Server) GetLocations(c *gin.Context) {
	var newLocation model.Location
	c.IndentedJSON(http.StatusOK, newLocation)
}
