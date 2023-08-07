package psqlstorage

import (
	"auth/internal/authApp/model"
	"errors"
	"github.com/jmoiron/sqlx"
	"time"
)

type RefreshTokenRepository struct {
	Tokens *sqlx.DB
}

func (r *RefreshTokenRepository) FindByUserId(id string) ([]*model.RefreshToken, error) {
	tokens := make([]*model.RefreshToken, 0)
	err := r.Tokens.Select(&tokens, "SELECT user_id, token, expire FROM refresh WHERE user_id = $1", id)
	if err != nil {
		return nil, err
	}
	return tokens, nil
}

func (r *RefreshTokenRepository) FindByToken(token string) (*model.RefreshToken, error) {
	refreshToken := make([]*model.RefreshToken, 0)
	err := r.Tokens.Select(&refreshToken, "SELECT token, user_id, expire FROM refresh WHERE token = $1", token)
	if err != nil {
		return nil, err
	}
	if len(refreshToken) != 1 {
		return nil, errors.New("error at find") //TODO NEW err
	}
	return refreshToken[0], nil
}

func (r *RefreshTokenRepository) ClearExpired() error {
	res, err := r.Tokens.Exec(`
			DELETE FROM refresh
			WHERE expire < $1`,
		time.Now(),
	)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	//if rowsAffected == 0 {
	//	return errors.New("errr") //TODO чё делать если пусто
	//}
	return nil
}

func (r *RefreshTokenRepository) Delete(token string) error {
	res, err := r.Tokens.Exec(`
			DELETE FROM refresh
			WHERE token = $1`,
		token,
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

func (r *RefreshTokenRepository) DeleteAllById(id string) error {
	res, err := r.Tokens.Exec(`
			DELETE FROM refresh
			WHERE user_id = $1`,
		id,
	)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	//if rowsAffected == 0 {
	//	return errors.New("errr") //TODO а что делать если нет?
	//}
	return nil
}

func (r *RefreshTokenRepository) Create(token *model.RefreshToken) error {
	res, err := r.Tokens.Exec(`
			INSERT INTO refresh (user_id, token, expire) 
			VALUES ($1, $2, $3)`,
		token.UserId, token.Token, token.Expire,
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
