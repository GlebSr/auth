package teststorage

import (
	"auth/internal/authApp/model"
	"auth/internal/authApp/storage"
)

type OauthTokenStorage struct {
	repository *OauthTokenRepository
}

func NewOauthTokenStorage() storage.OauthTokenStorage {
	return &OauthTokenStorage{}
}
func (s *OauthTokenStorage) OauthToken() storage.OauthTokenRepository {
	if s.repository != nil {
		return s.repository
	}
	s.repository = &OauthTokenRepository{
		Tokens: make([]*model.OauthToken, 0),
	}
	return s.repository
}
