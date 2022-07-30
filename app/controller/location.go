package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetLocations godoc
// @Summary      Список расположений камер
// @Tags         Расположение
// @Accept       json
// @Produce      json
// @Success      200  {array}   model.Location
// @Failure      400  {object}  controller.ActionError
// @Router       /locations [get]
// @Security 	 ApiKeyAuth
func (c *Config) GetLocations(g *gin.Context) {
	locations, err := c.service.Location().All()
	if err != nil {
		c.error(g, err)
		return
	}
	g.IndentedJSON(http.StatusOK, locations)
}
