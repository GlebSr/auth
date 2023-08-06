package model

import (
	"auth/internal/authApp/config"
	"github.com/google/uuid"
	"time"
)

type RefreshToken struct {
	Token  string
	UserId string
	Expiry time.Time
}

func NewRefreshToken(UserId string) *RefreshToken {
	return &RefreshToken{
		Token:  uuid.New().String(),
		UserId: UserId,
		Expiry: time.Now().Add(config.RefreshTokenLifeTime),
	}
}
