package model

import (
	"auth/internal/authApp/config"
	"errors"
	"time"
)

type OauthToken struct {
	UserId    string    `db:"user_id"`
	Service   string    `db:"service"`
	IsRefresh bool      `db:"is_refresh"`
	Token     string    `db:"token"`
	Expire    time.Time `db:"expire"`
}

func (t *OauthToken) Validate() error {
	if (len(t.UserId) == 0 || len(t.Token) == 0 ||
		time.Now().After(t.Expire)) || !ValidService(t.Service) {
		return errors.New("bad validation") //TODO new err
	}
	return nil
}

func ValidService(service string) bool {
	for _, serv := range config.Services {
		if serv == service {
			return true
		}
	}
	return false
}
