package teststorage

import (
	"auth/internal/authApp/model"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestRefreshTokenRepository_FindByToken(t *testing.T) {
	rep := RefreshTokenRepository{
		Tokens: make(map[string]*model.RefreshToken),
	}
	token := model.NewRefreshToken("user")
	rep.Tokens[token.Token] = token
	t1, err := rep.FindByToken(token.Token)
	assert.NoError(t, err)
	assert.Equal(t, token, t1)
	t2, err := rep.FindByToken("invalid")
	assert.Error(t, err)
	assert.Nil(t, t2)
}

func TestRefreshTokenRepository_FindByUserId(t *testing.T) {
	rep := RefreshTokenRepository{
		Tokens: make(map[string]*model.RefreshToken),
	}
	token := model.NewRefreshToken("user")
	rep.Tokens[token.Token] = token
	t1, err := rep.FindByUserId("user")
	assert.NoError(t, err)
	assert.Equal(t, 1, len(t1))
	t2, err := rep.FindByUserId("invalid")
	assert.NoError(t, err)
	assert.Equal(t, 0, len(t2))
}

func TestRefreshTokenRepository_Create(t *testing.T) {
	rep := RefreshTokenRepository{
		Tokens: make(map[string]*model.RefreshToken),
	}
	token := model.NewRefreshToken("user")
	err := rep.Create(token)
	assert.NoError(t, err)
	err = rep.Create(token)
	assert.Error(t, err)
}

func TestRefreshTokenRepository_Delete(t *testing.T) {
	rep := RefreshTokenRepository{
		Tokens: make(map[string]*model.RefreshToken),
	}
	token := model.NewRefreshToken("user")
	rep.Tokens[token.Token] = token
	err := rep.Delete(token.Token)
	assert.NoError(t, err)
	err = rep.Delete(token.Token)
	assert.Error(t, err)
}

func TestRefreshTokenRepository_DeleteAll(t *testing.T) {
	rep := RefreshTokenRepository{
		Tokens: make(map[string]*model.RefreshToken),
	}
	token1 := model.NewRefreshToken("user")
	token2 := model.NewRefreshToken("user")
	rep.Tokens[token1.Token] = token1
	rep.Tokens[token2.Token] = token2
	tokens1, err := rep.FindByUserId("user")
	assert.Equal(t, 2, len(tokens1))
	assert.NoError(t, err)
	err = rep.DeleteAll("user")
	assert.NoError(t, err)
	tokens2, err := rep.FindByUserId("user")
	assert.Equal(t, 0, len(tokens2))
	assert.NoError(t, err)
	err = rep.DeleteAll("user")
	assert.NoError(t, err)
}

func TestRefreshTokenRepository_ClearExpired(t *testing.T) {
	rep := RefreshTokenRepository{
		Tokens: make(map[string]*model.RefreshToken),
	}
	token1 := model.NewRefreshToken("user")
	token1.Expiry = time.Now()
	rep.Tokens[token1.Token] = token1
	time.Sleep(time.Millisecond)
	err := rep.ClearExpired()
	assert.NoError(t, err)
	_, err = rep.FindByToken(token1.Token)
	assert.Error(t, err)
}
