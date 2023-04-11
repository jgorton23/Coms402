package enthistory

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type history struct {
	ent.Schema
}

func (history) Fields() []ent.Field {
	return []ent.Field{
		field.Time("history_time").
			Default(time.Now).
			Immutable(),
		field.UUID("ref", uuid.UUID{}).
			Immutable().
			Optional(),
		field.Enum("operation").
			GoType(OpType("")).
			Immutable(),
	}
}
