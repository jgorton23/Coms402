package models

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// AttributeTypesToTemplates holds the schema definition for the AttributeTypesToTemplates entity.
type AttributeTypesToTemplates struct {
	ent.Schema
}

// Fields of the AttributeTypesToTemplates.
func (AttributeTypesToTemplates) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			StorageKey("UUID"),
		field.UUID("attributeTypeUUID", uuid.UUID{}),
		field.UUID("templateUUID", uuid.UUID{}),
	}
}

// Edges of the AttributeTypesToTemplates.
func (AttributeTypesToTemplates) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("attribute", AttributeType.Type).
			Required().
			Unique().
			Field("attributeTypeUUID"),
		edge.To("template", CertificationTemplate.Type).
			Required().
			Unique().
			Field("templateUUID"),
	}
}
