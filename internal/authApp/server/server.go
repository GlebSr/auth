package server

import (
	"auth/internal/authApp/store"
	"github.com/gorilla/mux"
)

type Server struct {
	router            mux.Router
	userStore         store.UserStore
	refreshTokenStore store.RefreshTokenStore
	oauthTokenStore   store.OauthTokenStore
	twoFactorStore    store.TwoFactorStore
}
