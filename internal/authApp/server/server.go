package server

import (
	"auth/internal/authApp/storage"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type server struct {
	router  *mux.Router
	storage *storage.Storage
}

type ctxKey int

func NewServer(storage *storage.Storage) *server {
	s := &server{
		router:  mux.NewRouter(),
		storage: storage,
	}
	s.configureRouter()
	return s
}

func (s *server) configureRouter() {

	v1Router := s.router.PathPrefix("/auth/v1").Subrouter()
	v1Router.HandleFunc("/register", s.handleRegistration()).Methods("POST")
	v1Router.HandleFunc("/login", s.handleLogin())
	v1Router.HandleFunc("/logout", s.handleLogout())
	v1Router.HandleFunc("/refresh", s.handleRefresh())
	googleV1 := v1Router.PathPrefix("/oauth/google").Subrouter()
	googleV1.HandleFunc("/login", s.handleGoogleLogin())
	googleV1.HandleFunc("/callback", s.handelGoogleCallBack())
	yandexV1 := v1Router.PathPrefix("/oauth/yandex").Subrouter()
	yandexV1.HandleFunc("/login", s.handleYandexLogin())
	yandexV1.HandleFunc("/callback", s.handelYandexCallBack())
	vkV1 := v1Router.PathPrefix("/oauth/vk").Subrouter()
	vkV1.HandleFunc("/login", s.handleVkLogin())
	vkV1.HandleFunc("/callback", s.handelVkCallBack())
	githubV1 := v1Router.PathPrefix("/oauth/github").Subrouter()
	githubV1.HandleFunc("/login", s.handleGithubLogin())
	githubV1.HandleFunc("/callback", s.handelGithubCallBack())

}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *server) respond(w http.ResponseWriter, r *http.Request, code int, data any) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
