package model

type CarMark struct {
	ID        int    `json:"id" example:"12"`
	Name      string `json:"name" example:"Toyota"`
	CreatedAt string `json:"created_at" example:"2022-07-28 11:23:55"`
}
