package store

type UserStore interface {
	User() UserRepository
}

type RefreshTokenStore interface {
	Token() RefreshTokenRepository
}

type OauthTokenStore interface {
	Token() OauthTokenRepository
}

type TwoFactorStore interface {
	Code() TwoFactorRepository
}
