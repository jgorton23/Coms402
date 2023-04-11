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

	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent/certificationtemplate"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent/company"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent/predicate"
)

// CertificationTemplateUpdate is the builder for updating CertificationTemplate entities.
type CertificationTemplateUpdate struct {
	config
	hooks    []Hook
	mutation *CertificationTemplateMutation
}

// Where appends a list predicates to the CertificationTemplateUpdate builder.
func (ctu *CertificationTemplateUpdate) Where(ps ...predicate.CertificationTemplate) *CertificationTemplateUpdate {
	ctu.mutation.Where(ps...)
	return ctu
}

// SetDescription sets the "description" field.
func (ctu *CertificationTemplateUpdate) SetDescription(s string) *CertificationTemplateUpdate {
	ctu.mutation.SetDescription(s)
	return ctu
}

// SetCompanyUUID sets the "companyUUID" field.
func (ctu *CertificationTemplateUpdate) SetCompanyUUID(u uuid.UUID) *CertificationTemplateUpdate {
	ctu.mutation.SetCompanyUUID(u)
	return ctu
}

// SetCompanyID sets the "company" edge to the Company entity by ID.
func (ctu *CertificationTemplateUpdate) SetCompanyID(id uuid.UUID) *CertificationTemplateUpdate {
	ctu.mutation.SetCompanyID(id)
	return ctu
}

// SetCompany sets the "company" edge to the Company entity.
func (ctu *CertificationTemplateUpdate) SetCompany(c *Company) *CertificationTemplateUpdate {
	return ctu.SetCompanyID(c.ID)
}

// Mutation returns the CertificationTemplateMutation object of the builder.
func (ctu *CertificationTemplateUpdate) Mutation() *CertificationTemplateMutation {
	return ctu.mutation
}

// ClearCompany clears the "company" edge to the Company entity.
func (ctu *CertificationTemplateUpdate) ClearCompany() *CertificationTemplateUpdate {
	ctu.mutation.ClearCompany()
	return ctu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ctu *CertificationTemplateUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, CertificationTemplateMutation](ctx, ctu.sqlSave, ctu.mutation, ctu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ctu *CertificationTemplateUpdate) SaveX(ctx context.Context) int {
	affected, err := ctu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ctu *CertificationTemplateUpdate) Exec(ctx context.Context) error {
	_, err := ctu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ctu *CertificationTemplateUpdate) ExecX(ctx context.Context) {
	if err := ctu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ctu *CertificationTemplateUpdate) check() error {
	if _, ok := ctu.mutation.CompanyID(); ctu.mutation.CompanyCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "CertificationTemplate.company"`)
	}
	return nil
}

func (ctu *CertificationTemplateUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ctu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(certificationtemplate.Table, certificationtemplate.Columns, sqlgraph.NewFieldSpec(certificationtemplate.FieldID, field.TypeUUID))
	if ps := ctu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ctu.mutation.Description(); ok {
		_spec.SetField(certificationtemplate.FieldDescription, field.TypeString, value)
	}
	if ctu.mutation.CompanyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   certificationtemplate.CompanyTable,
			Columns: []string{certificationtemplate.CompanyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(company.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ctu.mutation.CompanyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   certificationtemplate.CompanyTable,
			Columns: []string{certificationtemplate.CompanyColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, ctu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{certificationtemplate.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ctu.mutation.done = true
	return n, nil
}

// CertificationTemplateUpdateOne is the builder for updating a single CertificationTemplate entity.
type CertificationTemplateUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CertificationTemplateMutation
}

// SetDescription sets the "description" field.
func (ctuo *CertificationTemplateUpdateOne) SetDescription(s string) *CertificationTemplateUpdateOne {
	ctuo.mutation.SetDescription(s)
	return ctuo
}

// SetCompanyUUID sets the "companyUUID" field.
func (ctuo *CertificationTemplateUpdateOne) SetCompanyUUID(u uuid.UUID) *CertificationTemplateUpdateOne {
	ctuo.mutation.SetCompanyUUID(u)
	return ctuo
}

// SetCompanyID sets the "company" edge to the Company entity by ID.
func (ctuo *CertificationTemplateUpdateOne) SetCompanyID(id uuid.UUID) *CertificationTemplateUpdateOne {
	ctuo.mutation.SetCompanyID(id)
	return ctuo
}

// SetCompany sets the "company" edge to the Company entity.
func (ctuo *CertificationTemplateUpdateOne) SetCompany(c *Company) *CertificationTemplateUpdateOne {
	return ctuo.SetCompanyID(c.ID)
}

// Mutation returns the CertificationTemplateMutation object of the builder.
func (ctuo *CertificationTemplateUpdateOne) Mutation() *CertificationTemplateMutation {
	return ctuo.mutation
}

// ClearCompany clears the "company" edge to the Company entity.
func (ctuo *CertificationTemplateUpdateOne) ClearCompany() *CertificationTemplateUpdateOne {
	ctuo.mutation.ClearCompany()
	return ctuo
}

// Where appends a list predicates to the CertificationTemplateUpdate builder.
func (ctuo *CertificationTemplateUpdateOne) Where(ps ...predicate.CertificationTemplate) *CertificationTemplateUpdateOne {
	ctuo.mutation.Where(ps...)
	return ctuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ctuo *CertificationTemplateUpdateOne) Select(field string, fields ...string) *CertificationTemplateUpdateOne {
	ctuo.fields = append([]string{field}, fields...)
	return ctuo
}

// Save executes the query and returns the updated CertificationTemplate entity.
func (ctuo *CertificationTemplateUpdateOne) Save(ctx context.Context) (*CertificationTemplate, error) {
	return withHooks[*CertificationTemplate, CertificationTemplateMutation](ctx, ctuo.sqlSave, ctuo.mutation, ctuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ctuo *CertificationTemplateUpdateOne) SaveX(ctx context.Context) *CertificationTemplate {
	node, err := ctuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ctuo *CertificationTemplateUpdateOne) Exec(ctx context.Context) error {
	_, err := ctuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ctuo *CertificationTemplateUpdateOne) ExecX(ctx context.Context) {
	if err := ctuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ctuo *CertificationTemplateUpdateOne) check() error {
	if _, ok := ctuo.mutation.CompanyID(); ctuo.mutation.CompanyCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "CertificationTemplate.company"`)
	}
	return nil
}

func (ctuo *CertificationTemplateUpdateOne) sqlSave(ctx context.Context) (_node *CertificationTemplate, err error) {
	if err := ctuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(certificationtemplate.Table, certificationtemplate.Columns, sqlgraph.NewFieldSpec(certificationtemplate.FieldID, field.TypeUUID))
	id, ok := ctuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "CertificationTemplate.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ctuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, certificationtemplate.FieldID)
		for _, f := range fields {
			if !certificationtemplate.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != certificationtemplate.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ctuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ctuo.mutation.Description(); ok {
		_spec.SetField(certificationtemplate.FieldDescription, field.TypeString, value)
	}
	if ctuo.mutation.CompanyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   certificationtemplate.CompanyTable,
			Columns: []string{certificationtemplate.CompanyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(company.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ctuo.mutation.CompanyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   certificationtemplate.CompanyTable,
			Columns: []string{certificationtemplate.CompanyColumn},
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
	_node = &CertificationTemplate{config: ctuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ctuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{certificationtemplate.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ctuo.mutation.done = true
	return _node, nil
}
