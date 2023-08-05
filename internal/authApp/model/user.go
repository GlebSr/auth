package model

type User struct {
	Id                int
	Email             string
	Password          string
	EncryptedPassword string

	GoogleId string
	VkId     string
	YandexId string
}
