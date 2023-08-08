package model

import (
	"github.com/google/uuid"
	"github.com/pquerna/otp/totp"
	"image"
)

type TwoFactorCode struct {
	UserId       string   `db:"user_id"`
	SecretKey    string   `db:"secret_key"`
	ReserveCodes []string `db:"reserve_codes"`
}

func NewTwoFactorCode(id string) (*TwoFactorCode, image.Image, error) {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "nuhaibeb.ru",
		AccountName: id,
	})
	if err != nil {
		return nil, nil, err
	}
	qr, err := key.Image(256, 256)
	if err != nil {
		return nil, nil, err
	}
	codes := make([]string, 10)
	for i := 0; i < 10; i++ {
		codes[i] = uuid.New().String()
	}
	return &TwoFactorCode{
		UserId:       id,
		SecretKey:    key.Secret(),
		ReserveCodes: codes,
	}, qr, nil
}
