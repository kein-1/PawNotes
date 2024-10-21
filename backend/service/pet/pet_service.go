package pet

import (
	"backend/types"
)

type PetServiceStruct struct {
	dbRepo types.PetRepoInterface
}

func NewPetService(dbRepo types.PetRepoInterface) types.PetServiceInterface {
	return &PetServiceStruct{dbRepo: dbRepo}
}

func (p *PetServiceStruct) AddPet(pet types.Pet, userID int) error {
	if err := p.dbRepo.CreatePet(pet, userID); err != nil {
		return err
	}
	return nil
}

func (p *PetServiceStruct) RemovePet(petID int, userID int) error {
	if err := p.dbRepo.DeletePet(petID, userID); err != nil {
		return err
	}
	return nil

}
func (p *PetServiceStruct) UpdatePetAttribute(petID int, userID int, petPatch types.PetPatch) error {

	field := "name"
	value := 3

	if err := p.dbRepo.UpdatePet(petID, userID, field, value); err != nil {
		return err
	}
	return nil
}
