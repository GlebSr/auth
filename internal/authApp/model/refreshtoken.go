package model

import "time"

type RefreshToken struct {
	Token  string
	UserId int
	Expiry time.Time
}

//TODO token generator
