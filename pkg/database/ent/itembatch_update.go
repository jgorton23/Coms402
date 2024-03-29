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

	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent/company"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent/itembatch"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent/predicate"
)

// ItemBatchUpdate is the builder for updating ItemBatch entities.
type ItemBatchUpdate struct {
	config
	hooks    []Hook
	mutation *ItemBatchMutation
}

// Where appends a list predicates to the ItemBatchUpdate builder.
func (ibu *ItemBatchUpdate) Where(ps ...predicate.ItemBatch) *ItemBatchUpdate {
	ibu.mutation.Where(ps...)
	return ibu
}

// SetItemNumber sets the "itemNumber" field.
func (ibu *ItemBatchUpdate) SetItemNumber(s string) *ItemBatchUpdate {
	ibu.mutation.SetItemNumber(s)
	return ibu
}

// SetDescription sets the "description" field.
func (ibu *ItemBatchUpdate) SetDescription(s string) *ItemBatchUpdate {
	ibu.mutation.SetDescription(s)
	return ibu
}

// SetCompanyUUID sets the "companyUUID" field.
func (ibu *ItemBatchUpdate) SetCompanyUUID(u uuid.UUID) *ItemBatchUpdate {
	ibu.mutation.SetCompanyUUID(u)
	return ibu
}

// SetCompanyID sets the "company" edge to the Company entity by ID.
func (ibu *ItemBatchUpdate) SetCompanyID(id uuid.UUID) *ItemBatchUpdate {
	ibu.mutation.SetCompanyID(id)
	return ibu
}

// SetCompany sets the "company" edge to the Company entity.
func (ibu *ItemBatchUpdate) SetCompany(c *Company) *ItemBatchUpdate {
	return ibu.SetCompanyID(c.ID)
}

// Mutation returns the ItemBatchMutation object of the builder.
func (ibu *ItemBatchUpdate) Mutation() *ItemBatchMutation {
	return ibu.mutation
}

// ClearCompany clears the "company" edge to the Company entity.
func (ibu *ItemBatchUpdate) ClearCompany() *ItemBatchUpdate {
	ibu.mutation.ClearCompany()
	return ibu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ibu *ItemBatchUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, ItemBatchMutation](ctx, ibu.sqlSave, ibu.mutation, ibu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ibu *ItemBatchUpdate) SaveX(ctx context.Context) int {
	affected, err := ibu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ibu *ItemBatchUpdate) Exec(ctx context.Context) error {
	_, err := ibu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ibu *ItemBatchUpdate) ExecX(ctx context.Context) {
	if err := ibu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ibu *ItemBatchUpdate) check() error {
	if _, ok := ibu.mutation.CompanyID(); ibu.mutation.CompanyCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ItemBatch.company"`)
	}
	return nil
}

