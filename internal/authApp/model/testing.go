package model

import "testing"

func NewOauthToken(t *testing.T) *OauthToken {
	return &OauthToken{
		UserId:  "testUser",
		Service: "Yandex",
		Token:   "testToken",
	}
}

func NewUser(t *testing.T) *User {
	return &User{
		Id:                "test_id",
		Email:             "test@test.test",
		Password:          "password",
		EncryptedPassword: "encrrypted_password",
		Enabled2FA:        true,
		GoogleId:          "Google_id",
		VkId:              "vk_id",
		YandexId:          "yandex_id",
		GithubId:          "github_id",
	}
}
