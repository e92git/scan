package cron

import (
	"fmt"
	"scan/app/service"
	"time"

	"github.com/go-co-op/gocron"
)

var c *gocron.Scheduler

func CronStart(s *service.Config) {
	// выход если уже запущен крон
	if c != nil {
		return
	}
    // инициализируем объект планировщика
    c = gocron.NewScheduler(time.UTC)
    // добавляем одну задачу каждые 60 сек
    c.Every(60).Second().Do(task)
    // запускаем планировщик в фоне
    c.StartAsync()
}

func CronStop() {
	c.Stop()
}

func task() {
	fmt.Println("Hello, playground1")
	time.Sleep(1 * time.Second)
	fmt.Println("Hello, playground2")
}
