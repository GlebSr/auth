package teststorage

import (
	"auth/internal/authApp/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	rep := &UserRepository{
		Users: make(map[string]*model.User),
	}
	user := model.NewUser(t)
	err := rep.Create(user)
	assert.NoError(t, err)
	err = rep.Create(user)
	assert.Error(t, err)
	user.Email = "test@ya.ru"
	err = rep.Create(user)
	assert.Error(t, err)
}

func TestUserRepository_Delete(t *testing.T) {
	rep := &UserRepository{
		Users: make(map[string]*model.User),
	}
	err := rep.Delete("user")
	assert.Error(t, err)
	user := model.NewUser(t)
	rep.Users[user.Id] = user
	err = rep.Delete(user.Id)
	assert.NoError(t, err)
	err = rep.Delete(user.Id)
	assert.Error(t, err)
}

func TestUserRepository_Update(t *testing.T) {
	rep := &UserRepository{
		Users: make(map[string]*model.User),
	}
	user := model.NewUser(t)
	err := rep.Update(user)
	assert.Error(t, err)
	rep.Users[user.Id] = user
	err = rep.Update(user)
	assert.NoError(t, err)
	err = rep.Update(user)
	assert.Error(t, err)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	rep := &UserRepository{
		Users: make(map[string]*model.User),
	}
	user := model.NewUser(t)
	_, err := rep.FindByEmail(user.Email)
	assert.Error(t, err)
	rep.Users[user.Id] = user
	user_, err := rep.FindByEmail(user.Email)
	assert.NoError(t, err)
	assert.Equal(t, user, user_)
	prevEmail := user.Email
	rep.Users[user.Id].Email = "other@c.c"
	_, err = rep.FindByEmail(prevEmail)
	assert.Error(t, err)
}

func TestUserRepository_FindById(t *testing.T) {
	rep := &UserRepository{
		Users: make(map[string]*model.User),
	}
	user := model.NewUser(t)
	_, err := rep.FindById(user.Id)
	assert.Error(t, err)
	rep.Users[user.Id] = user
	user_, err := rep.FindById(user.Id)
	assert.NoError(t, err)
	assert.Equal(t, user, user_)
	_, err = rep.FindById("invalid")
	assert.Error(t, err)
}

func TestUserRepository_FindByOauthID(t *testing.T) {
	rep := &UserRepository{
		Users: make(map[string]*model.User),
	}
	user := model.NewUser(t)
	_, err := rep.FindByOauthID("yandex", "invalid")
	assert.Error(t, err)
	user.YandexId = "user_id"
	rep.Users[user.Id] = user
	user_, err := rep.FindByOauthID("yandex", "user_id")
	assert.NoError(t, err)
	assert.Equal(t, user, user_)
	_, err = rep.FindByOauthID("invalid", "user_id")
	assert.Error(t, err)
	_, err = rep.FindByOauthID("yandex", "invalid")
	assert.Error(t, err)
}
