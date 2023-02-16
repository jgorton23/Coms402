package models

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Session holds the schema definition for the Session entity.
type Session struct {
	ent.Schema
}

// Fields of the Session.
func (Session) Fields() []ent.Field {
	return []ent.Field{
		field.String("token").
			Unique(),
		field.Bytes("data").
			NotEmpty(),
		field.Time("expiry"),
	}
}

// Edges of the Session.
func (Session) Edges() []ent.Edge {
	return nil
}
