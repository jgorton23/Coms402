// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"

	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent/certificationtemplatehistory"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/enthistory"
)

// CertificationTemplateHistory is the model entity for the CertificationTemplateHistory schema.
type CertificationTemplateHistory struct {
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
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// CompanyUUID holds the value of the "companyUUID" field.
	CompanyUUID uuid.UUID `json:"companyUUID,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*CertificationTemplateHistory) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case certificationtemplatehistory.FieldOperation, certificationtemplatehistory.FieldUpdatedBy, certificationtemplatehistory.FieldDescription:
			values[i] = new(sql.NullString)
		case certificationtemplatehistory.FieldHistoryTime:
			values[i] = new(sql.NullTime)
		case certificationtemplatehistory.FieldID, certificationtemplatehistory.FieldRef, certificationtemplatehistory.FieldCompanyUUID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type CertificationTemplateHistory", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the CertificationTemplateHistory fields.
func (cth *CertificationTemplateHistory) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case certificationtemplatehistory.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				cth.ID = *value
			}
		case certificationtemplatehistory.FieldHistoryTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field history_time", values[i])
			} else if value.Valid {
				cth.HistoryTime = value.Time
			}
		case certificationtemplatehistory.FieldRef:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field ref", values[i])
			} else if value != nil {
				cth.Ref = *value
			}
		case certificationtemplatehistory.FieldOperation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field operation", values[i])
			} else if value.Valid {
				cth.Operation = enthistory.OpType(value.String)
			}
		case certificationtemplatehistory.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				cth.UpdatedBy = new(string)
				*cth.UpdatedBy = value.String
			}
		case certificationtemplatehistory.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				cth.Description = value.String
			}
		case certificationtemplatehistory.FieldCompanyUUID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field companyUUID", values[i])
			} else if value != nil {
				cth.CompanyUUID = *value
			}
		}
	}
	return nil
}

// Update returns a builder for updating this CertificationTemplateHistory.
// Note that you need to call CertificationTemplateHistory.Unwrap() before calling this method if this CertificationTemplateHistory
// was returned from a transaction, and the transaction was committed or rolled back.
func (cth *CertificationTemplateHistory) Update() *CertificationTemplateHistoryUpdateOne {
	return NewCertificationTemplateHistoryClient(cth.config).UpdateOne(cth)
}

// Unwrap unwraps the CertificationTemplateHistory entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (cth *CertificationTemplateHistory) Unwrap() *CertificationTemplateHistory {
	_tx, ok := cth.config.driver.(*txDriver)
	if !ok {
		panic("ent: CertificationTemplateHistory is not a transactional entity")
	}
	cth.config.driver = _tx.drv
	return cth
}

// String implements the fmt.Stringer.
func (cth *CertificationTemplateHistory) String() string {
	var builder strings.Builder
	builder.WriteString("CertificationTemplateHistory(")
	builder.WriteString(fmt.Sprintf("id=%v, ", cth.ID))
	builder.WriteString("history_time=")
	builder.WriteString(cth.HistoryTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("ref=")
	builder.WriteString(fmt.Sprintf("%v", cth.Ref))
	builder.WriteString(", ")
	builder.WriteString("operation=")
	builder.WriteString(fmt.Sprintf("%v", cth.Operation))
	builder.WriteString(", ")
	if v := cth.UpdatedBy; v != nil {
		builder.WriteString("updated_by=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(cth.Description)
	builder.WriteString(", ")
	builder.WriteString("companyUUID=")
	builder.WriteString(fmt.Sprintf("%v", cth.CompanyUUID))
	builder.WriteByte(')')
	return builder.String()
}

// CertificationTemplateHistories is a parsable slice of CertificationTemplateHistory.
type CertificationTemplateHistories []*CertificationTemplateHistory
