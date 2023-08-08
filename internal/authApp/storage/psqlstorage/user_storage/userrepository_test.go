package psqlstorage

import (
	"auth/internal/authApp/config"
	"auth/internal/authApp/model"
	"auth/internal/authApp/storage/psqlstorage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	db, teardown := psqlstorage.TestDB(t, config.TestDatabaseURL)
	defer teardown("users")
	rep := &UserRepository{
		Users: db,
	}
	user := model.NewUser(t)
	assert.NoError(t, rep.Create(user))
	assert.Error(t, rep.Create(user))
	user.Email = "test@ya.ru"
	assert.Error(t, rep.Create(user))
}

func TestUserRepository_Delete(t *testing.T) {
	db, teardown := psqlstorage.TestDB(t, config.TestDatabaseURL)
	defer teardown("users")
	rep := &UserRepository{
		Users: db,
	}
	assert.Error(t, rep.Delete("user"))
	user := model.NewUser(t)
	require.NoError(t, rep.Create(user))
	assert.NoError(t, rep.Delete(user.Id))
	assert.Error(t, rep.Delete(user.Id))
}

func TestUserRepository_Update(t *testing.T) {
	db, teardown := psqlstorage.TestDB(t, config.TestDatabaseURL)
	defer teardown("users")
	rep := &UserRepository{
		Users: db,
	}
	user := model.NewUser(t)
	assert.Error(t, rep.Update(user))
	require.NoError(t, rep.Create(user))
	assert.NoError(t, rep.Update(user))
	assert.NoError(t, rep.Update(user))
}

func TestUserRepository_FindByEmail(t *testing.T) {
	db, teardown := psqlstorage.TestDB(t, config.TestDatabaseURL)
	defer teardown("users")
	rep := &UserRepository{
		Users: db,
	}
	user := model.NewUser(t)
	_, err := rep.FindByEmail(user.Email)
	assert.Error(t, err)
	require.NoError(t, rep.Create(user))
	user_, err := rep.FindByEmail(user.Email)
	assert.NoError(t, err)
	assert.Equal(t, user, user_)
	_, err = rep.FindByEmail("other@other.other")
	assert.Error(t, err)
}

func TestUserRepository_FindById(t *testing.T) {
	db, teardown := psqlstorage.TestDB(t, config.TestDatabaseURL)
	defer teardown("users")
	rep := &UserRepository{
		Users: db,
	}
	user := model.NewUser(t)
	_, err := rep.FindById(user.Id)
	assert.Error(t, err)
	require.NoError(t, rep.Create(user))
	user_, err := rep.FindById(user.Id)
	assert.NoError(t, err)
	assert.Equal(t, user, user_)
	_, err = rep.FindById("invalid")
	assert.Error(t, err)
}

func TestUserRepository_FindByOauthID(t *testing.T) {
	db, teardown := psqlstorage.TestDB(t, config.TestDatabaseURL)
	defer teardown("users")
	rep := &UserRepository{
		Users: db,
	}
	user := model.NewUser(t)
	_, err := rep.FindByOauthID("yandex", "invalid")
	assert.Error(t, err)
	user.YandexId = "user_id"
	require.NoError(t, rep.Create(user))
	user_, err := rep.FindByOauthID("yandex", "user_id")
	assert.NoError(t, err)
	assert.Equal(t, user, user_)
	_, err = rep.FindByOauthID("invalid", "user_id")
	assert.Error(t, err)
	_, err = rep.FindByOauthID("yandex", "invalid")
	assert.Error(t, err)
}
