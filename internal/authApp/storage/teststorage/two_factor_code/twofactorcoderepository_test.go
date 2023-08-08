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
	assert.NoError(t, rep.Create(code))
	assert.Error(t, rep.Create(code))
}

func TestTwoFactorRepository_Delete(t *testing.T) {
	rep := &TwoFactorRepository{
		Codes: make(map[string]*model.TwoFactorCode),
	}
	code, _, _ := model.NewTwoFactorCode("user")
	rep.Codes["user"] = code
	assert.NoError(t, rep.Delete("user"))
	assert.Error(t, rep.Delete("user"))
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
