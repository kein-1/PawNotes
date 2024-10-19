package pet

import (
	"backend/types"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PetRepo struct {
	db *pgxpool.Pool
}

func NewPetRepo(db *pgxpool.Pool) types.PetRepoInterface {
	return &PetRepo{db: db}
}
