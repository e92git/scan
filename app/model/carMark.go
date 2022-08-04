package model

import "time"

type CarMark struct {
	ID        int       `json:"id" example:"12"`
	Name      string    `json:"name" example:"Toyota"`
	CreatedAt time.Time `json:"created_at" example:"2022-07-29T11:23:55+07:00"`
}
