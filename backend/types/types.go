package types

import (
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

type UserStore interface {
	CreateUser(user User)
}

type PetStore interface {
}

type NoteStore interface {
	CreateNote(note Note) error
}
