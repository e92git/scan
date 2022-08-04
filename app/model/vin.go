package model

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

type Vin struct {
	ID            int64     `json:"id" example:"7635"`
	Plate         string    `json:"plate" example:"О245КМ142"`
	Vin           *string   `json:"vin" example:"XTA219170K0330071"`
	Vin2          *string   `json:"vin2" example:"XTA219170K0330071"`
	Body          *string   `json:"body" example:"KGC100005240"`
	MarkId        *int      `json:"mark_id" example:"23"`
	ModelId       *int      `json:"model_id" example:"231"`
	Year          *int      `json:"year" example:"2012"`
	Response      *string   `json:"response" example:"{...}"`
	ResponseError *string   `json:"response_error" example:"400"`
	StatusId      int       `json:"status_id" example:"3"`
	AuthorUserId  int64     `json:"author_user_id" example:"234"`
	UpdatedAt     time.Time `json:"updated_at" example:"2022-07-23T11:23:55+07:00"`
	CreatedAt     time.Time `json:"created_at" example:"2022-07-28T11:23:55+07:00"`

	Author *User `json:"author" gorm:"foreignKey:AuthorUserId"`
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
	Success:     4, // Результат получен по API
}

// Validate ...
func (s *Vin) Validate() error {
	return validation.ValidateStruct(s,
		validation.Field(&s.Plate, validation.Required),
		validation.Field(&s.AuthorUserId, validation.Required),
		validation.Field(&s.StatusId, validation.Required),
		validation.Field(&s.UpdatedAt, validation.Required, validation.Date("2006-01-02 15:04:05")),
		validation.Field(&s.CreatedAt, validation.Required, validation.Date("2006-01-02 15:04:05")),
	)
}

func (s *Vin) NeedSend() bool {
	return s.StatusId == VinStatuses.Created || s.StatusId == VinStatuses.SendSuccess
}
