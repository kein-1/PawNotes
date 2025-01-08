// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/kein-1/pawnotes/ent/pet"
	"github.com/kein-1/pawnotes/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	petFields := schema.Pet{}.Fields()
	_ = petFields
	// petDescCreatedAt is the schema descriptor for created_at field.
	petDescCreatedAt := petFields[4].Descriptor()
	// pet.DefaultCreatedAt holds the default value on creation for the created_at field.
	pet.DefaultCreatedAt = petDescCreatedAt.Default.(time.Time)
}
