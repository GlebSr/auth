package psqlstorage

import (
	"auth/internal/authApp/config"
	"auth/internal/authApp/model"
	"auth/internal/authApp/storage/psqlstorage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestOauthTokenRepository_Delete(t *testing.T) {
	db, teardown := psqlstorage.TestDB(t, config.TestDatabaseURL)
	defer teardown("oauth")
	rep := OauthTokenRepository{
		Tokens: db,
	}
	tests := []struct {
		name    string
		token   func() *model.OauthToken
		isValid bool
	}{
		{
			name: "valid access token",
			token: func() *model.OauthToken {
				tok := model.NewOauthToken(t)
				return tok
			},
			isValid: true,
		},
		{
			name: "invalid refresh",
			token: func() *model.OauthToken {
				tok := model.NewOauthToken(t)
				tok.IsRefresh = true
				return tok
			},
			isValid: false,
		},
		{
			name: "other id",
			token: func() *model.OauthToken {
				tok := model.NewOauthToken(t)
				tok.UserId = "invalid"
				return tok
			},
			isValid: false,
		},
		{
			name: "other service",
			token: func() *model.OauthToken {
				tok := model.NewOauthToken(t)
				tok.Service = "invalid"
				return tok
			},
			isValid: false,
		},
		{
			name: "other service",
			token: func() *model.OauthToken {
				tok := model.NewOauthToken(t)
				tok.Service = "invalid"
				return tok
			},
			isValid: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tok := model.NewOauthToken(t)
			require.NoError(t, rep.Create(tok))
			if test.isValid {
				assert.NoError(t, rep.Delete(test.token().UserId, test.token().Service, test.token().IsRefresh))
			} else {
				assert.Error(t, rep.Delete(test.token().UserId, test.token().Service, test.token().IsRefresh))
				assert.NoError(t, rep.Delete(tok.UserId, tok.Service, tok.IsRefresh))
			}
		})
	}
}

func TestOauthTokenRepository_Delete2(t *testing.T) {
	db, teardown := psqlstorage.TestDB(t, config.TestDatabaseURL)
	defer teardown("oauth")
	rep := OauthTokenRepository{
		Tokens: db,
	}
	user := model.NewOauthToken(t)
	require.NoError(t, rep.Create(user))
	assert.NoError(t, rep.Delete(user.UserId, user.Service, user.IsRefresh))
	assert.Error(t, rep.Delete(user.UserId, user.Service, user.IsRefresh))
}

func TestOauthTokenRepository_Create(t *testing.T) {
	db, teardown := psqlstorage.TestDB(t, config.TestDatabaseURL)
	defer teardown("oauth")
	rep := OauthTokenRepository{
		Tokens: db,
	}
	assert.NoError(t, rep.Create(model.NewOauthToken(t)))
	assert.Error(t, rep.Create(model.NewOauthToken(t)))
	refresh := model.NewOauthToken(t)
	refresh.IsRefresh = true
	assert.NoError(t, rep.Create(refresh))
	assert.Error(t, rep.Create(refresh))
}

func TestOauthTokenRepository_Update(t *testing.T) {
	db, teardown := psqlstorage.TestDB(t, config.TestDatabaseURL)
	defer teardown("oauth")
	rep := OauthTokenRepository{
		Tokens: db,
	}
	require.NoError(t, rep.Create(model.NewOauthToken(t)))
	assert.NoError(t, rep.Update(model.NewOauthToken(t)))
	refresh := model.NewOauthToken(t)
	refresh.IsRefresh = true
	require.Error(t, rep.Update(refresh))
	other := model.NewOauthToken(t)
	other.UserId = "invalid"
	assert.Error(t, rep.Update(other))
}

func TestOauthTokenRepository_FindByUserId(t *testing.T) {
	db, teardown := psqlstorage.TestDB(t, config.TestDatabaseURL)
	defer teardown("oauth")
	rep := OauthTokenRepository{
		Tokens: db,
	}
	require.NoError(t, rep.Create(model.NewOauthToken(t)))
	toks1, err := rep.FindByUserId(model.NewOauthToken(t).UserId)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(toks1))
	refresh := model.NewOauthToken(t)
	refresh.IsRefresh = true
	other := model.NewOauthToken(t)
	other.Service = "github"
	require.NoError(t, rep.Create(refresh))
	require.NoError(t, rep.Create(other))
	toks2, err := rep.FindByUserId(refresh.UserId)
	assert.NoError(t, rep.Update(refresh))
	assert.Equal(t, 3, len(toks2))
	toks3, err := rep.FindByUserId("invalid")
	assert.Equal(t, 0, len(toks3))
}

func TestOauthTokenRepository_FindByUserIdAndService(t *testing.T) {
	db, teardown := psqlstorage.TestDB(t, config.TestDatabaseURL)
	defer teardown("oauth")
	rep := OauthTokenRepository{
		Tokens: db,
	}
	require.NoError(t, rep.Create(model.NewOauthToken(t)))
	toks1, err := rep.FindByUserIdAndService(model.NewOauthToken(t).UserId, model.NewOauthToken(t).Service)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(toks1))
	refresh := model.NewOauthToken(t)
	refresh.IsRefresh = true
	other := model.NewOauthToken(t)
	other.Service = "github"
	require.NoError(t, rep.Create(refresh))
	require.NoError(t, rep.Create(other))
	toks2, err := rep.FindByUserIdAndService(refresh.UserId, model.NewOauthToken(t).Service)
	assert.NoError(t, rep.Update(refresh))
	assert.Equal(t, 2, len(toks2))
	toks3, err := rep.FindByUserIdAndService("invalid", "yandex")
	assert.Equal(t, 0, len(toks3))
	toks4, err := rep.FindByUserIdAndService("invalid", "invalid")
	assert.Error(t, err)
	assert.Equal(t, 0, len(toks4))
}
