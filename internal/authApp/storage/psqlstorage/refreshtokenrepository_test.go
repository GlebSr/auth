package psqlstorage

import (
	"auth/internal/authApp/config"
	"auth/internal/authApp/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestRefreshTokenRepository_FindByToken(t *testing.T) {
	db, teardown := TestDB(t, config.TestDatabaseURL)
	defer teardown("refresh")
	rep := RefreshTokenRepository{
		Tokens: db,
	}
	token := model.NewRefreshToken("user")
	require.NoError(t, rep.Create(token))
	t1, err := rep.FindByToken(token.Token)
	assert.NoError(t, err)
	assert.Equal(t, token.UserId, t1.UserId)
	assert.Equal(t, token.Token, t1.Token)
	assert.True(t, t1.Expire.Equal(token.Expire))
	t2, err := rep.FindByToken("invalid")
	assert.Error(t, err)
	assert.Nil(t, t2)
}

func TestRefreshTokenRepository_FindByUserId(t *testing.T) {
	db, teardown := TestDB(t, config.TestDatabaseURL)
	defer teardown("refresh")
	rep := RefreshTokenRepository{
		Tokens: db,
	}
	token := model.NewRefreshToken("user")
	require.NoError(t, rep.Create(token))
	t1, err := rep.FindByUserId("user")
	assert.NoError(t, err)
	assert.Equal(t, 1, len(t1))
	t2, err := rep.FindByUserId("invalid")
	assert.NoError(t, err)
	assert.Equal(t, 0, len(t2))
}

func TestRefreshTokenRepository_Create(t *testing.T) {
	db, teardown := TestDB(t, config.TestDatabaseURL)
	defer teardown("refresh")
	rep := RefreshTokenRepository{
		Tokens: db,
	}
	token := model.NewRefreshToken("user")
	assert.NoError(t, rep.Create(token))
	assert.Error(t, rep.Create(token))
}

func TestRefreshTokenRepository_Delete(t *testing.T) {
	db, teardown := TestDB(t, config.TestDatabaseURL)
	defer teardown("refresh")
	rep := RefreshTokenRepository{
		Tokens: db,
	}
	token := model.NewRefreshToken("user")
	require.NoError(t, rep.Create(token))
	assert.NoError(t, rep.Delete(token.Token))
	assert.Error(t, rep.Delete(token.Token))
}

func TestRefreshTokenRepository_DeleteAll(t *testing.T) {
	db, teardown := TestDB(t, config.TestDatabaseURL)
	defer teardown("refresh")
	rep := RefreshTokenRepository{
		Tokens: db,
	}
	token1 := model.NewRefreshToken("user")
	token2 := model.NewRefreshToken("user")
	require.NoError(t, rep.Create(token1))
	require.NoError(t, rep.Create(token2))
	tokens1, err := rep.FindByUserId("user")
	assert.Equal(t, 2, len(tokens1))
	assert.NoError(t, err)
	assert.NoError(t, rep.DeleteAllById("user"))
	tokens2, err := rep.FindByUserId("user")
	assert.Equal(t, 0, len(tokens2))
	assert.NoError(t, err)
	assert.NoError(t, rep.DeleteAllById("user"))
}

func TestRefreshTokenRepository_ClearExpired(t *testing.T) {
	db, teardown := TestDB(t, config.TestDatabaseURL)
	defer teardown("refresh")
	rep := RefreshTokenRepository{
		Tokens: db,
	}
	token1 := model.NewRefreshToken("user")
	token1.Expire = time.Now()
	require.NoError(t, rep.Create(token1))
	time.Sleep(time.Millisecond)
	assert.NoError(t, rep.ClearExpired())
	_, err := rep.FindByToken(token1.Token)
	assert.Error(t, err)
}
