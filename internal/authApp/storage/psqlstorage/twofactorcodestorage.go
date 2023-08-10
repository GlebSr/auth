package psqlstorage

import (
	"auth/internal/authApp/storage"
	"github.com/jmoiron/sqlx"
)

type TwoFactorStorage struct {
	db         *sqlx.DB
	repository *TwoFactorRepository
}

func NewTwoFactorStorage(db *sqlx.DB) storage.TwoFactorStorage {
	return &TwoFactorStorage{
		db: db,
	}
}

func (s *TwoFactorStorage) TwoFactorCode() storage.TwoFactorRepository {
	if s.repository != nil {
		return s.repository
	}
	s.repository = &TwoFactorRepository{
		Codes: s.db,
	}
	return s.repository
}
