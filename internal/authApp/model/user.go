package model

type User struct {
	Id                int
	Email             string
	Password          string
	EncryptedPassword string
	Enabled2FA        bool

	GoogleId string
	VkId     string
	YandexId string
}
