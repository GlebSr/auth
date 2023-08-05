package server

import (
	"auth/internal/authApp/store"
	"github.com/gorilla/mux"
)

type server struct {
	router *mux.Router
	store  store.Store
}

type ctxKey int

func NewServer(store store.Store) *server {
	s := &server{
		router: mux.NewRouter(),
		store:  store,
	}
	s.configureRouter()
	return s
}

func (s *server) configureRouter() {

}
