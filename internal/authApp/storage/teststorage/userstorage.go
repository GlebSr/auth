package teststorage

import (
	"auth/internal/authApp/model"
	"auth/internal/authApp/storage"
)

type UserStorage struct {
	repository *UserRepository
}

func NewUserStorage() storage.UserStorage {
	return &UserStorage{}
}

func (s *UserStorage) User() storage.UserRepository {
	if s.repository != nil {
		return s.repository
	}
	s.repository = &UserRepository{
		Users: make(map[int]*model.User),
	}
	return s.repository
}
