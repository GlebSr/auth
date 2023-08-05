package teststore

import (
	"auth/internal/authApp/model"
	"auth/internal/authApp/store"
)

type UserStore struct {
	repository *UserRepository
}

func NewUserStore() store.UserStore {
	return &UserStore{}
}

func (s *UserStore) User() store.UserRepository {
	if s.repository != nil {
		return s.repository
	}
	s.repository = &UserRepository{
		Users: make(map[int]*model.User),
	}
	return s.repository
}
