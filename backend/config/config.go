package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	POSTGRES_USER     string
	POSTGRES_PASSWORD string
	POSTGRES_DB       string

	SECRET_KEY   string
	DATABASE_URL string
}

var ConfigValues = initConfig()

func initConfig() *Config {
	godotenv.Load()
	return &Config{
		POSTGRES_USER:     os.Getenv("POSTGRES_USER"),
		POSTGRES_PASSWORD: os.Getenv("POSTGRES_PASSWORD"),
		POSTGRES_DB:       os.Getenv("POSTGRES_DB"),
		SECRET_KEY:        os.Getenv("SECRET_KEY"),
		DATABASE_URL:      os.Getenv("DATABASE_URL"),
	}
}
