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
	userStorage         UserStorage
	refreshTokenStorage RefreshTokenStorage
	authTokenStorage    OauthTokenStorage
	twoFactorStorage    TwoFactorStorage
}

func (s *Storage) User() UserRepository {
	return s.userStorage.User()
}
func (s *Storage) RefreshToken() RefreshTokenRepository {
	return s.refreshTokenStorage.RefreshToken()
}
func (s *Storage) OauthToken() OauthTokenRepository {
	return s.authTokenStorage.OauthToken()
}
func (s *Storage) TwoFactorCode() TwoFactorRepository {
	return s.twoFactorStorage.TwoFactorCode()
}
