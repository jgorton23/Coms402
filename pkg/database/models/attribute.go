package models

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Attribute holds the schema definition for the Attribute entity.
type Attribute struct {
	ent.Schema
}

// Fields of the Attribute.
func (Attribute) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			StorageKey("UUID"),
		field.String("key"),
		field.String("value"),
		field.UUID("certUUID", uuid.UUID{}),
		field.UUID("attributeTypeUUID", uuid.UUID{}),
	}
}

// Edges of the Attribute.
func (Attribute) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("certification", Certification.Type).
			Required().
			Unique().
			Field("certUUID"),
		edge.To("attributeType", AttributeType.Type).
			Required().
			Unique().
			Field("attributeTypeUUID"),
	}
}
