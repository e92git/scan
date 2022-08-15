package main

import (
	"log"
	"scan/app/controller"
	_ "scan/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Дискаунтер автозачастей е92
// @version         1.0
// @description     Здесь представлены все методы для работы админстраторов и менеджеров магазинов.
// @description     Вопросы на info@e92.ru.

// @BasePath  /api/v1

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {

	c, err := controller.New()
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/api/v1")
	{
		// without user
		v1.GET("/locations", c.GetLocations)

		// auth User
		v1.Use(c.Auth())

		// "show_api" middleware
		v1.Use(c.ShowApiMiddleware())

		// "manager" middleware
		v1.Use(c.ManagerMiddleware())
		v1.POST("/scan", c.AddScan)
		v1.POST("/vin", c.VinByPlate)
	}

	err = r.Run(c.Addr())
	if err != nil {
		log.Fatal(err)
	}
}
