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

	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent/itembatch"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent/itembatchtoitembatch"
)

// ItemBatchToItemBatchCreate is the builder for creating a ItemBatchToItemBatch entity.
type ItemBatchToItemBatchCreate struct {
	config
	mutation *ItemBatchToItemBatchMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetChildUUID sets the "childUUID" field.
func (ibtibc *ItemBatchToItemBatchCreate) SetChildUUID(u uuid.UUID) *ItemBatchToItemBatchCreate {
	ibtibc.mutation.SetChildUUID(u)
	return ibtibc
}

// SetParentUUID sets the "parentUUID" field.
func (ibtibc *ItemBatchToItemBatchCreate) SetParentUUID(u uuid.UUID) *ItemBatchToItemBatchCreate {
	ibtibc.mutation.SetParentUUID(u)
	return ibtibc
}

// SetID sets the "id" field.
func (ibtibc *ItemBatchToItemBatchCreate) SetID(u uuid.UUID) *ItemBatchToItemBatchCreate {
	ibtibc.mutation.SetID(u)
	return ibtibc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ibtibc *ItemBatchToItemBatchCreate) SetNillableID(u *uuid.UUID) *ItemBatchToItemBatchCreate {
	if u != nil {
		ibtibc.SetID(*u)
	}
	return ibtibc
}

// SetParentID sets the "parent" edge to the ItemBatch entity by ID.
func (ibtibc *ItemBatchToItemBatchCreate) SetParentID(id uuid.UUID) *ItemBatchToItemBatchCreate {
	ibtibc.mutation.SetParentID(id)
	return ibtibc
}

// SetParent sets the "parent" edge to the ItemBatch entity.
func (ibtibc *ItemBatchToItemBatchCreate) SetParent(i *ItemBatch) *ItemBatchToItemBatchCreate {
	return ibtibc.SetParentID(i.ID)
}

// SetChildID sets the "child" edge to the ItemBatch entity by ID.
func (ibtibc *ItemBatchToItemBatchCreate) SetChildID(id uuid.UUID) *ItemBatchToItemBatchCreate {
	ibtibc.mutation.SetChildID(id)
	return ibtibc
}

// SetChild sets the "child" edge to the ItemBatch entity.
func (ibtibc *ItemBatchToItemBatchCreate) SetChild(i *ItemBatch) *ItemBatchToItemBatchCreate {
	return ibtibc.SetChildID(i.ID)
}

// Mutation returns the ItemBatchToItemBatchMutation object of the builder.
func (ibtibc *ItemBatchToItemBatchCreate) Mutation() *ItemBatchToItemBatchMutation {
	return ibtibc.mutation
}

// Save creates the ItemBatchToItemBatch in the database.
func (ibtibc *ItemBatchToItemBatchCreate) Save(ctx context.Context) (*ItemBatchToItemBatch, error) {
	ibtibc.defaults()
	return withHooks[*ItemBatchToItemBatch, ItemBatchToItemBatchMutation](ctx, ibtibc.sqlSave, ibtibc.mutation, ibtibc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ibtibc *ItemBatchToItemBatchCreate) SaveX(ctx context.Context) *ItemBatchToItemBatch {
	v, err := ibtibc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ibtibc *ItemBatchToItemBatchCreate) Exec(ctx context.Context) error {
	_, err := ibtibc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ibtibc *ItemBatchToItemBatchCreate) ExecX(ctx context.Context) {
	if err := ibtibc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ibtibc *ItemBatchToItemBatchCreate) defaults() {
	if _, ok := ibtibc.mutation.ID(); !ok {
		v := itembatchtoitembatch.DefaultID()
		ibtibc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ibtibc *ItemBatchToItemBatchCreate) check() error {
	if _, ok := ibtibc.mutation.ChildUUID(); !ok {
		return &ValidationError{Name: "childUUID", err: errors.New(`ent: missing required field "ItemBatchToItemBatch.childUUID"`)}
	}
	if _, ok := ibtibc.mutation.ParentUUID(); !ok {
		return &ValidationError{Name: "parentUUID", err: errors.New(`ent: missing required field "ItemBatchToItemBatch.parentUUID"`)}
	}
	if _, ok := ibtibc.mutation.ParentID(); !ok {
		return &ValidationError{Name: "parent", err: errors.New(`ent: missing required edge "ItemBatchToItemBatch.parent"`)}
	}
	if _, ok := ibtibc.mutation.ChildID(); !ok {
		return &ValidationError{Name: "child", err: errors.New(`ent: missing required edge "ItemBatchToItemBatch.child"`)}
	}
	return nil
}

func (ibtibc *ItemBatchToItemBatchCreate) sqlSave(ctx context.Context) (*ItemBatchToItemBatch, error) {
	if err := ibtibc.check(); err != nil {
		return nil, err
	}
	_node, _spec := ibtibc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ibtibc.driver, _spec); err != nil {
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
	ibtibc.mutation.id = &_node.ID
	ibtibc.mutation.done = true
	return _node, nil
}

func (ibtibc *ItemBatchToItemBatchCreate) createSpec() (*ItemBatchToItemBatch, *sqlgraph.CreateSpec) {
	var (
		_node = &ItemBatchToItemBatch{config: ibtibc.config}
		_spec = sqlgraph.NewCreateSpec(itembatchtoitembatch.Table, sqlgraph.NewFieldSpec(itembatchtoitembatch.FieldID, field.TypeUUID))
	)
	_spec.OnConflict = ibtibc.conflict
	if id, ok := ibtibc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if nodes := ibtibc.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   itembatchtoitembatch.ParentTable,
			Columns: []string{itembatchtoitembatch.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: itembatch.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ParentUUID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ibtibc.mutation.ChildIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   itembatchtoitembatch.ChildTable,
			Columns: []string{itembatchtoitembatch.ChildColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: itembatch.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ChildUUID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.ItemBatchToItemBatch.Create().
//		SetChildUUID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ItemBatchToItemBatchUpsert) {
//			SetChildUUID(v+v).
//		}).
//		Exec(ctx)
func (ibtibc *ItemBatchToItemBatchCreate) OnConflict(opts ...sql.ConflictOption) *ItemBatchToItemBatchUpsertOne {
	ibtibc.conflict = opts
	return &ItemBatchToItemBatchUpsertOne{
		create: ibtibc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.ItemBatchToItemBatch.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ibtibc *ItemBatchToItemBatchCreate) OnConflictColumns(columns ...string) *ItemBatchToItemBatchUpsertOne {
	ibtibc.conflict = append(ibtibc.conflict, sql.ConflictColumns(columns...))
	return &ItemBatchToItemBatchUpsertOne{
		create: ibtibc,
	}
}

type (
	// ItemBatchToItemBatchUpsertOne is the builder for "upsert"-ing
	//  one ItemBatchToItemBatch node.
	ItemBatchToItemBatchUpsertOne struct {
		create *ItemBatchToItemBatchCreate
	}

	// ItemBatchToItemBatchUpsert is the "OnConflict" setter.
	ItemBatchToItemBatchUpsert struct {
		*sql.UpdateSet
	}
)

// SetChildUUID sets the "childUUID" field.
func (u *ItemBatchToItemBatchUpsert) SetChildUUID(v uuid.UUID) *ItemBatchToItemBatchUpsert {
	u.Set(itembatchtoitembatch.FieldChildUUID, v)
	return u
}

// UpdateChildUUID sets the "childUUID" field to the value that was provided on create.
func (u *ItemBatchToItemBatchUpsert) UpdateChildUUID() *ItemBatchToItemBatchUpsert {
	u.SetExcluded(itembatchtoitembatch.FieldChildUUID)
	return u
}

// SetParentUUID sets the "parentUUID" field.
func (u *ItemBatchToItemBatchUpsert) SetParentUUID(v uuid.UUID) *ItemBatchToItemBatchUpsert {
	u.Set(itembatchtoitembatch.FieldParentUUID, v)
	return u
}

// UpdateParentUUID sets the "parentUUID" field to the value that was provided on create.
func (u *ItemBatchToItemBatchUpsert) UpdateParentUUID() *ItemBatchToItemBatchUpsert {
	u.SetExcluded(itembatchtoitembatch.FieldParentUUID)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.ItemBatchToItemBatch.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(itembatchtoitembatch.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *ItemBatchToItemBatchUpsertOne) UpdateNewValues() *ItemBatchToItemBatchUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(itembatchtoitembatch.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.ItemBatchToItemBatch.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *ItemBatchToItemBatchUpsertOne) Ignore() *ItemBatchToItemBatchUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ItemBatchToItemBatchUpsertOne) DoNothing() *ItemBatchToItemBatchUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ItemBatchToItemBatchCreate.OnConflict
// documentation for more info.
func (u *ItemBatchToItemBatchUpsertOne) Update(set func(*ItemBatchToItemBatchUpsert)) *ItemBatchToItemBatchUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ItemBatchToItemBatchUpsert{UpdateSet: update})
	}))
	return u
}

// SetChildUUID sets the "childUUID" field.
func (u *ItemBatchToItemBatchUpsertOne) SetChildUUID(v uuid.UUID) *ItemBatchToItemBatchUpsertOne {
	return u.Update(func(s *ItemBatchToItemBatchUpsert) {
		s.SetChildUUID(v)
	})
}

// UpdateChildUUID sets the "childUUID" field to the value that was provided on create.
func (u *ItemBatchToItemBatchUpsertOne) UpdateChildUUID() *ItemBatchToItemBatchUpsertOne {
	return u.Update(func(s *ItemBatchToItemBatchUpsert) {
		s.UpdateChildUUID()
	})
}

// SetParentUUID sets the "parentUUID" field.
func (u *ItemBatchToItemBatchUpsertOne) SetParentUUID(v uuid.UUID) *ItemBatchToItemBatchUpsertOne {
	return u.Update(func(s *ItemBatchToItemBatchUpsert) {
		s.SetParentUUID(v)
	})
}

// UpdateParentUUID sets the "parentUUID" field to the value that was provided on create.
func (u *ItemBatchToItemBatchUpsertOne) UpdateParentUUID() *ItemBatchToItemBatchUpsertOne {
	return u.Update(func(s *ItemBatchToItemBatchUpsert) {
		s.UpdateParentUUID()
	})
}

// Exec executes the query.
func (u *ItemBatchToItemBatchUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ItemBatchToItemBatchCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ItemBatchToItemBatchUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *ItemBatchToItemBatchUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: ItemBatchToItemBatchUpsertOne.ID is not supported by MySQL driver. Use ItemBatchToItemBatchUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *ItemBatchToItemBatchUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// ItemBatchToItemBatchCreateBulk is the builder for creating many ItemBatchToItemBatch entities in bulk.
type ItemBatchToItemBatchCreateBulk struct {
	config
	builders []*ItemBatchToItemBatchCreate
	conflict []sql.ConflictOption
}

// Save creates the ItemBatchToItemBatch entities in the database.
func (ibtibcb *ItemBatchToItemBatchCreateBulk) Save(ctx context.Context) ([]*ItemBatchToItemBatch, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ibtibcb.builders))
	nodes := make([]*ItemBatchToItemBatch, len(ibtibcb.builders))
	mutators := make([]Mutator, len(ibtibcb.builders))
	for i := range ibtibcb.builders {
		func(i int, root context.Context) {
			builder := ibtibcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ItemBatchToItemBatchMutation)
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
					_, err = mutators[i+1].Mutate(root, ibtibcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ibtibcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ibtibcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ibtibcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ibtibcb *ItemBatchToItemBatchCreateBulk) SaveX(ctx context.Context) []*ItemBatchToItemBatch {
	v, err := ibtibcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ibtibcb *ItemBatchToItemBatchCreateBulk) Exec(ctx context.Context) error {
	_, err := ibtibcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ibtibcb *ItemBatchToItemBatchCreateBulk) ExecX(ctx context.Context) {
	if err := ibtibcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.ItemBatchToItemBatch.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ItemBatchToItemBatchUpsert) {
//			SetChildUUID(v+v).
//		}).
//		Exec(ctx)
func (ibtibcb *ItemBatchToItemBatchCreateBulk) OnConflict(opts ...sql.ConflictOption) *ItemBatchToItemBatchUpsertBulk {
	ibtibcb.conflict = opts
	return &ItemBatchToItemBatchUpsertBulk{
		create: ibtibcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.ItemBatchToItemBatch.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ibtibcb *ItemBatchToItemBatchCreateBulk) OnConflictColumns(columns ...string) *ItemBatchToItemBatchUpsertBulk {
	ibtibcb.conflict = append(ibtibcb.conflict, sql.ConflictColumns(columns...))
	return &ItemBatchToItemBatchUpsertBulk{
		create: ibtibcb,
	}
}

// ItemBatchToItemBatchUpsertBulk is the builder for "upsert"-ing
// a bulk of ItemBatchToItemBatch nodes.
type ItemBatchToItemBatchUpsertBulk struct {
	create *ItemBatchToItemBatchCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.ItemBatchToItemBatch.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(itembatchtoitembatch.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *ItemBatchToItemBatchUpsertBulk) UpdateNewValues() *ItemBatchToItemBatchUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(itembatchtoitembatch.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.ItemBatchToItemBatch.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *ItemBatchToItemBatchUpsertBulk) Ignore() *ItemBatchToItemBatchUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ItemBatchToItemBatchUpsertBulk) DoNothing() *ItemBatchToItemBatchUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ItemBatchToItemBatchCreateBulk.OnConflict
// documentation for more info.
func (u *ItemBatchToItemBatchUpsertBulk) Update(set func(*ItemBatchToItemBatchUpsert)) *ItemBatchToItemBatchUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ItemBatchToItemBatchUpsert{UpdateSet: update})
	}))
	return u
}

