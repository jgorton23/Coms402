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
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent/predicate"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent/user"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent/userstocompany"
)

// UsersToCompanyUpdate is the builder for updating UsersToCompany entities.
type UsersToCompanyUpdate struct {
	config
	hooks    []Hook
	mutation *UsersToCompanyMutation
}

// Where appends a list predicates to the UsersToCompanyUpdate builder.
func (utcu *UsersToCompanyUpdate) Where(ps ...predicate.UsersToCompany) *UsersToCompanyUpdate {
	utcu.mutation.Where(ps...)
	return utcu
}

// SetCompanyUUID sets the "companyUUID" field.
func (utcu *UsersToCompanyUpdate) SetCompanyUUID(u uuid.UUID) *UsersToCompanyUpdate {
	utcu.mutation.SetCompanyUUID(u)
	return utcu
}

// SetUserUUID sets the "userUUID" field.
func (utcu *UsersToCompanyUpdate) SetUserUUID(i int) *UsersToCompanyUpdate {
	utcu.mutation.SetUserUUID(i)
	return utcu
}

// SetRoleType sets the "roleType" field.
func (utcu *UsersToCompanyUpdate) SetRoleType(s string) *UsersToCompanyUpdate {
	utcu.mutation.SetRoleType(s)
	return utcu
}

// SetApproved sets the "approved" field.
func (utcu *UsersToCompanyUpdate) SetApproved(b bool) *UsersToCompanyUpdate {
	utcu.mutation.SetApproved(b)
	return utcu
}

// SetUserID sets the "user" edge to the User entity by ID.
func (utcu *UsersToCompanyUpdate) SetUserID(id int) *UsersToCompanyUpdate {
	utcu.mutation.SetUserID(id)
	return utcu
}

// SetUser sets the "user" edge to the User entity.
func (utcu *UsersToCompanyUpdate) SetUser(u *User) *UsersToCompanyUpdate {
	return utcu.SetUserID(u.ID)
}

// SetCompanyID sets the "company" edge to the Company entity by ID.
func (utcu *UsersToCompanyUpdate) SetCompanyID(id uuid.UUID) *UsersToCompanyUpdate {
	utcu.mutation.SetCompanyID(id)
	return utcu
}

// SetCompany sets the "company" edge to the Company entity.
func (utcu *UsersToCompanyUpdate) SetCompany(c *Company) *UsersToCompanyUpdate {
	return utcu.SetCompanyID(c.ID)
}

