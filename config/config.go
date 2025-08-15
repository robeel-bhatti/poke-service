// Package config loads all the app-wide config for the app to run
package config

import (
	"os"
)

type Config struct {
	Port     string
	Env      string
	LogLevel string
}

// LoadConfig loads any required env vars
func LoadConfig() *Config {
	return &Config{
		Port:     os.Getenv("PORT"),
		Env:      os.Getenv("ENV"),
		LogLevel: os.Getenv("LOG_LEVEL"),
	}
}
