package pet

import (
	"backend/config"
	"backend/types"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
)

type PetHandler struct {
	petService types.PetServiceInterface
}

func NewPetHandler(p types.PetServiceInterface) *PetHandler {
	return &PetHandler{petService: p}
}

func (h *PetHandler) RegisterRoutes(r *chi.Mux) {

	tokenAuth := jwtauth.New("HS256", []byte(config.ConfigValues.SECRET_KEY), nil)

	r.Route("/api/pet", func(r chi.Router) {
		r.Group(func(r chi.Router) {

			r.Use(jwtauth.Verifier(tokenAuth))
			r.Use(jwtauth.Authenticator(tokenAuth))
			fmt.Println("No issues iwth token in pet route!")
			r.Post("/createPet", h.handleCreatePet)
			r.Delete("/deletePet", h.handleDeletePet)
			r.Put("/updatePet", h.handleUpdatingPet)

		})
	})

}

func (h *PetHandler) handleCreatePet(w http.ResponseWriter, r *http.Request) {
	fmt.Println("This route handles creating a pet!")
}

func (h *PetHandler) handleDeletePet(w http.ResponseWriter, r *http.Request) {
	fmt.Println("This route handles deleting a pet!")
}

func (h *PetHandler) handleUpdatingPet(w http.ResponseWriter, r *http.Request) {
	fmt.Println("This route handles updating a pet!")
}
