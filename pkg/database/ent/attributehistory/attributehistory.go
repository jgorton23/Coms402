// Code generated by ent, DO NOT EDIT.

package attributehistory

import (
	"fmt"
	"time"

	"github.com/google/uuid"

	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/enthistory"
)

const (
	// Label holds the string label denoting the attributehistory type in the database.
	Label = "attribute_history"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "UUID"
	// FieldHistoryTime holds the string denoting the history_time field in the database.
	FieldHistoryTime = "history_time"
	// FieldRef holds the string denoting the ref field in the database.
	FieldRef = "ref"
	// FieldOperation holds the string denoting the operation field in the database.
	FieldOperation = "operation"
	// FieldUpdatedBy holds the string denoting the updated_by field in the database.
	FieldUpdatedBy = "updated_by"
	// FieldKey holds the string denoting the key field in the database.
	FieldKey = "key"
	// FieldValue holds the string denoting the value field in the database.
	FieldValue = "value"
	// FieldCertUUID holds the string denoting the certuuid field in the database.
	FieldCertUUID = "cert_uuid"
	// FieldAttributeTypeUUID holds the string denoting the attributetypeuuid field in the database.
	FieldAttributeTypeUUID = "attribute_type_uuid"
	// Table holds the table name of the attributehistory in the database.
	Table = "attribute_history"
)

// Columns holds all SQL columns for attributehistory fields.
var Columns = []string{
	FieldID,
	FieldHistoryTime,
	FieldRef,
	FieldOperation,
	FieldUpdatedBy,
	FieldKey,
	FieldValue,
	FieldCertUUID,
	FieldAttributeTypeUUID,
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
	// DefaultHistoryTime holds the default value on creation for the "history_time" field.
	DefaultHistoryTime func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OperationValidator is a validator for the "operation" field enum values. It is called by the builders before save.
func OperationValidator(o enthistory.OpType) error {
	switch o.String() {
	case "INSERT", "UPDATE", "DELETE":
		return nil
	default:
		return fmt.Errorf("attributehistory: invalid enum value for operation field: %q", o)
	}
}
