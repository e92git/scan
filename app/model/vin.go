package model

import (
	"encoding/json"
	"errors"
	"time"
)

type Vin struct {
	ID            int64     `json:"id" example:"7635"`
	Plate         string    `json:"plate" example:"О245КМ142" validate:"required"`
	Vin           *string   `json:"vin" example:"XTA219170K0330071"`
	Vin2          *string   `json:"vin2" example:"XTA219170K0330071"`
	Body          *string   `json:"body" example:"KGC100005240"`
	MarkId        *int      `json:"mark_id" example:"23"`
	ModelId       *int      `json:"model_id" example:"231"`
	Year          *int      `json:"year" example:"2012"`
	Response      *string   `json:"response" example:"{...}"`
	ResponseError *string   `json:"response_error" example:"400"`
	StatusId      int       `json:"status_id" example:"3" validate:"required"`
	AuthorUserId  int64     `json:"author_user_id" example:"234" validate:"required"`
	UpdatedAt     time.Time `json:"updated_at" example:"2022-07-23T11:23:55+07:00"`
	CreatedAt     time.Time `json:"created_at" example:"2022-07-28T11:23:55+07:00"`

	Author *User `json:"-" gorm:"foreignKey:AuthorUserId"`
}

type responseItem struct {
	Uid        string `json:"uid"`
	IsNew      bool   `json:"isnew"`
	SuggestGet string `json:"suggest_get"`
}
type response struct {
	Data []responseItem `json:"data"`
}

func (v *Vin) GetAutocodeUid() (*string, error) {
	if v.Response == nil {
		return nil, errors.New("Vin.Response is nil")
	}
	r := response{}
	err := json.Unmarshal([]byte(*v.Response), &r)
	if err != nil {
		return nil, err
	}
	if len(r.Data) > 0 {
		uid := r.Data[0].Uid
		return &uid, nil
	}
	return nil, errors.New("Autocode uid not found")
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

func (s *Vin) NeedSend() bool {
	return s.StatusId == VinStatuses.Created || s.StatusId == VinStatuses.SendSuccess
}

func (s *Vin) IsSuccessStatus() bool {
	return s.StatusId == VinStatuses.Success
}

func (s *Vin) IsErrorStatus() bool {
	return s.StatusId == VinStatuses.SendError
}
