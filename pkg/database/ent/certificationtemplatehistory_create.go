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

	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent/certificationtemplatehistory"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/enthistory"
)

// CertificationTemplateHistoryCreate is the builder for creating a CertificationTemplateHistory entity.
type CertificationTemplateHistoryCreate struct {
	config
	mutation *CertificationTemplateHistoryMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetHistoryTime sets the "history_time" field.
func (cthc *CertificationTemplateHistoryCreate) SetHistoryTime(t time.Time) *CertificationTemplateHistoryCreate {
	cthc.mutation.SetHistoryTime(t)
	return cthc
}

// SetNillableHistoryTime sets the "history_time" field if the given value is not nil.
func (cthc *CertificationTemplateHistoryCreate) SetNillableHistoryTime(t *time.Time) *CertificationTemplateHistoryCreate {
	if t != nil {
		cthc.SetHistoryTime(*t)
	}
	return cthc
}

// SetRef sets the "ref" field.
func (cthc *CertificationTemplateHistoryCreate) SetRef(u uuid.UUID) *CertificationTemplateHistoryCreate {
	cthc.mutation.SetRef(u)
	return cthc
}

// SetNillableRef sets the "ref" field if the given value is not nil.
func (cthc *CertificationTemplateHistoryCreate) SetNillableRef(u *uuid.UUID) *CertificationTemplateHistoryCreate {
	if u != nil {
		cthc.SetRef(*u)
	}
	return cthc
}

// SetOperation sets the "operation" field.
func (cthc *CertificationTemplateHistoryCreate) SetOperation(et enthistory.OpType) *CertificationTemplateHistoryCreate {
	cthc.mutation.SetOperation(et)
	return cthc
}

// SetUpdatedBy sets the "updated_by" field.
func (cthc *CertificationTemplateHistoryCreate) SetUpdatedBy(s string) *CertificationTemplateHistoryCreate {
	cthc.mutation.SetUpdatedBy(s)
	return cthc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (cthc *CertificationTemplateHistoryCreate) SetNillableUpdatedBy(s *string) *CertificationTemplateHistoryCreate {
	if s != nil {
		cthc.SetUpdatedBy(*s)
	}
	return cthc
}

// SetDescription sets the "description" field.
func (cthc *CertificationTemplateHistoryCreate) SetDescription(s string) *CertificationTemplateHistoryCreate {
	cthc.mutation.SetDescription(s)
	return cthc
}

// SetCompanyUUID sets the "companyUUID" field.
func (cthc *CertificationTemplateHistoryCreate) SetCompanyUUID(u uuid.UUID) *CertificationTemplateHistoryCreate {
	cthc.mutation.SetCompanyUUID(u)
	return cthc
}

// SetID sets the "id" field.
func (cthc *CertificationTemplateHistoryCreate) SetID(u uuid.UUID) *CertificationTemplateHistoryCreate {
	cthc.mutation.SetID(u)
	return cthc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (cthc *CertificationTemplateHistoryCreate) SetNillableID(u *uuid.UUID) *CertificationTemplateHistoryCreate {
	if u != nil {
		cthc.SetID(*u)
	}
	return cthc
}

// Mutation returns the CertificationTemplateHistoryMutation object of the builder.
func (cthc *CertificationTemplateHistoryCreate) Mutation() *CertificationTemplateHistoryMutation {
	return cthc.mutation
}

// Save creates the CertificationTemplateHistory in the database.
func (cthc *CertificationTemplateHistoryCreate) Save(ctx context.Context) (*CertificationTemplateHistory, error) {
	cthc.defaults()
	return withHooks[*CertificationTemplateHistory, CertificationTemplateHistoryMutation](ctx, cthc.sqlSave, cthc.mutation, cthc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cthc *CertificationTemplateHistoryCreate) SaveX(ctx context.Context) *CertificationTemplateHistory {
	v, err := cthc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cthc *CertificationTemplateHistoryCreate) Exec(ctx context.Context) error {
	_, err := cthc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cthc *CertificationTemplateHistoryCreate) ExecX(ctx context.Context) {
	if err := cthc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cthc *CertificationTemplateHistoryCreate) defaults() {
	if _, ok := cthc.mutation.HistoryTime(); !ok {
		v := certificationtemplatehistory.DefaultHistoryTime()
		cthc.mutation.SetHistoryTime(v)
	}
	if _, ok := cthc.mutation.ID(); !ok {
		v := certificationtemplatehistory.DefaultID()
		cthc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cthc *CertificationTemplateHistoryCreate) check() error {
	if _, ok := cthc.mutation.HistoryTime(); !ok {
		return &ValidationError{Name: "history_time", err: errors.New(`ent: missing required field "CertificationTemplateHistory.history_time"`)}
	}
	if _, ok := cthc.mutation.Operation(); !ok {
		return &ValidationError{Name: "operation", err: errors.New(`ent: missing required field "CertificationTemplateHistory.operation"`)}
	}
	if v, ok := cthc.mutation.Operation(); ok {
		if err := certificationtemplatehistory.OperationValidator(v); err != nil {
			return &ValidationError{Name: "operation", err: fmt.Errorf(`ent: validator failed for field "CertificationTemplateHistory.operation": %w`, err)}
		}
	}
	if _, ok := cthc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "CertificationTemplateHistory.description"`)}
	}
	if _, ok := cthc.mutation.CompanyUUID(); !ok {
		return &ValidationError{Name: "companyUUID", err: errors.New(`ent: missing required field "CertificationTemplateHistory.companyUUID"`)}
	}
	return nil
}

func (cthc *CertificationTemplateHistoryCreate) sqlSave(ctx context.Context) (*CertificationTemplateHistory, error) {
	if err := cthc.check(); err != nil {
		return nil, err
	}
	_node, _spec := cthc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cthc.driver, _spec); err != nil {
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
	cthc.mutation.id = &_node.ID
	cthc.mutation.done = true
	return _node, nil
}

func (cthc *CertificationTemplateHistoryCreate) createSpec() (*CertificationTemplateHistory, *sqlgraph.CreateSpec) {
	var (
		_node = &CertificationTemplateHistory{config: cthc.config}
		_spec = sqlgraph.NewCreateSpec(certificationtemplatehistory.Table, sqlgraph.NewFieldSpec(certificationtemplatehistory.FieldID, field.TypeUUID))
	)
	_spec.OnConflict = cthc.conflict
	if id, ok := cthc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := cthc.mutation.HistoryTime(); ok {
		_spec.SetField(certificationtemplatehistory.FieldHistoryTime, field.TypeTime, value)
		_node.HistoryTime = value
	}
	if value, ok := cthc.mutation.Ref(); ok {
		_spec.SetField(certificationtemplatehistory.FieldRef, field.TypeUUID, value)
		_node.Ref = value
	}
	if value, ok := cthc.mutation.Operation(); ok {
		_spec.SetField(certificationtemplatehistory.FieldOperation, field.TypeEnum, value)
		_node.Operation = value
	}
	if value, ok := cthc.mutation.UpdatedBy(); ok {
		_spec.SetField(certificationtemplatehistory.FieldUpdatedBy, field.TypeString, value)
		_node.UpdatedBy = &value
	}
	if value, ok := cthc.mutation.Description(); ok {
		_spec.SetField(certificationtemplatehistory.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := cthc.mutation.CompanyUUID(); ok {
		_spec.SetField(certificationtemplatehistory.FieldCompanyUUID, field.TypeUUID, value)
		_node.CompanyUUID = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.CertificationTemplateHistory.Create().
//		SetHistoryTime(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.CertificationTemplateHistoryUpsert) {
//			SetHistoryTime(v+v).
//		}).
//		Exec(ctx)
func (cthc *CertificationTemplateHistoryCreate) OnConflict(opts ...sql.ConflictOption) *CertificationTemplateHistoryUpsertOne {
	cthc.conflict = opts
	return &CertificationTemplateHistoryUpsertOne{
		create: cthc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.CertificationTemplateHistory.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (cthc *CertificationTemplateHistoryCreate) OnConflictColumns(columns ...string) *CertificationTemplateHistoryUpsertOne {
	cthc.conflict = append(cthc.conflict, sql.ConflictColumns(columns...))
	return &CertificationTemplateHistoryUpsertOne{
		create: cthc,
	}
}

type (
	// CertificationTemplateHistoryUpsertOne is the builder for "upsert"-ing
	//  one CertificationTemplateHistory node.
	CertificationTemplateHistoryUpsertOne struct {
		create *CertificationTemplateHistoryCreate
	}

	// CertificationTemplateHistoryUpsert is the "OnConflict" setter.
	CertificationTemplateHistoryUpsert struct {
		*sql.UpdateSet
	}
)

// SetDescription sets the "description" field.
func (u *CertificationTemplateHistoryUpsert) SetDescription(v string) *CertificationTemplateHistoryUpsert {
	u.Set(certificationtemplatehistory.FieldDescription, v)
	return u
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *CertificationTemplateHistoryUpsert) UpdateDescription() *CertificationTemplateHistoryUpsert {
	u.SetExcluded(certificationtemplatehistory.FieldDescription)
	return u
}

// SetCompanyUUID sets the "companyUUID" field.
func (u *CertificationTemplateHistoryUpsert) SetCompanyUUID(v uuid.UUID) *CertificationTemplateHistoryUpsert {
	u.Set(certificationtemplatehistory.FieldCompanyUUID, v)
	return u
}

// UpdateCompanyUUID sets the "companyUUID" field to the value that was provided on create.
func (u *CertificationTemplateHistoryUpsert) UpdateCompanyUUID() *CertificationTemplateHistoryUpsert {
	u.SetExcluded(certificationtemplatehistory.FieldCompanyUUID)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.CertificationTemplateHistory.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(certificationtemplatehistory.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *CertificationTemplateHistoryUpsertOne) UpdateNewValues() *CertificationTemplateHistoryUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(certificationtemplatehistory.FieldID)
		}
		if _, exists := u.create.mutation.HistoryTime(); exists {
			s.SetIgnore(certificationtemplatehistory.FieldHistoryTime)
		}
		if _, exists := u.create.mutation.Ref(); exists {
			s.SetIgnore(certificationtemplatehistory.FieldRef)
		}
		if _, exists := u.create.mutation.Operation(); exists {
			s.SetIgnore(certificationtemplatehistory.FieldOperation)
		}
		if _, exists := u.create.mutation.UpdatedBy(); exists {
			s.SetIgnore(certificationtemplatehistory.FieldUpdatedBy)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.CertificationTemplateHistory.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *CertificationTemplateHistoryUpsertOne) Ignore() *CertificationTemplateHistoryUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *CertificationTemplateHistoryUpsertOne) DoNothing() *CertificationTemplateHistoryUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the CertificationTemplateHistoryCreate.OnConflict
// documentation for more info.
func (u *CertificationTemplateHistoryUpsertOne) Update(set func(*CertificationTemplateHistoryUpsert)) *CertificationTemplateHistoryUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&CertificationTemplateHistoryUpsert{UpdateSet: update})
	}))
	return u
}

// SetDescription sets the "description" field.
func (u *CertificationTemplateHistoryUpsertOne) SetDescription(v string) *CertificationTemplateHistoryUpsertOne {
	return u.Update(func(s *CertificationTemplateHistoryUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *CertificationTemplateHistoryUpsertOne) UpdateDescription() *CertificationTemplateHistoryUpsertOne {
	return u.Update(func(s *CertificationTemplateHistoryUpsert) {
		s.UpdateDescription()
	})
}

// SetCompanyUUID sets the "companyUUID" field.
func (u *CertificationTemplateHistoryUpsertOne) SetCompanyUUID(v uuid.UUID) *CertificationTemplateHistoryUpsertOne {
	return u.Update(func(s *CertificationTemplateHistoryUpsert) {
		s.SetCompanyUUID(v)
	})
}

// UpdateCompanyUUID sets the "companyUUID" field to the value that was provided on create.
func (u *CertificationTemplateHistoryUpsertOne) UpdateCompanyUUID() *CertificationTemplateHistoryUpsertOne {
	return u.Update(func(s *CertificationTemplateHistoryUpsert) {
		s.UpdateCompanyUUID()
	})
}

// Exec executes the query.
func (u *CertificationTemplateHistoryUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for CertificationTemplateHistoryCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *CertificationTemplateHistoryUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *CertificationTemplateHistoryUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: CertificationTemplateHistoryUpsertOne.ID is not supported by MySQL driver. Use CertificationTemplateHistoryUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *CertificationTemplateHistoryUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// CertificationTemplateHistoryCreateBulk is the builder for creating many CertificationTemplateHistory entities in bulk.
type CertificationTemplateHistoryCreateBulk struct {
	config
	builders []*CertificationTemplateHistoryCreate
	conflict []sql.ConflictOption
}

// Save creates the CertificationTemplateHistory entities in the database.
func (cthcb *CertificationTemplateHistoryCreateBulk) Save(ctx context.Context) ([]*CertificationTemplateHistory, error) {
	specs := make([]*sqlgraph.CreateSpec, len(cthcb.builders))
	nodes := make([]*CertificationTemplateHistory, len(cthcb.builders))
	mutators := make([]Mutator, len(cthcb.builders))
	for i := range cthcb.builders {
		func(i int, root context.Context) {
			builder := cthcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CertificationTemplateHistoryMutation)
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
					_, err = mutators[i+1].Mutate(root, cthcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = cthcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, cthcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, cthcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (cthcb *CertificationTemplateHistoryCreateBulk) SaveX(ctx context.Context) []*CertificationTemplateHistory {
	v, err := cthcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cthcb *CertificationTemplateHistoryCreateBulk) Exec(ctx context.Context) error {
	_, err := cthcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cthcb *CertificationTemplateHistoryCreateBulk) ExecX(ctx context.Context) {
	if err := cthcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.CertificationTemplateHistory.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.CertificationTemplateHistoryUpsert) {
//			SetHistoryTime(v+v).
//		}).
//		Exec(ctx)
func (cthcb *CertificationTemplateHistoryCreateBulk) OnConflict(opts ...sql.ConflictOption) *CertificationTemplateHistoryUpsertBulk {
	cthcb.conflict = opts
	return &CertificationTemplateHistoryUpsertBulk{
		create: cthcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.CertificationTemplateHistory.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (cthcb *CertificationTemplateHistoryCreateBulk) OnConflictColumns(columns ...string) *CertificationTemplateHistoryUpsertBulk {
	cthcb.conflict = append(cthcb.conflict, sql.ConflictColumns(columns...))
	return &CertificationTemplateHistoryUpsertBulk{
		create: cthcb,
	}
}

// CertificationTemplateHistoryUpsertBulk is the builder for "upsert"-ing
// a bulk of CertificationTemplateHistory nodes.
type CertificationTemplateHistoryUpsertBulk struct {
	create *CertificationTemplateHistoryCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.CertificationTemplateHistory.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(certificationtemplatehistory.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *CertificationTemplateHistoryUpsertBulk) UpdateNewValues() *CertificationTemplateHistoryUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(certificationtemplatehistory.FieldID)
			}
			if _, exists := b.mutation.HistoryTime(); exists {
				s.SetIgnore(certificationtemplatehistory.FieldHistoryTime)
			}
			if _, exists := b.mutation.Ref(); exists {
				s.SetIgnore(certificationtemplatehistory.FieldRef)
			}
			if _, exists := b.mutation.Operation(); exists {
				s.SetIgnore(certificationtemplatehistory.FieldOperation)
			}
			if _, exists := b.mutation.UpdatedBy(); exists {
				s.SetIgnore(certificationtemplatehistory.FieldUpdatedBy)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.CertificationTemplateHistory.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *CertificationTemplateHistoryUpsertBulk) Ignore() *CertificationTemplateHistoryUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *CertificationTemplateHistoryUpsertBulk) DoNothing() *CertificationTemplateHistoryUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the CertificationTemplateHistoryCreateBulk.OnConflict
// documentation for more info.
func (u *CertificationTemplateHistoryUpsertBulk) Update(set func(*CertificationTemplateHistoryUpsert)) *CertificationTemplateHistoryUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&CertificationTemplateHistoryUpsert{UpdateSet: update})
	}))
	return u
}

// SetDescription sets the "description" field.
func (u *CertificationTemplateHistoryUpsertBulk) SetDescription(v string) *CertificationTemplateHistoryUpsertBulk {
	return u.Update(func(s *CertificationTemplateHistoryUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *CertificationTemplateHistoryUpsertBulk) UpdateDescription() *CertificationTemplateHistoryUpsertBulk {
	return u.Update(func(s *CertificationTemplateHistoryUpsert) {
		s.UpdateDescription()
	})
}

// SetCompanyUUID sets the "companyUUID" field.
func (u *CertificationTemplateHistoryUpsertBulk) SetCompanyUUID(v uuid.UUID) *CertificationTemplateHistoryUpsertBulk {
	return u.Update(func(s *CertificationTemplateHistoryUpsert) {
		s.SetCompanyUUID(v)
	})
}

// UpdateCompanyUUID sets the "companyUUID" field to the value that was provided on create.
func (u *CertificationTemplateHistoryUpsertBulk) UpdateCompanyUUID() *CertificationTemplateHistoryUpsertBulk {
	return u.Update(func(s *CertificationTemplateHistoryUpsert) {
		s.UpdateCompanyUUID()
	})
}

// Exec executes the query.
func (u *CertificationTemplateHistoryUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the CertificationTemplateHistoryCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for CertificationTemplateHistoryCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *CertificationTemplateHistoryUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
