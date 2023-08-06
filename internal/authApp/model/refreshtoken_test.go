package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewRefreshToken(t *testing.T) {
	token := NewRefreshToken("test_id")
	assert.NotNil(t, token.Token)
	assert.NotNil(t, token.UserId)
	assert.True(t, time.Now().Before(token.Expiry))
}
