package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"

	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/enthistory"
)

// Session holds the schema definition for the Session entity.
type Session struct {
	ent.Schema
}

func (Session) Annotations() []schema.Annotation {
	return []schema.Annotation{
		enthistory.Annotations{
			// Tells enthistory to exclude history tables for this schema
			Exclude: true,
		},
	}
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
