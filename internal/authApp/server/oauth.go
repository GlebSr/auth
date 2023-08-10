package server

import (
	"auth/internal/authApp/config"
	"auth/internal/authApp/model"
	"auth/internal/authApp/storage"
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func createState(w http.ResponseWriter) string {
	var expiration = time.Now().Add(365 * 24 * time.Hour)
	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)
	cookie := &http.Cookie{Name: "oauth_state", Value: state, Expires: expiration}
	http.SetCookie(w, cookie)
	return state
}

func (s *server) handleGoogleLogin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("GoogleLoginHandler")
		u := config.GoogleOauth.AuthCodeURL(createState(w))
		http.Redirect(w, r, u, http.StatusTemporaryRedirect)
	}
}

func (s *server) handelGoogleCallBack() http.HandlerFunc {
	const oauthGoogleAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="
	type User struct {
		ID            string `json:"id"`
		Email         string `json:"email"`
		VerifiedEmail bool   `json:"verified_email"`
		Name          string `json:"name"`
		GivenName     string `json:"given_name"`
		FamilyName    string `json:"family_name"`
		Picture       string `json:"picture"`
		Locale        string `json:"locale"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("GoogleCallBackHandler")
		oauthState, _ := r.Cookie("oauth_state")
		if r.FormValue("state") != oauthState.Value {
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			return
		}

		token, err := config.GoogleOauth.Exchange(context.Background(), r.FormValue("code"))
		if err != nil {
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			return
		}
		response, err := http.Get(oauthGoogleAPI + token.AccessToken)
		if err != nil {
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			return
		}
		defer response.Body.Close()
		newUser := &User{}
		if err := json.NewDecoder(response.Body).Decode(newUser); err != nil {
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			return
		}
		user, err := s.storage.User().FindByOauthID("google", newUser.ID)
		if err != nil {
			if err != storage.ErrUserDoesNotExist {
				http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
				return
			}
			user = &model.User{
				GoogleId: newUser.ID,
			}
			err = s.storage.User().Create(user)
			if err != nil {
				http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
				return
			}
		}
		refreshTok := model.NewRefreshToken(user.Id)
		if err := s.storage.RefreshToken().Create(refreshTok); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		jwt, _ := GenerateJWT(user)
		cookieR := &http.Cookie{
			Name:     "refresh_token",
			Value:    refreshTok.Token,
			Path:     "/",
			HttpOnly: false,
			Secure:   false,
			Expires:  time.Now().Add(24 * time.Hour),
		}
		cookieJ := &http.Cookie{
			Name:     "jwt",
			Value:    jwt,
			Path:     "/",
			HttpOnly: false,
			Secure:   false,
			Expires:  time.Now().Add(24 * time.Hour),
		}
		http.SetCookie(w, cookieR)
		http.SetCookie(w, cookieJ)
		http.Redirect(w, r, "/set", http.StatusTemporaryRedirect)
		return

	}
}

func (s *server) handleYandexLogin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("YandexLoginHandler")
		u := config.YandexOauth.AuthCodeURL(createState(w))
		http.Redirect(w, r, u, http.StatusTemporaryRedirect)
	}
}

func (s *server) handelYandexCallBack() http.HandlerFunc {
	const oauthYandexAPI = "https://login.yandex.ru/info?format=json"
	type User struct {
		Login    string `json:"login"`
		ID       string `json:"id"`
		ClientID string `json:"client_id"`
		Psuid    string `json:"psuid"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("YandexCallBackHandler")
		oauthState, _ := r.Cookie("oauth_state")
		if r.FormValue("state") != oauthState.Value {
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			return
		}

		token, err := config.YandexOauth.Exchange(context.Background(), r.FormValue("code"))
		if err != nil {
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			return
		}

		req, err := http.NewRequest("GET", oauthYandexAPI, nil)
		if err != nil {
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			return
		}
		req.Header.Add("Authorization", fmt.Sprintf("OAuth %s", token.AccessToken))
		response, err := (&http.Client{}).Do(req)
		if err != nil {
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			return
		}
		defer response.Body.Close()
		newUser := &User{}
		if err := json.NewDecoder(response.Body).Decode(newUser); err != nil {
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			return
		}
		user, err := s.storage.User().FindByOauthID("yandex", newUser.ID)
		if err != nil {
			if err != storage.ErrUserDoesNotExist {
				http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
				return
			}
			user = &model.User{
				YandexId: newUser.ID,
			}
			err = s.storage.User().Create(user)
			if err != nil {
				http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
				return
			}
		}
		refreshTok := model.NewRefreshToken(user.Id)
		if err := s.storage.RefreshToken().Create(refreshTok); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		jwt, _ := GenerateJWT(user)
		cookieR := &http.Cookie{
			Name:     "refresh_token",
			Value:    refreshTok.Token,
			Path:     "/",
			HttpOnly: false,
			Secure:   false,
			Expires:  time.Now().Add(24 * time.Hour),
		}
		cookieJ := &http.Cookie{
			Name:     "jwt",
			Value:    jwt,
			Path:     "/",
			HttpOnly: false,
			Secure:   false,
			Expires:  time.Now().Add(24 * time.Hour),
		}
		http.SetCookie(w, cookieR)
		http.SetCookie(w, cookieJ)
		http.Redirect(w, r, "/set", http.StatusTemporaryRedirect)
		return

	}
}

