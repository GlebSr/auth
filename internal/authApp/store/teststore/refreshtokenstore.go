package teststore

import (
	"auth/internal/authApp/model"
	"auth/internal/authApp/store"
)

type RefreshTokenStore struct {
	repository *RefreshTokenRepository
}

func NewRefreshTokenStore() store.RefreshTokenStore {
	return &RefreshTokenStore{}
}

func (s *RefreshTokenStore) RefreshToken() store.RefreshTokenRepository {
	if s.repository != nil {
		return s.repository
	}
	s.repository = &RefreshTokenRepository{
		Tokens: make(map[string]*model.RefreshToken, 0),
	}
	return s.repository
}
