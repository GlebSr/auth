package teststore

import (
	"auth/internal/authApp/model"
	"auth/internal/authApp/store"
)

type TwoFactorStore struct {
	repository *TwoFactorRepository
}

func NewTwoFactorStore() store.TwoFactorStore {
	return &TwoFactorStore{}
}

func (t *TwoFactorStore) TwoFactorCode() store.TwoFactorRepository {
	if t.repository != nil {
		return t.repository
	}
	t.repository = &TwoFactorRepository{
		Codes: make(map[int]*model.TwoFactorCode),
	}
	return t.repository
}
