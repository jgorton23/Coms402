package models

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// CertificationTemplate holds the schema definition for the CertificationTemplate entity.
type CertificationTemplate struct {
	ent.Schema
}

// Fields of the CertificationTemplate.
func (CertificationTemplate) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			StorageKey("UUID"),
		field.String("description"),
		field.UUID("companyUUID", uuid.UUID{}),
	}
}

// Edges of the CertificationTemplate.
func (CertificationTemplate) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("company", Company.Type).
			Required().
			Unique().
			Field("companyUUID"),
	}
}
