// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"

	"github.com/MatthewBehnke/apis/pkg/database/ent/authorizationpolicy"
)

// AuthorizationPolicy is the model entity for the AuthorizationPolicy schema.
type AuthorizationPolicy struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Ptype holds the value of the "Ptype" field.
	Ptype string `json:"Ptype,omitempty"`
	// V0 holds the value of the "V0" field.
	V0 string `json:"V0,omitempty"`
	// V1 holds the value of the "V1" field.
	V1 string `json:"V1,omitempty"`
	// V2 holds the value of the "V2" field.
	V2 string `json:"V2,omitempty"`
	// V3 holds the value of the "V3" field.
	V3 string `json:"V3,omitempty"`
	// V4 holds the value of the "V4" field.
	V4 string `json:"V4,omitempty"`
	// V5 holds the value of the "V5" field.
	V5 string `json:"V5,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AuthorizationPolicy) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case authorizationpolicy.FieldID:
			values[i] = new(sql.NullInt64)
		case authorizationpolicy.FieldPtype, authorizationpolicy.FieldV0, authorizationpolicy.FieldV1, authorizationpolicy.FieldV2, authorizationpolicy.FieldV3, authorizationpolicy.FieldV4, authorizationpolicy.FieldV5:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type AuthorizationPolicy", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AuthorizationPolicy fields.
func (ap *AuthorizationPolicy) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case authorizationpolicy.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ap.ID = int(value.Int64)
		case authorizationpolicy.FieldPtype:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Ptype", values[i])
			} else if value.Valid {
				ap.Ptype = value.String
			}
		case authorizationpolicy.FieldV0:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field V0", values[i])
			} else if value.Valid {
				ap.V0 = value.String
			}
		case authorizationpolicy.FieldV1:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field V1", values[i])
			} else if value.Valid {
				ap.V1 = value.String
			}
		case authorizationpolicy.FieldV2:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field V2", values[i])
			} else if value.Valid {
				ap.V2 = value.String
			}
		case authorizationpolicy.FieldV3:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field V3", values[i])
			} else if value.Valid {
				ap.V3 = value.String
			}
		case authorizationpolicy.FieldV4:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field V4", values[i])
			} else if value.Valid {
				ap.V4 = value.String
			}
		case authorizationpolicy.FieldV5:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field V5", values[i])
			} else if value.Valid {
				ap.V5 = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this AuthorizationPolicy.
// Note that you need to call AuthorizationPolicy.Unwrap() before calling this method if this AuthorizationPolicy
// was returned from a transaction, and the transaction was committed or rolled back.
func (ap *AuthorizationPolicy) Update() *AuthorizationPolicyUpdateOne {
	return (&AuthorizationPolicyClient{config: ap.config}).UpdateOne(ap)
}

// Unwrap unwraps the AuthorizationPolicy entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ap *AuthorizationPolicy) Unwrap() *AuthorizationPolicy {
	_tx, ok := ap.config.driver.(*txDriver)
	if !ok {
		panic("ent: AuthorizationPolicy is not a transactional entity")
	}
	ap.config.driver = _tx.drv
	return ap
}

// String implements the fmt.Stringer.
func (ap *AuthorizationPolicy) String() string {
	var builder strings.Builder
	builder.WriteString("AuthorizationPolicy(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ap.ID))
	builder.WriteString("Ptype=")
	builder.WriteString(ap.Ptype)
	builder.WriteString(", ")
	builder.WriteString("V0=")
	builder.WriteString(ap.V0)
	builder.WriteString(", ")
	builder.WriteString("V1=")
	builder.WriteString(ap.V1)
	builder.WriteString(", ")
	builder.WriteString("V2=")
	builder.WriteString(ap.V2)
	builder.WriteString(", ")
	builder.WriteString("V3=")
	builder.WriteString(ap.V3)
	builder.WriteString(", ")
	builder.WriteString("V4=")
	builder.WriteString(ap.V4)
	builder.WriteString(", ")
	builder.WriteString("V5=")
	builder.WriteString(ap.V5)
	builder.WriteByte(')')
	return builder.String()
}

// AuthorizationPolicies is a parsable slice of AuthorizationPolicy.
type AuthorizationPolicies []*AuthorizationPolicy

func (ap AuthorizationPolicies) config(cfg config) {
	for _i := range ap {
		ap[_i].config = cfg
	}
}
