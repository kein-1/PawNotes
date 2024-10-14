package main

import (
	"backend/db"
	"backend/service/note"
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

	// create user store, inject db into it, then inject user store into handler
	userStore := user.NewUserStore(db)
	userHandler := user.NewUserHandler(userStore)
	userHandler.RegisterRoutes(router)

	// create note service; inject db into it, then inject notestore into note handler
	noteStore := note.NewNoteStore(db)
	noteHandler := note.NewNoteHandler(noteStore)
	noteHandler.RegisterRoutes(router)

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
