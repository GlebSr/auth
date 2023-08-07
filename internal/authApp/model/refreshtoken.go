package model

import (
	"auth/internal/authApp/config"
	"github.com/google/uuid"
	"time"
)

type RefreshToken struct {
	Token  string    `db:"token"`
	UserId string    `db:"user_id"`
	Expire time.Time `db:"expire"`
}

func NewRefreshToken(UserId string) *RefreshToken {
	return &RefreshToken{
		Token:  uuid.New().String(),
		UserId: UserId,
		Expire: time.Now().Add(config.RefreshTokenLifeTime),
	}
}
