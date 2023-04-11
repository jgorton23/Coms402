package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// ItemBatch holds the schema definition for the ItemBatch entity.
type ItemBatch struct {
	ent.Schema
}

// Fields of the ItemBatch.
func (ItemBatch) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			StorageKey("UUID"),
		field.String("itemNumber"),
		field.String("description"),
		field.UUID("companyUUID", uuid.UUID{}),
	}
}

// Edges of the ItemBatch.
func (ItemBatch) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("company", Company.Type).
			Required().
			Unique().
			Field("companyUUID"),
	}
}
