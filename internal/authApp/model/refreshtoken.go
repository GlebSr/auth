package model

import "time"

type RefreshToken struct {
	UserId int
	token  string
	expiry time.Time
}
