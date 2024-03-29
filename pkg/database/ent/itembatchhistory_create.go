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

	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent/itembatchhistory"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/enthistory"
)

// ItemBatchHistoryCreate is the builder for creating a ItemBatchHistory entity.
type ItemBatchHistoryCreate struct {
	config
	mutation *ItemBatchHistoryMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetHistoryTime sets the "history_time" field.
func (ibhc *ItemBatchHistoryCreate) SetHistoryTime(t time.Time) *ItemBatchHistoryCreate {
	ibhc.mutation.SetHistoryTime(t)
	return ibhc
}

// SetNillableHistoryTime sets the "history_time" field if the given value is not nil.
func (ibhc *ItemBatchHistoryCreate) SetNillableHistoryTime(t *time.Time) *ItemBatchHistoryCreate {
	if t != nil {
		ibhc.SetHistoryTime(*t)
	}
	return ibhc
}

// SetRef sets the "ref" field.
func (ibhc *ItemBatchHistoryCreate) SetRef(u uuid.UUID) *ItemBatchHistoryCreate {
	ibhc.mutation.SetRef(u)
	return ibhc
}

// SetNillableRef sets the "ref" field if the given value is not nil.
func (ibhc *ItemBatchHistoryCreate) SetNillableRef(u *uuid.UUID) *ItemBatchHistoryCreate {
	if u != nil {
		ibhc.SetRef(*u)
	}
	return ibhc
}

// SetOperation sets the "operation" field.
func (ibhc *ItemBatchHistoryCreate) SetOperation(et enthistory.OpType) *ItemBatchHistoryCreate {
	ibhc.mutation.SetOperation(et)
	return ibhc
}

// SetUpdatedBy sets the "updated_by" field.
func (ibhc *ItemBatchHistoryCreate) SetUpdatedBy(s string) *ItemBatchHistoryCreate {
	ibhc.mutation.SetUpdatedBy(s)
	return ibhc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (ibhc *ItemBatchHistoryCreate) SetNillableUpdatedBy(s *string) *ItemBatchHistoryCreate {
	if s != nil {
		ibhc.SetUpdatedBy(*s)
	}
	return ibhc
}

// SetItemNumber sets the "itemNumber" field.
func (ibhc *ItemBatchHistoryCreate) SetItemNumber(s string) *ItemBatchHistoryCreate {
	ibhc.mutation.SetItemNumber(s)
	return ibhc
}

// SetDescription sets the "description" field.
func (ibhc *ItemBatchHistoryCreate) SetDescription(s string) *ItemBatchHistoryCreate {
	ibhc.mutation.SetDescription(s)
	return ibhc
}

// SetCompanyUUID sets the "companyUUID" field.
func (ibhc *ItemBatchHistoryCreate) SetCompanyUUID(u uuid.UUID) *ItemBatchHistoryCreate {
	ibhc.mutation.SetCompanyUUID(u)
	return ibhc
}

// SetID sets the "id" field.
func (ibhc *ItemBatchHistoryCreate) SetID(u uuid.UUID) *ItemBatchHistoryCreate {
	ibhc.mutation.SetID(u)
	return ibhc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ibhc *ItemBatchHistoryCreate) SetNillableID(u *uuid.UUID) *ItemBatchHistoryCreate {
	if u != nil {
		ibhc.SetID(*u)
	}
	return ibhc
}

// Mutation returns the ItemBatchHistoryMutation object of the builder.
func (ibhc *ItemBatchHistoryCreate) Mutation() *ItemBatchHistoryMutation {
	return ibhc.mutation
}

// Save creates the ItemBatchHistory in the database.
func (ibhc *ItemBatchHistoryCreate) Save(ctx context.Context) (*ItemBatchHistory, error) {
	ibhc.defaults()
	return withHooks[*ItemBatchHistory, ItemBatchHistoryMutation](ctx, ibhc.sqlSave, ibhc.mutation, ibhc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ibhc *ItemBatchHistoryCreate) SaveX(ctx context.Context) *ItemBatchHistory {
	v, err := ibhc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ibhc *ItemBatchHistoryCreate) Exec(ctx context.Context) error {
	_, err := ibhc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ibhc *ItemBatchHistoryCreate) ExecX(ctx context.Context) {
	if err := ibhc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ibhc *ItemBatchHistoryCreate) defaults() {
	if _, ok := ibhc.mutation.HistoryTime(); !ok {
		v := itembatchhistory.DefaultHistoryTime()
		ibhc.mutation.SetHistoryTime(v)
	}
	if _, ok := ibhc.mutation.ID(); !ok {
		v := itembatchhistory.DefaultID()
		ibhc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ibhc *ItemBatchHistoryCreate) check() error {
	if _, ok := ibhc.mutation.HistoryTime(); !ok {
		return &ValidationError{Name: "history_time", err: errors.New(`ent: missing required field "ItemBatchHistory.history_time"`)}
	}
	if _, ok := ibhc.mutation.Operation(); !ok {
		return &ValidationError{Name: "operation", err: errors.New(`ent: missing required field "ItemBatchHistory.operation"`)}
	}
	if v, ok := ibhc.mutation.Operation(); ok {
		if err := itembatchhistory.OperationValidator(v); err != nil {
			return &ValidationError{Name: "operation", err: fmt.Errorf(`ent: validator failed for field "ItemBatchHistory.operation": %w`, err)}
		}
	}
	if _, ok := ibhc.mutation.ItemNumber(); !ok {
		return &ValidationError{Name: "itemNumber", err: errors.New(`ent: missing required field "ItemBatchHistory.itemNumber"`)}
	}
	if _, ok := ibhc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "ItemBatchHistory.description"`)}
	}
	if _, ok := ibhc.mutation.CompanyUUID(); !ok {
		return &ValidationError{Name: "companyUUID", err: errors.New(`ent: missing required field "ItemBatchHistory.companyUUID"`)}
	}
	return nil
}

func (ibhc *ItemBatchHistoryCreate) sqlSave(ctx context.Context) (*ItemBatchHistory, error) {
	if err := ibhc.check(); err != nil {
		return nil, err
	}
	_node, _spec := ibhc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ibhc.driver, _spec); err != nil {
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
	ibhc.mutation.id = &_node.ID
	ibhc.mutation.done = true
	return _node, nil
}

func (ibhc *ItemBatchHistoryCreate) createSpec() (*ItemBatchHistory, *sqlgraph.CreateSpec) {
	var (
		_node = &ItemBatchHistory{config: ibhc.config}
		_spec = sqlgraph.NewCreateSpec(itembatchhistory.Table, sqlgraph.NewFieldSpec(itembatchhistory.FieldID, field.TypeUUID))
	)
	_spec.OnConflict = ibhc.conflict
	if id, ok := ibhc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := ibhc.mutation.HistoryTime(); ok {
		_spec.SetField(itembatchhistory.FieldHistoryTime, field.TypeTime, value)
		_node.HistoryTime = value
	}
	if value, ok := ibhc.mutation.Ref(); ok {
		_spec.SetField(itembatchhistory.FieldRef, field.TypeUUID, value)
		_node.Ref = value
	}
	if value, ok := ibhc.mutation.Operation(); ok {
		_spec.SetField(itembatchhistory.FieldOperation, field.TypeEnum, value)
		_node.Operation = value
	}
	if value, ok := ibhc.mutation.UpdatedBy(); ok {
		_spec.SetField(itembatchhistory.FieldUpdatedBy, field.TypeString, value)
		_node.UpdatedBy = &value
	}
	if value, ok := ibhc.mutation.ItemNumber(); ok {
		_spec.SetField(itembatchhistory.FieldItemNumber, field.TypeString, value)
		_node.ItemNumber = value
	}
	if value, ok := ibhc.mutation.Description(); ok {
		_spec.SetField(itembatchhistory.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := ibhc.mutation.CompanyUUID(); ok {
		_spec.SetField(itembatchhistory.FieldCompanyUUID, field.TypeUUID, value)
		_node.CompanyUUID = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.ItemBatchHistory.Create().
//		SetHistoryTime(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ItemBatchHistoryUpsert) {
//			SetHistoryTime(v+v).
//		}).
//		Exec(ctx)
func (ibhc *ItemBatchHistoryCreate) OnConflict(opts ...sql.ConflictOption) *ItemBatchHistoryUpsertOne {
	ibhc.conflict = opts
	return &ItemBatchHistoryUpsertOne{
		create: ibhc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.ItemBatchHistory.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ibhc *ItemBatchHistoryCreate) OnConflictColumns(columns ...string) *ItemBatchHistoryUpsertOne {
	ibhc.conflict = append(ibhc.conflict, sql.ConflictColumns(columns...))
	return &ItemBatchHistoryUpsertOne{
		create: ibhc,
	}
}

type (
	// ItemBatchHistoryUpsertOne is the builder for "upsert"-ing
	//  one ItemBatchHistory node.
	ItemBatchHistoryUpsertOne struct {
		create *ItemBatchHistoryCreate
	}

	// ItemBatchHistoryUpsert is the "OnConflict" setter.
	ItemBatchHistoryUpsert struct {
		*sql.UpdateSet
	}
)

// SetItemNumber sets the "itemNumber" field.
func (u *ItemBatchHistoryUpsert) SetItemNumber(v string) *ItemBatchHistoryUpsert {
	u.Set(itembatchhistory.FieldItemNumber, v)
	return u
}

// UpdateItemNumber sets the "itemNumber" field to the value that was provided on create.
func (u *ItemBatchHistoryUpsert) UpdateItemNumber() *ItemBatchHistoryUpsert {
	u.SetExcluded(itembatchhistory.FieldItemNumber)
	return u
}

// SetDescription sets the "description" field.
func (u *ItemBatchHistoryUpsert) SetDescription(v string) *ItemBatchHistoryUpsert {
	u.Set(itembatchhistory.FieldDescription, v)
	return u
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *ItemBatchHistoryUpsert) UpdateDescription() *ItemBatchHistoryUpsert {
	u.SetExcluded(itembatchhistory.FieldDescription)
	return u
}

// SetCompanyUUID sets the "companyUUID" field.
func (u *ItemBatchHistoryUpsert) SetCompanyUUID(v uuid.UUID) *ItemBatchHistoryUpsert {
	u.Set(itembatchhistory.FieldCompanyUUID, v)
	return u
}

// UpdateCompanyUUID sets the "companyUUID" field to the value that was provided on create.
func (u *ItemBatchHistoryUpsert) UpdateCompanyUUID() *ItemBatchHistoryUpsert {
	u.SetExcluded(itembatchhistory.FieldCompanyUUID)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.ItemBatchHistory.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(itembatchhistory.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *ItemBatchHistoryUpsertOne) UpdateNewValues() *ItemBatchHistoryUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(itembatchhistory.FieldID)
		}
		if _, exists := u.create.mutation.HistoryTime(); exists {
			s.SetIgnore(itembatchhistory.FieldHistoryTime)
		}
		if _, exists := u.create.mutation.Ref(); exists {
			s.SetIgnore(itembatchhistory.FieldRef)
		}
		if _, exists := u.create.mutation.Operation(); exists {
			s.SetIgnore(itembatchhistory.FieldOperation)
		}
		if _, exists := u.create.mutation.UpdatedBy(); exists {
			s.SetIgnore(itembatchhistory.FieldUpdatedBy)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.ItemBatchHistory.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *ItemBatchHistoryUpsertOne) Ignore() *ItemBatchHistoryUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ItemBatchHistoryUpsertOne) DoNothing() *ItemBatchHistoryUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ItemBatchHistoryCreate.OnConflict
// documentation for more info.
func (u *ItemBatchHistoryUpsertOne) Update(set func(*ItemBatchHistoryUpsert)) *ItemBatchHistoryUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ItemBatchHistoryUpsert{UpdateSet: update})
	}))
	return u
}

// SetItemNumber sets the "itemNumber" field.
func (u *ItemBatchHistoryUpsertOne) SetItemNumber(v string) *ItemBatchHistoryUpsertOne {
	return u.Update(func(s *ItemBatchHistoryUpsert) {
		s.SetItemNumber(v)
	})
}

// UpdateItemNumber sets the "itemNumber" field to the value that was provided on create.
func (u *ItemBatchHistoryUpsertOne) UpdateItemNumber() *ItemBatchHistoryUpsertOne {
	return u.Update(func(s *ItemBatchHistoryUpsert) {
		s.UpdateItemNumber()
	})
}

// SetDescription sets the "description" field.
func (u *ItemBatchHistoryUpsertOne) SetDescription(v string) *ItemBatchHistoryUpsertOne {
	return u.Update(func(s *ItemBatchHistoryUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *ItemBatchHistoryUpsertOne) UpdateDescription() *ItemBatchHistoryUpsertOne {
	return u.Update(func(s *ItemBatchHistoryUpsert) {
		s.UpdateDescription()
	})
}

// SetCompanyUUID sets the "companyUUID" field.
func (u *ItemBatchHistoryUpsertOne) SetCompanyUUID(v uuid.UUID) *ItemBatchHistoryUpsertOne {
	return u.Update(func(s *ItemBatchHistoryUpsert) {
		s.SetCompanyUUID(v)
	})
}

// UpdateCompanyUUID sets the "companyUUID" field to the value that was provided on create.
func (u *ItemBatchHistoryUpsertOne) UpdateCompanyUUID() *ItemBatchHistoryUpsertOne {
	return u.Update(func(s *ItemBatchHistoryUpsert) {
		s.UpdateCompanyUUID()
	})
}

// Exec executes the query.
func (u *ItemBatchHistoryUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ItemBatchHistoryCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ItemBatchHistoryUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *ItemBatchHistoryUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: ItemBatchHistoryUpsertOne.ID is not supported by MySQL driver. Use ItemBatchHistoryUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *ItemBatchHistoryUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// ItemBatchHistoryCreateBulk is the builder for creating many ItemBatchHistory entities in bulk.
type ItemBatchHistoryCreateBulk struct {
	config
	builders []*ItemBatchHistoryCreate
	conflict []sql.ConflictOption
}

// Save creates the ItemBatchHistory entities in the database.
func (ibhcb *ItemBatchHistoryCreateBulk) Save(ctx context.Context) ([]*ItemBatchHistory, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ibhcb.builders))
	nodes := make([]*ItemBatchHistory, len(ibhcb.builders))
	mutators := make([]Mutator, len(ibhcb.builders))
	for i := range ibhcb.builders {
		func(i int, root context.Context) {
			builder := ibhcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ItemBatchHistoryMutation)
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
					_, err = mutators[i+1].Mutate(root, ibhcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ibhcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ibhcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ibhcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ibhcb *ItemBatchHistoryCreateBulk) SaveX(ctx context.Context) []*ItemBatchHistory {
	v, err := ibhcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ibhcb *ItemBatchHistoryCreateBulk) Exec(ctx context.Context) error {
	_, err := ibhcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ibhcb *ItemBatchHistoryCreateBulk) ExecX(ctx context.Context) {
	if err := ibhcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.ItemBatchHistory.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ItemBatchHistoryUpsert) {
//			SetHistoryTime(v+v).
//		}).
//		Exec(ctx)
func (ibhcb *ItemBatchHistoryCreateBulk) OnConflict(opts ...sql.ConflictOption) *ItemBatchHistoryUpsertBulk {
	ibhcb.conflict = opts
	return &ItemBatchHistoryUpsertBulk{
		create: ibhcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.ItemBatchHistory.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ibhcb *ItemBatchHistoryCreateBulk) OnConflictColumns(columns ...string) *ItemBatchHistoryUpsertBulk {
	ibhcb.conflict = append(ibhcb.conflict, sql.ConflictColumns(columns...))
	return &ItemBatchHistoryUpsertBulk{
		create: ibhcb,
	}
}

// ItemBatchHistoryUpsertBulk is the builder for "upsert"-ing
// a bulk of ItemBatchHistory nodes.
type ItemBatchHistoryUpsertBulk struct {
	create *ItemBatchHistoryCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.ItemBatchHistory.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(itembatchhistory.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *ItemBatchHistoryUpsertBulk) UpdateNewValues() *ItemBatchHistoryUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(itembatchhistory.FieldID)
			}
			if _, exists := b.mutation.HistoryTime(); exists {
				s.SetIgnore(itembatchhistory.FieldHistoryTime)
			}
			if _, exists := b.mutation.Ref(); exists {
				s.SetIgnore(itembatchhistory.FieldRef)
			}
			if _, exists := b.mutation.Operation(); exists {
				s.SetIgnore(itembatchhistory.FieldOperation)
			}
			if _, exists := b.mutation.UpdatedBy(); exists {
				s.SetIgnore(itembatchhistory.FieldUpdatedBy)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.ItemBatchHistory.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *ItemBatchHistoryUpsertBulk) Ignore() *ItemBatchHistoryUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ItemBatchHistoryUpsertBulk) DoNothing() *ItemBatchHistoryUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ItemBatchHistoryCreateBulk.OnConflict
// documentation for more info.
func (u *ItemBatchHistoryUpsertBulk) Update(set func(*ItemBatchHistoryUpsert)) *ItemBatchHistoryUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ItemBatchHistoryUpsert{UpdateSet: update})
	}))
	return u
}

// SetItemNumber sets the "itemNumber" field.
func (u *ItemBatchHistoryUpsertBulk) SetItemNumber(v string) *ItemBatchHistoryUpsertBulk {
	return u.Update(func(s *ItemBatchHistoryUpsert) {
		s.SetItemNumber(v)
	})
}

// UpdateItemNumber sets the "itemNumber" field to the value that was provided on create.
func (u *ItemBatchHistoryUpsertBulk) UpdateItemNumber() *ItemBatchHistoryUpsertBulk {
	return u.Update(func(s *ItemBatchHistoryUpsert) {
		s.UpdateItemNumber()
	})
}

// SetDescription sets the "description" field.
func (u *ItemBatchHistoryUpsertBulk) SetDescription(v string) *ItemBatchHistoryUpsertBulk {
	return u.Update(func(s *ItemBatchHistoryUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *ItemBatchHistoryUpsertBulk) UpdateDescription() *ItemBatchHistoryUpsertBulk {
	return u.Update(func(s *ItemBatchHistoryUpsert) {
		s.UpdateDescription()
	})
}

// SetCompanyUUID sets the "companyUUID" field.
func (u *ItemBatchHistoryUpsertBulk) SetCompanyUUID(v uuid.UUID) *ItemBatchHistoryUpsertBulk {
	return u.Update(func(s *ItemBatchHistoryUpsert) {
		s.SetCompanyUUID(v)
	})
}

// UpdateCompanyUUID sets the "companyUUID" field to the value that was provided on create.
func (u *ItemBatchHistoryUpsertBulk) UpdateCompanyUUID() *ItemBatchHistoryUpsertBulk {
	return u.Update(func(s *ItemBatchHistoryUpsert) {
		s.UpdateCompanyUUID()
	})
}

// Exec executes the query.
func (u *ItemBatchHistoryUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the ItemBatchHistoryCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ItemBatchHistoryCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ItemBatchHistoryUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
