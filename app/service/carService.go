package service

import (
	"errors"
	"scan/app/model"
	"scan/app/store"

	"gorm.io/gorm"
)

type CarService struct {
	store *store.Store
}

func NewCar(store *store.Store) *CarService {
	return &CarService{
		store: store,
	}
}

func (s *CarService) SaveMark(m *model.CarMark) error {
	return s.store.CarMark().Save(m)
}

func (s *CarService) SaveModel(m *model.CarModel) error {
	return s.store.CarModel().Save(m)
}

func (s *CarService) IsNeedInsertSynonymMark(m *model.CarMark, newSynonym *string) (bool, error) {
	if newSynonym == nil || m == nil {
		return false, nil
	}
	// Если такой синоним уже есть в этой марке
	if m.HasSynonym(*newSynonym) {
		return false, nil
	}
	// Если такой синоним уже есть в другой марке
	oldMark := &model.CarMark{}
	err := s.store.CarMark().FindByName(oldMark, *newSynonym)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return true, nil
		}
		return false, err
	}
	if oldMark.ID == m.ID {
		return true, nil
	}
	return false, nil
}

func (s *CarService) IsNeedInsertSynonymModel(m *model.CarModel, newSynonym *string) (bool, error) {
	if newSynonym == nil || m == nil {
		return false, nil
	}
	// Если такой синоним уже есть в этой марке
	if m.HasSynonym(*newSynonym) {
		return false, nil
	}
	// Если такой синоним уже есть в другой марке
	oldModel := &model.CarModel{}
	err := s.store.CarModel().FindByName(oldModel, m.MarkId, *newSynonym)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return true, nil
		}
		return false, err
	}
	if oldModel.ID == m.ID {
		return true, nil
	}
	return false, nil
}

func (s *CarService) FirstMark(name string) (*model.CarMark, error) {
	new := &model.CarMark{
		Name: name,
	}

	return new, s.store.CarMark().First(new)
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

func (s *CarService) FirstModel(markId int, name string) (*model.CarModel, error) {
	new := &model.CarModel{
		MarkId: markId,
		Name:   name,
	}

	return new, s.store.CarModel().First(new)
}

func (s *CarService) FirstModelByName(markId int, name string) (*model.CarModel, error) {
	m := &model.CarModel{}
	return m, s.store.CarModel().FindByName(m, markId, name)
}
