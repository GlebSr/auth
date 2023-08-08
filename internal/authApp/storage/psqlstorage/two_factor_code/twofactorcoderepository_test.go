package psqlstorage

import (
	"auth/internal/authApp/config"
	"auth/internal/authApp/model"
	"auth/internal/authApp/storage/psqlstorage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestTwoFactorRepository_Create(t *testing.T) {
	db, teardown := psqlstorage.TestDB(t, config.TestDatabaseURL)
	rep := &TwoFactorRepository{
		Codes: db,
	}
	defer teardown("two_factor")
	code, _, _ := model.NewTwoFactorCode("user")
	assert.NoError(t, rep.Create(code))
	assert.Error(t, rep.Create(code))
}

func TestTwoFactorRepository_Delete(t *testing.T) {
	db, teardown := psqlstorage.TestDB(t, config.TestDatabaseURL)
	rep := &TwoFactorRepository{
		Codes: db,
	}
	defer teardown("two_factor")
	code, _, _ := model.NewTwoFactorCode("user")
	require.NoError(t, rep.Create(code))
	assert.NoError(t, rep.Delete("user"))
	assert.Error(t, rep.Delete("user"))
}

func TestTwoFactorRepository_FindByUserId(t *testing.T) {
	db, teardown := psqlstorage.TestDB(t, config.TestDatabaseURL)
	rep := &TwoFactorRepository{
		Codes: db,
	}
	defer teardown("two_factor")
	code, _, _ := model.NewTwoFactorCode("user")
	_, err := rep.FindByUserId("user")
	assert.Error(t, err)
	require.NoError(t, rep.Create(code))
	_, err = rep.FindByUserId("user")
	assert.NoError(t, err)
}
