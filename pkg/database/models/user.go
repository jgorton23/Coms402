package models

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// User holds the schema definition for the User domain.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			StorageKey("UUID"),
		field.String("email").
			Unique(),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Time("updated_at").
			Default(time.Now()),
		field.String("password_hash").
			Optional(),
		// field.String("confirm_selector").
		// 	Optional(),
		// field.String("confirm_verifier").
		// 	Optional(),
		// field.Bool("confirmed").
		// 	Optional(),
		field.Int("attempt_count").
			Optional(),
		field.Time("last_attempt").
			Optional(),
		field.Time("locked").
			Optional(),
		// field.String("recover_selector").
		// 	Optional(),
		// field.String("recover_verifier").
		// 	Optional(),
		// field.Time("recover_token_expiry").
		// 	Optional(),
		field.String("role").
			Default("user"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
