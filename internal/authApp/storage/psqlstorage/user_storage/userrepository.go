package psqlstorage

import (
	"auth/internal/authApp/model"
	"auth/internal/authApp/storage"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"strings"
)

type UserRepository struct {
	Users *sqlx.DB
}

func (r *UserRepository) Create(user *model.User) error {
	if err := user.Validate(); err != nil {
		return err
	}
	if len(user.Id) == 0 {
		user.Id = uuid.New().String()
	}
	if len(user.Password) != 0 && len(user.EncryptedPassword) == 0 {
		if err := user.EncryptPassword(); err != nil {
			return err
		}
	}
	user.Sanitize()
	res, err := r.Users.Exec(`
			INSERT INTO users (id, email, encrypted_password, enabled2fa, google_id,  yandex_id,  vk_id,  github_id) 
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`,
		user.Id, user.Email, user.EncryptedPassword, user.Enabled2FA,
		user.GoogleId, user.YandexId, user.VkId, user.GithubId,
	)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("errr") //TODO new error
	}
	return nil
}

func (r *UserRepository) FindById(id string) (*model.User, error) {
	users := make([]*model.User, 0)
	err := r.Users.Select(
		&users, `
		SELECT id, email, encrypted_password, enabled2fa, google_id,  yandex_id,  vk_id,  github_id FROM users 
		WHERE id = $1`,
		id,
	)
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, storage.ErrUserDoesNotExist
	}
	return users[0], nil
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	users := make([]*model.User, 0)
	err := r.Users.Select(
		&users, `
		SELECT id, email, encrypted_password, enabled2fa, google_id,  yandex_id,  vk_id,  github_id FROM users 
		WHERE email = $1`,
		email,
	)
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, storage.ErrUserDoesNotExist
	}
	return users[0], nil
}

func (r *UserRepository) FindByOauthID(serviceName string, id string) (*model.User, error) {
	serviceName = strings.ToLower(serviceName)
	if !model.ValidService(serviceName) {
		return nil, storage.ErrServiceNotSupported
	}
	users := make([]*model.User, 0)
	query := fmt.Sprintf(`
		SELECT id, email, encrypted_password, enabled2fa,
		google_id,  yandex_id,  vk_id,  github_id FROM users  WHERE %s_id = $1`,
		serviceName)
	err := r.Users.Select(
		&users, query,
		id,
	)
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, storage.ErrUserDoesNotExist
	}
	return users[0], nil
}

func (r *UserRepository) Update(user *model.User) error {
	if err := user.Validate(); err != nil {
		return err
	}
	if len(user.Password) != 0 && len(user.EncryptedPassword) == 0 {
		if err := user.EncryptPassword(); err != nil {
			return err
		}
	}
	user.Sanitize()
	res, err := r.Users.Exec(`
			UPDATE users 
			SET email = $1, encrypted_password = $2, enabled2fa = $3,
		google_id = $4,  yandex_id = $5,  vk_id = $6,  github_id= $7
			WHERE id = $8`,
		user.Email, user.EncryptedPassword, user.Enabled2FA,
		user.GoogleId, user.YandexId, user.VkId, user.GithubId,
		user.Id,
	)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("errr") //TODO new error
	}
	return nil
}

func (r *UserRepository) Delete(id string) error {
	res, err := r.Users.Exec(`
			DELETE FROM users
			WHERE id = $1`,
		id,
	)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("errr") //TODO new error
	}
	return nil
}
