// Code generated by ent, DO NOT EDIT.

package authorizationpolicy

const (
	// Label holds the string label denoting the authorizationpolicy type in the database.
	Label = "authorization_policy"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldPtype holds the string denoting the ptype field in the database.
	FieldPtype = "ptype"
	// FieldV0 holds the string denoting the v0 field in the database.
	FieldV0 = "v0"
	// FieldV1 holds the string denoting the v1 field in the database.
	FieldV1 = "v1"
	// FieldV2 holds the string denoting the v2 field in the database.
	FieldV2 = "v2"
	// FieldV3 holds the string denoting the v3 field in the database.
	FieldV3 = "v3"
	// FieldV4 holds the string denoting the v4 field in the database.
	FieldV4 = "v4"
	// FieldV5 holds the string denoting the v5 field in the database.
	FieldV5 = "v5"
	// Table holds the table name of the authorizationpolicy in the database.
	Table = "authorization_policies"
)

// Columns holds all SQL columns for authorizationpolicy fields.
var Columns = []string{
	FieldID,
	FieldPtype,
	FieldV0,
	FieldV1,
	FieldV2,
	FieldV3,
	FieldV4,
	FieldV5,
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
	// DefaultPtype holds the default value on creation for the "Ptype" field.
	DefaultPtype string
	// DefaultV0 holds the default value on creation for the "V0" field.
	DefaultV0 string
	// DefaultV1 holds the default value on creation for the "V1" field.
	DefaultV1 string
	// DefaultV2 holds the default value on creation for the "V2" field.
	DefaultV2 string
	// DefaultV3 holds the default value on creation for the "V3" field.
	DefaultV3 string
	// DefaultV4 holds the default value on creation for the "V4" field.
	DefaultV4 string
	// DefaultV5 holds the default value on creation for the "V5" field.
	DefaultV5 string
)