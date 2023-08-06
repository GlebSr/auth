package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewTwoFactorCode(t *testing.T) {
	code, img, err := NewTwoFactorCode("test_id")
	assert.NotNil(t, code)
	assert.NotNil(t, img)
	assert.NoError(t, err)
}
