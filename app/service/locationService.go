package service

import (
	"scan/app/model"
	"scan/app/store"
)

type LocationService struct {
	store *store.Store
}

func NewLocation(store *store.Store) *LocationService {
	return &LocationService{
		store: store,
	}
}

func (s *LocationService) All() ([]model.Location, error) {
	return s.store.Location().All()
}

func (s *LocationService) FindByCode(code string) (*model.Location, error) {
	return s.store.Location().FindByCode(code)
}
