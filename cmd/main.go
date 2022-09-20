package main

import (
	"log"
	"scan/app/apiserver"
	"scan/app/controller"
	"scan/app/helper/cron"
	"scan/app/service"
	"scan/app/store"
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

	///// load
	// load config
	config, err := apiserver.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	// load store
	store, err := store.New(config)
	if err != nil {
		log.Fatal(err)
	}
	// load service
	service := service.New(config, store)
	// load controller
	controller := controller.New(config, store, service)


	///// run
	// run cron
	cron.CronStart(config, service)
	// run server
	err = controller.RunServer()
	if err != nil {
		log.Fatal(err)
	}
}
