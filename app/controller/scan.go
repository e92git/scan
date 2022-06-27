package controller

import (
	"fmt"
	"net/http"
	"scan/app/model"

	"github.com/gin-gonic/gin"
)

func (c *Config) AddScan(g *gin.Context) {

	var newScan model.Scan

	newScan.Plate = g.Query("plate")

	// s.Service.AddScan(newScan)

	fmt.Println(newScan)

	g.IndentedJSON(http.StatusCreated, newScan)
}
