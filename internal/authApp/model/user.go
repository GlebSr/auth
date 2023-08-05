package model

import (
	"errors"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id                int
	Email             string
	Password          string
	EncryptedPassword string
	Enabled2FA        bool

	GoogleId string
	VkId     string
	YandexId string
	GithubId string
}

func (u *User) Validate() error {
	errEmail := validation.ValidateStruct(u, validation.Field(&u.Email, validation.Required, is.Email))
	errPassword := validation.ValidateStruct(u, validation.Field(&u.Email, validation.Required, validation.Length(8, 72)))
	if u.GoogleId == "" && u.GithubId == "" && u.YandexId == "" && u.VkId == "" &&
		(errEmail != nil || (errPassword != nil && u.EncryptedPassword == "")) {
		return errors.New("validation error") //TODO new error
	}
	return nil
}

func (u *User) EncryptPassword() error {
	if len(u.Password) > 0 {
		enc, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if err != nil {
			return err //TODO new err
		}
		u.EncryptedPassword = string(enc)
	}
	return nil
}

func (u *User) Sanitize() {
	u.Password = ""
}
