package model

import (
	"time"
)

type VinStatus struct {
	ID        int       `json:"id" example:"4"`
	Name      string    `json:"name" example:"Результат успешно получен"`
	CreatedAt time.Time `json:"-" example:"2022-07-28T11:23:55.999+07:00"`
}

var VinStatuses = struct {
	InProcess       int
	CreatedDeferred int
	SendError       int
	SendSuccess     int
	Success         int
}{
	InProcess:       1, // Создан. Взят в работу.
	CreatedDeferred: 5, // Создан (отложенный запуск)
	SendError:       2, // API-запрос НЕ отправился. Ошибка.
	SendSuccess:     3, // API-запрос отправлен. Успех!
	Success:         4, // Результат успешно получен
}
