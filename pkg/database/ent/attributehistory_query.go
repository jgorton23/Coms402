// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"

	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent/attributehistory"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent/predicate"
)

// AttributeHistoryQuery is the builder for querying AttributeHistory entities.
type AttributeHistoryQuery struct {
	config
	ctx        *QueryContext
	order      []OrderFunc
	inters     []Interceptor
	predicates []predicate.AttributeHistory
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AttributeHistoryQuery builder.
func (ahq *AttributeHistoryQuery) Where(ps ...predicate.AttributeHistory) *AttributeHistoryQuery {
	ahq.predicates = append(ahq.predicates, ps...)
	return ahq
}

// Limit the number of records to be returned by this query.
func (ahq *AttributeHistoryQuery) Limit(limit int) *AttributeHistoryQuery {
	ahq.ctx.Limit = &limit
	return ahq
}

// Offset to start from.
func (ahq *AttributeHistoryQuery) Offset(offset int) *AttributeHistoryQuery {
	ahq.ctx.Offset = &offset
	return ahq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ahq *AttributeHistoryQuery) Unique(unique bool) *AttributeHistoryQuery {
	ahq.ctx.Unique = &unique
	return ahq
}

// Order specifies how the records should be ordered.
func (ahq *AttributeHistoryQuery) Order(o ...OrderFunc) *AttributeHistoryQuery {
	ahq.order = append(ahq.order, o...)
	return ahq
}

// First returns the first AttributeHistory entity from the query.
// Returns a *NotFoundError when no AttributeHistory was found.
func (ahq *AttributeHistoryQuery) First(ctx context.Context) (*AttributeHistory, error) {
	nodes, err := ahq.Limit(1).All(setContextOp(ctx, ahq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{attributehistory.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ahq *AttributeHistoryQuery) FirstX(ctx context.Context) *AttributeHistory {
	node, err := ahq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AttributeHistory ID from the query.
// Returns a *NotFoundError when no AttributeHistory ID was found.
func (ahq *AttributeHistoryQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = ahq.Limit(1).IDs(setContextOp(ctx, ahq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{attributehistory.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ahq *AttributeHistoryQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := ahq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AttributeHistory entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one AttributeHistory entity is found.
// Returns a *NotFoundError when no AttributeHistory entities are found.
func (ahq *AttributeHistoryQuery) Only(ctx context.Context) (*AttributeHistory, error) {
	nodes, err := ahq.Limit(2).All(setContextOp(ctx, ahq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{attributehistory.Label}
	default:
		return nil, &NotSingularError{attributehistory.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ahq *AttributeHistoryQuery) OnlyX(ctx context.Context) *AttributeHistory {
	node, err := ahq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AttributeHistory ID in the query.
// Returns a *NotSingularError when more than one AttributeHistory ID is found.
// Returns a *NotFoundError when no entities are found.
func (ahq *AttributeHistoryQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = ahq.Limit(2).IDs(setContextOp(ctx, ahq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{attributehistory.Label}
	default:
		err = &NotSingularError{attributehistory.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ahq *AttributeHistoryQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := ahq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AttributeHistories.
func (ahq *AttributeHistoryQuery) All(ctx context.Context) ([]*AttributeHistory, error) {
	ctx = setContextOp(ctx, ahq.ctx, "All")
	if err := ahq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*AttributeHistory, *AttributeHistoryQuery]()
	return withInterceptors[[]*AttributeHistory](ctx, ahq, qr, ahq.inters)
}

// AllX is like All, but panics if an error occurs.
func (ahq *AttributeHistoryQuery) AllX(ctx context.Context) []*AttributeHistory {
	nodes, err := ahq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AttributeHistory IDs.
func (ahq *AttributeHistoryQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if ahq.ctx.Unique == nil && ahq.path != nil {
		ahq.Unique(true)
	}
	ctx = setContextOp(ctx, ahq.ctx, "IDs")
	if err = ahq.Select(attributehistory.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ahq *AttributeHistoryQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := ahq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ahq *AttributeHistoryQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, ahq.ctx, "Count")
	if err := ahq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, ahq, querierCount[*AttributeHistoryQuery](), ahq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (ahq *AttributeHistoryQuery) CountX(ctx context.Context) int {
	count, err := ahq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ahq *AttributeHistoryQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, ahq.ctx, "Exist")
	switch _, err := ahq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (ahq *AttributeHistoryQuery) ExistX(ctx context.Context) bool {
	exist, err := ahq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AttributeHistoryQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ahq *AttributeHistoryQuery) Clone() *AttributeHistoryQuery {
	if ahq == nil {
		return nil
	}
	return &AttributeHistoryQuery{
		config:     ahq.config,
		ctx:        ahq.ctx.Clone(),
		order:      append([]OrderFunc{}, ahq.order...),
		inters:     append([]Interceptor{}, ahq.inters...),
		predicates: append([]predicate.AttributeHistory{}, ahq.predicates...),
		// clone intermediate query.
		sql:  ahq.sql.Clone(),
		path: ahq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		HistoryTime time.Time `json:"history_time,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.AttributeHistory.Query().
//		GroupBy(attributehistory.FieldHistoryTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (ahq *AttributeHistoryQuery) GroupBy(field string, fields ...string) *AttributeHistoryGroupBy {
	ahq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &AttributeHistoryGroupBy{build: ahq}
	grbuild.flds = &ahq.ctx.Fields
	grbuild.label = attributehistory.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		HistoryTime time.Time `json:"history_time,omitempty"`
//	}
//
//	client.AttributeHistory.Query().
//		Select(attributehistory.FieldHistoryTime).
//		Scan(ctx, &v)
func (ahq *AttributeHistoryQuery) Select(fields ...string) *AttributeHistorySelect {
	ahq.ctx.Fields = append(ahq.ctx.Fields, fields...)
	sbuild := &AttributeHistorySelect{AttributeHistoryQuery: ahq}
	sbuild.label = attributehistory.Label
	sbuild.flds, sbuild.scan = &ahq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a AttributeHistorySelect configured with the given aggregations.
func (ahq *AttributeHistoryQuery) Aggregate(fns ...AggregateFunc) *AttributeHistorySelect {
	return ahq.Select().Aggregate(fns...)
}

func (ahq *AttributeHistoryQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range ahq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, ahq); err != nil {
				return err
			}
		}
	}
	for _, f := range ahq.ctx.Fields {
		if !attributehistory.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ahq.path != nil {
		prev, err := ahq.path(ctx)
		if err != nil {
			return err
		}
		ahq.sql = prev
	}
	return nil
}

func (ahq *AttributeHistoryQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*AttributeHistory, error) {
	var (
		nodes = []*AttributeHistory{}
		_spec = ahq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*AttributeHistory).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &AttributeHistory{config: ahq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(ahq.modifiers) > 0 {
		_spec.Modifiers = ahq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ahq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (ahq *AttributeHistoryQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ahq.querySpec()
	if len(ahq.modifiers) > 0 {
		_spec.Modifiers = ahq.modifiers
	}
	_spec.Node.Columns = ahq.ctx.Fields
	if len(ahq.ctx.Fields) > 0 {
		_spec.Unique = ahq.ctx.Unique != nil && *ahq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, ahq.driver, _spec)
}

func (ahq *AttributeHistoryQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(attributehistory.Table, attributehistory.Columns, sqlgraph.NewFieldSpec(attributehistory.FieldID, field.TypeUUID))
	_spec.From = ahq.sql
	if unique := ahq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if ahq.path != nil {
		_spec.Unique = true
	}
	if fields := ahq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, attributehistory.FieldID)
		for i := range fields {
			if fields[i] != attributehistory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ahq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ahq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ahq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ahq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ahq *AttributeHistoryQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ahq.driver.Dialect())
	t1 := builder.Table(attributehistory.Table)
	columns := ahq.ctx.Fields
	if len(columns) == 0 {
		columns = attributehistory.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ahq.sql != nil {
		selector = ahq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ahq.ctx.Unique != nil && *ahq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range ahq.modifiers {
		m(selector)
	}
	for _, p := range ahq.predicates {
		p(selector)
	}
	for _, p := range ahq.order {
		p(selector)
	}
	if offset := ahq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ahq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (ahq *AttributeHistoryQuery) ForUpdate(opts ...sql.LockOption) *AttributeHistoryQuery {
	if ahq.driver.Dialect() == dialect.Postgres {
		ahq.Unique(false)
	}
	ahq.modifiers = append(ahq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return ahq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (ahq *AttributeHistoryQuery) ForShare(opts ...sql.LockOption) *AttributeHistoryQuery {
	if ahq.driver.Dialect() == dialect.Postgres {
		ahq.Unique(false)
	}
	ahq.modifiers = append(ahq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return ahq
}

// AttributeHistoryGroupBy is the group-by builder for AttributeHistory entities.
type AttributeHistoryGroupBy struct {
	selector
	build *AttributeHistoryQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ahgb *AttributeHistoryGroupBy) Aggregate(fns ...AggregateFunc) *AttributeHistoryGroupBy {
	ahgb.fns = append(ahgb.fns, fns...)
	return ahgb
}

// Scan applies the selector query and scans the result into the given value.
func (ahgb *AttributeHistoryGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ahgb.build.ctx, "GroupBy")
	if err := ahgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AttributeHistoryQuery, *AttributeHistoryGroupBy](ctx, ahgb.build, ahgb, ahgb.build.inters, v)
}

func (ahgb *AttributeHistoryGroupBy) sqlScan(ctx context.Context, root *AttributeHistoryQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ahgb.fns))
	for _, fn := range ahgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ahgb.flds)+len(ahgb.fns))
		for _, f := range *ahgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ahgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ahgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// AttributeHistorySelect is the builder for selecting fields of AttributeHistory entities.
type AttributeHistorySelect struct {
	*AttributeHistoryQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ahs *AttributeHistorySelect) Aggregate(fns ...AggregateFunc) *AttributeHistorySelect {
	ahs.fns = append(ahs.fns, fns...)
	return ahs
}

// Scan applies the selector query and scans the result into the given value.
func (ahs *AttributeHistorySelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ahs.ctx, "Select")
	if err := ahs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AttributeHistoryQuery, *AttributeHistorySelect](ctx, ahs.AttributeHistoryQuery, ahs, ahs.inters, v)
}

func (ahs *AttributeHistorySelect) sqlScan(ctx context.Context, root *AttributeHistoryQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ahs.fns))
	for _, fn := range ahs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ahs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ahs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
