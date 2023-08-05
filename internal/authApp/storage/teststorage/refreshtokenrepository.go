package teststorage

import (
	"auth/internal/authApp/model"
	"auth/internal/authApp/storage"
	"time"
)

type RefreshTokenRepository struct {
	Tokens map[string]*model.RefreshToken
}

func (r *RefreshTokenRepository) FindByUserId(id int) ([]*model.RefreshToken, error) {
	tokens := make([]*model.RefreshToken, 0)
	for _, t := range r.Tokens {
		if t.UserId == id {
			tokens = append(tokens, t)
		}
	}
	return tokens, nil
}

func (r *RefreshTokenRepository) FindByToken(token string) (*model.RefreshToken, error) {
	t, ok := r.Tokens[token]
	if ok {
		return t, nil
	}
	return nil, storage.ErrTokenDoesNotExist
}

func (r *RefreshTokenRepository) ClearExpired() error {
	for _, t := range r.Tokens {
		if time.Now().After(t.Expiry) {
			delete(r.Tokens, t.Token)
		}
	}
	return nil
}

func (r *RefreshTokenRepository) Delete(token string) error {
	_, ok := r.Tokens[token]
	if ok {
		delete(r.Tokens, token)
		return nil
	}
	return storage.ErrTokenDoesNotExist
}

func (r *RefreshTokenRepository) DeleteAll(id int) error {
	for _, t := range r.Tokens {
		if t.UserId == id {
			delete(r.Tokens, t.Token)
		}
	}
	return nil
}

func (r *RefreshTokenRepository) Create(token *model.RefreshToken) error {
	_, ok := r.Tokens[token.Token]
	if ok {
		return storage.ErrTokenAlreadyExist
	}
	r.Tokens[token.Token] = token
	return nil
}
