package models

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Certification holds the schema definition for the Certification entity.
type Certification struct {
	ent.Schema
}

// Fields of the Certification.
func (Certification) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			StorageKey("UUID"),
		field.String("primaryAttribute"),
		field.UUID("companyUUID", uuid.UUID{}),
		field.UUID("itemBatchUUID", uuid.UUID{}),
		field.UUID("imageUUID", uuid.UUID{}),
		field.UUID("templateUUID", uuid.UUID{}).
			Optional(),
	}
}

// Edges of the Certification.
func (Certification) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("company", Company.Type).
			Required().
			Unique().
			Field("companyUUID"),
		edge.To("itemBatch", ItemBatch.Type).
			Required().
			Unique().
			Field("itemBatchUUID"),
		edge.To("template", CertificationTemplate.Type).
			Unique().
			Field("templateUUID"),
	}
}
