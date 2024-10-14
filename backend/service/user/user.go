package user

import (
	"backend/types"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Store struct {
	db *pgxpool.Pool
}

// constructor
func NewUserStore(db *pgxpool.Pool) *Store {
	return &Store{db}
}

func (user_store *Store) CreateUser(user types.User) {
	fmt.Println("Creating user:", user)
}
