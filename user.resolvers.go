package pawnotes

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.62

import (
	"context"
	"fmt"

	"github.com/kein-1/pawnotes/ent"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input ent.CreateUserInput) (*ent.User, error) {
	return r.client.User.Create().SetInput(input).Save(ctx)
}

// DeleteUser is the resolver for the deleteUser field.
func (r *mutationResolver) DeleteUser(ctx context.Context, id int) (*ent.User, error) {
	panic(fmt.Errorf("not implemented: DeleteUser - deleteUser"))
}
