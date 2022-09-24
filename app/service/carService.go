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

func (s *CarService) FirstMarkByName(name string) (*model.CarMark, error) {
	mark := &model.CarMark{}
	return mark, s.store.CarMark().FindByName(mark, name)
}

func (s *CarService) FirstOrCreateModel(markId int, name string) (*model.CarModel, error) {
	new := &model.CarModel{
		MarkId: markId,
		Name:   name,
	}

	return new, s.store.CarModel().FirstOrCreate(new)
}

func (s *CarService) FirstModelByName(name string) (*model.CarModel, error) {
	m := &model.CarModel{}
	return m, s.store.CarModel().FindByName(m, name)
}