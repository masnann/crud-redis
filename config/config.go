package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	RedisAddress  string
	RedisPassword string
	DBHost        string
	DBPort        string
	DBPass        string
	DBUser        string
	DBName        string
	SSLMode       string
}

func LoadConfig() (Config, error) {
	var cfg Config

	err := godotenv.Load()
	if err != nil {
		return cfg, fmt.Errorf("error loading .env file: %v", err)
	}

	cfg.RedisAddress = os.Getenv("REDIS_ADDR")
	cfg.RedisPassword = os.Getenv("REDIS_PASSWORD")
	cfg.DBHost = os.Getenv("DB_HOST")
	cfg.DBPort = os.Getenv("DB_PORT")
	cfg.DBUser = os.Getenv("DB_USER")
	cfg.DBPass = os.Getenv("DB_PASSWORD")
	cfg.DBName = os.Getenv("DB_NAME")
	cfg.SSLMode = os.Getenv("DB_SSLMODE")

	return cfg, nil
}
