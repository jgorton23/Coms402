// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"

	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent/itembatchtoitembatchhistory"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/enthistory"
)

// ItemBatchToItemBatchHistory is the model entity for the ItemBatchToItemBatchHistory schema.
type ItemBatchToItemBatchHistory struct {
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
	// ChildUUID holds the value of the "childUUID" field.
	ChildUUID uuid.UUID `json:"childUUID,omitempty"`
	// ParentUUID holds the value of the "parentUUID" field.
	ParentUUID uuid.UUID `json:"parentUUID,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ItemBatchToItemBatchHistory) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case itembatchtoitembatchhistory.FieldOperation, itembatchtoitembatchhistory.FieldUpdatedBy:
			values[i] = new(sql.NullString)
		case itembatchtoitembatchhistory.FieldHistoryTime:
			values[i] = new(sql.NullTime)
		case itembatchtoitembatchhistory.FieldID, itembatchtoitembatchhistory.FieldRef, itembatchtoitembatchhistory.FieldChildUUID, itembatchtoitembatchhistory.FieldParentUUID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type ItemBatchToItemBatchHistory", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ItemBatchToItemBatchHistory fields.
func (ibtibh *ItemBatchToItemBatchHistory) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case itembatchtoitembatchhistory.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				ibtibh.ID = *value
			}
		case itembatchtoitembatchhistory.FieldHistoryTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field history_time", values[i])
			} else if value.Valid {
				ibtibh.HistoryTime = value.Time
			}
		case itembatchtoitembatchhistory.FieldRef:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field ref", values[i])
			} else if value != nil {
				ibtibh.Ref = *value
			}
		case itembatchtoitembatchhistory.FieldOperation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field operation", values[i])
			} else if value.Valid {
				ibtibh.Operation = enthistory.OpType(value.String)
			}
		case itembatchtoitembatchhistory.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				ibtibh.UpdatedBy = new(string)
				*ibtibh.UpdatedBy = value.String
			}
		case itembatchtoitembatchhistory.FieldChildUUID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field childUUID", values[i])
			} else if value != nil {
				ibtibh.ChildUUID = *value
			}
		case itembatchtoitembatchhistory.FieldParentUUID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field parentUUID", values[i])
			} else if value != nil {
				ibtibh.ParentUUID = *value
			}
		}
	}
	return nil
}

// Update returns a builder for updating this ItemBatchToItemBatchHistory.
// Note that you need to call ItemBatchToItemBatchHistory.Unwrap() before calling this method if this ItemBatchToItemBatchHistory
// was returned from a transaction, and the transaction was committed or rolled back.
func (ibtibh *ItemBatchToItemBatchHistory) Update() *ItemBatchToItemBatchHistoryUpdateOne {
	return NewItemBatchToItemBatchHistoryClient(ibtibh.config).UpdateOne(ibtibh)
}

// Unwrap unwraps the ItemBatchToItemBatchHistory entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ibtibh *ItemBatchToItemBatchHistory) Unwrap() *ItemBatchToItemBatchHistory {
	_tx, ok := ibtibh.config.driver.(*txDriver)
	if !ok {
		panic("ent: ItemBatchToItemBatchHistory is not a transactional entity")
	}
	ibtibh.config.driver = _tx.drv
	return ibtibh
}

// String implements the fmt.Stringer.
func (ibtibh *ItemBatchToItemBatchHistory) String() string {
	var builder strings.Builder
	builder.WriteString("ItemBatchToItemBatchHistory(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ibtibh.ID))
	builder.WriteString("history_time=")
	builder.WriteString(ibtibh.HistoryTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("ref=")
	builder.WriteString(fmt.Sprintf("%v", ibtibh.Ref))
	builder.WriteString(", ")
	builder.WriteString("operation=")
	builder.WriteString(fmt.Sprintf("%v", ibtibh.Operation))
	builder.WriteString(", ")
	if v := ibtibh.UpdatedBy; v != nil {
		builder.WriteString("updated_by=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("childUUID=")
	builder.WriteString(fmt.Sprintf("%v", ibtibh.ChildUUID))
	builder.WriteString(", ")
	builder.WriteString("parentUUID=")
	builder.WriteString(fmt.Sprintf("%v", ibtibh.ParentUUID))
	builder.WriteByte(')')
	return builder.String()
}

// ItemBatchToItemBatchHistories is a parsable slice of ItemBatchToItemBatchHistory.
type ItemBatchToItemBatchHistories []*ItemBatchToItemBatchHistory
