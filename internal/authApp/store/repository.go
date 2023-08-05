package store

import "auth/internal/authApp/model"

type UserRepository interface {
	Create(*model.User) error
	FindById(int) (*model.User, error)
	FindByEmail(string) (*model.User, error)
	FindByOauthID(serviceName string, Id string) (*model.User, error)
	Update(*model.User) error
	Delete(int) error
}

type RefreshTokenRepository interface {
	FindByUserId(int) ([]*model.RefreshToken, error)
	FindUser(string) (*model.User, error)
	FindToken(string) (*model.RefreshToken, error)
	ClearExpired() error
	Delete(string) error
	DeleteAll(int) error
	Create(int) (*model.RefreshToken, error)
}

type OauthTokenRepository interface {
	ClearExpired() error
	FindByUserId(int) ([]model.OauthToken, error)
	Create(*model.OauthToken) error
	Update(*model.OauthToken) error
	Delete(*model.OauthToken) error
	FindByUserIdAndService(int, string) ([]model.OauthToken, error)
}

type TwoFactorRepository interface {
	FindByUserId(int) (model.TwoFactorCode, error)
	Create(*model.TwoFactorCode) error
	Delete(int) error
}
