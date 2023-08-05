package teststore

import "auth/internal/authApp/model"

type OauthTokenRepository struct {
	Tokens []model.OauthToken
}
