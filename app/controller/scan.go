package controller

import (
	"example/hello/app/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)


func (s *Server) AddScan(c *gin.Context) {

	var newScan model.Scan

	newScan.Plate = c.Query("plate")

	s.Service.AddScan(newScan)

	fmt.Println(newScan)

	c.IndentedJSON(http.StatusCreated, newScan)
}
