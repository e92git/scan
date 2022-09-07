package controller

import (
	"scan/app/service"

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
	LocationId int64           `json:"location_id" example:"1" validate:"required"`
	Data       []service.Scans `json:"data" validate:"required"`
}

// AddScan godoc
// @Summary      Добавить отсканированные номера пачкой
// @Tags         Сканирование
// @Accept       json
// @Produce      json
// @Param 		 scan body AddScanBulkRequest true "Добавить сканирование"
// @Success      200
// @Failure      400  {object}  controller.ActionError
// @Router       /scan/bulk [post]
// @Security 	 ApiKeyAuth
func (c *Config) AddScanBulk(g *gin.Context) {
	req := &AddScanBulkRequest{}
	user, err := c.initRequest(g, req)
	if err != nil {
		c.error(g, err)
		return
	}

	err = c.service.Scan().CreateBulk(req.LocationId, &req.Data, user.ID)
	if err != nil {
		c.error(g, err)
		return
	}

	c.respond(g, nil)
}
