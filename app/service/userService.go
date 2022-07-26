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

func (s *UserService) FindByApiKey(apiKey string) (*model.User, error) {
	return s.store.User().FindByApiKey(apiKey)
}