func (s *server) handleVkLogin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("VkLoginHandler")
		u := config.VkOauth.AuthCodeURL(createState(w))
		http.Redirect(w, r, u, http.StatusTemporaryRedirect)
	}
}

func (s *server) handelVkCallBack() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("VkCallBackHandler")
		oauthState, _ := r.Cookie("oauth_state")
		if r.FormValue("state") != oauthState.Value {
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			return
		}

		token, err := config.VkOauth.Exchange(context.Background(), r.FormValue("code"))
		if err != nil {
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			return
		}
		user, err := s.storage.User().FindByOauthID("google", token.Extra("user_id").(string))
		if err != nil {
			if err != storage.ErrUserDoesNotExist {
				http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
				return
			}
			user = &model.User{
				VkId: token.Extra("user_id").(string),
			}
			err = s.storage.User().Create(user)
			if err != nil {
				http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
				return
			}
		}
		refreshTok := model.NewRefreshToken(user.Id)
		if err := s.storage.RefreshToken().Create(refreshTok); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		jwt, _ := GenerateJWT(user)
		cookieR := &http.Cookie{
			Name:     "refresh_token",
			Value:    refreshTok.Token,
			Path:     "/",
			HttpOnly: false,
			Secure:   false,
			Expires:  time.Now().Add(24 * time.Hour),
		}
		cookieJ := &http.Cookie{
			Name:     "jwt",
			Value:    jwt,
			Path:     "/",
			HttpOnly: false,
			Secure:   false,
			Expires:  time.Now().Add(24 * time.Hour),
		}
		http.SetCookie(w, cookieR)
		http.SetCookie(w, cookieJ)
		http.Redirect(w, r, "/set", http.StatusTemporaryRedirect)
		return

	}
}

func (s *server) handleGithubLogin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("GithubLoginHandler")
		u := config.GithubOauth.AuthCodeURL(createState(w))
		http.Redirect(w, r, u, http.StatusTemporaryRedirect)
	}
}

func (s *server) handelGithubCallBack() http.HandlerFunc {
	const oauthGithubAPI = "https://api.github.com/user"
	type User struct {
		ID int `json:"id"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("GithubCallBackHandler")
		oauthState, _ := r.Cookie("oauth_state")
		if r.FormValue("state") != oauthState.Value {
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			return
		}
		token, err := config.GithubOauth.Exchange(context.Background(), r.FormValue("code"))
		if err != nil {
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			return
		}
		req, err := http.NewRequest("GET", oauthGithubAPI, nil)
		if err != nil {
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			return
		}
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))
		response, err := (&http.Client{}).Do(req)
		if err != nil {
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			return
		}
		defer response.Body.Close()
		newUser := &User{}
		if err := json.NewDecoder(response.Body).Decode(&newUser); err != nil {
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			return
		}
		user, err := s.storage.User().FindByOauthID("github", string(newUser.ID))
		if err != nil {
			if err != storage.ErrUserDoesNotExist {
				http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
				return
			}
			user = &model.User{
				GithubId: string(newUser.ID),
			}
			err = s.storage.User().Create(user)
			if err != nil {
				http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
				return
			}
		}
		refreshTok := model.NewRefreshToken(user.Id)
		if err := s.storage.RefreshToken().Create(refreshTok); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		jwt, _ := GenerateJWT(user)
		cookieR := &http.Cookie{
			Name:     "refresh_token",
			Value:    refreshTok.Token,
			Path:     "/",
			HttpOnly: false,
			Secure:   false,
			Expires:  time.Now().Add(24 * time.Hour),
		}
		cookieJ := &http.Cookie{
			Name:     "jwt",
			Value:    jwt,
			Path:     "/",
			HttpOnly: false,
			Secure:   false,
			Expires:  time.Now().Add(24 * time.Hour),
		}
		http.SetCookie(w, cookieR)
		http.SetCookie(w, cookieJ)
		http.Redirect(w, r, "/set", http.StatusTemporaryRedirect)
		return
	}
}