func (ibu *ItemBatchUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ibu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(itembatch.Table, itembatch.Columns, sqlgraph.NewFieldSpec(itembatch.FieldID, field.TypeUUID))
	if ps := ibu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ibu.mutation.ItemNumber(); ok {
		_spec.SetField(itembatch.FieldItemNumber, field.TypeString, value)
	}
	if value, ok := ibu.mutation.Description(); ok {
		_spec.SetField(itembatch.FieldDescription, field.TypeString, value)
	}
	if ibu.mutation.CompanyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   itembatch.CompanyTable,
			Columns: []string{itembatch.CompanyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(company.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ibu.mutation.CompanyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   itembatch.CompanyTable,
			Columns: []string{itembatch.CompanyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(company.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ibu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{itembatch.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ibu.mutation.done = true
	return n, nil
}

// ItemBatchUpdateOne is the builder for updating a single ItemBatch entity.
type ItemBatchUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ItemBatchMutation
}

// SetItemNumber sets the "itemNumber" field.
func (ibuo *ItemBatchUpdateOne) SetItemNumber(s string) *ItemBatchUpdateOne {
	ibuo.mutation.SetItemNumber(s)
	return ibuo
}

// SetDescription sets the "description" field.
func (ibuo *ItemBatchUpdateOne) SetDescription(s string) *ItemBatchUpdateOne {
	ibuo.mutation.SetDescription(s)
	return ibuo
}

// SetCompanyUUID sets the "companyUUID" field.
func (ibuo *ItemBatchUpdateOne) SetCompanyUUID(u uuid.UUID) *ItemBatchUpdateOne {
	ibuo.mutation.SetCompanyUUID(u)
	return ibuo
}

// SetCompanyID sets the "company" edge to the Company entity by ID.
func (ibuo *ItemBatchUpdateOne) SetCompanyID(id uuid.UUID) *ItemBatchUpdateOne {
	ibuo.mutation.SetCompanyID(id)
	return ibuo
}

// SetCompany sets the "company" edge to the Company entity.
func (ibuo *ItemBatchUpdateOne) SetCompany(c *Company) *ItemBatchUpdateOne {
	return ibuo.SetCompanyID(c.ID)
}

// Mutation returns the ItemBatchMutation object of the builder.
func (ibuo *ItemBatchUpdateOne) Mutation() *ItemBatchMutation {
	return ibuo.mutation
}

// ClearCompany clears the "company" edge to the Company entity.
func (ibuo *ItemBatchUpdateOne) ClearCompany() *ItemBatchUpdateOne {
	ibuo.mutation.ClearCompany()
	return ibuo
}

// Where appends a list predicates to the ItemBatchUpdate builder.
func (ibuo *ItemBatchUpdateOne) Where(ps ...predicate.ItemBatch) *ItemBatchUpdateOne {
	ibuo.mutation.Where(ps...)
	return ibuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ibuo *ItemBatchUpdateOne) Select(field string, fields ...string) *ItemBatchUpdateOne {
	ibuo.fields = append([]string{field}, fields...)
	return ibuo
}

// Save executes the query and returns the updated ItemBatch entity.
func (ibuo *ItemBatchUpdateOne) Save(ctx context.Context) (*ItemBatch, error) {
	return withHooks[*ItemBatch, ItemBatchMutation](ctx, ibuo.sqlSave, ibuo.mutation, ibuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ibuo *ItemBatchUpdateOne) SaveX(ctx context.Context) *ItemBatch {
	node, err := ibuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ibuo *ItemBatchUpdateOne) Exec(ctx context.Context) error {
	_, err := ibuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ibuo *ItemBatchUpdateOne) ExecX(ctx context.Context) {
	if err := ibuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ibuo *ItemBatchUpdateOne) check() error {
	if _, ok := ibuo.mutation.CompanyID(); ibuo.mutation.CompanyCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ItemBatch.company"`)
	}
	return nil
}

func (ibuo *ItemBatchUpdateOne) sqlSave(ctx context.Context) (_node *ItemBatch, err error) {
	if err := ibuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(itembatch.Table, itembatch.Columns, sqlgraph.NewFieldSpec(itembatch.FieldID, field.TypeUUID))
	id, ok := ibuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ItemBatch.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ibuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, itembatch.FieldID)
		for _, f := range fields {
			if !itembatch.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != itembatch.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ibuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ibuo.mutation.ItemNumber(); ok {
		_spec.SetField(itembatch.FieldItemNumber, field.TypeString, value)
	}
	if value, ok := ibuo.mutation.Description(); ok {
		_spec.SetField(itembatch.FieldDescription, field.TypeString, value)
	}
	if ibuo.mutation.CompanyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   itembatch.CompanyTable,
			Columns: []string{itembatch.CompanyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(company.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ibuo.mutation.CompanyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   itembatch.CompanyTable,
			Columns: []string{itembatch.CompanyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(company.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &ItemBatch{config: ibuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ibuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{itembatch.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ibuo.mutation.done = true
	return _node, nil
}
