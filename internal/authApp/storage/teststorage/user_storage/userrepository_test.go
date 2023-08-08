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
	assert.NoError(t, rep.Create(user))
	assert.Error(t, rep.Create(user))
	user.Email = "test@ya.ru"
	assert.Error(t, rep.Create(user))
}

func TestUserRepository_Delete(t *testing.T) {
	rep := &UserRepository{
		Users: make(map[string]*model.User),
	}
	assert.Error(t, rep.Delete("user"))
	user := model.NewUser(t)
	rep.Users[user.Id] = user
	assert.NoError(t, rep.Delete(user.Id))
	assert.Error(t, rep.Delete(user.Id))
}

func TestUserRepository_Update(t *testing.T) {
	rep := &UserRepository{
		Users: make(map[string]*model.User),
	}
	user := model.NewUser(t)
	assert.Error(t, rep.Update(user))
	rep.Users[user.Id] = user
	assert.NoError(t, rep.Update(user))
	assert.NoError(t, rep.Update(user))
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
