package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Config) GetLocations(c *gin.Context) {
	locations, err := s.Service().Location().All()
	if err != nil {
		fmt.Println(err)
	}
	c.IndentedJSON(http.StatusOK, locations)
}
