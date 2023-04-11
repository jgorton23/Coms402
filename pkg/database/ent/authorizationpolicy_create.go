// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"

	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent/authorizationpolicy"
)

// AuthorizationPolicyCreate is the builder for creating a AuthorizationPolicy entity.
type AuthorizationPolicyCreate struct {
	config
	mutation *AuthorizationPolicyMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetPtype sets the "Ptype" field.
func (apc *AuthorizationPolicyCreate) SetPtype(s string) *AuthorizationPolicyCreate {
	apc.mutation.SetPtype(s)
	return apc
}

// SetNillablePtype sets the "Ptype" field if the given value is not nil.
func (apc *AuthorizationPolicyCreate) SetNillablePtype(s *string) *AuthorizationPolicyCreate {
	if s != nil {
		apc.SetPtype(*s)
	}
	return apc
}

// SetV0 sets the "V0" field.
func (apc *AuthorizationPolicyCreate) SetV0(s string) *AuthorizationPolicyCreate {
	apc.mutation.SetV0(s)
	return apc
}

// SetNillableV0 sets the "V0" field if the given value is not nil.
func (apc *AuthorizationPolicyCreate) SetNillableV0(s *string) *AuthorizationPolicyCreate {
	if s != nil {
		apc.SetV0(*s)
	}
	return apc
}

// SetV1 sets the "V1" field.
func (apc *AuthorizationPolicyCreate) SetV1(s string) *AuthorizationPolicyCreate {
	apc.mutation.SetV1(s)
	return apc
}

// SetNillableV1 sets the "V1" field if the given value is not nil.
func (apc *AuthorizationPolicyCreate) SetNillableV1(s *string) *AuthorizationPolicyCreate {
	if s != nil {
		apc.SetV1(*s)
	}
	return apc
}

// SetV2 sets the "V2" field.
func (apc *AuthorizationPolicyCreate) SetV2(s string) *AuthorizationPolicyCreate {
	apc.mutation.SetV2(s)
	return apc
}

// SetNillableV2 sets the "V2" field if the given value is not nil.
func (apc *AuthorizationPolicyCreate) SetNillableV2(s *string) *AuthorizationPolicyCreate {
	if s != nil {
		apc.SetV2(*s)
	}
	return apc
}

// SetV3 sets the "V3" field.
func (apc *AuthorizationPolicyCreate) SetV3(s string) *AuthorizationPolicyCreate {
	apc.mutation.SetV3(s)
	return apc
}

// SetNillableV3 sets the "V3" field if the given value is not nil.
func (apc *AuthorizationPolicyCreate) SetNillableV3(s *string) *AuthorizationPolicyCreate {
	if s != nil {
		apc.SetV3(*s)
	}
	return apc
}

// SetV4 sets the "V4" field.
func (apc *AuthorizationPolicyCreate) SetV4(s string) *AuthorizationPolicyCreate {
	apc.mutation.SetV4(s)
	return apc
}

// SetNillableV4 sets the "V4" field if the given value is not nil.
func (apc *AuthorizationPolicyCreate) SetNillableV4(s *string) *AuthorizationPolicyCreate {
	if s != nil {
		apc.SetV4(*s)
	}
	return apc
}

// SetV5 sets the "V5" field.
func (apc *AuthorizationPolicyCreate) SetV5(s string) *AuthorizationPolicyCreate {
	apc.mutation.SetV5(s)
	return apc
}

// SetNillableV5 sets the "V5" field if the given value is not nil.
func (apc *AuthorizationPolicyCreate) SetNillableV5(s *string) *AuthorizationPolicyCreate {
	if s != nil {
		apc.SetV5(*s)
	}
	return apc
}

// Mutation returns the AuthorizationPolicyMutation object of the builder.
func (apc *AuthorizationPolicyCreate) Mutation() *AuthorizationPolicyMutation {
	return apc.mutation
}

// Save creates the AuthorizationPolicy in the database.
func (apc *AuthorizationPolicyCreate) Save(ctx context.Context) (*AuthorizationPolicy, error) {
	apc.defaults()
	return withHooks[*AuthorizationPolicy, AuthorizationPolicyMutation](ctx, apc.sqlSave, apc.mutation, apc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (apc *AuthorizationPolicyCreate) SaveX(ctx context.Context) *AuthorizationPolicy {
	v, err := apc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (apc *AuthorizationPolicyCreate) Exec(ctx context.Context) error {
	_, err := apc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (apc *AuthorizationPolicyCreate) ExecX(ctx context.Context) {
	if err := apc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (apc *AuthorizationPolicyCreate) defaults() {
	if _, ok := apc.mutation.Ptype(); !ok {
		v := authorizationpolicy.DefaultPtype
		apc.mutation.SetPtype(v)
	}
	if _, ok := apc.mutation.V0(); !ok {
		v := authorizationpolicy.DefaultV0
		apc.mutation.SetV0(v)
	}
	if _, ok := apc.mutation.V1(); !ok {
		v := authorizationpolicy.DefaultV1
		apc.mutation.SetV1(v)
	}
	if _, ok := apc.mutation.V2(); !ok {
		v := authorizationpolicy.DefaultV2
		apc.mutation.SetV2(v)
	}
	if _, ok := apc.mutation.V3(); !ok {
		v := authorizationpolicy.DefaultV3
		apc.mutation.SetV3(v)
	}
	if _, ok := apc.mutation.V4(); !ok {
		v := authorizationpolicy.DefaultV4
		apc.mutation.SetV4(v)
	}
	if _, ok := apc.mutation.V5(); !ok {
		v := authorizationpolicy.DefaultV5
		apc.mutation.SetV5(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (apc *AuthorizationPolicyCreate) check() error {
	if _, ok := apc.mutation.Ptype(); !ok {
		return &ValidationError{Name: "Ptype", err: errors.New(`ent: missing required field "AuthorizationPolicy.Ptype"`)}
	}
	if _, ok := apc.mutation.V0(); !ok {
		return &ValidationError{Name: "V0", err: errors.New(`ent: missing required field "AuthorizationPolicy.V0"`)}
	}
	if _, ok := apc.mutation.V1(); !ok {
		return &ValidationError{Name: "V1", err: errors.New(`ent: missing required field "AuthorizationPolicy.V1"`)}
	}
	if _, ok := apc.mutation.V2(); !ok {
		return &ValidationError{Name: "V2", err: errors.New(`ent: missing required field "AuthorizationPolicy.V2"`)}
	}
	if _, ok := apc.mutation.V3(); !ok {
		return &ValidationError{Name: "V3", err: errors.New(`ent: missing required field "AuthorizationPolicy.V3"`)}
	}
	if _, ok := apc.mutation.V4(); !ok {
		return &ValidationError{Name: "V4", err: errors.New(`ent: missing required field "AuthorizationPolicy.V4"`)}
	}
	if _, ok := apc.mutation.V5(); !ok {
		return &ValidationError{Name: "V5", err: errors.New(`ent: missing required field "AuthorizationPolicy.V5"`)}
	}
	return nil
}

func (apc *AuthorizationPolicyCreate) sqlSave(ctx context.Context) (*AuthorizationPolicy, error) {
	if err := apc.check(); err != nil {
		return nil, err
	}
	_node, _spec := apc.createSpec()
	if err := sqlgraph.CreateNode(ctx, apc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	apc.mutation.id = &_node.ID
	apc.mutation.done = true
	return _node, nil
}

func (apc *AuthorizationPolicyCreate) createSpec() (*AuthorizationPolicy, *sqlgraph.CreateSpec) {
	var (
		_node = &AuthorizationPolicy{config: apc.config}
		_spec = sqlgraph.NewCreateSpec(authorizationpolicy.Table, sqlgraph.NewFieldSpec(authorizationpolicy.FieldID, field.TypeInt))
	)
	_spec.OnConflict = apc.conflict
	if value, ok := apc.mutation.Ptype(); ok {
		_spec.SetField(authorizationpolicy.FieldPtype, field.TypeString, value)
		_node.Ptype = value
	}
	if value, ok := apc.mutation.V0(); ok {
		_spec.SetField(authorizationpolicy.FieldV0, field.TypeString, value)
		_node.V0 = value
	}
	if value, ok := apc.mutation.V1(); ok {
		_spec.SetField(authorizationpolicy.FieldV1, field.TypeString, value)
		_node.V1 = value
	}
	if value, ok := apc.mutation.V2(); ok {
		_spec.SetField(authorizationpolicy.FieldV2, field.TypeString, value)
		_node.V2 = value
	}
	if value, ok := apc.mutation.V3(); ok {
		_spec.SetField(authorizationpolicy.FieldV3, field.TypeString, value)
		_node.V3 = value
	}
	if value, ok := apc.mutation.V4(); ok {
		_spec.SetField(authorizationpolicy.FieldV4, field.TypeString, value)
		_node.V4 = value
	}
	if value, ok := apc.mutation.V5(); ok {
		_spec.SetField(authorizationpolicy.FieldV5, field.TypeString, value)
		_node.V5 = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.AuthorizationPolicy.Create().
//		SetPtype(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AuthorizationPolicyUpsert) {
//			SetPtype(v+v).
//		}).
//		Exec(ctx)
func (apc *AuthorizationPolicyCreate) OnConflict(opts ...sql.ConflictOption) *AuthorizationPolicyUpsertOne {
	apc.conflict = opts
	return &AuthorizationPolicyUpsertOne{
		create: apc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.AuthorizationPolicy.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (apc *AuthorizationPolicyCreate) OnConflictColumns(columns ...string) *AuthorizationPolicyUpsertOne {
	apc.conflict = append(apc.conflict, sql.ConflictColumns(columns...))
	return &AuthorizationPolicyUpsertOne{
		create: apc,
	}
}

type (
	// AuthorizationPolicyUpsertOne is the builder for "upsert"-ing
	//  one AuthorizationPolicy node.
	AuthorizationPolicyUpsertOne struct {
		create *AuthorizationPolicyCreate
	}

	// AuthorizationPolicyUpsert is the "OnConflict" setter.
	AuthorizationPolicyUpsert struct {
		*sql.UpdateSet
	}
)

// SetPtype sets the "Ptype" field.
func (u *AuthorizationPolicyUpsert) SetPtype(v string) *AuthorizationPolicyUpsert {
	u.Set(authorizationpolicy.FieldPtype, v)
	return u
}

// UpdatePtype sets the "Ptype" field to the value that was provided on create.
func (u *AuthorizationPolicyUpsert) UpdatePtype() *AuthorizationPolicyUpsert {
	u.SetExcluded(authorizationpolicy.FieldPtype)
	return u
}

// SetV0 sets the "V0" field.
func (u *AuthorizationPolicyUpsert) SetV0(v string) *AuthorizationPolicyUpsert {
	u.Set(authorizationpolicy.FieldV0, v)
	return u
}

// UpdateV0 sets the "V0" field to the value that was provided on create.
func (u *AuthorizationPolicyUpsert) UpdateV0() *AuthorizationPolicyUpsert {
	u.SetExcluded(authorizationpolicy.FieldV0)
	return u
}

// SetV1 sets the "V1" field.
func (u *AuthorizationPolicyUpsert) SetV1(v string) *AuthorizationPolicyUpsert {
	u.Set(authorizationpolicy.FieldV1, v)
	return u
}

// UpdateV1 sets the "V1" field to the value that was provided on create.
func (u *AuthorizationPolicyUpsert) UpdateV1() *AuthorizationPolicyUpsert {
	u.SetExcluded(authorizationpolicy.FieldV1)
	return u
}

// SetV2 sets the "V2" field.
func (u *AuthorizationPolicyUpsert) SetV2(v string) *AuthorizationPolicyUpsert {
	u.Set(authorizationpolicy.FieldV2, v)
	return u
}

// UpdateV2 sets the "V2" field to the value that was provided on create.
func (u *AuthorizationPolicyUpsert) UpdateV2() *AuthorizationPolicyUpsert {
	u.SetExcluded(authorizationpolicy.FieldV2)
	return u
}

// SetV3 sets the "V3" field.
func (u *AuthorizationPolicyUpsert) SetV3(v string) *AuthorizationPolicyUpsert {
	u.Set(authorizationpolicy.FieldV3, v)
	return u
}

// UpdateV3 sets the "V3" field to the value that was provided on create.
func (u *AuthorizationPolicyUpsert) UpdateV3() *AuthorizationPolicyUpsert {
	u.SetExcluded(authorizationpolicy.FieldV3)
	return u
}

// SetV4 sets the "V4" field.
func (u *AuthorizationPolicyUpsert) SetV4(v string) *AuthorizationPolicyUpsert {
	u.Set(authorizationpolicy.FieldV4, v)
	return u
}

// UpdateV4 sets the "V4" field to the value that was provided on create.
func (u *AuthorizationPolicyUpsert) UpdateV4() *AuthorizationPolicyUpsert {
	u.SetExcluded(authorizationpolicy.FieldV4)
	return u
}

// SetV5 sets the "V5" field.
func (u *AuthorizationPolicyUpsert) SetV5(v string) *AuthorizationPolicyUpsert {
	u.Set(authorizationpolicy.FieldV5, v)
	return u
}

// UpdateV5 sets the "V5" field to the value that was provided on create.
func (u *AuthorizationPolicyUpsert) UpdateV5() *AuthorizationPolicyUpsert {
	u.SetExcluded(authorizationpolicy.FieldV5)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.AuthorizationPolicy.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *AuthorizationPolicyUpsertOne) UpdateNewValues() *AuthorizationPolicyUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.AuthorizationPolicy.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *AuthorizationPolicyUpsertOne) Ignore() *AuthorizationPolicyUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AuthorizationPolicyUpsertOne) DoNothing() *AuthorizationPolicyUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AuthorizationPolicyCreate.OnConflict
// documentation for more info.
func (u *AuthorizationPolicyUpsertOne) Update(set func(*AuthorizationPolicyUpsert)) *AuthorizationPolicyUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AuthorizationPolicyUpsert{UpdateSet: update})
	}))
	return u
}

// SetPtype sets the "Ptype" field.
func (u *AuthorizationPolicyUpsertOne) SetPtype(v string) *AuthorizationPolicyUpsertOne {
	return u.Update(func(s *AuthorizationPolicyUpsert) {
		s.SetPtype(v)
	})
}

// UpdatePtype sets the "Ptype" field to the value that was provided on create.
func (u *AuthorizationPolicyUpsertOne) UpdatePtype() *AuthorizationPolicyUpsertOne {
	return u.Update(func(s *AuthorizationPolicyUpsert) {
		s.UpdatePtype()
	})
}

// SetV0 sets the "V0" field.
func (u *AuthorizationPolicyUpsertOne) SetV0(v string) *AuthorizationPolicyUpsertOne {
	return u.Update(func(s *AuthorizationPolicyUpsert) {
		s.SetV0(v)
	})
}

// UpdateV0 sets the "V0" field to the value that was provided on create.
func (u *AuthorizationPolicyUpsertOne) UpdateV0() *AuthorizationPolicyUpsertOne {
	return u.Update(func(s *AuthorizationPolicyUpsert) {
		s.UpdateV0()
	})
}

// SetV1 sets the "V1" field.
func (u *AuthorizationPolicyUpsertOne) SetV1(v string) *AuthorizationPolicyUpsertOne {
	return u.Update(func(s *AuthorizationPolicyUpsert) {
		s.SetV1(v)
	})
}

// UpdateV1 sets the "V1" field to the value that was provided on create.
func (u *AuthorizationPolicyUpsertOne) UpdateV1() *AuthorizationPolicyUpsertOne {
	return u.Update(func(s *AuthorizationPolicyUpsert) {
		s.UpdateV1()
	})
}

// SetV2 sets the "V2" field.
func (u *AuthorizationPolicyUpsertOne) SetV2(v string) *AuthorizationPolicyUpsertOne {
	return u.Update(func(s *AuthorizationPolicyUpsert) {
		s.SetV2(v)
	})
}

// UpdateV2 sets the "V2" field to the value that was provided on create.
func (u *AuthorizationPolicyUpsertOne) UpdateV2() *AuthorizationPolicyUpsertOne {
	return u.Update(func(s *AuthorizationPolicyUpsert) {
		s.UpdateV2()
	})
}

// SetV3 sets the "V3" field.
func (u *AuthorizationPolicyUpsertOne) SetV3(v string) *AuthorizationPolicyUpsertOne {
	return u.Update(func(s *AuthorizationPolicyUpsert) {
		s.SetV3(v)
	})
}

// UpdateV3 sets the "V3" field to the value that was provided on create.
func (u *AuthorizationPolicyUpsertOne) UpdateV3() *AuthorizationPolicyUpsertOne {
	return u.Update(func(s *AuthorizationPolicyUpsert) {
		s.UpdateV3()
	})
}

// SetV4 sets the "V4" field.
func (u *AuthorizationPolicyUpsertOne) SetV4(v string) *AuthorizationPolicyUpsertOne {
	return u.Update(func(s *AuthorizationPolicyUpsert) {
		s.SetV4(v)
	})
}

// UpdateV4 sets the "V4" field to the value that was provided on create.
func (u *AuthorizationPolicyUpsertOne) UpdateV4() *AuthorizationPolicyUpsertOne {
	return u.Update(func(s *AuthorizationPolicyUpsert) {
		s.UpdateV4()
	})
}

// SetV5 sets the "V5" field.
func (u *AuthorizationPolicyUpsertOne) SetV5(v string) *AuthorizationPolicyUpsertOne {
	return u.Update(func(s *AuthorizationPolicyUpsert) {
		s.SetV5(v)
	})
}

// UpdateV5 sets the "V5" field to the value that was provided on create.
func (u *AuthorizationPolicyUpsertOne) UpdateV5() *AuthorizationPolicyUpsertOne {
	return u.Update(func(s *AuthorizationPolicyUpsert) {
		s.UpdateV5()
	})
}

// Exec executes the query.
func (u *AuthorizationPolicyUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AuthorizationPolicyCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AuthorizationPolicyUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *AuthorizationPolicyUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *AuthorizationPolicyUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// AuthorizationPolicyCreateBulk is the builder for creating many AuthorizationPolicy entities in bulk.
type AuthorizationPolicyCreateBulk struct {
	config
	builders []*AuthorizationPolicyCreate
	conflict []sql.ConflictOption
}

// Save creates the AuthorizationPolicy entities in the database.
func (apcb *AuthorizationPolicyCreateBulk) Save(ctx context.Context) ([]*AuthorizationPolicy, error) {
	specs := make([]*sqlgraph.CreateSpec, len(apcb.builders))
	nodes := make([]*AuthorizationPolicy, len(apcb.builders))
	mutators := make([]Mutator, len(apcb.builders))
	for i := range apcb.builders {
		func(i int, root context.Context) {
			builder := apcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AuthorizationPolicyMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, apcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = apcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, apcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, apcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (apcb *AuthorizationPolicyCreateBulk) SaveX(ctx context.Context) []*AuthorizationPolicy {
	v, err := apcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (apcb *AuthorizationPolicyCreateBulk) Exec(ctx context.Context) error {
	_, err := apcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (apcb *AuthorizationPolicyCreateBulk) ExecX(ctx context.Context) {
	if err := apcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.AuthorizationPolicy.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AuthorizationPolicyUpsert) {
//			SetPtype(v+v).
//		}).
//		Exec(ctx)
func (apcb *AuthorizationPolicyCreateBulk) OnConflict(opts ...sql.ConflictOption) *AuthorizationPolicyUpsertBulk {
	apcb.conflict = opts
	return &AuthorizationPolicyUpsertBulk{
		create: apcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.AuthorizationPolicy.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (apcb *AuthorizationPolicyCreateBulk) OnConflictColumns(columns ...string) *AuthorizationPolicyUpsertBulk {
	apcb.conflict = append(apcb.conflict, sql.ConflictColumns(columns...))
	return &AuthorizationPolicyUpsertBulk{
		create: apcb,
	}
}

// AuthorizationPolicyUpsertBulk is the builder for "upsert"-ing
// a bulk of AuthorizationPolicy nodes.
type AuthorizationPolicyUpsertBulk struct {
	create *AuthorizationPolicyCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.AuthorizationPolicy.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *AuthorizationPolicyUpsertBulk) UpdateNewValues() *AuthorizationPolicyUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.AuthorizationPolicy.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *AuthorizationPolicyUpsertBulk) Ignore() *AuthorizationPolicyUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AuthorizationPolicyUpsertBulk) DoNothing() *AuthorizationPolicyUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AuthorizationPolicyCreateBulk.OnConflict
// documentation for more info.
func (u *AuthorizationPolicyUpsertBulk) Update(set func(*AuthorizationPolicyUpsert)) *AuthorizationPolicyUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AuthorizationPolicyUpsert{UpdateSet: update})
	}))
	return u
}

// SetPtype sets the "Ptype" field.
func (u *AuthorizationPolicyUpsertBulk) SetPtype(v string) *AuthorizationPolicyUpsertBulk {
	return u.Update(func(s *AuthorizationPolicyUpsert) {
		s.SetPtype(v)
	})
}

// UpdatePtype sets the "Ptype" field to the value that was provided on create.
func (u *AuthorizationPolicyUpsertBulk) UpdatePtype() *AuthorizationPolicyUpsertBulk {
	return u.Update(func(s *AuthorizationPolicyUpsert) {
		s.UpdatePtype()
	})
}

// SetV0 sets the "V0" field.
func (u *AuthorizationPolicyUpsertBulk) SetV0(v string) *AuthorizationPolicyUpsertBulk {
	return u.Update(func(s *AuthorizationPolicyUpsert) {
		s.SetV0(v)
	})
}

// UpdateV0 sets the "V0" field to the value that was provided on create.
func (u *AuthorizationPolicyUpsertBulk) UpdateV0() *AuthorizationPolicyUpsertBulk {
	return u.Update(func(s *AuthorizationPolicyUpsert) {
		s.UpdateV0()
	})
}

// SetV1 sets the "V1" field.
func (u *AuthorizationPolicyUpsertBulk) SetV1(v string) *AuthorizationPolicyUpsertBulk {
	return u.Update(func(s *AuthorizationPolicyUpsert) {
		s.SetV1(v)
	})
}

// UpdateV1 sets the "V1" field to the value that was provided on create.
func (u *AuthorizationPolicyUpsertBulk) UpdateV1() *AuthorizationPolicyUpsertBulk {
	return u.Update(func(s *AuthorizationPolicyUpsert) {
		s.UpdateV1()
	})
}

// SetV2 sets the "V2" field.
func (u *AuthorizationPolicyUpsertBulk) SetV2(v string) *AuthorizationPolicyUpsertBulk {
	return u.Update(func(s *AuthorizationPolicyUpsert) {
		s.SetV2(v)
	})
}

// UpdateV2 sets the "V2" field to the value that was provided on create.
func (u *AuthorizationPolicyUpsertBulk) UpdateV2() *AuthorizationPolicyUpsertBulk {
	return u.Update(func(s *AuthorizationPolicyUpsert) {
		s.UpdateV2()
	})
}

// SetV3 sets the "V3" field.
func (u *AuthorizationPolicyUpsertBulk) SetV3(v string) *AuthorizationPolicyUpsertBulk {
	return u.Update(func(s *AuthorizationPolicyUpsert) {
		s.SetV3(v)
	})
}

// UpdateV3 sets the "V3" field to the value that was provided on create.
func (u *AuthorizationPolicyUpsertBulk) UpdateV3() *AuthorizationPolicyUpsertBulk {
	return u.Update(func(s *AuthorizationPolicyUpsert) {
		s.UpdateV3()
	})
}

// SetV4 sets the "V4" field.
func (u *AuthorizationPolicyUpsertBulk) SetV4(v string) *AuthorizationPolicyUpsertBulk {
	return u.Update(func(s *AuthorizationPolicyUpsert) {
		s.SetV4(v)
	})
}

// UpdateV4 sets the "V4" field to the value that was provided on create.
func (u *AuthorizationPolicyUpsertBulk) UpdateV4() *AuthorizationPolicyUpsertBulk {
	return u.Update(func(s *AuthorizationPolicyUpsert) {
		s.UpdateV4()
	})
}

// SetV5 sets the "V5" field.
func (u *AuthorizationPolicyUpsertBulk) SetV5(v string) *AuthorizationPolicyUpsertBulk {
	return u.Update(func(s *AuthorizationPolicyUpsert) {
		s.SetV5(v)
	})
}

// UpdateV5 sets the "V5" field to the value that was provided on create.
func (u *AuthorizationPolicyUpsertBulk) UpdateV5() *AuthorizationPolicyUpsertBulk {
	return u.Update(func(s *AuthorizationPolicyUpsert) {
		s.UpdateV5()
	})
}

// Exec executes the query.
func (u *AuthorizationPolicyUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the AuthorizationPolicyCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AuthorizationPolicyCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AuthorizationPolicyUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
