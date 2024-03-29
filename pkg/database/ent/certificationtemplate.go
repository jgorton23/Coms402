// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"

	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent/certificationtemplate"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent/company"
)

// CertificationTemplate is the model entity for the CertificationTemplate schema.
type CertificationTemplate struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// CompanyUUID holds the value of the "companyUUID" field.
	CompanyUUID uuid.UUID `json:"companyUUID,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CertificationTemplateQuery when eager-loading is set.
	Edges CertificationTemplateEdges `json:"edges"`
}

// CertificationTemplateEdges holds the relations/edges for other nodes in the graph.
type CertificationTemplateEdges struct {
	// Company holds the value of the company edge.
	Company *Company `json:"company,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// CompanyOrErr returns the Company value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CertificationTemplateEdges) CompanyOrErr() (*Company, error) {
	if e.loadedTypes[0] {
		if e.Company == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: company.Label}
		}
		return e.Company, nil
	}
	return nil, &NotLoadedError{edge: "company"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*CertificationTemplate) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case certificationtemplate.FieldDescription:
			values[i] = new(sql.NullString)
		case certificationtemplate.FieldID, certificationtemplate.FieldCompanyUUID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type CertificationTemplate", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the CertificationTemplate fields.
func (ct *CertificationTemplate) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case certificationtemplate.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				ct.ID = *value
			}
		case certificationtemplate.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				ct.Description = value.String
			}
		case certificationtemplate.FieldCompanyUUID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field companyUUID", values[i])
			} else if value != nil {
				ct.CompanyUUID = *value
			}
		}
	}
	return nil
}

// QueryCompany queries the "company" edge of the CertificationTemplate entity.
func (ct *CertificationTemplate) QueryCompany() *CompanyQuery {
	return NewCertificationTemplateClient(ct.config).QueryCompany(ct)
}

// Update returns a builder for updating this CertificationTemplate.
// Note that you need to call CertificationTemplate.Unwrap() before calling this method if this CertificationTemplate
// was returned from a transaction, and the transaction was committed or rolled back.
func (ct *CertificationTemplate) Update() *CertificationTemplateUpdateOne {
	return NewCertificationTemplateClient(ct.config).UpdateOne(ct)
}

// Unwrap unwraps the CertificationTemplate entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ct *CertificationTemplate) Unwrap() *CertificationTemplate {
	_tx, ok := ct.config.driver.(*txDriver)
	if !ok {
		panic("ent: CertificationTemplate is not a transactional entity")
	}
	ct.config.driver = _tx.drv
	return ct
}

// String implements the fmt.Stringer.
func (ct *CertificationTemplate) String() string {
	var builder strings.Builder
	builder.WriteString("CertificationTemplate(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ct.ID))
	builder.WriteString("description=")
	builder.WriteString(ct.Description)
	builder.WriteString(", ")
	builder.WriteString("companyUUID=")
	builder.WriteString(fmt.Sprintf("%v", ct.CompanyUUID))
	builder.WriteByte(')')
	return builder.String()
}

// CertificationTemplates is a parsable slice of CertificationTemplate.
type CertificationTemplates []*CertificationTemplate
