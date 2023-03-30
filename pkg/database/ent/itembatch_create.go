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

	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent/company"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent/itembatch"
)

// ItemBatchCreate is the builder for creating a ItemBatch entity.
type ItemBatchCreate struct {
	config
	mutation *ItemBatchMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetItemNumber sets the "itemNumber" field.
func (ibc *ItemBatchCreate) SetItemNumber(s string) *ItemBatchCreate {
	ibc.mutation.SetItemNumber(s)
	return ibc
}

// SetDescription sets the "description" field.
func (ibc *ItemBatchCreate) SetDescription(s string) *ItemBatchCreate {
	ibc.mutation.SetDescription(s)
	return ibc
}

// SetCompanyUUID sets the "companyUUID" field.
func (ibc *ItemBatchCreate) SetCompanyUUID(u uuid.UUID) *ItemBatchCreate {
	ibc.mutation.SetCompanyUUID(u)
	return ibc
}

// SetID sets the "id" field.
func (ibc *ItemBatchCreate) SetID(u uuid.UUID) *ItemBatchCreate {
	ibc.mutation.SetID(u)
	return ibc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ibc *ItemBatchCreate) SetNillableID(u *uuid.UUID) *ItemBatchCreate {
	if u != nil {
		ibc.SetID(*u)
	}
	return ibc
}

// SetCompanyID sets the "company" edge to the Company entity by ID.
func (ibc *ItemBatchCreate) SetCompanyID(id uuid.UUID) *ItemBatchCreate {
	ibc.mutation.SetCompanyID(id)
	return ibc
}

// SetCompany sets the "company" edge to the Company entity.
func (ibc *ItemBatchCreate) SetCompany(c *Company) *ItemBatchCreate {
	return ibc.SetCompanyID(c.ID)
}

// Mutation returns the ItemBatchMutation object of the builder.
func (ibc *ItemBatchCreate) Mutation() *ItemBatchMutation {
	return ibc.mutation
}

// Save creates the ItemBatch in the database.
func (ibc *ItemBatchCreate) Save(ctx context.Context) (*ItemBatch, error) {
	ibc.defaults()
	return withHooks[*ItemBatch, ItemBatchMutation](ctx, ibc.sqlSave, ibc.mutation, ibc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ibc *ItemBatchCreate) SaveX(ctx context.Context) *ItemBatch {
	v, err := ibc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ibc *ItemBatchCreate) Exec(ctx context.Context) error {
	_, err := ibc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ibc *ItemBatchCreate) ExecX(ctx context.Context) {
	if err := ibc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ibc *ItemBatchCreate) defaults() {
	if _, ok := ibc.mutation.ID(); !ok {
		v := itembatch.DefaultID()
		ibc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ibc *ItemBatchCreate) check() error {
	if _, ok := ibc.mutation.ItemNumber(); !ok {
		return &ValidationError{Name: "itemNumber", err: errors.New(`ent: missing required field "ItemBatch.itemNumber"`)}
	}
	if _, ok := ibc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "ItemBatch.description"`)}
	}
	if _, ok := ibc.mutation.CompanyUUID(); !ok {
		return &ValidationError{Name: "companyUUID", err: errors.New(`ent: missing required field "ItemBatch.companyUUID"`)}
	}
	if _, ok := ibc.mutation.CompanyID(); !ok {
		return &ValidationError{Name: "company", err: errors.New(`ent: missing required edge "ItemBatch.company"`)}
	}
	return nil
}

func (ibc *ItemBatchCreate) sqlSave(ctx context.Context) (*ItemBatch, error) {
	if err := ibc.check(); err != nil {
		return nil, err
	}
	_node, _spec := ibc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ibc.driver, _spec); err != nil {
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
	ibc.mutation.id = &_node.ID
	ibc.mutation.done = true
	return _node, nil
}

func (ibc *ItemBatchCreate) createSpec() (*ItemBatch, *sqlgraph.CreateSpec) {
	var (
		_node = &ItemBatch{config: ibc.config}
		_spec = sqlgraph.NewCreateSpec(itembatch.Table, sqlgraph.NewFieldSpec(itembatch.FieldID, field.TypeUUID))
	)
	_spec.OnConflict = ibc.conflict
	if id, ok := ibc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := ibc.mutation.ItemNumber(); ok {
		_spec.SetField(itembatch.FieldItemNumber, field.TypeString, value)
		_node.ItemNumber = value
	}
	if value, ok := ibc.mutation.Description(); ok {
		_spec.SetField(itembatch.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if nodes := ibc.mutation.CompanyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   itembatch.CompanyTable,
			Columns: []string{itembatch.CompanyColumn},
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
		_node.CompanyUUID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	httpclient.ItemBatch.Create().
//		SetItemNumber(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ItemBatchUpsert) {
//			SetItemNumber(v+v).
//		}).
//		Exec(ctx)
func (ibc *ItemBatchCreate) OnConflict(opts ...sql.ConflictOption) *ItemBatchUpsertOne {
	ibc.conflict = opts
	return &ItemBatchUpsertOne{
		create: ibc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	httpclient.ItemBatch.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ibc *ItemBatchCreate) OnConflictColumns(columns ...string) *ItemBatchUpsertOne {
	ibc.conflict = append(ibc.conflict, sql.ConflictColumns(columns...))
	return &ItemBatchUpsertOne{
		create: ibc,
	}
}

type (
	// ItemBatchUpsertOne is the builder for "upsert"-ing
	//  one ItemBatch node.
	ItemBatchUpsertOne struct {
		create *ItemBatchCreate
	}

	// ItemBatchUpsert is the "OnConflict" setter.
	ItemBatchUpsert struct {
		*sql.UpdateSet
	}
)

// SetItemNumber sets the "itemNumber" field.
func (u *ItemBatchUpsert) SetItemNumber(v string) *ItemBatchUpsert {
	u.Set(itembatch.FieldItemNumber, v)
	return u
}

// UpdateItemNumber sets the "itemNumber" field to the value that was provided on create.
func (u *ItemBatchUpsert) UpdateItemNumber() *ItemBatchUpsert {
	u.SetExcluded(itembatch.FieldItemNumber)
	return u
}

// SetDescription sets the "description" field.
func (u *ItemBatchUpsert) SetDescription(v string) *ItemBatchUpsert {
	u.Set(itembatch.FieldDescription, v)
	return u
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *ItemBatchUpsert) UpdateDescription() *ItemBatchUpsert {
	u.SetExcluded(itembatch.FieldDescription)
	return u
}

// SetCompanyUUID sets the "companyUUID" field.
func (u *ItemBatchUpsert) SetCompanyUUID(v uuid.UUID) *ItemBatchUpsert {
	u.Set(itembatch.FieldCompanyUUID, v)
	return u
}

// UpdateCompanyUUID sets the "companyUUID" field to the value that was provided on create.
func (u *ItemBatchUpsert) UpdateCompanyUUID() *ItemBatchUpsert {
	u.SetExcluded(itembatch.FieldCompanyUUID)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	httpclient.ItemBatch.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(itembatch.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *ItemBatchUpsertOne) UpdateNewValues() *ItemBatchUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(itembatch.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	httpclient.ItemBatch.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *ItemBatchUpsertOne) Ignore() *ItemBatchUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ItemBatchUpsertOne) DoNothing() *ItemBatchUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ItemBatchCreate.OnConflict
// documentation for more info.
func (u *ItemBatchUpsertOne) Update(set func(*ItemBatchUpsert)) *ItemBatchUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ItemBatchUpsert{UpdateSet: update})
	}))
	return u
}

// SetItemNumber sets the "itemNumber" field.
func (u *ItemBatchUpsertOne) SetItemNumber(v string) *ItemBatchUpsertOne {
	return u.Update(func(s *ItemBatchUpsert) {
		s.SetItemNumber(v)
	})
}

// UpdateItemNumber sets the "itemNumber" field to the value that was provided on create.
func (u *ItemBatchUpsertOne) UpdateItemNumber() *ItemBatchUpsertOne {
	return u.Update(func(s *ItemBatchUpsert) {
		s.UpdateItemNumber()
	})
}

// SetDescription sets the "description" field.
func (u *ItemBatchUpsertOne) SetDescription(v string) *ItemBatchUpsertOne {
	return u.Update(func(s *ItemBatchUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *ItemBatchUpsertOne) UpdateDescription() *ItemBatchUpsertOne {
	return u.Update(func(s *ItemBatchUpsert) {
		s.UpdateDescription()
	})
}

// SetCompanyUUID sets the "companyUUID" field.
func (u *ItemBatchUpsertOne) SetCompanyUUID(v uuid.UUID) *ItemBatchUpsertOne {
	return u.Update(func(s *ItemBatchUpsert) {
		s.SetCompanyUUID(v)
	})
}

// UpdateCompanyUUID sets the "companyUUID" field to the value that was provided on create.
func (u *ItemBatchUpsertOne) UpdateCompanyUUID() *ItemBatchUpsertOne {
	return u.Update(func(s *ItemBatchUpsert) {
		s.UpdateCompanyUUID()
	})
}

// Exec executes the query.
func (u *ItemBatchUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ItemBatchCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ItemBatchUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *ItemBatchUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: ItemBatchUpsertOne.ID is not supported by MySQL driver. Use ItemBatchUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *ItemBatchUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// ItemBatchCreateBulk is the builder for creating many ItemBatch entities in bulk.
type ItemBatchCreateBulk struct {
	config
	builders []*ItemBatchCreate
	conflict []sql.ConflictOption
}

// Save creates the ItemBatch entities in the database.
func (ibcb *ItemBatchCreateBulk) Save(ctx context.Context) ([]*ItemBatch, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ibcb.builders))
	nodes := make([]*ItemBatch, len(ibcb.builders))
	mutators := make([]Mutator, len(ibcb.builders))
	for i := range ibcb.builders {
		func(i int, root context.Context) {
			builder := ibcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ItemBatchMutation)
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
					_, err = mutators[i+1].Mutate(root, ibcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ibcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ibcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ibcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ibcb *ItemBatchCreateBulk) SaveX(ctx context.Context) []*ItemBatch {
	v, err := ibcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ibcb *ItemBatchCreateBulk) Exec(ctx context.Context) error {
	_, err := ibcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ibcb *ItemBatchCreateBulk) ExecX(ctx context.Context) {
	if err := ibcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	httpclient.ItemBatch.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ItemBatchUpsert) {
//			SetItemNumber(v+v).
//		}).
//		Exec(ctx)
func (ibcb *ItemBatchCreateBulk) OnConflict(opts ...sql.ConflictOption) *ItemBatchUpsertBulk {
	ibcb.conflict = opts
	return &ItemBatchUpsertBulk{
		create: ibcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	httpclient.ItemBatch.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ibcb *ItemBatchCreateBulk) OnConflictColumns(columns ...string) *ItemBatchUpsertBulk {
	ibcb.conflict = append(ibcb.conflict, sql.ConflictColumns(columns...))
	return &ItemBatchUpsertBulk{
		create: ibcb,
	}
}

// ItemBatchUpsertBulk is the builder for "upsert"-ing
// a bulk of ItemBatch nodes.
type ItemBatchUpsertBulk struct {
	create *ItemBatchCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	httpclient.ItemBatch.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(itembatch.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *ItemBatchUpsertBulk) UpdateNewValues() *ItemBatchUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(itembatch.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	httpclient.ItemBatch.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *ItemBatchUpsertBulk) Ignore() *ItemBatchUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ItemBatchUpsertBulk) DoNothing() *ItemBatchUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ItemBatchCreateBulk.OnConflict
// documentation for more info.
func (u *ItemBatchUpsertBulk) Update(set func(*ItemBatchUpsert)) *ItemBatchUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ItemBatchUpsert{UpdateSet: update})
	}))
	return u
}

// SetItemNumber sets the "itemNumber" field.
func (u *ItemBatchUpsertBulk) SetItemNumber(v string) *ItemBatchUpsertBulk {
	return u.Update(func(s *ItemBatchUpsert) {
		s.SetItemNumber(v)
	})
}

// UpdateItemNumber sets the "itemNumber" field to the value that was provided on create.
func (u *ItemBatchUpsertBulk) UpdateItemNumber() *ItemBatchUpsertBulk {
	return u.Update(func(s *ItemBatchUpsert) {
		s.UpdateItemNumber()
	})
}

// SetDescription sets the "description" field.
func (u *ItemBatchUpsertBulk) SetDescription(v string) *ItemBatchUpsertBulk {
	return u.Update(func(s *ItemBatchUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *ItemBatchUpsertBulk) UpdateDescription() *ItemBatchUpsertBulk {
	return u.Update(func(s *ItemBatchUpsert) {
		s.UpdateDescription()
	})
}

// SetCompanyUUID sets the "companyUUID" field.
func (u *ItemBatchUpsertBulk) SetCompanyUUID(v uuid.UUID) *ItemBatchUpsertBulk {
	return u.Update(func(s *ItemBatchUpsert) {
		s.SetCompanyUUID(v)
	})
}

// UpdateCompanyUUID sets the "companyUUID" field to the value that was provided on create.
func (u *ItemBatchUpsertBulk) UpdateCompanyUUID() *ItemBatchUpsertBulk {
	return u.Update(func(s *ItemBatchUpsert) {
		s.UpdateCompanyUUID()
	})
}

// Exec executes the query.
func (u *ItemBatchUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the ItemBatchCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ItemBatchCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ItemBatchUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
