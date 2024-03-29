// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"

	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent/attribute"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent/attributetype"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent/certification"
)

// AttributeCreate is the builder for creating a Attribute entity.
type AttributeCreate struct {
	config
	mutation *AttributeMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetKey sets the "key" field.
func (ac *AttributeCreate) SetKey(s string) *AttributeCreate {
	ac.mutation.SetKey(s)
	return ac
}

// SetValue sets the "value" field.
func (ac *AttributeCreate) SetValue(s string) *AttributeCreate {
	ac.mutation.SetValue(s)
	return ac
}

// SetCertUUID sets the "certUUID" field.
func (ac *AttributeCreate) SetCertUUID(u uuid.UUID) *AttributeCreate {
	ac.mutation.SetCertUUID(u)
	return ac
}

// SetAttributeTypeUUID sets the "attributeTypeUUID" field.
func (ac *AttributeCreate) SetAttributeTypeUUID(u uuid.UUID) *AttributeCreate {
	ac.mutation.SetAttributeTypeUUID(u)
	return ac
}

// SetID sets the "id" field.
func (ac *AttributeCreate) SetID(u uuid.UUID) *AttributeCreate {
	ac.mutation.SetID(u)
	return ac
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ac *AttributeCreate) SetNillableID(u *uuid.UUID) *AttributeCreate {
	if u != nil {
		ac.SetID(*u)
	}
	return ac
}

// SetCertificationID sets the "certification" edge to the Certification entity by ID.
func (ac *AttributeCreate) SetCertificationID(id uuid.UUID) *AttributeCreate {
	ac.mutation.SetCertificationID(id)
	return ac
}

// SetCertification sets the "certification" edge to the Certification entity.
func (ac *AttributeCreate) SetCertification(c *Certification) *AttributeCreate {
	return ac.SetCertificationID(c.ID)
}

// SetAttributeTypeID sets the "attributeType" edge to the AttributeType entity by ID.
func (ac *AttributeCreate) SetAttributeTypeID(id uuid.UUID) *AttributeCreate {
	ac.mutation.SetAttributeTypeID(id)
	return ac
}

// SetAttributeType sets the "attributeType" edge to the AttributeType entity.
func (ac *AttributeCreate) SetAttributeType(a *AttributeType) *AttributeCreate {
	return ac.SetAttributeTypeID(a.ID)
}

// Mutation returns the AttributeMutation object of the builder.
func (ac *AttributeCreate) Mutation() *AttributeMutation {
	return ac.mutation
}

// Save creates the Attribute in the database.
func (ac *AttributeCreate) Save(ctx context.Context) (*Attribute, error) {
	ac.defaults()
	return withHooks[*Attribute, AttributeMutation](ctx, ac.sqlSave, ac.mutation, ac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ac *AttributeCreate) SaveX(ctx context.Context) *Attribute {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *AttributeCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *AttributeCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ac *AttributeCreate) defaults() {
	if _, ok := ac.mutation.ID(); !ok {
		v := attribute.DefaultID()
		ac.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *AttributeCreate) check() error {
	if _, ok := ac.mutation.Key(); !ok {
		return &ValidationError{Name: "key", err: errors.New(`ent: missing required field "Attribute.key"`)}
	}
	if _, ok := ac.mutation.Value(); !ok {
		return &ValidationError{Name: "value", err: errors.New(`ent: missing required field "Attribute.value"`)}
	}
	if _, ok := ac.mutation.CertUUID(); !ok {
		return &ValidationError{Name: "certUUID", err: errors.New(`ent: missing required field "Attribute.certUUID"`)}
	}
	if _, ok := ac.mutation.AttributeTypeUUID(); !ok {
		return &ValidationError{Name: "attributeTypeUUID", err: errors.New(`ent: missing required field "Attribute.attributeTypeUUID"`)}
	}
	if _, ok := ac.mutation.CertificationID(); !ok {
		return &ValidationError{Name: "certification", err: errors.New(`ent: missing required edge "Attribute.certification"`)}
	}
	if _, ok := ac.mutation.AttributeTypeID(); !ok {
		return &ValidationError{Name: "attributeType", err: errors.New(`ent: missing required edge "Attribute.attributeType"`)}
	}
	return nil
}

func (ac *AttributeCreate) sqlSave(ctx context.Context) (*Attribute, error) {
	if err := ac.check(); err != nil {
		return nil, err
	}
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	ac.mutation.id = &_node.ID
	ac.mutation.done = true
	return _node, nil
}

func (ac *AttributeCreate) createSpec() (*Attribute, *sqlgraph.CreateSpec) {
	var (
		_node = &Attribute{config: ac.config}
		_spec = sqlgraph.NewCreateSpec(attribute.Table, sqlgraph.NewFieldSpec(attribute.FieldID, field.TypeUUID))
	)
	_spec.OnConflict = ac.conflict
	if id, ok := ac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := ac.mutation.Key(); ok {
		_spec.SetField(attribute.FieldKey, field.TypeString, value)
		_node.Key = value
	}
	if value, ok := ac.mutation.Value(); ok {
		_spec.SetField(attribute.FieldValue, field.TypeString, value)
		_node.Value = value
	}
	if nodes := ac.mutation.CertificationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   attribute.CertificationTable,
			Columns: []string{attribute.CertificationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(certification.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.CertUUID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.AttributeTypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   attribute.AttributeTypeTable,
			Columns: []string{attribute.AttributeTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(attributetype.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.AttributeTypeUUID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Attribute.Create().
//		SetKey(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AttributeUpsert) {
//			SetKey(v+v).
//		}).
//		Exec(ctx)
func (ac *AttributeCreate) OnConflict(opts ...sql.ConflictOption) *AttributeUpsertOne {
	ac.conflict = opts
	return &AttributeUpsertOne{
		create: ac,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Attribute.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ac *AttributeCreate) OnConflictColumns(columns ...string) *AttributeUpsertOne {
	ac.conflict = append(ac.conflict, sql.ConflictColumns(columns...))
	return &AttributeUpsertOne{
		create: ac,
	}
}

type (
	// AttributeUpsertOne is the builder for "upsert"-ing
	//  one Attribute node.
	AttributeUpsertOne struct {
		create *AttributeCreate
	}

	// AttributeUpsert is the "OnConflict" setter.
	AttributeUpsert struct {
		*sql.UpdateSet
	}
)

// SetKey sets the "key" field.
func (u *AttributeUpsert) SetKey(v string) *AttributeUpsert {
	u.Set(attribute.FieldKey, v)
	return u
}

// UpdateKey sets the "key" field to the value that was provided on create.
func (u *AttributeUpsert) UpdateKey() *AttributeUpsert {
	u.SetExcluded(attribute.FieldKey)
	return u
}

// SetValue sets the "value" field.
func (u *AttributeUpsert) SetValue(v string) *AttributeUpsert {
	u.Set(attribute.FieldValue, v)
	return u
}

// UpdateValue sets the "value" field to the value that was provided on create.
func (u *AttributeUpsert) UpdateValue() *AttributeUpsert {
	u.SetExcluded(attribute.FieldValue)
	return u
}

// SetCertUUID sets the "certUUID" field.
func (u *AttributeUpsert) SetCertUUID(v uuid.UUID) *AttributeUpsert {
	u.Set(attribute.FieldCertUUID, v)
	return u
}

// UpdateCertUUID sets the "certUUID" field to the value that was provided on create.
func (u *AttributeUpsert) UpdateCertUUID() *AttributeUpsert {
	u.SetExcluded(attribute.FieldCertUUID)
	return u
}

// SetAttributeTypeUUID sets the "attributeTypeUUID" field.
func (u *AttributeUpsert) SetAttributeTypeUUID(v uuid.UUID) *AttributeUpsert {
	u.Set(attribute.FieldAttributeTypeUUID, v)
	return u
}

// UpdateAttributeTypeUUID sets the "attributeTypeUUID" field to the value that was provided on create.
func (u *AttributeUpsert) UpdateAttributeTypeUUID() *AttributeUpsert {
	u.SetExcluded(attribute.FieldAttributeTypeUUID)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Attribute.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(attribute.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *AttributeUpsertOne) UpdateNewValues() *AttributeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(attribute.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Attribute.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *AttributeUpsertOne) Ignore() *AttributeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AttributeUpsertOne) DoNothing() *AttributeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AttributeCreate.OnConflict
// documentation for more info.
func (u *AttributeUpsertOne) Update(set func(*AttributeUpsert)) *AttributeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AttributeUpsert{UpdateSet: update})
	}))
	return u
}

// SetKey sets the "key" field.
func (u *AttributeUpsertOne) SetKey(v string) *AttributeUpsertOne {
	return u.Update(func(s *AttributeUpsert) {
		s.SetKey(v)
	})
}

// UpdateKey sets the "key" field to the value that was provided on create.
func (u *AttributeUpsertOne) UpdateKey() *AttributeUpsertOne {
	return u.Update(func(s *AttributeUpsert) {
		s.UpdateKey()
	})
}

// SetValue sets the "value" field.
func (u *AttributeUpsertOne) SetValue(v string) *AttributeUpsertOne {
	return u.Update(func(s *AttributeUpsert) {
		s.SetValue(v)
	})
}

// UpdateValue sets the "value" field to the value that was provided on create.
func (u *AttributeUpsertOne) UpdateValue() *AttributeUpsertOne {
	return u.Update(func(s *AttributeUpsert) {
		s.UpdateValue()
	})
}

// SetCertUUID sets the "certUUID" field.
func (u *AttributeUpsertOne) SetCertUUID(v uuid.UUID) *AttributeUpsertOne {
	return u.Update(func(s *AttributeUpsert) {
		s.SetCertUUID(v)
	})
}

// UpdateCertUUID sets the "certUUID" field to the value that was provided on create.
func (u *AttributeUpsertOne) UpdateCertUUID() *AttributeUpsertOne {
	return u.Update(func(s *AttributeUpsert) {
		s.UpdateCertUUID()
	})
}

// SetAttributeTypeUUID sets the "attributeTypeUUID" field.
func (u *AttributeUpsertOne) SetAttributeTypeUUID(v uuid.UUID) *AttributeUpsertOne {
	return u.Update(func(s *AttributeUpsert) {
		s.SetAttributeTypeUUID(v)
	})
}

// UpdateAttributeTypeUUID sets the "attributeTypeUUID" field to the value that was provided on create.
func (u *AttributeUpsertOne) UpdateAttributeTypeUUID() *AttributeUpsertOne {
	return u.Update(func(s *AttributeUpsert) {
		s.UpdateAttributeTypeUUID()
	})
}

// Exec executes the query.
func (u *AttributeUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AttributeCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AttributeUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *AttributeUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: AttributeUpsertOne.ID is not supported by MySQL driver. Use AttributeUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *AttributeUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// AttributeCreateBulk is the builder for creating many Attribute entities in bulk.
type AttributeCreateBulk struct {
	config
	builders []*AttributeCreate
	conflict []sql.ConflictOption
}

// Save creates the Attribute entities in the database.
func (acb *AttributeCreateBulk) Save(ctx context.Context) ([]*Attribute, error) {
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Attribute, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AttributeMutation)
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
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = acb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
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
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *AttributeCreateBulk) SaveX(ctx context.Context) []*Attribute {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *AttributeCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *AttributeCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Attribute.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AttributeUpsert) {
//			SetKey(v+v).
//		}).
//		Exec(ctx)
func (acb *AttributeCreateBulk) OnConflict(opts ...sql.ConflictOption) *AttributeUpsertBulk {
	acb.conflict = opts
	return &AttributeUpsertBulk{
		create: acb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Attribute.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (acb *AttributeCreateBulk) OnConflictColumns(columns ...string) *AttributeUpsertBulk {
	acb.conflict = append(acb.conflict, sql.ConflictColumns(columns...))
	return &AttributeUpsertBulk{
		create: acb,
	}
}

// AttributeUpsertBulk is the builder for "upsert"-ing
// a bulk of Attribute nodes.
type AttributeUpsertBulk struct {
	create *AttributeCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Attribute.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(attribute.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *AttributeUpsertBulk) UpdateNewValues() *AttributeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(attribute.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Attribute.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *AttributeUpsertBulk) Ignore() *AttributeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AttributeUpsertBulk) DoNothing() *AttributeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AttributeCreateBulk.OnConflict
// documentation for more info.
func (u *AttributeUpsertBulk) Update(set func(*AttributeUpsert)) *AttributeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AttributeUpsert{UpdateSet: update})
	}))
	return u
}

// SetKey sets the "key" field.
func (u *AttributeUpsertBulk) SetKey(v string) *AttributeUpsertBulk {
	return u.Update(func(s *AttributeUpsert) {
		s.SetKey(v)
	})
}

// UpdateKey sets the "key" field to the value that was provided on create.
func (u *AttributeUpsertBulk) UpdateKey() *AttributeUpsertBulk {
	return u.Update(func(s *AttributeUpsert) {
		s.UpdateKey()
	})
}

// SetValue sets the "value" field.
func (u *AttributeUpsertBulk) SetValue(v string) *AttributeUpsertBulk {
	return u.Update(func(s *AttributeUpsert) {
		s.SetValue(v)
	})
}

// UpdateValue sets the "value" field to the value that was provided on create.
func (u *AttributeUpsertBulk) UpdateValue() *AttributeUpsertBulk {
	return u.Update(func(s *AttributeUpsert) {
		s.UpdateValue()
	})
}

// SetCertUUID sets the "certUUID" field.
func (u *AttributeUpsertBulk) SetCertUUID(v uuid.UUID) *AttributeUpsertBulk {
	return u.Update(func(s *AttributeUpsert) {
		s.SetCertUUID(v)
	})
}

// UpdateCertUUID sets the "certUUID" field to the value that was provided on create.
func (u *AttributeUpsertBulk) UpdateCertUUID() *AttributeUpsertBulk {
	return u.Update(func(s *AttributeUpsert) {
		s.UpdateCertUUID()
	})
}

// SetAttributeTypeUUID sets the "attributeTypeUUID" field.
func (u *AttributeUpsertBulk) SetAttributeTypeUUID(v uuid.UUID) *AttributeUpsertBulk {
	return u.Update(func(s *AttributeUpsert) {
		s.SetAttributeTypeUUID(v)
	})
}

// UpdateAttributeTypeUUID sets the "attributeTypeUUID" field to the value that was provided on create.
func (u *AttributeUpsertBulk) UpdateAttributeTypeUUID() *AttributeUpsertBulk {
	return u.Update(func(s *AttributeUpsert) {
		s.UpdateAttributeTypeUUID()
	})
}

// Exec executes the query.
func (u *AttributeUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the AttributeCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AttributeCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AttributeUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
