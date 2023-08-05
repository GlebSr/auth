package teststorage

import (
	"auth/internal/authApp/model"
	"auth/internal/authApp/storage"
)

type TwoFactorStorage struct {
	repository *TwoFactorRepository
}

func NewTwoFactorStorage() storage.TwoFactorStorage {
	return &TwoFactorStorage{}
}

func (t *TwoFactorStorage) TwoFactorCode() storage.TwoFactorRepository {
	if t.repository != nil {
		return t.repository
	}
	t.repository = &TwoFactorRepository{
		Codes: make(map[int]*model.TwoFactorCode),
	}
	return t.repository
}
