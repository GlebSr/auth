package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestOauthToken_Validate(t *testing.T) {
	tests := []struct {
		name    string
		token   func() *OauthToken
		isValid bool
	}{
		{
			name: "valid refresh",
			token: func() *OauthToken {
				tok := NewOauthToken(t)
				tok.IsRefresh = true
				return tok
			},
			isValid: true,
		},
		{
			name: "valid access token",
			token: func() *OauthToken {
				tok := NewOauthToken(t)
				tok.IsRefresh = false
				return tok
			},
			isValid: true,
		},
		{
			name: "invalid service refresh",
			token: func() *OauthToken {
				tok := NewOauthToken(t)
				tok.IsRefresh = true
				tok.Service = "invalid"
				return tok
			},
			isValid: false,
		},
		{
			name: "invalid service access token",
			token: func() *OauthToken {
				tok := NewOauthToken(t)
				tok.IsRefresh = false
				tok.Service = "invalid"
				return tok
			},
			isValid: false,
		},
		{
			name: "empty userId refresh",
			token: func() *OauthToken {
				tok := NewOauthToken(t)
				tok.IsRefresh = true
				tok.UserId = ""
				return tok
			},
			isValid: false,
		},
		{
			name: "empty userId access token",
			token: func() *OauthToken {
				tok := NewOauthToken(t)
				tok.IsRefresh = false
				tok.UserId = ""
				return tok
			},
			isValid: false,
		},
		{
			name: "empty token refresh",
			token: func() *OauthToken {
				tok := NewOauthToken(t)
				tok.IsRefresh = true
				tok.Token = ""
				return tok
			},
			isValid: false,
		},
		{
			name: "empty token access token",
			token: func() *OauthToken {
				tok := NewOauthToken(t)
				tok.IsRefresh = false
				tok.Token = ""
				return tok
			},
			isValid: false,
		},
		{
			name: "expired refresh",
			token: func() *OauthToken {
				tok := NewOauthToken(t)
				tok.IsRefresh = true
				tok.Expire = time.Now()
				return tok
			},
			isValid: false,
		},
		{
			name: "expired access token",
			token: func() *OauthToken {
				tok := NewOauthToken(t)
				tok.IsRefresh = false
				tok.Expire = time.Now()
				return tok
			},
			isValid: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.isValid {
				assert.NoError(t, test.token().Validate())
			} else {
				assert.Error(t, test.token().Validate())
			}
		})
	}
}

func TestValidService(t *testing.T) {
	assert.Equal(t, false, ValidService("Invalid"))
	assert.Equal(t, false, ValidService("Vk"))
	assert.Equal(t, true, ValidService("yandex"))
}
