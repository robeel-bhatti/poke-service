// Package config loads all the app-wide config for the app to run
package config

import (
	"os"
)

type Config struct {
	port     string
	env      string
	logLevel string
}

// LoadConfig loads any required env vars
func LoadConfig() *Config {
	return &Config{
		port:     os.Getenv("PORT"),
		env:      os.Getenv("ENV"),
		logLevel: os.Getenv("LOG_LEVEL"),
	}
}
