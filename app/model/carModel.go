package model

import (
	"encoding/json"
	"scan/app/helper"
	"time"
)

type CarModel struct {
	ID           int       `json:"id" example:"123"`
	MarkId       int       `json:"mark_id" example:"12" validate:"required"`
	Name         string    `json:"name" example:"Prius" validate:"required"`
	NameInTires  *string   `json:"name_in_tires" example:"Prius"`
	NameSynonyms *string   `json:"name_synonyms" example:"[Prius,Приус,PRIUS II]"`
	CreatedAt    time.Time `json:"-" example:"2022-07-29T11:23:55.999+07:00"`
}

func (m *CarModel) GetSynonyms() *[]string {
	var s *[]string
	if m.NameSynonyms != nil {
		_ = json.Unmarshal([]byte(*m.NameSynonyms), &s)
	}
	return s
}

func (m *CarModel) HasSynonym(synonym string) bool {
	return helper.InArray(m.GetSynonyms(), synonym)
}

func (m *CarModel) AddSynonym(synonym string) bool {
	if m.HasSynonym(synonym) {
		return false
	}
	var synonyms []string
	s := m.GetSynonyms()
	if s != nil {
		synonyms = *s
	}
	synonyms = append(synonyms, synonym)
	j, _ := json.Marshal(synonyms)
	r := string(j)
	m.NameSynonyms = &r

	return true
}
