package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// AttributeType holds the schema definition for the AttributeType entity.
type AttributeType struct {
	ent.Schema
}

// Fields of the AttributeType.
func (AttributeType) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			StorageKey("UUID"),
		field.String("key"),
		field.UUID("companyUUID", uuid.UUID{}),
	}
}

// Edges of the AttributeType.
func (AttributeType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("company", Company.Type).
			Required().
			Unique().
			Field("companyUUID"),
	}
}
