package note

import (
	"backend/types"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Actually responsible for interacting with the db

type DataStore struct {
	db *pgxpool.Pool
}

// Returns an instance of the Store struct, which holds a reference to the
// db connection pool.
func NewNoteStore(db *pgxpool.Pool) types.NoteStore {
	return &DataStore{db}
}

func (s *DataStore) CreateNote(note types.Note) error {
	query := "INSERT INTO notes (title, content, owner_id, pet_id, created_at) VALUES ($1,$2,$3,$4,$5)"

	_, err := s.db.Exec(context.Background(), query, note.Title, note.Content, note.OwnerID, note.PetID, note.CreatedAt)
	return err
}
