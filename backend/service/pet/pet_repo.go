package pet

import (
	"backend/types"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PetRepo struct {
	db *pgxpool.Pool
}

func NewPetRepo(db *pgxpool.Pool) types.PetRepoInterface {
	return &PetRepo{db: db}
}

func (p *PetRepo) CreatePet(pet types.Pet, userID int) error {

	fmt.Println("pet time:", pet.DOB.Time)
	fmt.Println("pet time:", pet.DOB)

	query := "INSERT INTO pets (name, breed, weight, dob, owner_id) VALUES ($1,$2,$3,$4, $5)"
	_, err := p.db.Exec(context.Background(), query, pet.Name, pet.Breed, pet.Weight, pet.DOB.Time, userID)
	if err != nil {
		return err
	}
	return nil
}

func (p *PetRepo) DeletePet(petID int, userID int) error {
	query := "DELETE FROM pets WHERE petID = $1 and owner_id = $2"
	commandTag, err := p.db.Exec(context.Background(), query, petID, userID)
	if commandTag.RowsAffected() == 0 {
		return types.ErrNoRecord
	}
	if err != nil {
		return fmt.Errorf("Error deleting the pet with id %d. Error: %w", petID, err)
	}
	return nil
}

func (p *PetRepo) UpdatePet(petID int, userID int, column string, value any) error {

	return nil
}
