package config

import (
	"os"

	"github.com/pivovarit/hexagonal-go/app/config/profile"
)

var (
	appProfile = profile.Profile(getEnvOr("profile", profile.DEFAULT.AsString()))
)

type AppConfiguration struct {
}

func (c *AppConfiguration) GetProfile() profile.Profile {
	return appProfile
}

func getEnvOr(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}
