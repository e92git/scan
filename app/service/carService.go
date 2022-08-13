package service

import (
	"scan/app/model"
	"scan/app/store"
)

type CarService struct {
	store *store.Store
}

func NewCar(store *store.Store) *CarService {
	return &CarService{
		store: store,
	}
}

func (s *CarService) FirstOrCreateMark(name string) (*model.CarMark, error) {
	new := &model.CarMark{
		Name: name,
	}

	return new, s.store.CarMark().FirstOrCreate(new)
}

func (s *CarService) FirstOrCreateModel(markId int, name string) (*model.CarModel, error) {
	new := &model.CarModel{
		MarkId: markId,
		Name:   name,
	}

	return new, s.store.CarModel().FirstOrCreate(new)
}
