package cron

import (
	"scan/app/apiserver"
	"scan/app/service"
	"time"

	"github.com/go-co-op/gocron"
)

var cron *gocron.Scheduler

func CronStart( conf *apiserver.Config, s *service.Config) {
	// выход, если отключен запуск крона в настройках .env
	if !conf.RunCron {
		return
	}	
	// выход, если уже запущен крон
	if cron != nil {
		return
	}
    // инициализируем объект планировщика
    cron = gocron.NewScheduler(time.UTC)
    // добавляем одну задачу каждые 5 мин
    cron.Every(5).Minute().Do(s.Vin().CronFindDeffered)
    // запускаем планировщик в фоне
    cron.StartAsync()
}

func CronStop() {
	if cron == nil {
		return
	}
	cron.Stop()
}
