package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type AddScanRequest struct {
	Place     string `json:"place" example:"pokrovka" validate:"required"`
	Plate     string `json:"plate" example:"M343TT123" validate:"required"`
	ScannedAt string `json:"scanned_at" example:"2022-07-23 11:23:55" validate:"required"`
}

// AddScan godoc
// @Summary      Добавить отсканированный номер
// @Tags         Сканирование
// @Accept       json
// @Produce      json
// @Param 		 scan body AddScanRequest true "Добавить сканирование"
// @Success      200  {array}   model.Scan
// @Failure      400  {object}  controller.ActionError
// @Router       /scan [post]
// @Security 	 ApiKeyAuth
func (c *Config) AddScan(g *gin.Context) {
	req := &AddScanRequest{}
	user, err := c.initRequest(g, req)
	if err != nil {
		c.error(g, err)
		return
	}

	res, err := c.service.Scan().AddScanWithPrepare(req.Place, req.Plate, req.ScannedAt, user.ID)
	if err != nil {
		c.error(g, err)
		return
	}

	c.respond(g, res)
}

type AddScanBulkRequest struct {
	LocationId   int    `json:"location_id" example:"1" validate:"required"`
	PlateAndDate string `json:"plate_and_date" example:"Т237АС142	2022-07-06 10:31:12 Т182АС142	2022-07-06 10:29:40" validate:"required"`
}

func (c *Config) AddScanBulk(g *gin.Context) {
	req := &AddScanBulkRequest{}
	user, err := c.initRequest(g, req)
	if err != nil {
		c.error(g, err)
		return
	}

	res := "ff"
	fmt.Println(user)

	// res, err := c.service.Scan().AddScanWithPrepare(req.Place, req.Plate, req.ScannedAt, user.ID)
	// if err != nil {
	// 	c.error(g, err)
	// 	return
	// }

	c.respond(g, res)
}