// Mutation returns the UsersToCompanyMutation object of the builder.
func (utcu *UsersToCompanyUpdate) Mutation() *UsersToCompanyMutation {
	return utcu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (utcu *UsersToCompanyUpdate) ClearUser() *UsersToCompanyUpdate {
	utcu.mutation.ClearUser()
	return utcu
}

// ClearCompany clears the "company" edge to the Company entity.
func (utcu *UsersToCompanyUpdate) ClearCompany() *UsersToCompanyUpdate {
	utcu.mutation.ClearCompany()
	return utcu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (utcu *UsersToCompanyUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, UsersToCompanyMutation](ctx, utcu.sqlSave, utcu.mutation, utcu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (utcu *UsersToCompanyUpdate) SaveX(ctx context.Context) int {
	affected, err := utcu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (utcu *UsersToCompanyUpdate) Exec(ctx context.Context) error {
	_, err := utcu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (utcu *UsersToCompanyUpdate) ExecX(ctx context.Context) {
	if err := utcu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (utcu *UsersToCompanyUpdate) check() error {
	if _, ok := utcu.mutation.UserID(); utcu.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "UsersToCompany.user"`)
	}
	if _, ok := utcu.mutation.CompanyID(); utcu.mutation.CompanyCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "UsersToCompany.company"`)
	}
	return nil
}

func (utcu *UsersToCompanyUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := utcu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(userstocompany.Table, userstocompany.Columns, sqlgraph.NewFieldSpec(userstocompany.FieldID, field.TypeUUID))
	if ps := utcu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := utcu.mutation.RoleType(); ok {
		_spec.SetField(userstocompany.FieldRoleType, field.TypeString, value)
	}
	if value, ok := utcu.mutation.Approved(); ok {
		_spec.SetField(userstocompany.FieldApproved, field.TypeBool, value)
	}
	if utcu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   userstocompany.UserTable,
			Columns: []string{userstocompany.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := utcu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   userstocompany.UserTable,
			Columns: []string{userstocompany.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if utcu.mutation.CompanyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   userstocompany.CompanyTable,
			Columns: []string{userstocompany.CompanyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: company.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := utcu.mutation.CompanyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   userstocompany.CompanyTable,
			Columns: []string{userstocompany.CompanyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: company.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, utcu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userstocompany.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	utcu.mutation.done = true
	return n, nil
}

// UsersToCompanyUpdateOne is the builder for updating a single UsersToCompany entity.
type UsersToCompanyUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UsersToCompanyMutation
}

// SetCompanyUUID sets the "companyUUID" field.
func (utcuo *UsersToCompanyUpdateOne) SetCompanyUUID(u uuid.UUID) *UsersToCompanyUpdateOne {
	utcuo.mutation.SetCompanyUUID(u)
	return utcuo
}

// SetUserUUID sets the "userUUID" field.
func (utcuo *UsersToCompanyUpdateOne) SetUserUUID(i int) *UsersToCompanyUpdateOne {
	utcuo.mutation.SetUserUUID(i)
	return utcuo
}

// SetRoleType sets the "roleType" field.
func (utcuo *UsersToCompanyUpdateOne) SetRoleType(s string) *UsersToCompanyUpdateOne {
	utcuo.mutation.SetRoleType(s)
	return utcuo
}

// SetApproved sets the "approved" field.
func (utcuo *UsersToCompanyUpdateOne) SetApproved(b bool) *UsersToCompanyUpdateOne {
	utcuo.mutation.SetApproved(b)
	return utcuo
}

// SetUserID sets the "user" edge to the User entity by ID.
func (utcuo *UsersToCompanyUpdateOne) SetUserID(id int) *UsersToCompanyUpdateOne {
	utcuo.mutation.SetUserID(id)
	return utcuo
}

// SetUser sets the "user" edge to the User entity.
func (utcuo *UsersToCompanyUpdateOne) SetUser(u *User) *UsersToCompanyUpdateOne {
	return utcuo.SetUserID(u.ID)
}

// SetCompanyID sets the "company" edge to the Company entity by ID.
func (utcuo *UsersToCompanyUpdateOne) SetCompanyID(id uuid.UUID) *UsersToCompanyUpdateOne {
	utcuo.mutation.SetCompanyID(id)
	return utcuo
}

// SetCompany sets the "company" edge to the Company entity.
func (utcuo *UsersToCompanyUpdateOne) SetCompany(c *Company) *UsersToCompanyUpdateOne {
	return utcuo.SetCompanyID(c.ID)
}

// Mutation returns the UsersToCompanyMutation object of the builder.
func (utcuo *UsersToCompanyUpdateOne) Mutation() *UsersToCompanyMutation {
	return utcuo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (utcuo *UsersToCompanyUpdateOne) ClearUser() *UsersToCompanyUpdateOne {
	utcuo.mutation.ClearUser()
	return utcuo
}

// ClearCompany clears the "company" edge to the Company entity.
func (utcuo *UsersToCompanyUpdateOne) ClearCompany() *UsersToCompanyUpdateOne {
	utcuo.mutation.ClearCompany()
	return utcuo
}

// Where appends a list predicates to the UsersToCompanyUpdate builder.
func (utcuo *UsersToCompanyUpdateOne) Where(ps ...predicate.UsersToCompany) *UsersToCompanyUpdateOne {
	utcuo.mutation.Where(ps...)
	return utcuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (utcuo *UsersToCompanyUpdateOne) Select(field string, fields ...string) *UsersToCompanyUpdateOne {
	utcuo.fields = append([]string{field}, fields...)
	return utcuo
}

// Save executes the query and returns the updated UsersToCompany entity.
func (utcuo *UsersToCompanyUpdateOne) Save(ctx context.Context) (*UsersToCompany, error) {
	return withHooks[*UsersToCompany, UsersToCompanyMutation](ctx, utcuo.sqlSave, utcuo.mutation, utcuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (utcuo *UsersToCompanyUpdateOne) SaveX(ctx context.Context) *UsersToCompany {
	node, err := utcuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (utcuo *UsersToCompanyUpdateOne) Exec(ctx context.Context) error {
	_, err := utcuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (utcuo *UsersToCompanyUpdateOne) ExecX(ctx context.Context) {
	if err := utcuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (utcuo *UsersToCompanyUpdateOne) check() error {
	if _, ok := utcuo.mutation.UserID(); utcuo.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "UsersToCompany.user"`)
	}
	if _, ok := utcuo.mutation.CompanyID(); utcuo.mutation.CompanyCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "UsersToCompany.company"`)
	}
	return nil
}

func (utcuo *UsersToCompanyUpdateOne) sqlSave(ctx context.Context) (_node *UsersToCompany, err error) {
	if err := utcuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(userstocompany.Table, userstocompany.Columns, sqlgraph.NewFieldSpec(userstocompany.FieldID, field.TypeUUID))
	id, ok := utcuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "UsersToCompany.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := utcuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, userstocompany.FieldID)
		for _, f := range fields {
			if !userstocompany.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != userstocompany.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := utcuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := utcuo.mutation.RoleType(); ok {
		_spec.SetField(userstocompany.FieldRoleType, field.TypeString, value)
	}
	if value, ok := utcuo.mutation.Approved(); ok {
		_spec.SetField(userstocompany.FieldApproved, field.TypeBool, value)
	}
	if utcuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   userstocompany.UserTable,
			Columns: []string{userstocompany.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := utcuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   userstocompany.UserTable,
			Columns: []string{userstocompany.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if utcuo.mutation.CompanyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   userstocompany.CompanyTable,
			Columns: []string{userstocompany.CompanyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: company.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := utcuo.mutation.CompanyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   userstocompany.CompanyTable,
			Columns: []string{userstocompany.CompanyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: company.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &UsersToCompany{config: utcuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, utcuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userstocompany.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	utcuo.mutation.done = true
	return _node, nil
}
