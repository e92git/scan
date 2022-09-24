package model

import (
	"encoding/json"
	"scan/app/helper"
	"time"
)

type CarMark struct {
	ID           int       `json:"id" example:"12"`
	Name         string    `json:"name" example:"Toyota" validate:"required"`
	NameInTires  *string   `json:"name_in_tires" example:"Toyota"`
	NameSynonyms *string   `json:"name_synonyms" example:"[Toyota,Тойота]"`
	CreatedAt    time.Time `json:"-" example:"2022-07-29T11:23:55.999+07:00"`
}

func (m *CarMark) GetSynonyms() *[]string {
	var s *[]string
    _ = json.Unmarshal([]byte(*m.NameSynonyms), &s)
	return s
}

func (m *CarMark) HasSynonym(synonym string) bool {
	return helper.InArray(m.GetSynonyms(), synonym)
}

func (m *CarMark) AddSynonym(synonym string) bool {
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
