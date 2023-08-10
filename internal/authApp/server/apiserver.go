package server

import (
	"auth/internal/authApp/config"
	"auth/internal/authApp/storage"
	"auth/internal/authApp/storage/psqlstorage"
	"github.com/jmoiron/sqlx"
	"github.com/rs/cors"
	"log"
	"net/http"
)

func Start() error {
	log.Println("Server started")
	defer log.Println("Server finished")
	db, err := newDB(config.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	store := &storage.Storage{
		UserStorage:         psqlstorage.NewUserStorage(db),
		RefreshTokenStorage: psqlstorage.NewRefreshTokenStorage(db),
		OauthTokenStorage:   psqlstorage.NewOauthTokenStorage(db),
		TwoFactorStorage:    psqlstorage.NewTwoFactorStorage(db),
	}

	srv := NewServer(store)
	return http.ListenAndServe(config.BindAddr, cors.Default().Handler(srv))
}

func newDB(databaseURL string) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
