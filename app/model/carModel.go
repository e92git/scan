package model

import "time"

type CarModel struct {
	ID           int       `json:"id" example:"123"`
	MarkId       int       `json:"mark_id" example:"12" validate:"required"`
	Name         string    `json:"name" example:"Prius" validate:"required"`
	NameInTires  *string   `json:"name_in_tires" example:"Prius"`
	NameSynonyms *string   `json:"name_synonyms" example:"[Prius,Приус,PRIUS II]"`
	CreatedAt    time.Time `json:"-" example:"2022-07-29T11:23:55.999+07:00"`
}
