package storage

type UserStorage interface {
	User() UserRepository
}

type RefreshTokenStorage interface {
	RefreshToken() RefreshTokenRepository
}

type OauthTokenStorage interface {
	OauthToken() OauthTokenRepository
}

type TwoFactorStorage interface {
	TwoFactorCode() TwoFactorRepository
}

type Storage struct {
	UserStorage         UserStorage
	RefreshTokenStorage RefreshTokenStorage
	OauthTokenStorage   OauthTokenStorage
	TwoFactorStorage    TwoFactorStorage
}

func (s *Storage) User() UserRepository {
	return s.UserStorage.User()
}
func (s *Storage) RefreshToken() RefreshTokenRepository {
	return s.RefreshTokenStorage.RefreshToken()
}
func (s *Storage) OauthToken() OauthTokenRepository {
	return s.OauthTokenStorage.OauthToken()
}
func (s *Storage) TwoFactorCode() TwoFactorRepository {
	return s.TwoFactorStorage.TwoFactorCode()
}
