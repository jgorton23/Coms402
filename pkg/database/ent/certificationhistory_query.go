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

	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent/certificationhistory"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent/predicate"
)

// CertificationHistoryQuery is the builder for querying CertificationHistory entities.
type CertificationHistoryQuery struct {
	config
	ctx        *QueryContext
	order      []OrderFunc
	inters     []Interceptor
	predicates []predicate.CertificationHistory
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CertificationHistoryQuery builder.
func (chq *CertificationHistoryQuery) Where(ps ...predicate.CertificationHistory) *CertificationHistoryQuery {
	chq.predicates = append(chq.predicates, ps...)
	return chq
}

// Limit the number of records to be returned by this query.
func (chq *CertificationHistoryQuery) Limit(limit int) *CertificationHistoryQuery {
	chq.ctx.Limit = &limit
	return chq
}

// Offset to start from.
func (chq *CertificationHistoryQuery) Offset(offset int) *CertificationHistoryQuery {
	chq.ctx.Offset = &offset
	return chq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (chq *CertificationHistoryQuery) Unique(unique bool) *CertificationHistoryQuery {
	chq.ctx.Unique = &unique
	return chq
}

// Order specifies how the records should be ordered.
func (chq *CertificationHistoryQuery) Order(o ...OrderFunc) *CertificationHistoryQuery {
	chq.order = append(chq.order, o...)
	return chq
}

// First returns the first CertificationHistory entity from the query.
// Returns a *NotFoundError when no CertificationHistory was found.
func (chq *CertificationHistoryQuery) First(ctx context.Context) (*CertificationHistory, error) {
	nodes, err := chq.Limit(1).All(setContextOp(ctx, chq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{certificationhistory.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (chq *CertificationHistoryQuery) FirstX(ctx context.Context) *CertificationHistory {
	node, err := chq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first CertificationHistory ID from the query.
// Returns a *NotFoundError when no CertificationHistory ID was found.
func (chq *CertificationHistoryQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = chq.Limit(1).IDs(setContextOp(ctx, chq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{certificationhistory.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (chq *CertificationHistoryQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := chq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single CertificationHistory entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one CertificationHistory entity is found.
// Returns a *NotFoundError when no CertificationHistory entities are found.
func (chq *CertificationHistoryQuery) Only(ctx context.Context) (*CertificationHistory, error) {
	nodes, err := chq.Limit(2).All(setContextOp(ctx, chq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{certificationhistory.Label}
	default:
		return nil, &NotSingularError{certificationhistory.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (chq *CertificationHistoryQuery) OnlyX(ctx context.Context) *CertificationHistory {
	node, err := chq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only CertificationHistory ID in the query.
// Returns a *NotSingularError when more than one CertificationHistory ID is found.
// Returns a *NotFoundError when no entities are found.
func (chq *CertificationHistoryQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = chq.Limit(2).IDs(setContextOp(ctx, chq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{certificationhistory.Label}
	default:
		err = &NotSingularError{certificationhistory.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (chq *CertificationHistoryQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := chq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of CertificationHistories.
func (chq *CertificationHistoryQuery) All(ctx context.Context) ([]*CertificationHistory, error) {
	ctx = setContextOp(ctx, chq.ctx, "All")
	if err := chq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*CertificationHistory, *CertificationHistoryQuery]()
	return withInterceptors[[]*CertificationHistory](ctx, chq, qr, chq.inters)
}

// AllX is like All, but panics if an error occurs.
func (chq *CertificationHistoryQuery) AllX(ctx context.Context) []*CertificationHistory {
	nodes, err := chq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of CertificationHistory IDs.
func (chq *CertificationHistoryQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if chq.ctx.Unique == nil && chq.path != nil {
		chq.Unique(true)
	}
	ctx = setContextOp(ctx, chq.ctx, "IDs")
	if err = chq.Select(certificationhistory.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (chq *CertificationHistoryQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := chq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (chq *CertificationHistoryQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, chq.ctx, "Count")
	if err := chq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, chq, querierCount[*CertificationHistoryQuery](), chq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (chq *CertificationHistoryQuery) CountX(ctx context.Context) int {
	count, err := chq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (chq *CertificationHistoryQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, chq.ctx, "Exist")
	switch _, err := chq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (chq *CertificationHistoryQuery) ExistX(ctx context.Context) bool {
	exist, err := chq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CertificationHistoryQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (chq *CertificationHistoryQuery) Clone() *CertificationHistoryQuery {
	if chq == nil {
		return nil
	}
	return &CertificationHistoryQuery{
		config:     chq.config,
		ctx:        chq.ctx.Clone(),
		order:      append([]OrderFunc{}, chq.order...),
		inters:     append([]Interceptor{}, chq.inters...),
		predicates: append([]predicate.CertificationHistory{}, chq.predicates...),
		// clone intermediate query.
		sql:  chq.sql.Clone(),
		path: chq.path,
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
//	client.CertificationHistory.Query().
//		GroupBy(certificationhistory.FieldHistoryTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (chq *CertificationHistoryQuery) GroupBy(field string, fields ...string) *CertificationHistoryGroupBy {
	chq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &CertificationHistoryGroupBy{build: chq}
	grbuild.flds = &chq.ctx.Fields
	grbuild.label = certificationhistory.Label
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
//	client.CertificationHistory.Query().
//		Select(certificationhistory.FieldHistoryTime).
//		Scan(ctx, &v)
func (chq *CertificationHistoryQuery) Select(fields ...string) *CertificationHistorySelect {
	chq.ctx.Fields = append(chq.ctx.Fields, fields...)
	sbuild := &CertificationHistorySelect{CertificationHistoryQuery: chq}
	sbuild.label = certificationhistory.Label
	sbuild.flds, sbuild.scan = &chq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a CertificationHistorySelect configured with the given aggregations.
func (chq *CertificationHistoryQuery) Aggregate(fns ...AggregateFunc) *CertificationHistorySelect {
	return chq.Select().Aggregate(fns...)
}

func (chq *CertificationHistoryQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range chq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, chq); err != nil {
				return err
			}
		}
	}
	for _, f := range chq.ctx.Fields {
		if !certificationhistory.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if chq.path != nil {
		prev, err := chq.path(ctx)
		if err != nil {
			return err
		}
		chq.sql = prev
	}
	return nil
}

func (chq *CertificationHistoryQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*CertificationHistory, error) {
	var (
		nodes = []*CertificationHistory{}
		_spec = chq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*CertificationHistory).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &CertificationHistory{config: chq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(chq.modifiers) > 0 {
		_spec.Modifiers = chq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, chq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (chq *CertificationHistoryQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := chq.querySpec()
	if len(chq.modifiers) > 0 {
		_spec.Modifiers = chq.modifiers
	}
	_spec.Node.Columns = chq.ctx.Fields
	if len(chq.ctx.Fields) > 0 {
		_spec.Unique = chq.ctx.Unique != nil && *chq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, chq.driver, _spec)
}

func (chq *CertificationHistoryQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(certificationhistory.Table, certificationhistory.Columns, sqlgraph.NewFieldSpec(certificationhistory.FieldID, field.TypeUUID))
	_spec.From = chq.sql
	if unique := chq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if chq.path != nil {
		_spec.Unique = true
	}
	if fields := chq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, certificationhistory.FieldID)
		for i := range fields {
			if fields[i] != certificationhistory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := chq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := chq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := chq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := chq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (chq *CertificationHistoryQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(chq.driver.Dialect())
	t1 := builder.Table(certificationhistory.Table)
	columns := chq.ctx.Fields
	if len(columns) == 0 {
		columns = certificationhistory.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if chq.sql != nil {
		selector = chq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if chq.ctx.Unique != nil && *chq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range chq.modifiers {
		m(selector)
	}
	for _, p := range chq.predicates {
		p(selector)
	}
	for _, p := range chq.order {
		p(selector)
	}
	if offset := chq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := chq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (chq *CertificationHistoryQuery) ForUpdate(opts ...sql.LockOption) *CertificationHistoryQuery {
	if chq.driver.Dialect() == dialect.Postgres {
		chq.Unique(false)
	}
	chq.modifiers = append(chq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return chq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (chq *CertificationHistoryQuery) ForShare(opts ...sql.LockOption) *CertificationHistoryQuery {
	if chq.driver.Dialect() == dialect.Postgres {
		chq.Unique(false)
	}
	chq.modifiers = append(chq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return chq
}

// CertificationHistoryGroupBy is the group-by builder for CertificationHistory entities.
type CertificationHistoryGroupBy struct {
	selector
	build *CertificationHistoryQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (chgb *CertificationHistoryGroupBy) Aggregate(fns ...AggregateFunc) *CertificationHistoryGroupBy {
	chgb.fns = append(chgb.fns, fns...)
	return chgb
}

// Scan applies the selector query and scans the result into the given value.
func (chgb *CertificationHistoryGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, chgb.build.ctx, "GroupBy")
	if err := chgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CertificationHistoryQuery, *CertificationHistoryGroupBy](ctx, chgb.build, chgb, chgb.build.inters, v)
}

func (chgb *CertificationHistoryGroupBy) sqlScan(ctx context.Context, root *CertificationHistoryQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(chgb.fns))
	for _, fn := range chgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*chgb.flds)+len(chgb.fns))
		for _, f := range *chgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*chgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := chgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// CertificationHistorySelect is the builder for selecting fields of CertificationHistory entities.
type CertificationHistorySelect struct {
	*CertificationHistoryQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (chs *CertificationHistorySelect) Aggregate(fns ...AggregateFunc) *CertificationHistorySelect {
	chs.fns = append(chs.fns, fns...)
	return chs
}

// Scan applies the selector query and scans the result into the given value.
func (chs *CertificationHistorySelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, chs.ctx, "Select")
	if err := chs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CertificationHistoryQuery, *CertificationHistorySelect](ctx, chs.CertificationHistoryQuery, chs, chs.inters, v)
}

func (chs *CertificationHistorySelect) sqlScan(ctx context.Context, root *CertificationHistoryQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(chs.fns))
	for _, fn := range chs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*chs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := chs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
