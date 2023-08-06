package teststorage

import (
	"auth/internal/authApp/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTwoFactorRepository_Create(t *testing.T) {
	rep := &TwoFactorRepository{
		Codes: make(map[string]*model.TwoFactorCode),
	}
	code, _, _ := model.NewTwoFactorCode("user")
	err := rep.Create(code)
	assert.NoError(t, err)
	err = rep.Create(code)
	assert.Error(t, err)
}

func TestTwoFactorRepository_Delete(t *testing.T) {
	rep := &TwoFactorRepository{
		Codes: make(map[string]*model.TwoFactorCode),
	}
	code, _, _ := model.NewTwoFactorCode("user")
	rep.Codes["user"] = code
	err := rep.Delete("user")
	assert.NoError(t, err)
	err = rep.Delete("user")
	assert.Error(t, err)
}

func TestTwoFactorRepository_FindByUserId(t *testing.T) {
	rep := &TwoFactorRepository{
		Codes: make(map[string]*model.TwoFactorCode),
	}
	code, _, _ := model.NewTwoFactorCode("user")
	_, err := rep.FindByUserId("user")
	assert.Error(t, err)
	rep.Codes["user"] = code
	_, err = rep.FindByUserId("user")
	assert.NoError(t, err)
}
