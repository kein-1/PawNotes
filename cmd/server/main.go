package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/kein-1/pawnotes"
	"github.com/kein-1/pawnotes/ent"
	"github.com/kein-1/pawnotes/ent/migrate"
	auth "github.com/kein-1/pawnotes/internal/jwt"

	// "github.com/kein-1/pawnotes/ent"
	// "github.com/kein-1/pawnotes/ent/migrate"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql/handler"
	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	"github.com/kein-1/pawnotes/config"
)

// Open new connection
func Open() *ent.Client {
	uri := config.ConfigValues.DATABASE_URI
	db, err := sql.Open("pgx", uri)
	if err != nil {
		log.Fatal(err)
	}

	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.Postgres, db)

	return ent.NewClient(ent.Driver(drv))
}

func SetupServer() (*chi.Mux, error) {

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"}, // change in production
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	return r, nil

}

func main() {
	// Create ent.Client and run the schema migration.
	client := Open()

	if err := client.Schema.Create(
		context.Background(),
		migrate.WithGlobalUniqueID(true),
	); err != nil {
		log.Fatal("opening ent client", err)
	}

	fmt.Println("Db successfully connected")

	r, err := SetupServer()
	if err != nil {
		log.Fatal("Error setting up server")
	}

	tokenStr, err := auth.GenerateToken(8589934592)
	if err != nil {
		log.Fatal("Error creating token", err.Error())
	}
	fmt.Println("Token str:", tokenStr)

	// Configure the server and start listening on :8080.
	srv := handler.NewDefaultServer(pawnotes.NewSchema(client))
	r.Handle("/graphql", srv)
	log.Println("listening on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("http. server terminated", err)
	}
}
