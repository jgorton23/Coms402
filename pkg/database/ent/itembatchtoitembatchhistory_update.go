// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"

	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent/itembatchtoitembatchhistory"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent/predicate"
)

// ItemBatchToItemBatchHistoryUpdate is the builder for updating ItemBatchToItemBatchHistory entities.
type ItemBatchToItemBatchHistoryUpdate struct {
	config
	hooks    []Hook
	mutation *ItemBatchToItemBatchHistoryMutation
}

// Where appends a list predicates to the ItemBatchToItemBatchHistoryUpdate builder.
func (ibtibhu *ItemBatchToItemBatchHistoryUpdate) Where(ps ...predicate.ItemBatchToItemBatchHistory) *ItemBatchToItemBatchHistoryUpdate {
	ibtibhu.mutation.Where(ps...)
	return ibtibhu
}

// SetChildUUID sets the "childUUID" field.
func (ibtibhu *ItemBatchToItemBatchHistoryUpdate) SetChildUUID(u uuid.UUID) *ItemBatchToItemBatchHistoryUpdate {
	ibtibhu.mutation.SetChildUUID(u)
	return ibtibhu
}

// SetParentUUID sets the "parentUUID" field.
func (ibtibhu *ItemBatchToItemBatchHistoryUpdate) SetParentUUID(u uuid.UUID) *ItemBatchToItemBatchHistoryUpdate {
	ibtibhu.mutation.SetParentUUID(u)
	return ibtibhu
}

// Mutation returns the ItemBatchToItemBatchHistoryMutation object of the builder.
func (ibtibhu *ItemBatchToItemBatchHistoryUpdate) Mutation() *ItemBatchToItemBatchHistoryMutation {
	return ibtibhu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ibtibhu *ItemBatchToItemBatchHistoryUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, ItemBatchToItemBatchHistoryMutation](ctx, ibtibhu.sqlSave, ibtibhu.mutation, ibtibhu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ibtibhu *ItemBatchToItemBatchHistoryUpdate) SaveX(ctx context.Context) int {
	affected, err := ibtibhu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ibtibhu *ItemBatchToItemBatchHistoryUpdate) Exec(ctx context.Context) error {
	_, err := ibtibhu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ibtibhu *ItemBatchToItemBatchHistoryUpdate) ExecX(ctx context.Context) {
	if err := ibtibhu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ibtibhu *ItemBatchToItemBatchHistoryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(itembatchtoitembatchhistory.Table, itembatchtoitembatchhistory.Columns, sqlgraph.NewFieldSpec(itembatchtoitembatchhistory.FieldID, field.TypeUUID))
	if ps := ibtibhu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if ibtibhu.mutation.RefCleared() {
		_spec.ClearField(itembatchtoitembatchhistory.FieldRef, field.TypeUUID)
	}
	if ibtibhu.mutation.UpdatedByCleared() {
		_spec.ClearField(itembatchtoitembatchhistory.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := ibtibhu.mutation.ChildUUID(); ok {
		_spec.SetField(itembatchtoitembatchhistory.FieldChildUUID, field.TypeUUID, value)
	}
	if value, ok := ibtibhu.mutation.ParentUUID(); ok {
		_spec.SetField(itembatchtoitembatchhistory.FieldParentUUID, field.TypeUUID, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ibtibhu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{itembatchtoitembatchhistory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ibtibhu.mutation.done = true
	return n, nil
}

// ItemBatchToItemBatchHistoryUpdateOne is the builder for updating a single ItemBatchToItemBatchHistory entity.
type ItemBatchToItemBatchHistoryUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ItemBatchToItemBatchHistoryMutation
}

// SetChildUUID sets the "childUUID" field.
func (ibtibhuo *ItemBatchToItemBatchHistoryUpdateOne) SetChildUUID(u uuid.UUID) *ItemBatchToItemBatchHistoryUpdateOne {
	ibtibhuo.mutation.SetChildUUID(u)
	return ibtibhuo
}

// SetParentUUID sets the "parentUUID" field.
func (ibtibhuo *ItemBatchToItemBatchHistoryUpdateOne) SetParentUUID(u uuid.UUID) *ItemBatchToItemBatchHistoryUpdateOne {
	ibtibhuo.mutation.SetParentUUID(u)
	return ibtibhuo
}

// Mutation returns the ItemBatchToItemBatchHistoryMutation object of the builder.
func (ibtibhuo *ItemBatchToItemBatchHistoryUpdateOne) Mutation() *ItemBatchToItemBatchHistoryMutation {
	return ibtibhuo.mutation
}

// Where appends a list predicates to the ItemBatchToItemBatchHistoryUpdate builder.
func (ibtibhuo *ItemBatchToItemBatchHistoryUpdateOne) Where(ps ...predicate.ItemBatchToItemBatchHistory) *ItemBatchToItemBatchHistoryUpdateOne {
	ibtibhuo.mutation.Where(ps...)
	return ibtibhuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ibtibhuo *ItemBatchToItemBatchHistoryUpdateOne) Select(field string, fields ...string) *ItemBatchToItemBatchHistoryUpdateOne {
	ibtibhuo.fields = append([]string{field}, fields...)
	return ibtibhuo
}

// Save executes the query and returns the updated ItemBatchToItemBatchHistory entity.
func (ibtibhuo *ItemBatchToItemBatchHistoryUpdateOne) Save(ctx context.Context) (*ItemBatchToItemBatchHistory, error) {
	return withHooks[*ItemBatchToItemBatchHistory, ItemBatchToItemBatchHistoryMutation](ctx, ibtibhuo.sqlSave, ibtibhuo.mutation, ibtibhuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ibtibhuo *ItemBatchToItemBatchHistoryUpdateOne) SaveX(ctx context.Context) *ItemBatchToItemBatchHistory {
	node, err := ibtibhuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ibtibhuo *ItemBatchToItemBatchHistoryUpdateOne) Exec(ctx context.Context) error {
	_, err := ibtibhuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ibtibhuo *ItemBatchToItemBatchHistoryUpdateOne) ExecX(ctx context.Context) {
	if err := ibtibhuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ibtibhuo *ItemBatchToItemBatchHistoryUpdateOne) sqlSave(ctx context.Context) (_node *ItemBatchToItemBatchHistory, err error) {
	_spec := sqlgraph.NewUpdateSpec(itembatchtoitembatchhistory.Table, itembatchtoitembatchhistory.Columns, sqlgraph.NewFieldSpec(itembatchtoitembatchhistory.FieldID, field.TypeUUID))
	id, ok := ibtibhuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ItemBatchToItemBatchHistory.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ibtibhuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, itembatchtoitembatchhistory.FieldID)
		for _, f := range fields {
			if !itembatchtoitembatchhistory.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != itembatchtoitembatchhistory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ibtibhuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if ibtibhuo.mutation.RefCleared() {
		_spec.ClearField(itembatchtoitembatchhistory.FieldRef, field.TypeUUID)
	}
	if ibtibhuo.mutation.UpdatedByCleared() {
		_spec.ClearField(itembatchtoitembatchhistory.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := ibtibhuo.mutation.ChildUUID(); ok {
		_spec.SetField(itembatchtoitembatchhistory.FieldChildUUID, field.TypeUUID, value)
	}
	if value, ok := ibtibhuo.mutation.ParentUUID(); ok {
		_spec.SetField(itembatchtoitembatchhistory.FieldParentUUID, field.TypeUUID, value)
	}
	_node = &ItemBatchToItemBatchHistory{config: ibtibhuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ibtibhuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{itembatchtoitembatchhistory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ibtibhuo.mutation.done = true
	return _node, nil
}