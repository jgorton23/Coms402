// Code generated by ent, DO NOT EDIT.

package itembatch

import (
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the itembatch type in the database.
	Label = "item_batch"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "UUID"
	// FieldItemNumber holds the string denoting the itemnumber field in the database.
	FieldItemNumber = "item_number"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldCompanyUUID holds the string denoting the companyuuid field in the database.
	FieldCompanyUUID = "company_uuid"
	// EdgeCompany holds the string denoting the company edge name in mutations.
	EdgeCompany = "company"
	// Table holds the table name of the itembatch in the database.
	Table = "item_batches"
	// CompanyTable is the table that holds the company relation/edge.
	CompanyTable = "item_batches"
	// CompanyInverseTable is the table name for the Company entity.
	// It exists in this package in order to avoid circular dependency with the "company" package.
	CompanyInverseTable = "companies"
	// CompanyColumn is the table column denoting the company relation/edge.
	CompanyColumn = "company_uuid"
)

// Columns holds all SQL columns for itembatch fields.
var Columns = []string{
	FieldID,
	FieldItemNumber,
	FieldDescription,
	FieldCompanyUUID,
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
