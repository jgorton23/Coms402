package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Company holds the schema definition for the Company entity.
type Company struct {
	ent.Schema
}

// Fields of the Company.
func (Company) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			StorageKey("UUID").
			Unique(),
		field.String("name"),
	}
}

// Edges of the Company.
func (Company) Edges() []ent.Edge {
	return nil
}
