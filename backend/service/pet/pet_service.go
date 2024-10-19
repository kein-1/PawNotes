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
