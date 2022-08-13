package model

import "time"

type CarMark struct {
	ID        int       `json:"id" example:"12"`
	Name      string    `json:"name" example:"Toyota" validate:"required"`
	CreatedAt time.Time `json:"-" example:"2022-07-29T11:23:55.999+07:00"`
}
