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

	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent/predicate"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent/userstocompanyhistory"
)

// UsersToCompanyHistoryUpdate is the builder for updating UsersToCompanyHistory entities.
type UsersToCompanyHistoryUpdate struct {
	config
	hooks    []Hook
	mutation *UsersToCompanyHistoryMutation
}

// Where appends a list predicates to the UsersToCompanyHistoryUpdate builder.
func (utchu *UsersToCompanyHistoryUpdate) Where(ps ...predicate.UsersToCompanyHistory) *UsersToCompanyHistoryUpdate {
	utchu.mutation.Where(ps...)
	return utchu
}

// SetCompanyUUID sets the "companyUUID" field.
func (utchu *UsersToCompanyHistoryUpdate) SetCompanyUUID(u uuid.UUID) *UsersToCompanyHistoryUpdate {
	utchu.mutation.SetCompanyUUID(u)
	return utchu
}

// SetUserUUID sets the "userUUID" field.
func (utchu *UsersToCompanyHistoryUpdate) SetUserUUID(u uuid.UUID) *UsersToCompanyHistoryUpdate {
	utchu.mutation.SetUserUUID(u)
	return utchu
}

// SetRoleType sets the "roleType" field.
func (utchu *UsersToCompanyHistoryUpdate) SetRoleType(s string) *UsersToCompanyHistoryUpdate {
	utchu.mutation.SetRoleType(s)
	return utchu
}

// SetApproved sets the "approved" field.
func (utchu *UsersToCompanyHistoryUpdate) SetApproved(b bool) *UsersToCompanyHistoryUpdate {
	utchu.mutation.SetApproved(b)
	return utchu
}

