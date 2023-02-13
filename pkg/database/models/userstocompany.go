package models

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// UsersToCompany holds the schema definition for the UsersToCompany entity.
type UsersToCompany struct {
	ent.Schema
}

// Fields of the UsersToCompany.
func (UsersToCompany) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			StorageKey("UUID").
			Unique(),
		field.UUID("companyUUID", uuid.UUID{}),
		// This field currentlly is forced to be an Int instead of UUID, because the user
		// primary key is an int and it has some methods dependent on it (GetById)
		field.Int("userUUID"),
		field.String("roleType"),
		field.Bool("approved"),
	}
}

// Edges of the UsersToCompany.
func (UsersToCompany) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).
			Required().
			Unique().
			Field("userUUID"),
		edge.To("company", Company.Type).
			Required().
			Unique().
			Field("companyUUID"),
	}
}
