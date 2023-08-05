package model

import "time"

type OauthToken struct {
	UserId    int
	Service   string
	IsRefresh bool
	Token     string
	expire    time.Time
}
