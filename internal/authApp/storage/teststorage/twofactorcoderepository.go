package teststorage

import (
	"auth/internal/authApp/model"
	"auth/internal/authApp/storage"
)

type TwoFactorRepository struct {
	Codes map[string]*model.TwoFactorCode
}

func (r *TwoFactorRepository) FindByUserId(id string) (*model.TwoFactorCode, error) {
	code, ok := r.Codes[id]
	if ok {
		return code, nil
	}
	return nil, storage.ErrTokenDoesNotExist
}

func (r *TwoFactorRepository) Create(code *model.TwoFactorCode) error {
	_, ok := r.Codes[code.UserId]
	if ok {
		return storage.ErrTokenAlreadyExist
	}
	r.Codes[code.UserId] = code
	return nil
}

func (r *TwoFactorRepository) Delete(id string) error {
	_, ok := r.Codes[id]
	if ok {
		delete(r.Codes, id)
		return nil
	}
	return storage.ErrTokenDoesNotExist
}
