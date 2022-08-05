package controller

import (
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

// func (c *Config) AddScanGet(g *gin.Context) {
// 	newScan, err := c.service.Scan().Create(
// 		g.Query("place"),
// 		g.Query("plate"),
// 		g.Query("datetime"),
// 	)
// 	if err != nil {
// 		c.error(g, err)
// 		return
// 	}

// 	c.respond(g, newScan)
// }
