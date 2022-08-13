package model

import "time"

type CarModel struct {
	ID        int       `json:"id" example:"123"`
	MarkId    int       `json:"mark_id" example:"12"`
	Name      string    `json:"name" example:"Kalina"`
	CreatedAt time.Time `json:"created_at" example:"2022-07-29T11:23:55+07:00"`
}
