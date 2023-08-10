package psqlstorage

import (
	"auth/internal/authApp/storage"
	"github.com/jmoiron/sqlx"
)

type UserStorage struct {
	db         *sqlx.DB
	repository *UserRepository
}

func NewUserStorage(db *sqlx.DB) storage.UserStorage {
	return &UserStorage{
		db: db,
	}
}

func (s *UserStorage) User() storage.UserRepository {
	if s.repository != nil {
		return s.repository
	}
	s.repository = &UserRepository{
		Users: s.db,
	}
	return s.repository
}
