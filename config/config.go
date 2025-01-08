package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DATABASE_URI string
	JWT_SECRET   string
}

var ConfigValues = initConfig()

func initConfig() *Config {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	return &Config{
		DATABASE_URI: os.Getenv("DATABASE_URI"),
		JWT_SECRET:   os.Getenv("JWT_SECRET"),
	}
}
