package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

// This is method in go to act as the "constructor"
func NewStoragePostgres() (*pgxpool.Pool, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbURL := os.Getenv("DATABASE_URL")
	fmt.Println("db url:", dbURL)

	dbpool, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		log.Fatal("error in establishing a db connection", err)
	}

	return dbpool, nil
}
