// Code generated by ent, DO NOT EDIT.

package attributetypestotemplates

import (
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the attributetypestotemplates type in the database.
	Label = "attribute_types_to_templates"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "UUID"
	// FieldAttributeTypeUUID holds the string denoting the attributetypeuuid field in the database.
	FieldAttributeTypeUUID = "attribute_type_uuid"
	// FieldTemplateUUID holds the string denoting the templateuuid field in the database.
	FieldTemplateUUID = "template_uuid"
	// EdgeAttribute holds the string denoting the attribute edge name in mutations.
	EdgeAttribute = "attribute"
	// EdgeTemplate holds the string denoting the template edge name in mutations.
	EdgeTemplate = "template"
	// Table holds the table name of the attributetypestotemplates in the database.
	Table = "attribute_types_to_templates"
	// AttributeTable is the table that holds the attribute relation/edge.
	AttributeTable = "attribute_types_to_templates"
	// AttributeInverseTable is the table name for the AttributeType entity.
	// It exists in this package in order to avoid circular dependency with the "attributetype" package.
	AttributeInverseTable = "attribute_types"
	// AttributeColumn is the table column denoting the attribute relation/edge.
	AttributeColumn = "attribute_type_uuid"
	// TemplateTable is the table that holds the template relation/edge.
	TemplateTable = "attribute_types_to_templates"
	// TemplateInverseTable is the table name for the CertificationTemplate entity.
	// It exists in this package in order to avoid circular dependency with the "certificationtemplate" package.
	TemplateInverseTable = "certification_templates"
	// TemplateColumn is the table column denoting the template relation/edge.
	TemplateColumn = "template_uuid"
)

// Columns holds all SQL columns for attributetypestotemplates fields.
var Columns = []string{
	FieldID,
	FieldAttributeTypeUUID,
	FieldTemplateUUID,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)