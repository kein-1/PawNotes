package main

import (
	"backend/db"
	"backend/service/pet"
	"backend/service/user"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
)

func setUpServer(db *pgxpool.Pool) (*chi.Mux, error) {

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	// create user repo(interface), inject db into it
	// create new service(interface), create handler (controller)
	// inject it into handler

	userRepo := user.NewUserRepo(db)
	userService := user.NewUserServiceStruct(userRepo)
	userHandler := user.NewUserHandler(userService)
	userHandler.RegisterRoutes(router)

	// create note service; inject db into it, then inject notestore into note handler
	// noteStore := note.NewNoteStore(db)
	// noteHandler := note.NewNoteHandler(noteStore)
	// noteHandler.RegisterRoutes(router)

	petRepo := pet.NewPetRepo(db)
	petService := pet.NewPetService(petRepo)
	petHandler := pet.NewPetHandler(petService)
	petHandler.RegisterRoutes(router)
	return router, nil
}

func main() {
	db, err := db.NewStoragePostgres()
	if err != nil {
		log.Fatal("Error in db connection")
	}

	router, err := setUpServer(db)
	if err != nil {
		return
	}
	http.ListenAndServe(":8080", router)

}
