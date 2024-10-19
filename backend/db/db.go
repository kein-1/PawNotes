package db

import (
	"backend/config"
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

// This is method in go to act as the "constructor"
func NewStoragePostgres() (*pgxpool.Pool, error) {

	dbURL := config.ConfigValues.DATABASE_URL

	dbpool, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		log.Fatal("error in establishing a db connection", err)
	}

	return dbpool, nil
}
