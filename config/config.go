package config

import (
	"os"
)

type Config struct {
	Port     string
	Env      string
	LogLevel string
}

func LoadConfig() *Config {
	return &Config{
		Port:     os.Getenv("PORT"),
		Env:      os.Getenv("ENV"),
		LogLevel: os.Getenv("LOG_LEVEL"),
	}
}
