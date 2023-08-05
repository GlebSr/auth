package teststore

import (
	"auth/internal/authApp/model"
	"errors"
)

type TwoFactorRepository struct {
	Codes map[int]*model.TwoFactorCode
}

func (r *TwoFactorRepository) FindByUserId(id int) (*model.TwoFactorCode, error) {
	code, ok := r.Codes[id]
	if ok {
		return code, nil
	}
	return nil, errors.New("code not exist") //TODO new err
}

func (r *TwoFactorRepository) Create(code *model.TwoFactorCode) error {
	_, ok := r.Codes[code.UserId]
	if ok {
		return errors.New("code exist") //TODO new err
	}
	r.Codes[code.UserId] = code
	return nil
}

func (r *TwoFactorRepository) Delete(id int) error {
	_, ok := r.Codes[id]
	if ok {
		delete(r.Codes, id)
		return nil
	}
	return errors.New("code not exist") //TODO new err
}
