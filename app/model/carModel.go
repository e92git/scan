package model

type CarModel struct {
	ID        int    `json:"id" example:"123"`
	MarkId    string `json:"mark_id" example:"12"`
	Name      string `json:"name" example:"Kalina"`
	CreatedAt string `json:"created_at" example:"2022-07-28 11:23:55"`
}
