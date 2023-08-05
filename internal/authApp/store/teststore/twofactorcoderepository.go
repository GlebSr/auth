package teststore

import "auth/internal/authApp/model"

type TwoFactorRepository struct {
	Codes map[int]model.TwoFactorCode
}
