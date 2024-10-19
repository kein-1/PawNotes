package user

import (
	"backend/types"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserDBRepo struct {
	db *pgxpool.Pool
}

// constructor
func NewUserRepo(db *pgxpool.Pool) types.UserDBRepoInterface {
	return &UserDBRepo{db: db}
}

func (r *UserDBRepo) CreateUser(user types.User) error {

	query := "INSERT INTO users (first_name, last_name, email, password, created_at) VALUES ($1,$2,$3,$4,$5)"
	_, err := r.db.Exec(context.Background(), query, user.FirstName, user.LastName, user.Email, user.Password, user.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserDBRepo) GetUserByEmail(email string) (*types.User, error) {

	var user types.User

	query := "SELECT * FROM users WHERE users.email = $1"
	err := r.db.QueryRow(context.Background(), query, email).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			return &user, types.ErrNoRecord
		}
		return &user, err
	}
	fmt.Println("User found:", user)
	return &user, nil
}

func (r *UserDBRepo) CheckEmail(user types.User) error {
	email := user.Email
	var dbEmail string
	err := r.db.QueryRow(context.Background(),
		"SELECT email FROM users WHERE email = $1", email).Scan(&dbEmail)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil
		}
		return fmt.Errorf("Error querying the databse for email %w", err)
	}
	// user exists since a record was found
	return types.ErrUserExists
}

func (r *UserDBRepo) DeleteUser(id int) error {
	query := "DELETE FROM users WHERE users.id = $1"
	commandTag, err := r.db.Exec(context.Background(), query, id)
	if err != nil {
		return fmt.Errorf("Error deleting your account %w", err)
	}
	if commandTag.RowsAffected() == 0 {
		return types.ErrNoRecord
	}
	return nil

}
