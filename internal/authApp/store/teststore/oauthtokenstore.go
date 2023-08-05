package teststore

import (
	"auth/internal/authApp/model"
	"auth/internal/authApp/store"
)

type OauthTokenStore struct {
	repository *OauthTokenRepository
}

func NewOauthTokenStore() store.OauthTokenStore {
	return &OauthTokenStore{}
}
func (s *OauthTokenStore) OauthToken() store.OauthTokenRepository {
	if s.repository != nil {
		return s.repository
	}
	s.repository = &OauthTokenRepository{
		Tokens: make([]*model.OauthToken, 0),
	}
	return s.repository
}
