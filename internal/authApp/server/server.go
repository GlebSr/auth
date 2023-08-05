package server

import (
	"auth/internal/authApp/storage"
	"github.com/gorilla/mux"
)

type server struct {
	router  *mux.Router
	storage storage.Storage
}

type ctxKey int

func NewServer(storage storage.Storage) *server {
	s := &server{
		router:  mux.NewRouter(),
		storage: storage,
	}
	s.configureRouter()
	return s
}

func (s *server) configureRouter() {

}
