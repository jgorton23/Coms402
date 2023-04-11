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

	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent/predicate"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent/userstocompanyhistory"
)

// UsersToCompanyHistoryQuery is the builder for querying UsersToCompanyHistory entities.
type UsersToCompanyHistoryQuery struct {
	config
	ctx        *QueryContext
	order      []OrderFunc
	inters     []Interceptor
	predicates []predicate.UsersToCompanyHistory
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the UsersToCompanyHistoryQuery builder.
func (utchq *UsersToCompanyHistoryQuery) Where(ps ...predicate.UsersToCompanyHistory) *UsersToCompanyHistoryQuery {
	utchq.predicates = append(utchq.predicates, ps...)
	return utchq
}

// Limit the number of records to be returned by this query.
func (utchq *UsersToCompanyHistoryQuery) Limit(limit int) *UsersToCompanyHistoryQuery {
	utchq.ctx.Limit = &limit
	return utchq
}

// Offset to start from.
func (utchq *UsersToCompanyHistoryQuery) Offset(offset int) *UsersToCompanyHistoryQuery {
	utchq.ctx.Offset = &offset
	return utchq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (utchq *UsersToCompanyHistoryQuery) Unique(unique bool) *UsersToCompanyHistoryQuery {
	utchq.ctx.Unique = &unique
	return utchq
}

// Order specifies how the records should be ordered.
func (utchq *UsersToCompanyHistoryQuery) Order(o ...OrderFunc) *UsersToCompanyHistoryQuery {
	utchq.order = append(utchq.order, o...)
	return utchq
}

// First returns the first UsersToCompanyHistory entity from the query.
// Returns a *NotFoundError when no UsersToCompanyHistory was found.
func (utchq *UsersToCompanyHistoryQuery) First(ctx context.Context) (*UsersToCompanyHistory, error) {
	nodes, err := utchq.Limit(1).All(setContextOp(ctx, utchq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{userstocompanyhistory.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (utchq *UsersToCompanyHistoryQuery) FirstX(ctx context.Context) *UsersToCompanyHistory {
	node, err := utchq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first UsersToCompanyHistory ID from the query.
// Returns a *NotFoundError when no UsersToCompanyHistory ID was found.
func (utchq *UsersToCompanyHistoryQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = utchq.Limit(1).IDs(setContextOp(ctx, utchq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{userstocompanyhistory.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (utchq *UsersToCompanyHistoryQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := utchq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single UsersToCompanyHistory entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one UsersToCompanyHistory entity is found.
// Returns a *NotFoundError when no UsersToCompanyHistory entities are found.
func (utchq *UsersToCompanyHistoryQuery) Only(ctx context.Context) (*UsersToCompanyHistory, error) {
	nodes, err := utchq.Limit(2).All(setContextOp(ctx, utchq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{userstocompanyhistory.Label}
	default:
		return nil, &NotSingularError{userstocompanyhistory.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (utchq *UsersToCompanyHistoryQuery) OnlyX(ctx context.Context) *UsersToCompanyHistory {
	node, err := utchq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only UsersToCompanyHistory ID in the query.
// Returns a *NotSingularError when more than one UsersToCompanyHistory ID is found.
// Returns a *NotFoundError when no entities are found.
func (utchq *UsersToCompanyHistoryQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = utchq.Limit(2).IDs(setContextOp(ctx, utchq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{userstocompanyhistory.Label}
	default:
		err = &NotSingularError{userstocompanyhistory.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (utchq *UsersToCompanyHistoryQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := utchq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of UsersToCompanyHistories.
func (utchq *UsersToCompanyHistoryQuery) All(ctx context.Context) ([]*UsersToCompanyHistory, error) {
	ctx = setContextOp(ctx, utchq.ctx, "All")
	if err := utchq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*UsersToCompanyHistory, *UsersToCompanyHistoryQuery]()
	return withInterceptors[[]*UsersToCompanyHistory](ctx, utchq, qr, utchq.inters)
}

// AllX is like All, but panics if an error occurs.
func (utchq *UsersToCompanyHistoryQuery) AllX(ctx context.Context) []*UsersToCompanyHistory {
	nodes, err := utchq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of UsersToCompanyHistory IDs.
func (utchq *UsersToCompanyHistoryQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if utchq.ctx.Unique == nil && utchq.path != nil {
		utchq.Unique(true)
	}
	ctx = setContextOp(ctx, utchq.ctx, "IDs")
	if err = utchq.Select(userstocompanyhistory.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (utchq *UsersToCompanyHistoryQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := utchq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (utchq *UsersToCompanyHistoryQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, utchq.ctx, "Count")
	if err := utchq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, utchq, querierCount[*UsersToCompanyHistoryQuery](), utchq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (utchq *UsersToCompanyHistoryQuery) CountX(ctx context.Context) int {
	count, err := utchq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (utchq *UsersToCompanyHistoryQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, utchq.ctx, "Exist")
	switch _, err := utchq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (utchq *UsersToCompanyHistoryQuery) ExistX(ctx context.Context) bool {
	exist, err := utchq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UsersToCompanyHistoryQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (utchq *UsersToCompanyHistoryQuery) Clone() *UsersToCompanyHistoryQuery {
	if utchq == nil {
		return nil
	}
	return &UsersToCompanyHistoryQuery{
		config:     utchq.config,
		ctx:        utchq.ctx.Clone(),
		order:      append([]OrderFunc{}, utchq.order...),
		inters:     append([]Interceptor{}, utchq.inters...),
		predicates: append([]predicate.UsersToCompanyHistory{}, utchq.predicates...),
		// clone intermediate query.
		sql:  utchq.sql.Clone(),
		path: utchq.path,
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
//	client.UsersToCompanyHistory.Query().
//		GroupBy(userstocompanyhistory.FieldHistoryTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (utchq *UsersToCompanyHistoryQuery) GroupBy(field string, fields ...string) *UsersToCompanyHistoryGroupBy {
	utchq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &UsersToCompanyHistoryGroupBy{build: utchq}
	grbuild.flds = &utchq.ctx.Fields
	grbuild.label = userstocompanyhistory.Label
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
//	client.UsersToCompanyHistory.Query().
//		Select(userstocompanyhistory.FieldHistoryTime).
//		Scan(ctx, &v)
func (utchq *UsersToCompanyHistoryQuery) Select(fields ...string) *UsersToCompanyHistorySelect {
	utchq.ctx.Fields = append(utchq.ctx.Fields, fields...)
	sbuild := &UsersToCompanyHistorySelect{UsersToCompanyHistoryQuery: utchq}
	sbuild.label = userstocompanyhistory.Label
	sbuild.flds, sbuild.scan = &utchq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a UsersToCompanyHistorySelect configured with the given aggregations.
func (utchq *UsersToCompanyHistoryQuery) Aggregate(fns ...AggregateFunc) *UsersToCompanyHistorySelect {
	return utchq.Select().Aggregate(fns...)
}

func (utchq *UsersToCompanyHistoryQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range utchq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, utchq); err != nil {
				return err
			}
		}
	}
	for _, f := range utchq.ctx.Fields {
		if !userstocompanyhistory.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if utchq.path != nil {
		prev, err := utchq.path(ctx)
		if err != nil {
			return err
		}
		utchq.sql = prev
	}
	return nil
}

func (utchq *UsersToCompanyHistoryQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*UsersToCompanyHistory, error) {
	var (
		nodes = []*UsersToCompanyHistory{}
		_spec = utchq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*UsersToCompanyHistory).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &UsersToCompanyHistory{config: utchq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(utchq.modifiers) > 0 {
		_spec.Modifiers = utchq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, utchq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (utchq *UsersToCompanyHistoryQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := utchq.querySpec()
	if len(utchq.modifiers) > 0 {
		_spec.Modifiers = utchq.modifiers
	}
	_spec.Node.Columns = utchq.ctx.Fields
	if len(utchq.ctx.Fields) > 0 {
		_spec.Unique = utchq.ctx.Unique != nil && *utchq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, utchq.driver, _spec)
}

func (utchq *UsersToCompanyHistoryQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(userstocompanyhistory.Table, userstocompanyhistory.Columns, sqlgraph.NewFieldSpec(userstocompanyhistory.FieldID, field.TypeUUID))
	_spec.From = utchq.sql
	if unique := utchq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if utchq.path != nil {
		_spec.Unique = true
	}
	if fields := utchq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, userstocompanyhistory.FieldID)
		for i := range fields {
			if fields[i] != userstocompanyhistory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := utchq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := utchq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := utchq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := utchq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (utchq *UsersToCompanyHistoryQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(utchq.driver.Dialect())
	t1 := builder.Table(userstocompanyhistory.Table)
	columns := utchq.ctx.Fields
	if len(columns) == 0 {
		columns = userstocompanyhistory.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if utchq.sql != nil {
		selector = utchq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if utchq.ctx.Unique != nil && *utchq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range utchq.modifiers {
		m(selector)
	}
	for _, p := range utchq.predicates {
		p(selector)
	}
	for _, p := range utchq.order {
		p(selector)
	}
	if offset := utchq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := utchq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (utchq *UsersToCompanyHistoryQuery) ForUpdate(opts ...sql.LockOption) *UsersToCompanyHistoryQuery {
	if utchq.driver.Dialect() == dialect.Postgres {
		utchq.Unique(false)
	}
	utchq.modifiers = append(utchq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return utchq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (utchq *UsersToCompanyHistoryQuery) ForShare(opts ...sql.LockOption) *UsersToCompanyHistoryQuery {
	if utchq.driver.Dialect() == dialect.Postgres {
		utchq.Unique(false)
	}
	utchq.modifiers = append(utchq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return utchq
}

// UsersToCompanyHistoryGroupBy is the group-by builder for UsersToCompanyHistory entities.
type UsersToCompanyHistoryGroupBy struct {
	selector
	build *UsersToCompanyHistoryQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (utchgb *UsersToCompanyHistoryGroupBy) Aggregate(fns ...AggregateFunc) *UsersToCompanyHistoryGroupBy {
	utchgb.fns = append(utchgb.fns, fns...)
	return utchgb
}

// Scan applies the selector query and scans the result into the given value.
func (utchgb *UsersToCompanyHistoryGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, utchgb.build.ctx, "GroupBy")
	if err := utchgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UsersToCompanyHistoryQuery, *UsersToCompanyHistoryGroupBy](ctx, utchgb.build, utchgb, utchgb.build.inters, v)
}

func (utchgb *UsersToCompanyHistoryGroupBy) sqlScan(ctx context.Context, root *UsersToCompanyHistoryQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(utchgb.fns))
	for _, fn := range utchgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*utchgb.flds)+len(utchgb.fns))
		for _, f := range *utchgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*utchgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := utchgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// UsersToCompanyHistorySelect is the builder for selecting fields of UsersToCompanyHistory entities.
type UsersToCompanyHistorySelect struct {
	*UsersToCompanyHistoryQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (utchs *UsersToCompanyHistorySelect) Aggregate(fns ...AggregateFunc) *UsersToCompanyHistorySelect {
	utchs.fns = append(utchs.fns, fns...)
	return utchs
}

// Scan applies the selector query and scans the result into the given value.
func (utchs *UsersToCompanyHistorySelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, utchs.ctx, "Select")
	if err := utchs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UsersToCompanyHistoryQuery, *UsersToCompanyHistorySelect](ctx, utchs.UsersToCompanyHistoryQuery, utchs, utchs.inters, v)
}

func (utchs *UsersToCompanyHistorySelect) sqlScan(ctx context.Context, root *UsersToCompanyHistoryQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(utchs.fns))
	for _, fn := range utchs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*utchs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := utchs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}