package config

import "time"

var (
	RefreshTokenLifeTime = time.Hour * 24 * 15
	Services             = []string{"yandex", "google", "vk", "github"}
)