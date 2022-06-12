package service

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)


type scan struct {
    ID              int64  `json:"id,omitempty"`
    LocationId      int64  `json:"location_id"`
    Plate           string `json:"plate"`
    VinId           int64  `json:"vin_id,omitempty"`
    ScannedAt       string `json:"scanned_at"`
    CreatedAt       string `json:"created_at"`
}

func (s *Server) AddScan(c *gin.Context) {

	var newScan scan

	newScan.Plate = c.Query("plate")

	fmt.Println(newScan)

	c.IndentedJSON(http.StatusCreated, newScan)
}