package config

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/vk"
	"golang.org/x/oauth2/yandex"
	"time"
)

var (
	RefreshTokenLifeTime = time.Hour * 24 * 15
	Services             = []string{"yandex", "google", "vk", "github"}
	TestDatabaseURL      = "host=localhost user=*user* password=*password* dbname=*auth_test* sslmode=disable"
	PrivateRSAKey        = ([]byte)(`-----BEGIN PRIVATE KEY-----
........................
-----END PUBLIC KEY-----`)
	DatabaseURL = "host=localhost user=*user* password=*password* dbname=*auth_dev* sslmode=disable"
	BindAddr    = ":8080"
	GoogleOauth = &oauth2.Config{
		RedirectURL:  "https://nuhaibeb.ru/auth/v1/oauth/google/callback",
		ClientID:     "...",
		ClientSecret: "...",
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
	YandexOauth = &oauth2.Config{
		RedirectURL:  "https://nuhaibeb.ru/auth/v1/oauth/yandex/callback",
		ClientID:     "...",
		ClientSecret: "...",
		Scopes:       []string{""},
		Endpoint:     yandex.Endpoint,
	}
	VkOauth = &oauth2.Config{
		RedirectURL:  "https://nuhaibeb.ru/auth/v1/oauth/vk/callback",
		ClientID:     "...",
		ClientSecret: "...",
		Scopes:       []string{"5309444"},
		Endpoint:     vk.Endpoint,
	}
	GithubOauth = &oauth2.Config{
		RedirectURL:  "https://nuhaibeb.ru/auth/v1/oauth/github/callback",
		ClientID:     "...",
		ClientSecret: "...",
		Scopes:       []string{""},
		Endpoint:     github.Endpoint,
	}
)
