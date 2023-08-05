package teststore

import "auth/internal/authApp/model"

type RefreshTokenRepository struct {
	Tokens []model.RefreshToken
}
