package model

import (
	"testing"
	"time"
)

func init() {
	time.Local = time.FixedZone("", 0)
}

func NewOauthToken(t *testing.T) *OauthToken {
	return &OauthToken{
		UserId:  "test_user",
		Service: "yandex",
		Token:   "test_token",
		Expire:  time.Now().Add(time.Hour),
	}
}

func NewUser(t *testing.T) *User {
	return &User{
		Id:                "test_id",
		Email:             "test@test.test",
		Password:          "password",
		EncryptedPassword: "encrypted_password",
		Enabled2FA:        true,
		GoogleId:          "google_id",
		VkId:              "vk_id",
		YandexId:          "yandex_id",
		GithubId:          "github_id",
	}
}
