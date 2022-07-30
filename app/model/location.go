package model

type Location struct {
	ID   int64  `json:"id" example:"12"`
	Code string `json:"code" example:"pokrovka"`
	Name string `json:"name" example:"Красноярск Покровка"`
}
