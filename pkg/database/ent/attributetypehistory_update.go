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

	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent/attributetypehistory"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent/predicate"
)

// AttributeTypeHistoryUpdate is the builder for updating AttributeTypeHistory entities.
type AttributeTypeHistoryUpdate struct {
	config
	hooks    []Hook
	mutation *AttributeTypeHistoryMutation
}

// Where appends a list predicates to the AttributeTypeHistoryUpdate builder.
func (athu *AttributeTypeHistoryUpdate) Where(ps ...predicate.AttributeTypeHistory) *AttributeTypeHistoryUpdate {
	athu.mutation.Where(ps...)
	return athu
}

// SetKey sets the "key" field.
func (athu *AttributeTypeHistoryUpdate) SetKey(s string) *AttributeTypeHistoryUpdate {
	athu.mutation.SetKey(s)
	return athu
}

// SetCompanyUUID sets the "companyUUID" field.
func (athu *AttributeTypeHistoryUpdate) SetCompanyUUID(u uuid.UUID) *AttributeTypeHistoryUpdate {
	athu.mutation.SetCompanyUUID(u)
	return athu
}

// Mutation returns the AttributeTypeHistoryMutation object of the builder.
func (athu *AttributeTypeHistoryUpdate) Mutation() *AttributeTypeHistoryMutation {
	return athu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (athu *AttributeTypeHistoryUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, AttributeTypeHistoryMutation](ctx, athu.sqlSave, athu.mutation, athu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (athu *AttributeTypeHistoryUpdate) SaveX(ctx context.Context) int {
	affected, err := athu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (athu *AttributeTypeHistoryUpdate) Exec(ctx context.Context) error {
	_, err := athu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (athu *AttributeTypeHistoryUpdate) ExecX(ctx context.Context) {
	if err := athu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (athu *AttributeTypeHistoryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(attributetypehistory.Table, attributetypehistory.Columns, sqlgraph.NewFieldSpec(attributetypehistory.FieldID, field.TypeUUID))
	if ps := athu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if athu.mutation.RefCleared() {
		_spec.ClearField(attributetypehistory.FieldRef, field.TypeUUID)
	}
	if athu.mutation.UpdatedByCleared() {
		_spec.ClearField(attributetypehistory.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := athu.mutation.Key(); ok {
		_spec.SetField(attributetypehistory.FieldKey, field.TypeString, value)
	}
	if value, ok := athu.mutation.CompanyUUID(); ok {
		_spec.SetField(attributetypehistory.FieldCompanyUUID, field.TypeUUID, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, athu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{attributetypehistory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	athu.mutation.done = true
	return n, nil
}

// AttributeTypeHistoryUpdateOne is the builder for updating a single AttributeTypeHistory entity.
type AttributeTypeHistoryUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AttributeTypeHistoryMutation
}

// SetKey sets the "key" field.
func (athuo *AttributeTypeHistoryUpdateOne) SetKey(s string) *AttributeTypeHistoryUpdateOne {
	athuo.mutation.SetKey(s)
	return athuo
}

// SetCompanyUUID sets the "companyUUID" field.
func (athuo *AttributeTypeHistoryUpdateOne) SetCompanyUUID(u uuid.UUID) *AttributeTypeHistoryUpdateOne {
	athuo.mutation.SetCompanyUUID(u)
	return athuo
}

// Mutation returns the AttributeTypeHistoryMutation object of the builder.
func (athuo *AttributeTypeHistoryUpdateOne) Mutation() *AttributeTypeHistoryMutation {
	return athuo.mutation
}

// Where appends a list predicates to the AttributeTypeHistoryUpdate builder.
func (athuo *AttributeTypeHistoryUpdateOne) Where(ps ...predicate.AttributeTypeHistory) *AttributeTypeHistoryUpdateOne {
	athuo.mutation.Where(ps...)
	return athuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (athuo *AttributeTypeHistoryUpdateOne) Select(field string, fields ...string) *AttributeTypeHistoryUpdateOne {
	athuo.fields = append([]string{field}, fields...)
	return athuo
}

// Save executes the query and returns the updated AttributeTypeHistory entity.
func (athuo *AttributeTypeHistoryUpdateOne) Save(ctx context.Context) (*AttributeTypeHistory, error) {
	return withHooks[*AttributeTypeHistory, AttributeTypeHistoryMutation](ctx, athuo.sqlSave, athuo.mutation, athuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (athuo *AttributeTypeHistoryUpdateOne) SaveX(ctx context.Context) *AttributeTypeHistory {
	node, err := athuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (athuo *AttributeTypeHistoryUpdateOne) Exec(ctx context.Context) error {
	_, err := athuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (athuo *AttributeTypeHistoryUpdateOne) ExecX(ctx context.Context) {
	if err := athuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (athuo *AttributeTypeHistoryUpdateOne) sqlSave(ctx context.Context) (_node *AttributeTypeHistory, err error) {
	_spec := sqlgraph.NewUpdateSpec(attributetypehistory.Table, attributetypehistory.Columns, sqlgraph.NewFieldSpec(attributetypehistory.FieldID, field.TypeUUID))
	id, ok := athuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "AttributeTypeHistory.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := athuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, attributetypehistory.FieldID)
		for _, f := range fields {
			if !attributetypehistory.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != attributetypehistory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := athuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if athuo.mutation.RefCleared() {
		_spec.ClearField(attributetypehistory.FieldRef, field.TypeUUID)
	}
	if athuo.mutation.UpdatedByCleared() {
		_spec.ClearField(attributetypehistory.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := athuo.mutation.Key(); ok {
		_spec.SetField(attributetypehistory.FieldKey, field.TypeString, value)
	}
	if value, ok := athuo.mutation.CompanyUUID(); ok {
		_spec.SetField(attributetypehistory.FieldCompanyUUID, field.TypeUUID, value)
	}
	_node = &AttributeTypeHistory{config: athuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, athuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{attributetypehistory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	athuo.mutation.done = true
	return _node, nil
}
