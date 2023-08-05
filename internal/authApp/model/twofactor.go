package model

type TwoFactorCode struct {
	UserId       int
	SecretKey    string
	ReserveCodes []string
}
