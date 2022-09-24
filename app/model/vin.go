package model

import (
	"encoding/json"
	"errors"
	"time"
)

type Vin struct {
	ID            int64     `json:"id" example:"7635"`
	Plate         string    `json:"plate" example:"О245КМ142" validate:"required|regex:^[A-Z0-9]+$"`
	Vin           *string   `json:"vin" example:"XTA219170K0330071"`
	Vin2          *string   `json:"vin2" example:"XTA219170K0330071"`
	Body          *string   `json:"body" example:"KGC100005240"`
	MarkId        *int      `json:"-" example:"12"`
	ModelId       *int      `json:"-" example:"123"`
	Year          *int      `json:"year" example:"2012"`
	Response      *string   `json:"response" example:"{...}"`
	ResponseCloud *string   `json:"response_cloud" example:"{...}"`
	ResponseError *string   `json:"response_error" example:"400: bad request"`
	StatusId      int       `json:"-" example:"4" validate:"required"`
	AuthorUserId  int64     `json:"-" example:"234" validate:"required"`
	UpdatedAt     time.Time `json:"updated_at" example:"2022-07-23T11:23:55.999+07:00"`
	CreatedAt     time.Time `json:"created_at" example:"2022-07-28T11:23:55.999+07:00"`

	Mark   *CarMark   `json:"mark" gorm:"foreignKey:MarkId"`
	Model  *CarModel  `json:"model" gorm:"foreignKey:ModelId"`
	Status *VinStatus `json:"status" gorm:"foreignKey:StatusId"`
	Author *User      `json:"author" gorm:"foreignKey:AuthorUserId"`
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

func (s *Vin) IsSuccessStatus() bool {
	return s.StatusId == VinStatuses.Success
}

func (s *Vin) IsErrorStatus() bool {
	return s.StatusId == VinStatuses.SendError
}

func (s *Vin) IsEmptyVin() bool {
	return s.Vin == nil && s.Vin2 == nil && s.Body == nil
}

func (s *Vin) IsEmptyCar() bool {
	return s.MarkId == nil || s.ModelId == nil || s.Year == nil
}
