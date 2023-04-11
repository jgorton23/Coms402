// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"

	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent/itembatchtoitembatchhistory"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/enthistory"
)

// ItemBatchToItemBatchHistoryCreate is the builder for creating a ItemBatchToItemBatchHistory entity.
type ItemBatchToItemBatchHistoryCreate struct {
	config
	mutation *ItemBatchToItemBatchHistoryMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetHistoryTime sets the "history_time" field.
func (ibtibhc *ItemBatchToItemBatchHistoryCreate) SetHistoryTime(t time.Time) *ItemBatchToItemBatchHistoryCreate {
	ibtibhc.mutation.SetHistoryTime(t)
	return ibtibhc
}

// SetNillableHistoryTime sets the "history_time" field if the given value is not nil.
func (ibtibhc *ItemBatchToItemBatchHistoryCreate) SetNillableHistoryTime(t *time.Time) *ItemBatchToItemBatchHistoryCreate {
	if t != nil {
		ibtibhc.SetHistoryTime(*t)
	}
	return ibtibhc
}

// SetRef sets the "ref" field.
func (ibtibhc *ItemBatchToItemBatchHistoryCreate) SetRef(u uuid.UUID) *ItemBatchToItemBatchHistoryCreate {
	ibtibhc.mutation.SetRef(u)
	return ibtibhc
}

// SetNillableRef sets the "ref" field if the given value is not nil.
func (ibtibhc *ItemBatchToItemBatchHistoryCreate) SetNillableRef(u *uuid.UUID) *ItemBatchToItemBatchHistoryCreate {
	if u != nil {
		ibtibhc.SetRef(*u)
	}
	return ibtibhc
}

// SetOperation sets the "operation" field.
func (ibtibhc *ItemBatchToItemBatchHistoryCreate) SetOperation(et enthistory.OpType) *ItemBatchToItemBatchHistoryCreate {
	ibtibhc.mutation.SetOperation(et)
	return ibtibhc
}

// SetUpdatedBy sets the "updated_by" field.
func (ibtibhc *ItemBatchToItemBatchHistoryCreate) SetUpdatedBy(s string) *ItemBatchToItemBatchHistoryCreate {
	ibtibhc.mutation.SetUpdatedBy(s)
	return ibtibhc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (ibtibhc *ItemBatchToItemBatchHistoryCreate) SetNillableUpdatedBy(s *string) *ItemBatchToItemBatchHistoryCreate {
	if s != nil {
		ibtibhc.SetUpdatedBy(*s)
	}
	return ibtibhc
}

// SetChildUUID sets the "childUUID" field.
func (ibtibhc *ItemBatchToItemBatchHistoryCreate) SetChildUUID(u uuid.UUID) *ItemBatchToItemBatchHistoryCreate {
	ibtibhc.mutation.SetChildUUID(u)
	return ibtibhc
}

// SetParentUUID sets the "parentUUID" field.
func (ibtibhc *ItemBatchToItemBatchHistoryCreate) SetParentUUID(u uuid.UUID) *ItemBatchToItemBatchHistoryCreate {
	ibtibhc.mutation.SetParentUUID(u)
	return ibtibhc
}

// SetID sets the "id" field.
func (ibtibhc *ItemBatchToItemBatchHistoryCreate) SetID(u uuid.UUID) *ItemBatchToItemBatchHistoryCreate {
	ibtibhc.mutation.SetID(u)
	return ibtibhc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ibtibhc *ItemBatchToItemBatchHistoryCreate) SetNillableID(u *uuid.UUID) *ItemBatchToItemBatchHistoryCreate {
	if u != nil {
		ibtibhc.SetID(*u)
	}
	return ibtibhc
}

// Mutation returns the ItemBatchToItemBatchHistoryMutation object of the builder.
func (ibtibhc *ItemBatchToItemBatchHistoryCreate) Mutation() *ItemBatchToItemBatchHistoryMutation {
	return ibtibhc.mutation
}

// Save creates the ItemBatchToItemBatchHistory in the database.
func (ibtibhc *ItemBatchToItemBatchHistoryCreate) Save(ctx context.Context) (*ItemBatchToItemBatchHistory, error) {
	ibtibhc.defaults()
	return withHooks[*ItemBatchToItemBatchHistory, ItemBatchToItemBatchHistoryMutation](ctx, ibtibhc.sqlSave, ibtibhc.mutation, ibtibhc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ibtibhc *ItemBatchToItemBatchHistoryCreate) SaveX(ctx context.Context) *ItemBatchToItemBatchHistory {
	v, err := ibtibhc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ibtibhc *ItemBatchToItemBatchHistoryCreate) Exec(ctx context.Context) error {
	_, err := ibtibhc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ibtibhc *ItemBatchToItemBatchHistoryCreate) ExecX(ctx context.Context) {
	if err := ibtibhc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ibtibhc *ItemBatchToItemBatchHistoryCreate) defaults() {
	if _, ok := ibtibhc.mutation.HistoryTime(); !ok {
		v := itembatchtoitembatchhistory.DefaultHistoryTime()
		ibtibhc.mutation.SetHistoryTime(v)
	}
	if _, ok := ibtibhc.mutation.ID(); !ok {
		v := itembatchtoitembatchhistory.DefaultID()
		ibtibhc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ibtibhc *ItemBatchToItemBatchHistoryCreate) check() error {
	if _, ok := ibtibhc.mutation.HistoryTime(); !ok {
		return &ValidationError{Name: "history_time", err: errors.New(`ent: missing required field "ItemBatchToItemBatchHistory.history_time"`)}
	}
	if _, ok := ibtibhc.mutation.Operation(); !ok {
		return &ValidationError{Name: "operation", err: errors.New(`ent: missing required field "ItemBatchToItemBatchHistory.operation"`)}
	}
	if v, ok := ibtibhc.mutation.Operation(); ok {
		if err := itembatchtoitembatchhistory.OperationValidator(v); err != nil {
			return &ValidationError{Name: "operation", err: fmt.Errorf(`ent: validator failed for field "ItemBatchToItemBatchHistory.operation": %w`, err)}
		}
	}
	if _, ok := ibtibhc.mutation.ChildUUID(); !ok {
		return &ValidationError{Name: "childUUID", err: errors.New(`ent: missing required field "ItemBatchToItemBatchHistory.childUUID"`)}
	}
	if _, ok := ibtibhc.mutation.ParentUUID(); !ok {
		return &ValidationError{Name: "parentUUID", err: errors.New(`ent: missing required field "ItemBatchToItemBatchHistory.parentUUID"`)}
	}
	return nil
}

func (ibtibhc *ItemBatchToItemBatchHistoryCreate) sqlSave(ctx context.Context) (*ItemBatchToItemBatchHistory, error) {
	if err := ibtibhc.check(); err != nil {
		return nil, err
	}
	_node, _spec := ibtibhc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ibtibhc.driver, _spec); err != nil {
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
	ibtibhc.mutation.id = &_node.ID
	ibtibhc.mutation.done = true
	return _node, nil
}

func (ibtibhc *ItemBatchToItemBatchHistoryCreate) createSpec() (*ItemBatchToItemBatchHistory, *sqlgraph.CreateSpec) {
	var (
		_node = &ItemBatchToItemBatchHistory{config: ibtibhc.config}
		_spec = sqlgraph.NewCreateSpec(itembatchtoitembatchhistory.Table, sqlgraph.NewFieldSpec(itembatchtoitembatchhistory.FieldID, field.TypeUUID))
	)
	_spec.OnConflict = ibtibhc.conflict
	if id, ok := ibtibhc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := ibtibhc.mutation.HistoryTime(); ok {
		_spec.SetField(itembatchtoitembatchhistory.FieldHistoryTime, field.TypeTime, value)
		_node.HistoryTime = value
	}
	if value, ok := ibtibhc.mutation.Ref(); ok {
		_spec.SetField(itembatchtoitembatchhistory.FieldRef, field.TypeUUID, value)
		_node.Ref = value
	}
	if value, ok := ibtibhc.mutation.Operation(); ok {
		_spec.SetField(itembatchtoitembatchhistory.FieldOperation, field.TypeEnum, value)
		_node.Operation = value
	}
	if value, ok := ibtibhc.mutation.UpdatedBy(); ok {
		_spec.SetField(itembatchtoitembatchhistory.FieldUpdatedBy, field.TypeString, value)
		_node.UpdatedBy = &value
	}
	if value, ok := ibtibhc.mutation.ChildUUID(); ok {
		_spec.SetField(itembatchtoitembatchhistory.FieldChildUUID, field.TypeUUID, value)
		_node.ChildUUID = value
	}
	if value, ok := ibtibhc.mutation.ParentUUID(); ok {
		_spec.SetField(itembatchtoitembatchhistory.FieldParentUUID, field.TypeUUID, value)
		_node.ParentUUID = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.ItemBatchToItemBatchHistory.Create().
//		SetHistoryTime(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ItemBatchToItemBatchHistoryUpsert) {
//			SetHistoryTime(v+v).
//		}).
//		Exec(ctx)
func (ibtibhc *ItemBatchToItemBatchHistoryCreate) OnConflict(opts ...sql.ConflictOption) *ItemBatchToItemBatchHistoryUpsertOne {
	ibtibhc.conflict = opts
	return &ItemBatchToItemBatchHistoryUpsertOne{
		create: ibtibhc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.ItemBatchToItemBatchHistory.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ibtibhc *ItemBatchToItemBatchHistoryCreate) OnConflictColumns(columns ...string) *ItemBatchToItemBatchHistoryUpsertOne {
	ibtibhc.conflict = append(ibtibhc.conflict, sql.ConflictColumns(columns...))
	return &ItemBatchToItemBatchHistoryUpsertOne{
		create: ibtibhc,
	}
}

type (
	// ItemBatchToItemBatchHistoryUpsertOne is the builder for "upsert"-ing
	//  one ItemBatchToItemBatchHistory node.
	ItemBatchToItemBatchHistoryUpsertOne struct {
		create *ItemBatchToItemBatchHistoryCreate
	}

	// ItemBatchToItemBatchHistoryUpsert is the "OnConflict" setter.
	ItemBatchToItemBatchHistoryUpsert struct {
		*sql.UpdateSet
	}
)

// SetChildUUID sets the "childUUID" field.
func (u *ItemBatchToItemBatchHistoryUpsert) SetChildUUID(v uuid.UUID) *ItemBatchToItemBatchHistoryUpsert {
	u.Set(itembatchtoitembatchhistory.FieldChildUUID, v)
	return u
}

// UpdateChildUUID sets the "childUUID" field to the value that was provided on create.
func (u *ItemBatchToItemBatchHistoryUpsert) UpdateChildUUID() *ItemBatchToItemBatchHistoryUpsert {
	u.SetExcluded(itembatchtoitembatchhistory.FieldChildUUID)
	return u
}

// SetParentUUID sets the "parentUUID" field.
func (u *ItemBatchToItemBatchHistoryUpsert) SetParentUUID(v uuid.UUID) *ItemBatchToItemBatchHistoryUpsert {
	u.Set(itembatchtoitembatchhistory.FieldParentUUID, v)
	return u
}

// UpdateParentUUID sets the "parentUUID" field to the value that was provided on create.
func (u *ItemBatchToItemBatchHistoryUpsert) UpdateParentUUID() *ItemBatchToItemBatchHistoryUpsert {
	u.SetExcluded(itembatchtoitembatchhistory.FieldParentUUID)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.ItemBatchToItemBatchHistory.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(itembatchtoitembatchhistory.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *ItemBatchToItemBatchHistoryUpsertOne) UpdateNewValues() *ItemBatchToItemBatchHistoryUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(itembatchtoitembatchhistory.FieldID)
		}
		if _, exists := u.create.mutation.HistoryTime(); exists {
			s.SetIgnore(itembatchtoitembatchhistory.FieldHistoryTime)
		}
		if _, exists := u.create.mutation.Ref(); exists {
			s.SetIgnore(itembatchtoitembatchhistory.FieldRef)
		}
		if _, exists := u.create.mutation.Operation(); exists {
			s.SetIgnore(itembatchtoitembatchhistory.FieldOperation)
		}
		if _, exists := u.create.mutation.UpdatedBy(); exists {
			s.SetIgnore(itembatchtoitembatchhistory.FieldUpdatedBy)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.ItemBatchToItemBatchHistory.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *ItemBatchToItemBatchHistoryUpsertOne) Ignore() *ItemBatchToItemBatchHistoryUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ItemBatchToItemBatchHistoryUpsertOne) DoNothing() *ItemBatchToItemBatchHistoryUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ItemBatchToItemBatchHistoryCreate.OnConflict
// documentation for more info.
func (u *ItemBatchToItemBatchHistoryUpsertOne) Update(set func(*ItemBatchToItemBatchHistoryUpsert)) *ItemBatchToItemBatchHistoryUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ItemBatchToItemBatchHistoryUpsert{UpdateSet: update})
	}))
	return u
}

// SetChildUUID sets the "childUUID" field.
func (u *ItemBatchToItemBatchHistoryUpsertOne) SetChildUUID(v uuid.UUID) *ItemBatchToItemBatchHistoryUpsertOne {
	return u.Update(func(s *ItemBatchToItemBatchHistoryUpsert) {
		s.SetChildUUID(v)
	})
}

// UpdateChildUUID sets the "childUUID" field to the value that was provided on create.
func (u *ItemBatchToItemBatchHistoryUpsertOne) UpdateChildUUID() *ItemBatchToItemBatchHistoryUpsertOne {
	return u.Update(func(s *ItemBatchToItemBatchHistoryUpsert) {
		s.UpdateChildUUID()
	})
}

// SetParentUUID sets the "parentUUID" field.
func (u *ItemBatchToItemBatchHistoryUpsertOne) SetParentUUID(v uuid.UUID) *ItemBatchToItemBatchHistoryUpsertOne {
	return u.Update(func(s *ItemBatchToItemBatchHistoryUpsert) {
		s.SetParentUUID(v)
	})
}

// UpdateParentUUID sets the "parentUUID" field to the value that was provided on create.
func (u *ItemBatchToItemBatchHistoryUpsertOne) UpdateParentUUID() *ItemBatchToItemBatchHistoryUpsertOne {
	return u.Update(func(s *ItemBatchToItemBatchHistoryUpsert) {
		s.UpdateParentUUID()
	})
}

// Exec executes the query.
func (u *ItemBatchToItemBatchHistoryUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ItemBatchToItemBatchHistoryCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ItemBatchToItemBatchHistoryUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *ItemBatchToItemBatchHistoryUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: ItemBatchToItemBatchHistoryUpsertOne.ID is not supported by MySQL driver. Use ItemBatchToItemBatchHistoryUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *ItemBatchToItemBatchHistoryUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// ItemBatchToItemBatchHistoryCreateBulk is the builder for creating many ItemBatchToItemBatchHistory entities in bulk.
type ItemBatchToItemBatchHistoryCreateBulk struct {
	config
	builders []*ItemBatchToItemBatchHistoryCreate
	conflict []sql.ConflictOption
}

// Save creates the ItemBatchToItemBatchHistory entities in the database.
func (ibtibhcb *ItemBatchToItemBatchHistoryCreateBulk) Save(ctx context.Context) ([]*ItemBatchToItemBatchHistory, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ibtibhcb.builders))
	nodes := make([]*ItemBatchToItemBatchHistory, len(ibtibhcb.builders))
	mutators := make([]Mutator, len(ibtibhcb.builders))
	for i := range ibtibhcb.builders {
		func(i int, root context.Context) {
			builder := ibtibhcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ItemBatchToItemBatchHistoryMutation)
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
					_, err = mutators[i+1].Mutate(root, ibtibhcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ibtibhcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ibtibhcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ibtibhcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ibtibhcb *ItemBatchToItemBatchHistoryCreateBulk) SaveX(ctx context.Context) []*ItemBatchToItemBatchHistory {
	v, err := ibtibhcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ibtibhcb *ItemBatchToItemBatchHistoryCreateBulk) Exec(ctx context.Context) error {
	_, err := ibtibhcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ibtibhcb *ItemBatchToItemBatchHistoryCreateBulk) ExecX(ctx context.Context) {
	if err := ibtibhcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.ItemBatchToItemBatchHistory.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ItemBatchToItemBatchHistoryUpsert) {
//			SetHistoryTime(v+v).
//		}).
//		Exec(ctx)
func (ibtibhcb *ItemBatchToItemBatchHistoryCreateBulk) OnConflict(opts ...sql.ConflictOption) *ItemBatchToItemBatchHistoryUpsertBulk {
	ibtibhcb.conflict = opts
	return &ItemBatchToItemBatchHistoryUpsertBulk{
		create: ibtibhcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.ItemBatchToItemBatchHistory.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ibtibhcb *ItemBatchToItemBatchHistoryCreateBulk) OnConflictColumns(columns ...string) *ItemBatchToItemBatchHistoryUpsertBulk {
	ibtibhcb.conflict = append(ibtibhcb.conflict, sql.ConflictColumns(columns...))
	return &ItemBatchToItemBatchHistoryUpsertBulk{
		create: ibtibhcb,
	}
}

// ItemBatchToItemBatchHistoryUpsertBulk is the builder for "upsert"-ing
// a bulk of ItemBatchToItemBatchHistory nodes.
type ItemBatchToItemBatchHistoryUpsertBulk struct {
	create *ItemBatchToItemBatchHistoryCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.ItemBatchToItemBatchHistory.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(itembatchtoitembatchhistory.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *ItemBatchToItemBatchHistoryUpsertBulk) UpdateNewValues() *ItemBatchToItemBatchHistoryUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(itembatchtoitembatchhistory.FieldID)
			}
			if _, exists := b.mutation.HistoryTime(); exists {
				s.SetIgnore(itembatchtoitembatchhistory.FieldHistoryTime)
			}
			if _, exists := b.mutation.Ref(); exists {
				s.SetIgnore(itembatchtoitembatchhistory.FieldRef)
			}
			if _, exists := b.mutation.Operation(); exists {
				s.SetIgnore(itembatchtoitembatchhistory.FieldOperation)
			}
			if _, exists := b.mutation.UpdatedBy(); exists {
				s.SetIgnore(itembatchtoitembatchhistory.FieldUpdatedBy)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.ItemBatchToItemBatchHistory.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *ItemBatchToItemBatchHistoryUpsertBulk) Ignore() *ItemBatchToItemBatchHistoryUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ItemBatchToItemBatchHistoryUpsertBulk) DoNothing() *ItemBatchToItemBatchHistoryUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ItemBatchToItemBatchHistoryCreateBulk.OnConflict
// documentation for more info.
func (u *ItemBatchToItemBatchHistoryUpsertBulk) Update(set func(*ItemBatchToItemBatchHistoryUpsert)) *ItemBatchToItemBatchHistoryUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ItemBatchToItemBatchHistoryUpsert{UpdateSet: update})
	}))
	return u
}