// SetChildUUID sets the "childUUID" field.
func (u *ItemBatchToItemBatchUpsertBulk) SetChildUUID(v uuid.UUID) *ItemBatchToItemBatchUpsertBulk {
	return u.Update(func(s *ItemBatchToItemBatchUpsert) {
		s.SetChildUUID(v)
	})
}

// UpdateChildUUID sets the "childUUID" field to the value that was provided on create.
func (u *ItemBatchToItemBatchUpsertBulk) UpdateChildUUID() *ItemBatchToItemBatchUpsertBulk {
	return u.Update(func(s *ItemBatchToItemBatchUpsert) {
		s.UpdateChildUUID()
	})
}

// SetParentUUID sets the "parentUUID" field.
func (u *ItemBatchToItemBatchUpsertBulk) SetParentUUID(v uuid.UUID) *ItemBatchToItemBatchUpsertBulk {
	return u.Update(func(s *ItemBatchToItemBatchUpsert) {
		s.SetParentUUID(v)
	})
}

// UpdateParentUUID sets the "parentUUID" field to the value that was provided on create.
func (u *ItemBatchToItemBatchUpsertBulk) UpdateParentUUID() *ItemBatchToItemBatchUpsertBulk {
	return u.Update(func(s *ItemBatchToItemBatchUpsert) {
		s.UpdateParentUUID()
	})
}

// Exec executes the query.
func (u *ItemBatchToItemBatchUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the ItemBatchToItemBatchCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ItemBatchToItemBatchCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ItemBatchToItemBatchUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}