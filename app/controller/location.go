package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func (s *Server) GetLocations(c *gin.Context) {
	locations, _ := s.Service.GetAll()
	c.IndentedJSON(http.StatusOK, locations)
}