// SetChildUUID sets the "childUUID" field.
func (u *ItemBatchToItemBatchHistoryUpsertBulk) SetChildUUID(v uuid.UUID) *ItemBatchToItemBatchHistoryUpsertBulk {
	return u.Update(func(s *ItemBatchToItemBatchHistoryUpsert) {
		s.SetChildUUID(v)
	})
}

// UpdateChildUUID sets the "childUUID" field to the value that was provided on create.
func (u *ItemBatchToItemBatchHistoryUpsertBulk) UpdateChildUUID() *ItemBatchToItemBatchHistoryUpsertBulk {
	return u.Update(func(s *ItemBatchToItemBatchHistoryUpsert) {
		s.UpdateChildUUID()
	})
}

// SetParentUUID sets the "parentUUID" field.
func (u *ItemBatchToItemBatchHistoryUpsertBulk) SetParentUUID(v uuid.UUID) *ItemBatchToItemBatchHistoryUpsertBulk {
	return u.Update(func(s *ItemBatchToItemBatchHistoryUpsert) {
		s.SetParentUUID(v)
	})
}

// UpdateParentUUID sets the "parentUUID" field to the value that was provided on create.
func (u *ItemBatchToItemBatchHistoryUpsertBulk) UpdateParentUUID() *ItemBatchToItemBatchHistoryUpsertBulk {
	return u.Update(func(s *ItemBatchToItemBatchHistoryUpsert) {
		s.UpdateParentUUID()
	})
}

// Exec executes the query.
func (u *ItemBatchToItemBatchHistoryUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the ItemBatchToItemBatchHistoryCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ItemBatchToItemBatchHistoryCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ItemBatchToItemBatchHistoryUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}