package main

import (
	"log"
	"scan/app/controller"
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

	// all routers
	c.SetUpRouters()

	err = c.RunServer()
	if err != nil {
		log.Fatal(err)
	}
}
