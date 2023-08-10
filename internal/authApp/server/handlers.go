package server

import (
	"auth/internal/authApp/model"
	"auth/internal/authApp/storage"
	"encoding/json"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

func (s *server) handleRegistration() http.HandlerFunc {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("RegistrationHandler")
		if _, err := r.Cookie("refresh"); err == nil {
			s.error(w, r, http.StatusBadRequest, errors.New("user loggined")) //TODO new err
			return
		}
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		users, err := s.storage.User().FindByEmail(req.Email)
		if users != nil || err != storage.ErrUserDoesNotExist {
			s.error(w, r, http.StatusConflict, storage.ErrUserAlreadyExist)
			return
		}

		user := &model.User{
			Email:    req.Email,
			Password: req.Password,
		}
		if err := s.storage.User().Create(user); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		refreshTok := model.NewRefreshToken(user.Id)
		if err := s.storage.RefreshToken().Create(refreshTok); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		jwt, err := GenerateJWT(user)
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		response := struct {
			JwtTok       string `json:"jwt"`
			RefreshToken string `json:"refresh_token"`
		}{
			JwtTok:       jwt,
			RefreshToken: refreshTok.Token,
		}
		s.respond(w, r, http.StatusOK, response)
	}
}

func (s *server) handleLogin() http.HandlerFunc {

	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("LoginHandler")
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		user, err := s.storage.User().FindByEmail(req.Email)
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		if user == nil {
			s.error(w, r, http.StatusBadRequest, storage.ErrUserDoesNotExist)
			return
		}

		if err := bcrypt.CompareHashAndPassword(
			[]byte(user.EncryptedPassword), []byte(req.Password)); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		refreshTok := model.NewRefreshToken(user.Id)
		if err := s.storage.RefreshToken().Create(refreshTok); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		jwt, err := GenerateJWT(user)
		response := struct {
			JwtTok       string `json:"jwt"`
			RefreshToken string `json:"refresh_token"`
		}{
			JwtTok:       jwt,
			RefreshToken: refreshTok.Token,
		}
		s.respond(w, r, http.StatusOK, response)
	}
}

func (s *server) handleLogout() http.HandlerFunc {
	type request struct {
		JwtTok       string `json:"jwt"`
		RefreshToken string `json:"refresh_token"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("LogoutHandler")
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		if err := s.storage.RefreshToken().Delete(req.RefreshToken); err != nil {
			s.error(w, r, http.StatusUnauthorized, err)
			return
		}
		s.respond(w, r, http.StatusOK, nil)
	}
}

func (s *server) handleRefresh() http.HandlerFunc {
	type request struct {
		RefreshToken string `json:"refresh_token"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("RefreshHandler")
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		token, err := s.storage.RefreshToken().FindByToken(req.RefreshToken)
		user, err := s.storage.User().FindById(token.UserId)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}
		jwt, err := GenerateJWT(user)
		response := struct {
			JwtTok string `json:"jwt"`
		}{
			JwtTok: jwt,
		}
		s.respond(w, r, http.StatusOK, response)
	}
}
