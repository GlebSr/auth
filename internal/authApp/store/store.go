package store

type UserStore interface {
	User() UserRepository
}

type RefreshTokenStore interface {
	RefreshToken() RefreshTokenRepository
}

type OauthTokenStore interface {
	OauthToken() OauthTokenRepository
}

type TwoFactorStore interface {
	TwoFactorCode() TwoFactorRepository
}

type Store struct {
	userStore         UserStore
	refreshTokenStore RefreshTokenStore
	authTokenStore    OauthTokenStore
	twoFactorStore    TwoFactorStore
}

func (s *Store) User() UserRepository {
	return s.userStore.User()
}
func (s *Store) RefreshToken() RefreshTokenRepository {
	return s.refreshTokenStore.RefreshToken()
}
func (s *Store) OauthToken() OauthTokenRepository {
	return s.authTokenStore.OauthToken()
}
func (s *Store) TwoFactorCode() TwoFactorRepository {
	return s.twoFactorStore.TwoFactorCode()
}
