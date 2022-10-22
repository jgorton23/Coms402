package models

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// AuthorizationPolicy holds the schema definition for the AuthorizationPolicy entity.
type AuthorizationPolicy struct {
	ent.Schema
}

// Fields of the AuthorizationPolicy.
func (AuthorizationPolicy) Fields() []ent.Field {
	return []ent.Field{
		field.String("Ptype").Default(""),
		field.String("V0").Default(""),
		field.String("V1").Default(""),
		field.String("V2").Default(""),
		field.String("V3").Default(""),
		field.String("V4").Default(""),
		field.String("V5").Default(""),
	}
}

// Edges of the AuthorizationPolicy.
func (AuthorizationPolicy) Edges() []ent.Edge {
	return nil
}

func (AuthorizationPolicy) Index() []ent.Index {
	return []ent.Index{
		index.Fields("Ptype", "V0", "V1", "V2", "V3", "V4", "V5").Unique(),
	}
}
