package types

import (
	"errors"
	"strings"
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

type UserPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=3,max=15"`
}

type Pet struct {
	Id        int        `json:"id"`
	Name      string     `json:"name" validate:"required,alpha"`
	Breed     string     `json:"breed" validate:"required,alpha"`
	Weight    float64    `json:"weight" validate:"required,number"`
	DOB       CustomTime `json:"dob"`
	OwnerID   int        `json:"ownerID"`
	CreatedAt time.Time  `json:"-"`
}

type PetPatch struct {
	Name   *string     `json:"name"`
	Breed  *string     `json:"breed"`
	Weight *float64    `json:"weight"`
	DOB    *CustomTime `json:"dob"`
}

type Note struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	OwnerID   int       `json:"ownerID"`
	PetID     int       `json:"petID"`
	CreatedAt time.Time `json:"createdAt"`
}

// =========

// used for custom json decoding of a time format since Go requires
// the input to be a certain format. Using below, we can pass in
// a time format in "YYYY-MM-DD"
type CustomTime struct {
	time.Time
}

func (t *CustomTime) UnmarshalJSON(b []byte) (err error) {

	dateStr := strings.Trim(string(b), `"`)
	date, err := time.Parse(time.DateOnly, string(dateStr))
	if err != nil {
		return err
	}
	t.Time = date
	return nil
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
	AddPet(pet Pet, userID int) error
	RemovePet(petID int, userID int) error
	UpdatePetAttribute(petID int, userID int, petPatch PetPatch) error
}

type PetRepoInterface interface {
	CreatePet(pet Pet, userID int) error
	DeletePet(petID int, userID int) error
	UpdatePet(petID int, userID int, field string, value any) error
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
	ErrUnauthorized           = errors.New("Unauthorized to make this request")
	ErrInsertionError         = errors.New("Error adding this into the database")
)
