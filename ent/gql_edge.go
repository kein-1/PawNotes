// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

func (pe *Pet) Owner(ctx context.Context) (*User, error) {
	result, err := pe.Edges.OwnerOrErr()
	if IsNotLoaded(err) {
		result, err = pe.QueryOwner().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (u *User) Pets(ctx context.Context) (result []*Pet, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = u.NamedPets(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = u.Edges.PetsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = u.QueryPets().All(ctx)
	}
	return result, err
}