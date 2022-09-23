package controller

import (
	"github.com/gin-gonic/gin"
)

// GetTireAnalytics godoc
// @Summary      Аналитика для закупки шин
// @Tags         Шины
// @Accept       json
// @Produce      json
// @Success      200  {object}  model.TireAnalyticsResponse
// @Failure      400  {object}  controller.ActionError
// @Router       /tire/analytics [get]
// @Security 	 ApiKeyAuth
func (c *Config) GetTireAnalytics(g *gin.Context) {
	r, err := c.service.Tire().GetTireAnalytics()
	if err != nil {
		c.error(g, err)
		return
	}
	c.respond(g, r)
}


// GetTireSync godoc
// @Summary      Синхронизировать марки и модели с таблицей tires
// @Tags         Шины
// @Accept       json
// @Produce      json
// @Success      200  {object}  service.GetTireSyncResponse
// @Failure      400  {object}  controller.ActionError
// @Router       /tire/sync [get]
// @Security 	 ApiKeyAuth
func (c *Config) GetTireSync(g *gin.Context) {
	r, err := c.service.Tire().GetTireSync()
	if err != nil {
		c.error(g, err)
		return
	}
	c.respond(g, r)
}
