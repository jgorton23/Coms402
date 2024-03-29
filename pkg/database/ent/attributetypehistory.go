// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"

	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent/attributetypehistory"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/enthistory"
)

// AttributeTypeHistory is the model entity for the AttributeTypeHistory schema.
type AttributeTypeHistory struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// HistoryTime holds the value of the "history_time" field.
	HistoryTime time.Time `json:"history_time,omitempty"`
	// Ref holds the value of the "ref" field.
	Ref uuid.UUID `json:"ref,omitempty"`
	// Operation holds the value of the "operation" field.
	Operation enthistory.OpType `json:"operation,omitempty"`
	// UpdatedBy holds the value of the "updated_by" field.
	UpdatedBy *string `json:"updated_by,omitempty"`
	// Key holds the value of the "key" field.
	Key string `json:"key,omitempty"`
	// CompanyUUID holds the value of the "companyUUID" field.
	CompanyUUID uuid.UUID `json:"companyUUID,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AttributeTypeHistory) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case attributetypehistory.FieldOperation, attributetypehistory.FieldUpdatedBy, attributetypehistory.FieldKey:
			values[i] = new(sql.NullString)
		case attributetypehistory.FieldHistoryTime:
			values[i] = new(sql.NullTime)
		case attributetypehistory.FieldID, attributetypehistory.FieldRef, attributetypehistory.FieldCompanyUUID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type AttributeTypeHistory", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AttributeTypeHistory fields.
func (ath *AttributeTypeHistory) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case attributetypehistory.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				ath.ID = *value
			}
		case attributetypehistory.FieldHistoryTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field history_time", values[i])
			} else if value.Valid {
				ath.HistoryTime = value.Time
			}
		case attributetypehistory.FieldRef:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field ref", values[i])
			} else if value != nil {
				ath.Ref = *value
			}
		case attributetypehistory.FieldOperation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field operation", values[i])
			} else if value.Valid {
				ath.Operation = enthistory.OpType(value.String)
			}
		case attributetypehistory.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				ath.UpdatedBy = new(string)
				*ath.UpdatedBy = value.String
			}
		case attributetypehistory.FieldKey:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field key", values[i])
			} else if value.Valid {
				ath.Key = value.String
			}
		case attributetypehistory.FieldCompanyUUID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field companyUUID", values[i])
			} else if value != nil {
				ath.CompanyUUID = *value
			}
		}
	}
	return nil
}

// Update returns a builder for updating this AttributeTypeHistory.
// Note that you need to call AttributeTypeHistory.Unwrap() before calling this method if this AttributeTypeHistory
// was returned from a transaction, and the transaction was committed or rolled back.
func (ath *AttributeTypeHistory) Update() *AttributeTypeHistoryUpdateOne {
	return NewAttributeTypeHistoryClient(ath.config).UpdateOne(ath)
}

// Unwrap unwraps the AttributeTypeHistory entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ath *AttributeTypeHistory) Unwrap() *AttributeTypeHistory {
	_tx, ok := ath.config.driver.(*txDriver)
	if !ok {
		panic("ent: AttributeTypeHistory is not a transactional entity")
	}
	ath.config.driver = _tx.drv
	return ath
}

// String implements the fmt.Stringer.
func (ath *AttributeTypeHistory) String() string {
	var builder strings.Builder
	builder.WriteString("AttributeTypeHistory(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ath.ID))
	builder.WriteString("history_time=")
	builder.WriteString(ath.HistoryTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("ref=")
	builder.WriteString(fmt.Sprintf("%v", ath.Ref))
	builder.WriteString(", ")
	builder.WriteString("operation=")
	builder.WriteString(fmt.Sprintf("%v", ath.Operation))
	builder.WriteString(", ")
	if v := ath.UpdatedBy; v != nil {
		builder.WriteString("updated_by=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("key=")
	builder.WriteString(ath.Key)
	builder.WriteString(", ")
	builder.WriteString("companyUUID=")
	builder.WriteString(fmt.Sprintf("%v", ath.CompanyUUID))
	builder.WriteByte(')')
	return builder.String()
}

// AttributeTypeHistories is a parsable slice of AttributeTypeHistory.
type AttributeTypeHistories []*AttributeTypeHistory
