package storage

import "auth/internal/authApp/model"

type UserRepository interface {
	Create(*model.User) error
	FindById(string) (*model.User, error)
	FindByEmail(string) (*model.User, error)
	FindByOauthID(serviceName string, id string) (*model.User, error)
	Update(*model.User) error
	Delete(string) error
}

type RefreshTokenRepository interface {
	FindByUserId(string) ([]*model.RefreshToken, error)
	FindByToken(string) (*model.RefreshToken, error)
	ClearExpired() error
	Delete(string) error
	DeleteAll(string) error
	Create(*model.RefreshToken) error
}

type OauthTokenRepository interface {
	FindByUserId(string) ([]*model.OauthToken, error)
	Create(*model.OauthToken) error
	Update(*model.OauthToken) error
	Delete(*model.OauthToken) error
	FindByUserIdAndService(string, string) ([]*model.OauthToken, error)
}

type TwoFactorRepository interface {
	FindByUserId(string) (*model.TwoFactorCode, error)
	Create(*model.TwoFactorCode) error
	Delete(string) error
}
