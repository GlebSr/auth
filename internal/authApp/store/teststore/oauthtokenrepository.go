package teststore

import (
	"auth/internal/authApp/model"
	"errors"
)

type OauthTokenRepository struct {
	Tokens []*model.OauthToken
}

func (r *OauthTokenRepository) FindByUserId(id int) ([]*model.OauthToken, error) {
	users := make([]*model.OauthToken, 0)
	for _, token := range r.Tokens {
		if token.UserId == id {
			users = append(users, token)
		}
	}
	return users, nil
}
func (r *OauthTokenRepository) Create(token *model.OauthToken) error {
	if err := token.Validate(); err != nil {
		return err
	}
	tokens, err := r.FindByUserIdAndService(token.UserId, token.Service)
	if err != nil {
		return err
	}
	for _, t := range tokens {
		if t.UserId == token.UserId && t.Service == token.Service && t.IsRefresh == token.IsRefresh {
			return errors.New("token exist") //TODO new err
		}
	}
	r.Tokens = append(r.Tokens, token)
	return nil
}

func (r *OauthTokenRepository) Update(token *model.OauthToken) error {
	if err := token.Validate(); err != nil {
		return err
	}
	tokens, err := r.FindByUserIdAndService(token.UserId, token.Service)
	if err != nil {
		return err
	}
	for _, t := range tokens {
		if t.UserId == token.UserId && t.Service == token.Service && t.IsRefresh == token.IsRefresh {
			t = token
			return nil
		}
	}
	return errors.New("token not exist") //TODO new err
}

func (r *OauthTokenRepository) Delete(token *model.OauthToken) error {
	fl := false
	for pos, t := range r.Tokens {
		if fl {
			r.Tokens[pos-1] = t
		} else {
			if t.UserId == token.UserId && t.Service == token.Service && t.IsRefresh == token.IsRefresh {
				fl = true
			}
		}
	}
	if !fl {
		return errors.New("user not exist") //TODO new err
	}
	r.Tokens = r.Tokens[:len(r.Tokens)-1]
	return nil
}

func (r *OauthTokenRepository) FindByUserIdAndService(id int, service string) ([]*model.OauthToken, error) {
	users := make([]*model.OauthToken, 0)
	for _, token := range r.Tokens {
		if token.UserId == id && token.Service == service {
			users = append(users, token)
		}
	}
	return users, nil
}
