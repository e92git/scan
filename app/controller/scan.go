package controller

import (
	// "scan/app/model"

	"github.com/gin-gonic/gin"
)

func (c *Config) AddScan(g *gin.Context) {
	// loc, err := c.store.Location().FindByCode(g.Query("place"))
	// if err != nil {
	// 	c.error(g, err)
	// 	return
	// }

	// newScan := &model.Scan{
	// 	LocationId: loc.ID,
	// 	Plate:      g.Query("plate"),
	// 	ScannedAt:  g.Query("datetime"),
	// }

	// if err := c.store.Scan().Create(newScan); err != nil {
	// 	c.error(g, err)
	// 	return
	// }

	// c.respond(g, newScan)
}
