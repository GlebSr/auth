package teststorage

import (
	"auth/internal/authApp/model"
	"auth/internal/authApp/storage"
)

type RefreshTokenStorage struct {
	repository *RefreshTokenRepository
}

func NewRefreshTokenStorage() storage.RefreshTokenStorage {
	return &RefreshTokenStorage{}
}

func (s *RefreshTokenStorage) RefreshToken() storage.RefreshTokenRepository {
	if s.repository != nil {
		return s.repository
	}
	s.repository = &RefreshTokenRepository{
		Tokens: make(map[string]*model.RefreshToken, 0),
	}
	return s.repository
}
