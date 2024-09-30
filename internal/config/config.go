package config

import (
	"os"
)

type Config struct {
	Port string
}

func NewConfig() *Config {
	return &Config{
		Port: getEnv("PORT", "3030"),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
