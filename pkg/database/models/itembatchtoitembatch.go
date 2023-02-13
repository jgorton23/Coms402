package models

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// ItemBatchToItemBatch holds the schema definition for the ItemBatchToItemBatch entity.
type ItemBatchToItemBatch struct {
	ent.Schema
}

// Fields of the ItemBatchToItemBatch.
func (ItemBatchToItemBatch) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			StorageKey("UUID"),
		field.UUID("childUUID", uuid.UUID{}),
		field.UUID("parentUUID", uuid.UUID{}),
	}
}

// Edges of the ItemBatchToItemBatch.
func (ItemBatchToItemBatch) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("parent", ItemBatch.Type).
			Required().
			Unique().
			Field("parentUUID"),
		edge.To("child", ItemBatch.Type).
			Required().
			Unique().
			Field("childUUID"),
	}
}
