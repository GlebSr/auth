package model

import (
	"errors"
	"time"
)

type OauthToken struct {
	UserId    string
	Service   string
	IsRefresh bool
	Token     string
	expire    time.Time
}

func (t *OauthToken) Validate() error {
	if (len(t.UserId) == 0 || len(t.Token) == 0 || time.Now().After(t.expire)) ||
		t.Service != "Google" && t.Service != "Yandex" && t.Service != "Vk" && t.Service != "Github" {
		return errors.New("bad validation") //TODO new err
	}
	return nil
}
