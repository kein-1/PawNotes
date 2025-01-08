// Code generated by ent, DO NOT EDIT.

package user

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/kein-1/pawnotes/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.User {
	return predicate.User(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.User {
	return predicate.User(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.User {
	return predicate.User(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.User {
	return predicate.User(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.User {
	return predicate.User(sql.FieldLTE(FieldID, id))
}

// First applies equality check predicate on the "first" field. It's identical to FirstEQ.
func First(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldFirst, v))
}

// Last applies equality check predicate on the "last" field. It's identical to LastEQ.
func Last(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldLast, v))
}

// Email applies equality check predicate on the "email" field. It's identical to EmailEQ.
func Email(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldEmail, v))
}

// FirstEQ applies the EQ predicate on the "first" field.
func FirstEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldFirst, v))
}

// FirstNEQ applies the NEQ predicate on the "first" field.
func FirstNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldFirst, v))
}

// FirstIn applies the In predicate on the "first" field.
func FirstIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldFirst, vs...))
}

// FirstNotIn applies the NotIn predicate on the "first" field.
func FirstNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldFirst, vs...))
}

// FirstGT applies the GT predicate on the "first" field.
func FirstGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldFirst, v))
}

// FirstGTE applies the GTE predicate on the "first" field.
func FirstGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldFirst, v))
}

// FirstLT applies the LT predicate on the "first" field.
func FirstLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldFirst, v))
}

// FirstLTE applies the LTE predicate on the "first" field.
func FirstLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldFirst, v))
}

// FirstContains applies the Contains predicate on the "first" field.
func FirstContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldFirst, v))
}

// FirstHasPrefix applies the HasPrefix predicate on the "first" field.
func FirstHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldFirst, v))
}

// FirstHasSuffix applies the HasSuffix predicate on the "first" field.
func FirstHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldFirst, v))
}

// FirstEqualFold applies the EqualFold predicate on the "first" field.
func FirstEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldFirst, v))
}

// FirstContainsFold applies the ContainsFold predicate on the "first" field.
func FirstContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldFirst, v))
}

// LastEQ applies the EQ predicate on the "last" field.
func LastEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldLast, v))
}

// LastNEQ applies the NEQ predicate on the "last" field.
func LastNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldLast, v))
}

// LastIn applies the In predicate on the "last" field.
func LastIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldLast, vs...))
}

// LastNotIn applies the NotIn predicate on the "last" field.
func LastNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldLast, vs...))
}

// LastGT applies the GT predicate on the "last" field.
func LastGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldLast, v))
}

// LastGTE applies the GTE predicate on the "last" field.
func LastGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldLast, v))
}

// LastLT applies the LT predicate on the "last" field.
func LastLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldLast, v))
}

// LastLTE applies the LTE predicate on the "last" field.
func LastLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldLast, v))
}

// LastContains applies the Contains predicate on the "last" field.
func LastContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldLast, v))
}

// LastHasPrefix applies the HasPrefix predicate on the "last" field.
func LastHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldLast, v))
}

// LastHasSuffix applies the HasSuffix predicate on the "last" field.
func LastHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldLast, v))
}

// LastEqualFold applies the EqualFold predicate on the "last" field.
func LastEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldLast, v))
}

// LastContainsFold applies the ContainsFold predicate on the "last" field.
func LastContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldLast, v))
}

// EmailEQ applies the EQ predicate on the "email" field.
func EmailEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldEmail, v))
}

// EmailNEQ applies the NEQ predicate on the "email" field.
func EmailNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldEmail, v))
}

// EmailIn applies the In predicate on the "email" field.
func EmailIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldEmail, vs...))
}

// EmailNotIn applies the NotIn predicate on the "email" field.
func EmailNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldEmail, vs...))
}

// EmailGT applies the GT predicate on the "email" field.
func EmailGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldEmail, v))
}

// EmailGTE applies the GTE predicate on the "email" field.
func EmailGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldEmail, v))
}

// EmailLT applies the LT predicate on the "email" field.
func EmailLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldEmail, v))
}

// EmailLTE applies the LTE predicate on the "email" field.
func EmailLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldEmail, v))
}

// EmailContains applies the Contains predicate on the "email" field.
func EmailContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldEmail, v))
}

// EmailHasPrefix applies the HasPrefix predicate on the "email" field.
func EmailHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldEmail, v))
}

// EmailHasSuffix applies the HasSuffix predicate on the "email" field.
func EmailHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldEmail, v))
}

// EmailEqualFold applies the EqualFold predicate on the "email" field.
func EmailEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldEmail, v))
}

// EmailContainsFold applies the ContainsFold predicate on the "email" field.
func EmailContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldEmail, v))
}

// HasPets applies the HasEdge predicate on the "pets" edge.
func HasPets() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PetsTable, PetsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPetsWith applies the HasEdge predicate on the "pets" edge with a given conditions (other predicates).
func HasPetsWith(preds ...predicate.Pet) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newPetsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.User) predicate.User {
	return predicate.User(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.User) predicate.User {
	return predicate.User(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.User) predicate.User {
	return predicate.User(sql.NotPredicates(p))
}