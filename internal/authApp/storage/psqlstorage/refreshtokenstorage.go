package psqlstorage

import (
	"auth/internal/authApp/storage"
	"github.com/jmoiron/sqlx"
)

type RefreshTokenStorage struct {
	db         *sqlx.DB
	repository *RefreshTokenRepository
}

func NewRefreshTokenStorage(db *sqlx.DB) storage.RefreshTokenStorage {
	return &RefreshTokenStorage{
		db: db,
	}
}

func (s *RefreshTokenStorage) RefreshToken() storage.RefreshTokenRepository {
	if s.repository != nil {
		return s.repository
	}
	s.repository = &RefreshTokenRepository{
		Tokens: s.db,
	}
	return s.repository
}
