package service

import (
	"scan/app/model"
	"scan/app/store"
)

type UserService struct {
	store *store.Store
}

func NewUser(store *store.Store) *UserService {
	return &UserService{
		store: store,
	}
}

func (s *UserService) FindBySession(session string) (*model.User, error) {
	return s.store.User().FindBySession(session)
}
