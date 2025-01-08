// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"
)

// CreatePetInput represents a mutation input for creating pets.
type CreatePetInput struct {
	Name      string
	Age       int
	Breed     string
	Weight    float64
	CreatedAt *time.Time
	Dob       time.Time
	OwnerID   *int
}

// Mutate applies the CreatePetInput on the PetMutation builder.
func (i *CreatePetInput) Mutate(m *PetMutation) {
	m.SetName(i.Name)
	m.SetAge(i.Age)
	m.SetBreed(i.Breed)
	m.SetWeight(i.Weight)
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	m.SetDob(i.Dob)
	if v := i.OwnerID; v != nil {
		m.SetOwnerID(*v)
	}
}

// SetInput applies the change-set in the CreatePetInput on the PetCreate builder.
func (c *PetCreate) SetInput(i CreatePetInput) *PetCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdatePetInput represents a mutation input for updating pets.
type UpdatePetInput struct {
	Name       *string
	Age        *int
	Breed      *string
	Weight     *float64
	CreatedAt  *time.Time
	Dob        *time.Time
	ClearOwner bool
	OwnerID    *int
}

// Mutate applies the UpdatePetInput on the PetMutation builder.
func (i *UpdatePetInput) Mutate(m *PetMutation) {
	if v := i.Name; v != nil {
		m.SetName(*v)
	}
	if v := i.Age; v != nil {
		m.SetAge(*v)
	}
	if v := i.Breed; v != nil {
		m.SetBreed(*v)
	}
	if v := i.Weight; v != nil {
		m.SetWeight(*v)
	}
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.Dob; v != nil {
		m.SetDob(*v)
	}
	if i.ClearOwner {
		m.ClearOwner()
	}
	if v := i.OwnerID; v != nil {
		m.SetOwnerID(*v)
	}
}

// SetInput applies the change-set in the UpdatePetInput on the PetUpdate builder.
func (c *PetUpdate) SetInput(i UpdatePetInput) *PetUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdatePetInput on the PetUpdateOne builder.
func (c *PetUpdateOne) SetInput(i UpdatePetInput) *PetUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateUserInput represents a mutation input for creating users.
type CreateUserInput struct {
	First  string
	Last   string
	Email  string
	PetIDs []int
}

// Mutate applies the CreateUserInput on the UserMutation builder.
func (i *CreateUserInput) Mutate(m *UserMutation) {
	m.SetFirst(i.First)
	m.SetLast(i.Last)
	m.SetEmail(i.Email)
	if v := i.PetIDs; len(v) > 0 {
		m.AddPetIDs(v...)
	}
}

// SetInput applies the change-set in the CreateUserInput on the UserCreate builder.
func (c *UserCreate) SetInput(i CreateUserInput) *UserCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateUserInput represents a mutation input for updating users.
type UpdateUserInput struct {
	First        *string
	Last         *string
	Email        *string
	ClearPets    bool
	AddPetIDs    []int
	RemovePetIDs []int
}

// Mutate applies the UpdateUserInput on the UserMutation builder.
func (i *UpdateUserInput) Mutate(m *UserMutation) {
	if v := i.First; v != nil {
		m.SetFirst(*v)
	}
	if v := i.Last; v != nil {
		m.SetLast(*v)
	}
	if v := i.Email; v != nil {
		m.SetEmail(*v)
	}
	if i.ClearPets {
		m.ClearPets()
	}
	if v := i.AddPetIDs; len(v) > 0 {
		m.AddPetIDs(v...)
	}
	if v := i.RemovePetIDs; len(v) > 0 {
		m.RemovePetIDs(v...)
	}
}

// SetInput applies the change-set in the UpdateUserInput on the UserUpdate builder.
func (c *UserUpdate) SetInput(i UpdateUserInput) *UserUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateUserInput on the UserUpdateOne builder.
func (c *UserUpdateOne) SetInput(i UpdateUserInput) *UserUpdateOne {
	i.Mutate(c.Mutation())
	return c
}
