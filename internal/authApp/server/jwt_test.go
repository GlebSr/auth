package server

import (
	"auth/internal/authApp/model"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestJWT(t *testing.T) {
	user := model.NewUser(t)
	token, err := GenerateJWT(user)
	log.Println(token)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
	valid, err := ValidateJWT(token)
	assert.NoError(t, err)
	assert.True(t, valid)
}
