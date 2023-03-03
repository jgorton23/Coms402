// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"

	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent/itembatch"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent/predicate"
)

// ItemBatchDelete is the builder for deleting a ItemBatch entity.
type ItemBatchDelete struct {
	config
	hooks    []Hook
	mutation *ItemBatchMutation
}

// Where appends a list predicates to the ItemBatchDelete builder.
func (ibd *ItemBatchDelete) Where(ps ...predicate.ItemBatch) *ItemBatchDelete {
	ibd.mutation.Where(ps...)
	return ibd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ibd *ItemBatchDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, ItemBatchMutation](ctx, ibd.sqlExec, ibd.mutation, ibd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ibd *ItemBatchDelete) ExecX(ctx context.Context) int {
	n, err := ibd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ibd *ItemBatchDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(itembatch.Table, sqlgraph.NewFieldSpec(itembatch.FieldID, field.TypeUUID))
	if ps := ibd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ibd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ibd.mutation.done = true
	return affected, err
}

// ItemBatchDeleteOne is the builder for deleting a single ItemBatch entity.
type ItemBatchDeleteOne struct {
	ibd *ItemBatchDelete
}

// Where appends a list predicates to the ItemBatchDelete builder.
func (ibdo *ItemBatchDeleteOne) Where(ps ...predicate.ItemBatch) *ItemBatchDeleteOne {
	ibdo.ibd.mutation.Where(ps...)
	return ibdo
}

// Exec executes the deletion query.
func (ibdo *ItemBatchDeleteOne) Exec(ctx context.Context) error {
	n, err := ibdo.ibd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{itembatch.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ibdo *ItemBatchDeleteOne) ExecX(ctx context.Context) {
	if err := ibdo.Exec(ctx); err != nil {
		panic(err)
	}
}