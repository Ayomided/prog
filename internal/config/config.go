package config

import (
	"os"
)

type Config struct {
	Port         string
	StaticPath   string
	StaticPathOG string
}

func NewConfig() *Config {
	return &Config{
		Port:         getEnv("PORT", "8080"),
		StaticPath:   getEnv("STATIC", "static"),
		StaticPathOG: getEnv("STATICOG", "static/og-images"),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
