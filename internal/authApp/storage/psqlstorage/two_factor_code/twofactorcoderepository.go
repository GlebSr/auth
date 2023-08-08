package psqlstorage

import (
	"auth/internal/authApp/model"
	"auth/internal/authApp/storage"
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type TwoFactorRepository struct {
	Codes *sqlx.DB
}

type raw2FA struct {
	UserId       string         `db:"user_id"`
	SecretKey    string         `db:"secret_key"`
	ReserveCodes pq.StringArray `db:"reserve_codes"`
}

func (r *TwoFactorRepository) FindByUserId(id string) (*model.TwoFactorCode, error) {
	codes := make([]*raw2FA, 0)
	err := r.Codes.Select(&codes, "SELECT user_id, secret_key, reserve_codes FROM two_factor WHERE user_id = $1", id)
	if err != nil {
		return nil, err
	}
	if len(codes) == 0 {
		return nil, storage.ErrTokenDoesNotExist
	}
	return &model.TwoFactorCode{
		UserId:       codes[0].UserId,
		SecretKey:    codes[0].SecretKey,
		ReserveCodes: ([]string)(codes[0].ReserveCodes),
	}, err
}

func (r *TwoFactorRepository) Create(code *model.TwoFactorCode) error {
	res, err := r.Codes.Exec(`
			INSERT INTO two_factor (user_id, secret_key, reserve_codes) 
			VALUES ($1, $2, $3)`,
		code.UserId, code.SecretKey, (pq.StringArray)(code.ReserveCodes),
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

func (r *TwoFactorRepository) Delete(id string) error {
	res, err := r.Codes.Exec(`
			DELETE FROM two_factor
			WHERE user_id = $1`,
		id,
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
