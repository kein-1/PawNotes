package pet

import (
	"backend/config"
	"backend/types"
	"backend/utils"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/go-playground/validator/v10"
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
			r.Post("/createPet", h.handleCreatePet)
			r.Delete("/deletePet/{id}", h.handleDeletePet)
			r.Patch("/updatePet/{id}", h.handleUpdatingPet)

		})
	})

}

func (h *PetHandler) handleCreatePet(w http.ResponseWriter, r *http.Request) {

	userID, err := utils.ExtractIDFromJWT(r)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	var pet types.Pet
	if err := utils.ParseJSON(r, &pet); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	fmt.Println("printing pet:", pet)
	fmt.Println("printing pett:", pet.DOB)
	fmt.Println("printing pettt:", pet.DOB.Time)

	if err := utils.Validate.Struct(pet); err != nil {
		validationErrors, ok := err.(validator.ValidationErrors)
		if ok {
			utils.WriteValidationError(w, http.StatusBadRequest, validationErrors)
			return
		}
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := h.petService.AddPet(pet, userID); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	utils.SendJSON(w, 200, map[string]string{"message": "Pet added!"})
}

func (h *PetHandler) handleDeletePet(w http.ResponseWriter, r *http.Request) {

	userID, err := utils.ExtractIDFromJWT(r)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	petID_ := chi.URLParam(r, "id")
	petID, err := strconv.Atoi(petID_)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Error conversion: %w", w))
		return
	}

	err = h.petService.RemovePet(petID, userID)
	if err != nil {
		if errors.Is(err, types.ErrNoRecord) {
			utils.WriteError(w, http.StatusNotFound, fmt.Errorf("Failed to delete pet; no pet with this ID was found"))
		} else {
			utils.WriteError(w, http.StatusInternalServerError, err)
		}
	}
	utils.SendJSON(w, 200, map[string]string{"message": "Pet removed!"})
}

func (h *PetHandler) handleUpdatingPet(w http.ResponseWriter, r *http.Request) {

	userID, err := utils.ExtractIDFromJWT(r)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	petID_ := chi.URLParam(r, "id")
	petID, err := strconv.Atoi(petID_)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Error conversion: %w", w))
		return
	}

	var pet types.PetPatch
	if err := utils.ParseJSON(r, &pet); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	fmt.Println("The payload is:", pet)

	if err := h.petService.UpdatePetAttribute(petID, userID, pet); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.SendJSON(w, 200, map[string]string{"message": "Pet attribute updated!"})

}
