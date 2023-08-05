package teststore

import (
	"auth/internal/authApp/model"
	"errors"
)

type UserRepository struct {
	Users map[int]*model.User
}

func (r *UserRepository) Create(user *model.User) error {
	if err := user.Validate(); err != nil {
		return err
	}
	for _, us := range r.Users {
		if us.Email == user.Email {
			return errors.New("user exist") //TODO new error
		}
	}
	user.Id = len(r.Users)
	_, ok := r.Users[user.Id]
	for ok {
		user.Id += 1
		_, ok = r.Users[user.Id]
	}

	if err := user.EncryptPassword(); err != nil {
		return err
	}
	user.Sanitize()
	r.Users[user.Id] = user
	return nil
}

func (r *UserRepository) FindById(id int) (*model.User, error) {
	user, ok := r.Users[id]
	if ok {
		return user, nil
	}
	return nil, errors.New("user doesnt exist") //TODO new error
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	for _, us := range r.Users {
		if us.Email == email {
			return us, nil
		}
	}
	return nil, errors.New("user doesnt exist") //TODO new error
}

func (r *UserRepository) FindByOauthID(serviceName string, id string) (*model.User, error) {
	switch serviceName {
	case "Google":
		for _, us := range r.Users {
			if us.GoogleId == id {
				return us, nil
			}
		}
	case "Yandex":
		for _, us := range r.Users {
			if us.YandexId == id {
				return us, nil
			}
		}
	case "Github":
		for _, us := range r.Users {
			if us.GithubId == id {
				return us, nil
			}
		}
	case "Vk":
		for _, us := range r.Users {
			if us.VkId == id {
				return us, nil
			}
		}
	}
	return nil, errors.New("service not supported") //TODO new error
}

func (r *UserRepository) Update(user *model.User) error {
	if err := user.Validate(); err != nil {
		return err
	}
	_, ok := r.Users[user.Id]
	if !ok {
		return errors.New("user doesnt exist") //TODO new error
	}
	if err := user.EncryptPassword(); err != nil {
		return err
	}
	user.Sanitize()
	r.Users[user.Id] = user
	return nil
}

func (r *UserRepository) Delete(id int) error {
	_, ok := r.Users[id]
	if !ok {
		return errors.New("user doesnt exist") //TODO new error
	}
	delete(r.Users, id)
	return nil
}
