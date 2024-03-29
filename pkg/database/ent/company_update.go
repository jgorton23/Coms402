// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"

	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent/company"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent/predicate"
)

// CompanyUpdate is the builder for updating Company entities.
type CompanyUpdate struct {
	config
	hooks    []Hook
	mutation *CompanyMutation
}

// Where appends a list predicates to the CompanyUpdate builder.
func (cu *CompanyUpdate) Where(ps ...predicate.Company) *CompanyUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetName sets the "name" field.
func (cu *CompanyUpdate) SetName(s string) *CompanyUpdate {
	cu.mutation.SetName(s)
	return cu
}

// Mutation returns the CompanyMutation object of the builder.
func (cu *CompanyUpdate) Mutation() *CompanyMutation {
	return cu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CompanyUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, CompanyMutation](ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CompanyUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CompanyUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CompanyUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cu *CompanyUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(company.Table, company.Columns, sqlgraph.NewFieldSpec(company.FieldID, field.TypeUUID))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.Name(); ok {
		_spec.SetField(company.FieldName, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{company.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// CompanyUpdateOne is the builder for updating a single Company entity.
type CompanyUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CompanyMutation
}

// SetName sets the "name" field.
func (cuo *CompanyUpdateOne) SetName(s string) *CompanyUpdateOne {
	cuo.mutation.SetName(s)
	return cuo
}

// Mutation returns the CompanyMutation object of the builder.
func (cuo *CompanyUpdateOne) Mutation() *CompanyMutation {
	return cuo.mutation
}

// Where appends a list predicates to the CompanyUpdate builder.
func (cuo *CompanyUpdateOne) Where(ps ...predicate.Company) *CompanyUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CompanyUpdateOne) Select(field string, fields ...string) *CompanyUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Company entity.
func (cuo *CompanyUpdateOne) Save(ctx context.Context) (*Company, error) {
	return withHooks[*Company, CompanyMutation](ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CompanyUpdateOne) SaveX(ctx context.Context) *Company {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CompanyUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CompanyUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cuo *CompanyUpdateOne) sqlSave(ctx context.Context) (_node *Company, err error) {
	_spec := sqlgraph.NewUpdateSpec(company.Table, company.Columns, sqlgraph.NewFieldSpec(company.FieldID, field.TypeUUID))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Company.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, company.FieldID)
		for _, f := range fields {
			if !company.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != company.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.Name(); ok {
		_spec.SetField(company.FieldName, field.TypeString, value)
	}
	_node = &Company{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{company.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}
