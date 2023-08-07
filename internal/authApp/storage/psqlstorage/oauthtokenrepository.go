package psqlstorage

import (
	"auth/internal/authApp/model"
	"auth/internal/authApp/storage"
	"errors"
	"github.com/jmoiron/sqlx"
	"strings"
)

type OauthTokenRepository struct {
	Tokens *sqlx.DB
}

func (r *OauthTokenRepository) FindByUserId(id string) ([]*model.OauthToken, error) {
	tokens := make([]*model.OauthToken, 0)
	err := r.Tokens.Select(&tokens, "SELECT user_id, service, is_refresh, token, expire FROM oauth WHERE user_id = $1", id)
	if err != nil {
		return nil, err
	}
	return tokens, nil
}

func (r *OauthTokenRepository) Create(token *model.OauthToken) error { //TODO мб можно переделать без поиска
	if err := token.Validate(); err != nil {
		return err
	}
	tokens, err := r.FindByUserIdAndService(token.UserId, token.Service)
	if err != nil {
		return err
	}
	for _, t := range tokens {
		if t.IsRefresh == token.IsRefresh {
			return storage.ErrTokenAlreadyExist
		}
	}
	res, err := r.Tokens.Exec(`
			INSERT INTO oauth (user_id, service, is_refresh, token, expire) 
			VALUES ($1, $2, $3, $4, $5)`,
		token.UserId, token.Service, token.IsRefresh, token.Token, token.Expire,
	)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("errr") //TODO new error
	}
	return nil
}

func (r *OauthTokenRepository) Update(token *model.OauthToken) error {
	if err := token.Validate(); err != nil {
		return err
	}
	res, err := r.Tokens.Exec(`
			UPDATE oauth 
			SET token = $1, expire = $2
			WHERE user_id = $3 AND service = $4 AND is_refresh = $5`,
		token.Token, token.Expire, token.UserId, token.Service, token.IsRefresh,
	)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("errr") //TODO new error
	}
	return nil
}

func (r *OauthTokenRepository) Delete(token *model.OauthToken) error {
	res, err := r.Tokens.Exec(`
			DELETE FROM oauth
			WHERE user_id = $1 AND service = $2 AND is_refresh = $3`,
		token.UserId, token.Service, token.IsRefresh,
	)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("errr") //TODO new error
	}
	return nil
}

func (r *OauthTokenRepository) FindByUserIdAndService(id string, service string) ([]*model.OauthToken, error) {
	service = strings.ToLower(service)
	if !model.ValidService(service) {
		return nil, storage.ErrServiceNotSupported
	}
	tokens := make([]*model.OauthToken, 0)
	err := r.Tokens.Select(&tokens, `
		SELECT user_id, service, is_refresh, token, expire 
		FROM oauth 
		WHERE user_id = $1 AND service = $2`, id, service)
	if err != nil {
		return nil, err
	}
	return tokens, nil
}
