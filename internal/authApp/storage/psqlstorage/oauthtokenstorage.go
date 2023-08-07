package psqlstorage

import (
	"auth/internal/authApp/storage"
	"github.com/jmoiron/sqlx"
)

type OauthTokenStorage struct {
	db         *sqlx.DB
	repository *OauthTokenRepository
}

func NewOauthTokenStorage(db *sqlx.DB) storage.OauthTokenStorage {
	return &OauthTokenStorage{
		db: db,
	}
}
func (s *OauthTokenStorage) OauthToken() storage.OauthTokenRepository {
	if s.repository != nil {
		return s.repository
	}
	s.repository = &OauthTokenRepository{
		Tokens: s.db,
	}
	return s.repository
}
