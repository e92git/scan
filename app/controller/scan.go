package controller

import (
	"github.com/gin-gonic/gin"
)

type AddScanRequest struct {
	Place    string `json:"place" example:"pokrovka"`
	Plate    string `json:"plate" example:"M343TT123"`
	Datetime string `json:"datetime" example:"2022-07-23 11:23:55"`
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
	scan := &AddScanRequest{}
	if err := g.BindJSON(scan); err != nil {
		c.error(g, err)
		return
	}

	newScan, err := c.service.Scan().Create(
		scan.Place,
		scan.Plate,
		scan.Datetime,
	)
	if err != nil {
		c.error(g, err)
		return
	}

	c.respond(g, newScan)
}

func (c *Config) AddScanGet(g *gin.Context) {
	newScan, err := c.service.Scan().Create(
		g.Query("place"),
		g.Query("plate"),
		g.Query("datetime"),
	)
	if err != nil {
		c.error(g, err)
		return
	}

	c.respond(g, newScan)
}
