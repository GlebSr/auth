package server

import (
	"auth/internal/authApp/config"
	"auth/internal/authApp/model"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type Claims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

func GenerateJWT(user *model.User) (string, error) {
	//fmt.Println(config.PrivateRSAKey)
	claims := &Claims{
		UserID: user.Id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(5 * time.Minute)),
		},
	}
	key, err := jwt.ParseRSAPrivateKeyFromPEM(config.PrivateRSAKey)
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS512, claims)
	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", err
	}
	return tokenString, nil

}

func ValidateJWT(jtoken string) (bool, error) {
	claims := &Claims{}
	key, err := jwt.ParseRSAPublicKeyFromPEM(config.PublicRSAKey)
	if err != nil {
		return false, err
	}
	tkn, err := jwt.ParseWithClaims(jtoken, claims, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		return false, err
	}
	if !tkn.Valid {
		return false, nil
	}
	return true, nil
}
