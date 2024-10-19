package types

import (
	"errors"
	"time"
)

type User struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firstName" validate:"required"`
	LastName  string    `json:"lastName"  validate:"required"`
	Email     string    `json:"email" validate:"required,email"`
	Password  string    `json:"password" validate:"required,min=3,max=15"`
	CreatedAt time.Time `json:"createdAt"`
}

type Pet struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Breed     string    `json:"breed"`
	Weight    float64   `json:"weight"`
	OwnerID   int       `json:"ownerID"`
	CreatedAt time.Time `json:"createdAt"`
}

type Note struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	OwnerID   int       `json:"ownerID"`
	PetID     int       `json:"petID"`
	CreatedAt time.Time `json:"createdAt"`
}

type UserPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=3,max=15"`
}

// =========

type UserServiceInterface interface {
	LoginUser(user UserPayload) (string, error)
	RegisterUser(user User) error
	RemoveUser(id int) error
}

type UserDBRepoInterface interface {
	CreateUser(user User) error
	GetUserByEmail(email string) (*User, error)
	CheckEmail(user User) error
	DeleteUser(id int) error
}

// =========

type PetServiceInterface interface {
}

type PetRepoInterface interface {
}

type NoteStore interface {
	CreateNote(note Note) error
}

// =========

var (
	ErrNoRecord               = errors.New("No record found")
	ErrInvalidUserCredentials = errors.New("Incorrect username or password")
	ErrUserExists             = errors.New("A user with that email exists. Please use a different email")
	ErrUserDoesNotExist       = errors.New("Cannot delete account. A user with that email does not exist")
)
