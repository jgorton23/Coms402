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

	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent/attributetype"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent/attributetypestotemplates"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent/certificationtemplate"
)

// AttributeTypesToTemplatesCreate is the builder for creating a AttributeTypesToTemplates entity.
type AttributeTypesToTemplatesCreate struct {
	config
	mutation *AttributeTypesToTemplatesMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetAttributeTypeUUID sets the "attributeTypeUUID" field.
func (atttc *AttributeTypesToTemplatesCreate) SetAttributeTypeUUID(u uuid.UUID) *AttributeTypesToTemplatesCreate {
	atttc.mutation.SetAttributeTypeUUID(u)
	return atttc
}

// SetTemplateUUID sets the "templateUUID" field.
func (atttc *AttributeTypesToTemplatesCreate) SetTemplateUUID(u uuid.UUID) *AttributeTypesToTemplatesCreate {
	atttc.mutation.SetTemplateUUID(u)
	return atttc
}

// SetID sets the "id" field.
func (atttc *AttributeTypesToTemplatesCreate) SetID(u uuid.UUID) *AttributeTypesToTemplatesCreate {
	atttc.mutation.SetID(u)
	return atttc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (atttc *AttributeTypesToTemplatesCreate) SetNillableID(u *uuid.UUID) *AttributeTypesToTemplatesCreate {
	if u != nil {
		atttc.SetID(*u)
	}
	return atttc
}

// SetAttributeID sets the "attribute" edge to the AttributeType entity by ID.
func (atttc *AttributeTypesToTemplatesCreate) SetAttributeID(id uuid.UUID) *AttributeTypesToTemplatesCreate {
	atttc.mutation.SetAttributeID(id)
	return atttc
}

// SetAttribute sets the "attribute" edge to the AttributeType entity.
func (atttc *AttributeTypesToTemplatesCreate) SetAttribute(a *AttributeType) *AttributeTypesToTemplatesCreate {
	return atttc.SetAttributeID(a.ID)
}

// SetTemplateID sets the "template" edge to the CertificationTemplate entity by ID.
func (atttc *AttributeTypesToTemplatesCreate) SetTemplateID(id uuid.UUID) *AttributeTypesToTemplatesCreate {
	atttc.mutation.SetTemplateID(id)
	return atttc
}

// SetTemplate sets the "template" edge to the CertificationTemplate entity.
func (atttc *AttributeTypesToTemplatesCreate) SetTemplate(c *CertificationTemplate) *AttributeTypesToTemplatesCreate {
	return atttc.SetTemplateID(c.ID)
}

// Mutation returns the AttributeTypesToTemplatesMutation object of the builder.
func (atttc *AttributeTypesToTemplatesCreate) Mutation() *AttributeTypesToTemplatesMutation {
	return atttc.mutation
}

// Save creates the AttributeTypesToTemplates in the database.
func (atttc *AttributeTypesToTemplatesCreate) Save(ctx context.Context) (*AttributeTypesToTemplates, error) {
	atttc.defaults()
	return withHooks[*AttributeTypesToTemplates, AttributeTypesToTemplatesMutation](ctx, atttc.sqlSave, atttc.mutation, atttc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (atttc *AttributeTypesToTemplatesCreate) SaveX(ctx context.Context) *AttributeTypesToTemplates {
	v, err := atttc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (atttc *AttributeTypesToTemplatesCreate) Exec(ctx context.Context) error {
	_, err := atttc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (atttc *AttributeTypesToTemplatesCreate) ExecX(ctx context.Context) {
	if err := atttc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (atttc *AttributeTypesToTemplatesCreate) defaults() {
	if _, ok := atttc.mutation.ID(); !ok {
		v := attributetypestotemplates.DefaultID()
		atttc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (atttc *AttributeTypesToTemplatesCreate) check() error {
	if _, ok := atttc.mutation.AttributeTypeUUID(); !ok {
		return &ValidationError{Name: "attributeTypeUUID", err: errors.New(`ent: missing required field "AttributeTypesToTemplates.attributeTypeUUID"`)}
	}
	if _, ok := atttc.mutation.TemplateUUID(); !ok {
		return &ValidationError{Name: "templateUUID", err: errors.New(`ent: missing required field "AttributeTypesToTemplates.templateUUID"`)}
	}
	if _, ok := atttc.mutation.AttributeID(); !ok {
		return &ValidationError{Name: "attribute", err: errors.New(`ent: missing required edge "AttributeTypesToTemplates.attribute"`)}
	}
	if _, ok := atttc.mutation.TemplateID(); !ok {
		return &ValidationError{Name: "template", err: errors.New(`ent: missing required edge "AttributeTypesToTemplates.template"`)}
	}
	return nil
}

func (atttc *AttributeTypesToTemplatesCreate) sqlSave(ctx context.Context) (*AttributeTypesToTemplates, error) {
	if err := atttc.check(); err != nil {
		return nil, err
	}
	_node, _spec := atttc.createSpec()
	if err := sqlgraph.CreateNode(ctx, atttc.driver, _spec); err != nil {
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
	atttc.mutation.id = &_node.ID
	atttc.mutation.done = true
	return _node, nil
}

func (atttc *AttributeTypesToTemplatesCreate) createSpec() (*AttributeTypesToTemplates, *sqlgraph.CreateSpec) {
	var (
		_node = &AttributeTypesToTemplates{config: atttc.config}
		_spec = sqlgraph.NewCreateSpec(attributetypestotemplates.Table, sqlgraph.NewFieldSpec(attributetypestotemplates.FieldID, field.TypeUUID))
	)
	_spec.OnConflict = atttc.conflict
	if id, ok := atttc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if nodes := atttc.mutation.AttributeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   attributetypestotemplates.AttributeTable,
			Columns: []string{attributetypestotemplates.AttributeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: attributetype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.AttributeTypeUUID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := atttc.mutation.TemplateIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   attributetypestotemplates.TemplateTable,
			Columns: []string{attributetypestotemplates.TemplateColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: certificationtemplate.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.TemplateUUID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.AttributeTypesToTemplates.Create().
//		SetAttributeTypeUUID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AttributeTypesToTemplatesUpsert) {
//			SetAttributeTypeUUID(v+v).
//		}).
//		Exec(ctx)
func (atttc *AttributeTypesToTemplatesCreate) OnConflict(opts ...sql.ConflictOption) *AttributeTypesToTemplatesUpsertOne {
	atttc.conflict = opts
	return &AttributeTypesToTemplatesUpsertOne{
		create: atttc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.AttributeTypesToTemplates.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (atttc *AttributeTypesToTemplatesCreate) OnConflictColumns(columns ...string) *AttributeTypesToTemplatesUpsertOne {
	atttc.conflict = append(atttc.conflict, sql.ConflictColumns(columns...))
	return &AttributeTypesToTemplatesUpsertOne{
		create: atttc,
	}
}

type (
	// AttributeTypesToTemplatesUpsertOne is the builder for "upsert"-ing
	//  one AttributeTypesToTemplates node.
	AttributeTypesToTemplatesUpsertOne struct {
		create *AttributeTypesToTemplatesCreate
	}

	// AttributeTypesToTemplatesUpsert is the "OnConflict" setter.
	AttributeTypesToTemplatesUpsert struct {
		*sql.UpdateSet
	}
)

// SetAttributeTypeUUID sets the "attributeTypeUUID" field.
func (u *AttributeTypesToTemplatesUpsert) SetAttributeTypeUUID(v uuid.UUID) *AttributeTypesToTemplatesUpsert {
	u.Set(attributetypestotemplates.FieldAttributeTypeUUID, v)
	return u
}

// UpdateAttributeTypeUUID sets the "attributeTypeUUID" field to the value that was provided on create.
func (u *AttributeTypesToTemplatesUpsert) UpdateAttributeTypeUUID() *AttributeTypesToTemplatesUpsert {
	u.SetExcluded(attributetypestotemplates.FieldAttributeTypeUUID)
	return u
}

// SetTemplateUUID sets the "templateUUID" field.
func (u *AttributeTypesToTemplatesUpsert) SetTemplateUUID(v uuid.UUID) *AttributeTypesToTemplatesUpsert {
	u.Set(attributetypestotemplates.FieldTemplateUUID, v)
	return u
}

// UpdateTemplateUUID sets the "templateUUID" field to the value that was provided on create.
func (u *AttributeTypesToTemplatesUpsert) UpdateTemplateUUID() *AttributeTypesToTemplatesUpsert {
	u.SetExcluded(attributetypestotemplates.FieldTemplateUUID)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.AttributeTypesToTemplates.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(attributetypestotemplates.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *AttributeTypesToTemplatesUpsertOne) UpdateNewValues() *AttributeTypesToTemplatesUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(attributetypestotemplates.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.AttributeTypesToTemplates.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *AttributeTypesToTemplatesUpsertOne) Ignore() *AttributeTypesToTemplatesUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AttributeTypesToTemplatesUpsertOne) DoNothing() *AttributeTypesToTemplatesUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AttributeTypesToTemplatesCreate.OnConflict
// documentation for more info.
func (u *AttributeTypesToTemplatesUpsertOne) Update(set func(*AttributeTypesToTemplatesUpsert)) *AttributeTypesToTemplatesUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AttributeTypesToTemplatesUpsert{UpdateSet: update})
	}))
	return u
}

// SetAttributeTypeUUID sets the "attributeTypeUUID" field.
func (u *AttributeTypesToTemplatesUpsertOne) SetAttributeTypeUUID(v uuid.UUID) *AttributeTypesToTemplatesUpsertOne {
	return u.Update(func(s *AttributeTypesToTemplatesUpsert) {
		s.SetAttributeTypeUUID(v)
	})
}

// UpdateAttributeTypeUUID sets the "attributeTypeUUID" field to the value that was provided on create.
func (u *AttributeTypesToTemplatesUpsertOne) UpdateAttributeTypeUUID() *AttributeTypesToTemplatesUpsertOne {
	return u.Update(func(s *AttributeTypesToTemplatesUpsert) {
		s.UpdateAttributeTypeUUID()
	})
}

// SetTemplateUUID sets the "templateUUID" field.
func (u *AttributeTypesToTemplatesUpsertOne) SetTemplateUUID(v uuid.UUID) *AttributeTypesToTemplatesUpsertOne {
	return u.Update(func(s *AttributeTypesToTemplatesUpsert) {
		s.SetTemplateUUID(v)
	})
}

// UpdateTemplateUUID sets the "templateUUID" field to the value that was provided on create.
func (u *AttributeTypesToTemplatesUpsertOne) UpdateTemplateUUID() *AttributeTypesToTemplatesUpsertOne {
	return u.Update(func(s *AttributeTypesToTemplatesUpsert) {
		s.UpdateTemplateUUID()
	})
}

// Exec executes the query.
func (u *AttributeTypesToTemplatesUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AttributeTypesToTemplatesCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AttributeTypesToTemplatesUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *AttributeTypesToTemplatesUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: AttributeTypesToTemplatesUpsertOne.ID is not supported by MySQL driver. Use AttributeTypesToTemplatesUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *AttributeTypesToTemplatesUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// AttributeTypesToTemplatesCreateBulk is the builder for creating many AttributeTypesToTemplates entities in bulk.
type AttributeTypesToTemplatesCreateBulk struct {
	config
	builders []*AttributeTypesToTemplatesCreate
	conflict []sql.ConflictOption
}

// Save creates the AttributeTypesToTemplates entities in the database.
func (atttcb *AttributeTypesToTemplatesCreateBulk) Save(ctx context.Context) ([]*AttributeTypesToTemplates, error) {
	specs := make([]*sqlgraph.CreateSpec, len(atttcb.builders))
	nodes := make([]*AttributeTypesToTemplates, len(atttcb.builders))
	mutators := make([]Mutator, len(atttcb.builders))
	for i := range atttcb.builders {
		func(i int, root context.Context) {
			builder := atttcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AttributeTypesToTemplatesMutation)
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
					_, err = mutators[i+1].Mutate(root, atttcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = atttcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, atttcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, atttcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (atttcb *AttributeTypesToTemplatesCreateBulk) SaveX(ctx context.Context) []*AttributeTypesToTemplates {
	v, err := atttcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (atttcb *AttributeTypesToTemplatesCreateBulk) Exec(ctx context.Context) error {
	_, err := atttcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (atttcb *AttributeTypesToTemplatesCreateBulk) ExecX(ctx context.Context) {
	if err := atttcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.AttributeTypesToTemplates.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AttributeTypesToTemplatesUpsert) {
//			SetAttributeTypeUUID(v+v).
//		}).
//		Exec(ctx)
func (atttcb *AttributeTypesToTemplatesCreateBulk) OnConflict(opts ...sql.ConflictOption) *AttributeTypesToTemplatesUpsertBulk {
	atttcb.conflict = opts
	return &AttributeTypesToTemplatesUpsertBulk{
		create: atttcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.AttributeTypesToTemplates.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (atttcb *AttributeTypesToTemplatesCreateBulk) OnConflictColumns(columns ...string) *AttributeTypesToTemplatesUpsertBulk {
	atttcb.conflict = append(atttcb.conflict, sql.ConflictColumns(columns...))
	return &AttributeTypesToTemplatesUpsertBulk{
		create: atttcb,
	}
}

// AttributeTypesToTemplatesUpsertBulk is the builder for "upsert"-ing
// a bulk of AttributeTypesToTemplates nodes.
type AttributeTypesToTemplatesUpsertBulk struct {
	create *AttributeTypesToTemplatesCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.AttributeTypesToTemplates.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(attributetypestotemplates.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *AttributeTypesToTemplatesUpsertBulk) UpdateNewValues() *AttributeTypesToTemplatesUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(attributetypestotemplates.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.AttributeTypesToTemplates.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *AttributeTypesToTemplatesUpsertBulk) Ignore() *AttributeTypesToTemplatesUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AttributeTypesToTemplatesUpsertBulk) DoNothing() *AttributeTypesToTemplatesUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AttributeTypesToTemplatesCreateBulk.OnConflict
// documentation for more info.
func (u *AttributeTypesToTemplatesUpsertBulk) Update(set func(*AttributeTypesToTemplatesUpsert)) *AttributeTypesToTemplatesUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AttributeTypesToTemplatesUpsert{UpdateSet: update})
	}))
	return u
}

// SetAttributeTypeUUID sets the "attributeTypeUUID" field.
func (u *AttributeTypesToTemplatesUpsertBulk) SetAttributeTypeUUID(v uuid.UUID) *AttributeTypesToTemplatesUpsertBulk {
	return u.Update(func(s *AttributeTypesToTemplatesUpsert) {
		s.SetAttributeTypeUUID(v)
	})
}

// UpdateAttributeTypeUUID sets the "attributeTypeUUID" field to the value that was provided on create.
func (u *AttributeTypesToTemplatesUpsertBulk) UpdateAttributeTypeUUID() *AttributeTypesToTemplatesUpsertBulk {
	return u.Update(func(s *AttributeTypesToTemplatesUpsert) {
		s.UpdateAttributeTypeUUID()
	})
}

// SetTemplateUUID sets the "templateUUID" field.
func (u *AttributeTypesToTemplatesUpsertBulk) SetTemplateUUID(v uuid.UUID) *AttributeTypesToTemplatesUpsertBulk {
	return u.Update(func(s *AttributeTypesToTemplatesUpsert) {
		s.SetTemplateUUID(v)
	})
}

// UpdateTemplateUUID sets the "templateUUID" field to the value that was provided on create.
func (u *AttributeTypesToTemplatesUpsertBulk) UpdateTemplateUUID() *AttributeTypesToTemplatesUpsertBulk {
	return u.Update(func(s *AttributeTypesToTemplatesUpsert) {
		s.UpdateTemplateUUID()
	})
}

// Exec executes the query.
func (u *AttributeTypesToTemplatesUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the AttributeTypesToTemplatesCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AttributeTypesToTemplatesCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AttributeTypesToTemplatesUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}