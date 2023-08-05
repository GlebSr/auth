package store

import "auth/internal/authApp/model"

type UserRepository interface {
	Create(*model.User) error
	FindById(int) (*model.User, error)
	FindByEmail(string) (*model.User, error)
	FindByOauthID(serviceName string, id string) (*model.User, error)
	Update(*model.User) error
	Delete(int) error
}

type RefreshTokenRepository interface {
	FindByUserId(int) ([]*model.RefreshToken, error)
	FindByToken(string) (*model.RefreshToken, error)
	ClearExpired() error
	Delete(string) error
	DeleteAll(int) error
	Create(*model.RefreshToken) error
}

type OauthTokenRepository interface {
	FindByUserId(int) ([]*model.OauthToken, error)
	Create(*model.OauthToken) error
	Update(*model.OauthToken) error
	Delete(*model.OauthToken) error
	FindByUserIdAndService(int, string) ([]*model.OauthToken, error)
}

type TwoFactorRepository interface {
	FindByUserId(int) (*model.TwoFactorCode, error)
	Create(*model.TwoFactorCode) error
	Delete(int) error
}
