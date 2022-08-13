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
	Created     int
	SendError   int
	SendSuccess int
	Success     int
}{
	Created:     1, // Создан
	SendError:   2, // API-запрос НЕ отправился. Ошибка.
	SendSuccess: 3, // API-запрос отправлен. Успех!
	Success:     4, // Результат успешно получен
}
