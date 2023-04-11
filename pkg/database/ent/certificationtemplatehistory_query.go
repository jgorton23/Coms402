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

	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent/certificationtemplatehistory"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent/predicate"
)

// CertificationTemplateHistoryQuery is the builder for querying CertificationTemplateHistory entities.
type CertificationTemplateHistoryQuery struct {
	config
	ctx        *QueryContext
	order      []OrderFunc
	inters     []Interceptor
	predicates []predicate.CertificationTemplateHistory
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CertificationTemplateHistoryQuery builder.
func (cthq *CertificationTemplateHistoryQuery) Where(ps ...predicate.CertificationTemplateHistory) *CertificationTemplateHistoryQuery {
	cthq.predicates = append(cthq.predicates, ps...)
	return cthq
}

// Limit the number of records to be returned by this query.
func (cthq *CertificationTemplateHistoryQuery) Limit(limit int) *CertificationTemplateHistoryQuery {
	cthq.ctx.Limit = &limit
	return cthq
}

// Offset to start from.
func (cthq *CertificationTemplateHistoryQuery) Offset(offset int) *CertificationTemplateHistoryQuery {
	cthq.ctx.Offset = &offset
	return cthq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (cthq *CertificationTemplateHistoryQuery) Unique(unique bool) *CertificationTemplateHistoryQuery {
	cthq.ctx.Unique = &unique
	return cthq
}

// Order specifies how the records should be ordered.
func (cthq *CertificationTemplateHistoryQuery) Order(o ...OrderFunc) *CertificationTemplateHistoryQuery {
	cthq.order = append(cthq.order, o...)
	return cthq
}

// First returns the first CertificationTemplateHistory entity from the query.
// Returns a *NotFoundError when no CertificationTemplateHistory was found.
func (cthq *CertificationTemplateHistoryQuery) First(ctx context.Context) (*CertificationTemplateHistory, error) {
	nodes, err := cthq.Limit(1).All(setContextOp(ctx, cthq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{certificationtemplatehistory.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cthq *CertificationTemplateHistoryQuery) FirstX(ctx context.Context) *CertificationTemplateHistory {
	node, err := cthq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first CertificationTemplateHistory ID from the query.
// Returns a *NotFoundError when no CertificationTemplateHistory ID was found.
func (cthq *CertificationTemplateHistoryQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = cthq.Limit(1).IDs(setContextOp(ctx, cthq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{certificationtemplatehistory.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (cthq *CertificationTemplateHistoryQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := cthq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single CertificationTemplateHistory entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one CertificationTemplateHistory entity is found.
// Returns a *NotFoundError when no CertificationTemplateHistory entities are found.
func (cthq *CertificationTemplateHistoryQuery) Only(ctx context.Context) (*CertificationTemplateHistory, error) {
	nodes, err := cthq.Limit(2).All(setContextOp(ctx, cthq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{certificationtemplatehistory.Label}
	default:
		return nil, &NotSingularError{certificationtemplatehistory.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cthq *CertificationTemplateHistoryQuery) OnlyX(ctx context.Context) *CertificationTemplateHistory {
	node, err := cthq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only CertificationTemplateHistory ID in the query.
// Returns a *NotSingularError when more than one CertificationTemplateHistory ID is found.
// Returns a *NotFoundError when no entities are found.
func (cthq *CertificationTemplateHistoryQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = cthq.Limit(2).IDs(setContextOp(ctx, cthq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{certificationtemplatehistory.Label}
	default:
		err = &NotSingularError{certificationtemplatehistory.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (cthq *CertificationTemplateHistoryQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := cthq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of CertificationTemplateHistories.
func (cthq *CertificationTemplateHistoryQuery) All(ctx context.Context) ([]*CertificationTemplateHistory, error) {
	ctx = setContextOp(ctx, cthq.ctx, "All")
	if err := cthq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*CertificationTemplateHistory, *CertificationTemplateHistoryQuery]()
	return withInterceptors[[]*CertificationTemplateHistory](ctx, cthq, qr, cthq.inters)
}

// AllX is like All, but panics if an error occurs.
func (cthq *CertificationTemplateHistoryQuery) AllX(ctx context.Context) []*CertificationTemplateHistory {
	nodes, err := cthq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of CertificationTemplateHistory IDs.
func (cthq *CertificationTemplateHistoryQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if cthq.ctx.Unique == nil && cthq.path != nil {
		cthq.Unique(true)
	}
	ctx = setContextOp(ctx, cthq.ctx, "IDs")
	if err = cthq.Select(certificationtemplatehistory.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (cthq *CertificationTemplateHistoryQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := cthq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (cthq *CertificationTemplateHistoryQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, cthq.ctx, "Count")
	if err := cthq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, cthq, querierCount[*CertificationTemplateHistoryQuery](), cthq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (cthq *CertificationTemplateHistoryQuery) CountX(ctx context.Context) int {
	count, err := cthq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cthq *CertificationTemplateHistoryQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, cthq.ctx, "Exist")
	switch _, err := cthq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (cthq *CertificationTemplateHistoryQuery) ExistX(ctx context.Context) bool {
	exist, err := cthq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CertificationTemplateHistoryQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cthq *CertificationTemplateHistoryQuery) Clone() *CertificationTemplateHistoryQuery {
	if cthq == nil {
		return nil
	}
	return &CertificationTemplateHistoryQuery{
		config:     cthq.config,
		ctx:        cthq.ctx.Clone(),
		order:      append([]OrderFunc{}, cthq.order...),
		inters:     append([]Interceptor{}, cthq.inters...),
		predicates: append([]predicate.CertificationTemplateHistory{}, cthq.predicates...),
		// clone intermediate query.
		sql:  cthq.sql.Clone(),
		path: cthq.path,
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
//	client.CertificationTemplateHistory.Query().
//		GroupBy(certificationtemplatehistory.FieldHistoryTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (cthq *CertificationTemplateHistoryQuery) GroupBy(field string, fields ...string) *CertificationTemplateHistoryGroupBy {
	cthq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &CertificationTemplateHistoryGroupBy{build: cthq}
	grbuild.flds = &cthq.ctx.Fields
	grbuild.label = certificationtemplatehistory.Label
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
//	client.CertificationTemplateHistory.Query().
//		Select(certificationtemplatehistory.FieldHistoryTime).
//		Scan(ctx, &v)
func (cthq *CertificationTemplateHistoryQuery) Select(fields ...string) *CertificationTemplateHistorySelect {
	cthq.ctx.Fields = append(cthq.ctx.Fields, fields...)
	sbuild := &CertificationTemplateHistorySelect{CertificationTemplateHistoryQuery: cthq}
	sbuild.label = certificationtemplatehistory.Label
	sbuild.flds, sbuild.scan = &cthq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a CertificationTemplateHistorySelect configured with the given aggregations.
func (cthq *CertificationTemplateHistoryQuery) Aggregate(fns ...AggregateFunc) *CertificationTemplateHistorySelect {
	return cthq.Select().Aggregate(fns...)
}

func (cthq *CertificationTemplateHistoryQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range cthq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, cthq); err != nil {
				return err
			}
		}
	}
	for _, f := range cthq.ctx.Fields {
		if !certificationtemplatehistory.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if cthq.path != nil {
		prev, err := cthq.path(ctx)
		if err != nil {
			return err
		}
		cthq.sql = prev
	}
	return nil
}

func (cthq *CertificationTemplateHistoryQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*CertificationTemplateHistory, error) {
	var (
		nodes = []*CertificationTemplateHistory{}
		_spec = cthq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*CertificationTemplateHistory).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &CertificationTemplateHistory{config: cthq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(cthq.modifiers) > 0 {
		_spec.Modifiers = cthq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, cthq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (cthq *CertificationTemplateHistoryQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := cthq.querySpec()
	if len(cthq.modifiers) > 0 {
		_spec.Modifiers = cthq.modifiers
	}
	_spec.Node.Columns = cthq.ctx.Fields
	if len(cthq.ctx.Fields) > 0 {
		_spec.Unique = cthq.ctx.Unique != nil && *cthq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, cthq.driver, _spec)
}

func (cthq *CertificationTemplateHistoryQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(certificationtemplatehistory.Table, certificationtemplatehistory.Columns, sqlgraph.NewFieldSpec(certificationtemplatehistory.FieldID, field.TypeUUID))
	_spec.From = cthq.sql
	if unique := cthq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if cthq.path != nil {
		_spec.Unique = true
	}
	if fields := cthq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, certificationtemplatehistory.FieldID)
		for i := range fields {
			if fields[i] != certificationtemplatehistory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := cthq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := cthq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := cthq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := cthq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (cthq *CertificationTemplateHistoryQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(cthq.driver.Dialect())
	t1 := builder.Table(certificationtemplatehistory.Table)
	columns := cthq.ctx.Fields
	if len(columns) == 0 {
		columns = certificationtemplatehistory.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if cthq.sql != nil {
		selector = cthq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if cthq.ctx.Unique != nil && *cthq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range cthq.modifiers {
		m(selector)
	}
	for _, p := range cthq.predicates {
		p(selector)
	}
	for _, p := range cthq.order {
		p(selector)
	}
	if offset := cthq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := cthq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (cthq *CertificationTemplateHistoryQuery) ForUpdate(opts ...sql.LockOption) *CertificationTemplateHistoryQuery {
	if cthq.driver.Dialect() == dialect.Postgres {
		cthq.Unique(false)
	}
	cthq.modifiers = append(cthq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return cthq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (cthq *CertificationTemplateHistoryQuery) ForShare(opts ...sql.LockOption) *CertificationTemplateHistoryQuery {
	if cthq.driver.Dialect() == dialect.Postgres {
		cthq.Unique(false)
	}
	cthq.modifiers = append(cthq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return cthq
}

// CertificationTemplateHistoryGroupBy is the group-by builder for CertificationTemplateHistory entities.
type CertificationTemplateHistoryGroupBy struct {
	selector
	build *CertificationTemplateHistoryQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cthgb *CertificationTemplateHistoryGroupBy) Aggregate(fns ...AggregateFunc) *CertificationTemplateHistoryGroupBy {
	cthgb.fns = append(cthgb.fns, fns...)
	return cthgb
}

// Scan applies the selector query and scans the result into the given value.
func (cthgb *CertificationTemplateHistoryGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cthgb.build.ctx, "GroupBy")
	if err := cthgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CertificationTemplateHistoryQuery, *CertificationTemplateHistoryGroupBy](ctx, cthgb.build, cthgb, cthgb.build.inters, v)
}

func (cthgb *CertificationTemplateHistoryGroupBy) sqlScan(ctx context.Context, root *CertificationTemplateHistoryQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(cthgb.fns))
	for _, fn := range cthgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*cthgb.flds)+len(cthgb.fns))
		for _, f := range *cthgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*cthgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cthgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// CertificationTemplateHistorySelect is the builder for selecting fields of CertificationTemplateHistory entities.
type CertificationTemplateHistorySelect struct {
	*CertificationTemplateHistoryQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (cths *CertificationTemplateHistorySelect) Aggregate(fns ...AggregateFunc) *CertificationTemplateHistorySelect {
	cths.fns = append(cths.fns, fns...)
	return cths
}

// Scan applies the selector query and scans the result into the given value.
func (cths *CertificationTemplateHistorySelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cths.ctx, "Select")
	if err := cths.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CertificationTemplateHistoryQuery, *CertificationTemplateHistorySelect](ctx, cths.CertificationTemplateHistoryQuery, cths, cths.inters, v)
}

func (cths *CertificationTemplateHistorySelect) sqlScan(ctx context.Context, root *CertificationTemplateHistoryQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(cths.fns))
	for _, fn := range cths.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*cths.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cths.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
