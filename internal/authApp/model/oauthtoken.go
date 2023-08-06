package model

import (
	"auth/internal/authApp/config"
	"errors"
	"time"
)

type OauthToken struct {
	UserId    string
	Service   string
	IsRefresh bool
	Token     string
	Expire    time.Time
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
