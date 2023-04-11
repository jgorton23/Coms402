// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"

	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent/companyhistory"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent/predicate"
)

// CompanyHistoryUpdate is the builder for updating CompanyHistory entities.
type CompanyHistoryUpdate struct {
	config
	hooks    []Hook
	mutation *CompanyHistoryMutation
}

// Where appends a list predicates to the CompanyHistoryUpdate builder.
func (chu *CompanyHistoryUpdate) Where(ps ...predicate.CompanyHistory) *CompanyHistoryUpdate {
	chu.mutation.Where(ps...)
	return chu
}

// SetName sets the "name" field.
func (chu *CompanyHistoryUpdate) SetName(s string) *CompanyHistoryUpdate {
	chu.mutation.SetName(s)
	return chu
}

// Mutation returns the CompanyHistoryMutation object of the builder.
func (chu *CompanyHistoryUpdate) Mutation() *CompanyHistoryMutation {
	return chu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (chu *CompanyHistoryUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, CompanyHistoryMutation](ctx, chu.sqlSave, chu.mutation, chu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (chu *CompanyHistoryUpdate) SaveX(ctx context.Context) int {
	affected, err := chu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (chu *CompanyHistoryUpdate) Exec(ctx context.Context) error {
	_, err := chu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (chu *CompanyHistoryUpdate) ExecX(ctx context.Context) {
	if err := chu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (chu *CompanyHistoryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(companyhistory.Table, companyhistory.Columns, sqlgraph.NewFieldSpec(companyhistory.FieldID, field.TypeUUID))
	if ps := chu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if chu.mutation.RefCleared() {
		_spec.ClearField(companyhistory.FieldRef, field.TypeUUID)
	}
	if chu.mutation.UpdatedByCleared() {
		_spec.ClearField(companyhistory.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := chu.mutation.Name(); ok {
		_spec.SetField(companyhistory.FieldName, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, chu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{companyhistory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	chu.mutation.done = true
	return n, nil
}

// CompanyHistoryUpdateOne is the builder for updating a single CompanyHistory entity.
type CompanyHistoryUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CompanyHistoryMutation
}

// SetName sets the "name" field.
func (chuo *CompanyHistoryUpdateOne) SetName(s string) *CompanyHistoryUpdateOne {
	chuo.mutation.SetName(s)
	return chuo
}

// Mutation returns the CompanyHistoryMutation object of the builder.
func (chuo *CompanyHistoryUpdateOne) Mutation() *CompanyHistoryMutation {
	return chuo.mutation
}

// Where appends a list predicates to the CompanyHistoryUpdate builder.
func (chuo *CompanyHistoryUpdateOne) Where(ps ...predicate.CompanyHistory) *CompanyHistoryUpdateOne {
	chuo.mutation.Where(ps...)
	return chuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (chuo *CompanyHistoryUpdateOne) Select(field string, fields ...string) *CompanyHistoryUpdateOne {
	chuo.fields = append([]string{field}, fields...)
	return chuo
}

// Save executes the query and returns the updated CompanyHistory entity.
func (chuo *CompanyHistoryUpdateOne) Save(ctx context.Context) (*CompanyHistory, error) {
	return withHooks[*CompanyHistory, CompanyHistoryMutation](ctx, chuo.sqlSave, chuo.mutation, chuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (chuo *CompanyHistoryUpdateOne) SaveX(ctx context.Context) *CompanyHistory {
	node, err := chuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (chuo *CompanyHistoryUpdateOne) Exec(ctx context.Context) error {
	_, err := chuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (chuo *CompanyHistoryUpdateOne) ExecX(ctx context.Context) {
	if err := chuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (chuo *CompanyHistoryUpdateOne) sqlSave(ctx context.Context) (_node *CompanyHistory, err error) {
	_spec := sqlgraph.NewUpdateSpec(companyhistory.Table, companyhistory.Columns, sqlgraph.NewFieldSpec(companyhistory.FieldID, field.TypeUUID))
	id, ok := chuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "CompanyHistory.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := chuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, companyhistory.FieldID)
		for _, f := range fields {
			if !companyhistory.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != companyhistory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := chuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if chuo.mutation.RefCleared() {
		_spec.ClearField(companyhistory.FieldRef, field.TypeUUID)
	}
	if chuo.mutation.UpdatedByCleared() {
		_spec.ClearField(companyhistory.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := chuo.mutation.Name(); ok {
		_spec.SetField(companyhistory.FieldName, field.TypeString, value)
	}
	_node = &CompanyHistory{config: chuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, chuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{companyhistory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	chuo.mutation.done = true
	return _node, nil
}