// Mutation returns the UsersToCompanyHistoryMutation object of the builder.
func (utchu *UsersToCompanyHistoryUpdate) Mutation() *UsersToCompanyHistoryMutation {
	return utchu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (utchu *UsersToCompanyHistoryUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, UsersToCompanyHistoryMutation](ctx, utchu.sqlSave, utchu.mutation, utchu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (utchu *UsersToCompanyHistoryUpdate) SaveX(ctx context.Context) int {
	affected, err := utchu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (utchu *UsersToCompanyHistoryUpdate) Exec(ctx context.Context) error {
	_, err := utchu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (utchu *UsersToCompanyHistoryUpdate) ExecX(ctx context.Context) {
	if err := utchu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (utchu *UsersToCompanyHistoryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(userstocompanyhistory.Table, userstocompanyhistory.Columns, sqlgraph.NewFieldSpec(userstocompanyhistory.FieldID, field.TypeUUID))
	if ps := utchu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if utchu.mutation.RefCleared() {
		_spec.ClearField(userstocompanyhistory.FieldRef, field.TypeUUID)
	}
	if utchu.mutation.UpdatedByCleared() {
		_spec.ClearField(userstocompanyhistory.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := utchu.mutation.CompanyUUID(); ok {
		_spec.SetField(userstocompanyhistory.FieldCompanyUUID, field.TypeUUID, value)
	}
	if value, ok := utchu.mutation.UserUUID(); ok {
		_spec.SetField(userstocompanyhistory.FieldUserUUID, field.TypeUUID, value)
	}
	if value, ok := utchu.mutation.RoleType(); ok {
		_spec.SetField(userstocompanyhistory.FieldRoleType, field.TypeString, value)
	}
	if value, ok := utchu.mutation.Approved(); ok {
		_spec.SetField(userstocompanyhistory.FieldApproved, field.TypeBool, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, utchu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userstocompanyhistory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	utchu.mutation.done = true
	return n, nil
}

// UsersToCompanyHistoryUpdateOne is the builder for updating a single UsersToCompanyHistory entity.
type UsersToCompanyHistoryUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UsersToCompanyHistoryMutation
}

// SetCompanyUUID sets the "companyUUID" field.
func (utchuo *UsersToCompanyHistoryUpdateOne) SetCompanyUUID(u uuid.UUID) *UsersToCompanyHistoryUpdateOne {
	utchuo.mutation.SetCompanyUUID(u)
	return utchuo
}

// SetUserUUID sets the "userUUID" field.
func (utchuo *UsersToCompanyHistoryUpdateOne) SetUserUUID(u uuid.UUID) *UsersToCompanyHistoryUpdateOne {
	utchuo.mutation.SetUserUUID(u)
	return utchuo
}

// SetRoleType sets the "roleType" field.
func (utchuo *UsersToCompanyHistoryUpdateOne) SetRoleType(s string) *UsersToCompanyHistoryUpdateOne {
	utchuo.mutation.SetRoleType(s)
	return utchuo
}

// SetApproved sets the "approved" field.
func (utchuo *UsersToCompanyHistoryUpdateOne) SetApproved(b bool) *UsersToCompanyHistoryUpdateOne {
	utchuo.mutation.SetApproved(b)
	return utchuo
}

// Mutation returns the UsersToCompanyHistoryMutation object of the builder.
func (utchuo *UsersToCompanyHistoryUpdateOne) Mutation() *UsersToCompanyHistoryMutation {
	return utchuo.mutation
}

// Where appends a list predicates to the UsersToCompanyHistoryUpdate builder.
func (utchuo *UsersToCompanyHistoryUpdateOne) Where(ps ...predicate.UsersToCompanyHistory) *UsersToCompanyHistoryUpdateOne {
	utchuo.mutation.Where(ps...)
	return utchuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (utchuo *UsersToCompanyHistoryUpdateOne) Select(field string, fields ...string) *UsersToCompanyHistoryUpdateOne {
	utchuo.fields = append([]string{field}, fields...)
	return utchuo
}

// Save executes the query and returns the updated UsersToCompanyHistory entity.
func (utchuo *UsersToCompanyHistoryUpdateOne) Save(ctx context.Context) (*UsersToCompanyHistory, error) {
	return withHooks[*UsersToCompanyHistory, UsersToCompanyHistoryMutation](ctx, utchuo.sqlSave, utchuo.mutation, utchuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (utchuo *UsersToCompanyHistoryUpdateOne) SaveX(ctx context.Context) *UsersToCompanyHistory {
	node, err := utchuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (utchuo *UsersToCompanyHistoryUpdateOne) Exec(ctx context.Context) error {
	_, err := utchuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (utchuo *UsersToCompanyHistoryUpdateOne) ExecX(ctx context.Context) {
	if err := utchuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (utchuo *UsersToCompanyHistoryUpdateOne) sqlSave(ctx context.Context) (_node *UsersToCompanyHistory, err error) {
	_spec := sqlgraph.NewUpdateSpec(userstocompanyhistory.Table, userstocompanyhistory.Columns, sqlgraph.NewFieldSpec(userstocompanyhistory.FieldID, field.TypeUUID))
	id, ok := utchuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "UsersToCompanyHistory.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := utchuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, userstocompanyhistory.FieldID)
		for _, f := range fields {
			if !userstocompanyhistory.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != userstocompanyhistory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := utchuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if utchuo.mutation.RefCleared() {
		_spec.ClearField(userstocompanyhistory.FieldRef, field.TypeUUID)
	}
	if utchuo.mutation.UpdatedByCleared() {
		_spec.ClearField(userstocompanyhistory.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := utchuo.mutation.CompanyUUID(); ok {
		_spec.SetField(userstocompanyhistory.FieldCompanyUUID, field.TypeUUID, value)
	}
	if value, ok := utchuo.mutation.UserUUID(); ok {
		_spec.SetField(userstocompanyhistory.FieldUserUUID, field.TypeUUID, value)
	}
	if value, ok := utchuo.mutation.RoleType(); ok {
		_spec.SetField(userstocompanyhistory.FieldRoleType, field.TypeString, value)
	}
	if value, ok := utchuo.mutation.Approved(); ok {
		_spec.SetField(userstocompanyhistory.FieldApproved, field.TypeBool, value)
	}
	_node = &UsersToCompanyHistory{config: utchuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, utchuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userstocompanyhistory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	utchuo.mutation.done = true
	return _node, nil
}
