package note

import (
	"backend/types"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

type NoteHandler struct {
	noteStore types.NoteStore
}

// kind of like the view model in Swift's MVVM; this holds a reference to the
// serivce, which is the "Manager" in a sense that intereacts with the db and
// handles the logic
// Similarly we could put something like a network manager here
func NewNoteHandler(noteStore types.NoteStore) *NoteHandler {
	return &NoteHandler{noteStore: noteStore}
}

func (h *NoteHandler) RegisterRoutes(router *chi.Mux) {
	router.Post("/api/note/createNote", h.handleCreateNote)
}

func (h *NoteHandler) handleCreateNote(w http.ResponseWriter, r *http.Request) {

	note := types.Note{
		Title:     "My first note",
		Content:   "This is some long text",
		OwnerID:   1,
		PetID:     1,
		CreatedAt: time.Now(),
	}

	err := h.noteStore.CreateNote(note)
	if err != nil {
		fmt.Println("Error in inserting a note", err)
	}
}
