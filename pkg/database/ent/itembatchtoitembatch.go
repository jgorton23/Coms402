// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/google/uuid"

	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent/itembatch"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent/itembatchtoitembatch"
)

// ItemBatchToItemBatch is the model entity for the ItemBatchToItemBatch schema.
type ItemBatchToItemBatch struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// ChildUUID holds the value of the "childUUID" field.
	ChildUUID uuid.UUID `json:"childUUID,omitempty"`
	// ParentUUID holds the value of the "parentUUID" field.
	ParentUUID uuid.UUID `json:"parentUUID,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ItemBatchToItemBatchQuery when eager-loading is set.
	Edges ItemBatchToItemBatchEdges `json:"edges"`
}

// ItemBatchToItemBatchEdges holds the relations/edges for other nodes in the graph.
type ItemBatchToItemBatchEdges struct {
	// Parent holds the value of the parent edge.
	Parent *ItemBatch `json:"parent,omitempty"`
	// Child holds the value of the child edge.
	Child *ItemBatch `json:"child,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ItemBatchToItemBatchEdges) ParentOrErr() (*ItemBatch, error) {
	if e.loadedTypes[0] {
		if e.Parent == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: itembatch.Label}
		}
		return e.Parent, nil
	}
	return nil, &NotLoadedError{edge: "parent"}
}

// ChildOrErr returns the Child value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ItemBatchToItemBatchEdges) ChildOrErr() (*ItemBatch, error) {
	if e.loadedTypes[1] {
		if e.Child == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: itembatch.Label}
		}
		return e.Child, nil
	}
	return nil, &NotLoadedError{edge: "child"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ItemBatchToItemBatch) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case itembatchtoitembatch.FieldID, itembatchtoitembatch.FieldChildUUID, itembatchtoitembatch.FieldParentUUID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type ItemBatchToItemBatch", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ItemBatchToItemBatch fields.
func (ibtib *ItemBatchToItemBatch) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case itembatchtoitembatch.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				ibtib.ID = *value
			}
		case itembatchtoitembatch.FieldChildUUID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field childUUID", values[i])
			} else if value != nil {
				ibtib.ChildUUID = *value
			}
		case itembatchtoitembatch.FieldParentUUID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field parentUUID", values[i])
			} else if value != nil {
				ibtib.ParentUUID = *value
			}
		}
	}
	return nil
}

// QueryParent queries the "parent" edge of the ItemBatchToItemBatch entity.
func (ibtib *ItemBatchToItemBatch) QueryParent() *ItemBatchQuery {
	return NewItemBatchToItemBatchClient(ibtib.config).QueryParent(ibtib)
}

// QueryChild queries the "child" edge of the ItemBatchToItemBatch entity.
func (ibtib *ItemBatchToItemBatch) QueryChild() *ItemBatchQuery {
	return NewItemBatchToItemBatchClient(ibtib.config).QueryChild(ibtib)
}

// Update returns a builder for updating this ItemBatchToItemBatch.
// Note that you need to call ItemBatchToItemBatch.Unwrap() before calling this method if this ItemBatchToItemBatch
// was returned from a transaction, and the transaction was committed or rolled back.
func (ibtib *ItemBatchToItemBatch) Update() *ItemBatchToItemBatchUpdateOne {
	return NewItemBatchToItemBatchClient(ibtib.config).UpdateOne(ibtib)
}

// Unwrap unwraps the ItemBatchToItemBatch entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ibtib *ItemBatchToItemBatch) Unwrap() *ItemBatchToItemBatch {
	_tx, ok := ibtib.config.driver.(*txDriver)
	if !ok {
		panic("ent: ItemBatchToItemBatch is not a transactional entity")
	}
	ibtib.config.driver = _tx.drv
	return ibtib
}

// String implements the fmt.Stringer.
func (ibtib *ItemBatchToItemBatch) String() string {
	var builder strings.Builder
	builder.WriteString("ItemBatchToItemBatch(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ibtib.ID))
	builder.WriteString("childUUID=")
	builder.WriteString(fmt.Sprintf("%v", ibtib.ChildUUID))
	builder.WriteString(", ")
	builder.WriteString("parentUUID=")
	builder.WriteString(fmt.Sprintf("%v", ibtib.ParentUUID))
	builder.WriteByte(')')
	return builder.String()
}

// ItemBatchToItemBatches is a parsable slice of ItemBatchToItemBatch.
type ItemBatchToItemBatches []*ItemBatchToItemBatch