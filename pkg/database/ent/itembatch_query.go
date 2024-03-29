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

	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent/company"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent/itembatch"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent/predicate"
)

// ItemBatchQuery is the builder for querying ItemBatch entities.
type ItemBatchQuery struct {
	config
	ctx         *QueryContext
	order       []OrderFunc
	inters      []Interceptor
	predicates  []predicate.ItemBatch
	withCompany *CompanyQuery
	modifiers   []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ItemBatchQuery builder.
func (ibq *ItemBatchQuery) Where(ps ...predicate.ItemBatch) *ItemBatchQuery {
	ibq.predicates = append(ibq.predicates, ps...)
	return ibq
}

// Limit the number of records to be returned by this query.
func (ibq *ItemBatchQuery) Limit(limit int) *ItemBatchQuery {
	ibq.ctx.Limit = &limit
	return ibq
}

// Offset to start from.
func (ibq *ItemBatchQuery) Offset(offset int) *ItemBatchQuery {
	ibq.ctx.Offset = &offset
	return ibq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ibq *ItemBatchQuery) Unique(unique bool) *ItemBatchQuery {
	ibq.ctx.Unique = &unique
	return ibq
}

// Order specifies how the records should be ordered.
func (ibq *ItemBatchQuery) Order(o ...OrderFunc) *ItemBatchQuery {
	ibq.order = append(ibq.order, o...)
	return ibq
}

// QueryCompany chains the current query on the "company" edge.
func (ibq *ItemBatchQuery) QueryCompany() *CompanyQuery {
	query := (&CompanyClient{config: ibq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ibq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ibq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(itembatch.Table, itembatch.FieldID, selector),
			sqlgraph.To(company.Table, company.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, itembatch.CompanyTable, itembatch.CompanyColumn),
		)
		fromU = sqlgraph.SetNeighbors(ibq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ItemBatch entity from the query.
// Returns a *NotFoundError when no ItemBatch was found.
func (ibq *ItemBatchQuery) First(ctx context.Context) (*ItemBatch, error) {
	nodes, err := ibq.Limit(1).All(setContextOp(ctx, ibq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{itembatch.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ibq *ItemBatchQuery) FirstX(ctx context.Context) *ItemBatch {
	node, err := ibq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ItemBatch ID from the query.
// Returns a *NotFoundError when no ItemBatch ID was found.
func (ibq *ItemBatchQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = ibq.Limit(1).IDs(setContextOp(ctx, ibq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{itembatch.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ibq *ItemBatchQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := ibq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ItemBatch entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ItemBatch entity is found.
// Returns a *NotFoundError when no ItemBatch entities are found.
func (ibq *ItemBatchQuery) Only(ctx context.Context) (*ItemBatch, error) {
	nodes, err := ibq.Limit(2).All(setContextOp(ctx, ibq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{itembatch.Label}
	default:
		return nil, &NotSingularError{itembatch.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ibq *ItemBatchQuery) OnlyX(ctx context.Context) *ItemBatch {
	node, err := ibq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ItemBatch ID in the query.
// Returns a *NotSingularError when more than one ItemBatch ID is found.
// Returns a *NotFoundError when no entities are found.
func (ibq *ItemBatchQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = ibq.Limit(2).IDs(setContextOp(ctx, ibq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{itembatch.Label}
	default:
		err = &NotSingularError{itembatch.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ibq *ItemBatchQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := ibq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ItemBatches.
func (ibq *ItemBatchQuery) All(ctx context.Context) ([]*ItemBatch, error) {
	ctx = setContextOp(ctx, ibq.ctx, "All")
	if err := ibq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ItemBatch, *ItemBatchQuery]()
	return withInterceptors[[]*ItemBatch](ctx, ibq, qr, ibq.inters)
}

// AllX is like All, but panics if an error occurs.
func (ibq *ItemBatchQuery) AllX(ctx context.Context) []*ItemBatch {
	nodes, err := ibq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ItemBatch IDs.
func (ibq *ItemBatchQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if ibq.ctx.Unique == nil && ibq.path != nil {
		ibq.Unique(true)
	}
	ctx = setContextOp(ctx, ibq.ctx, "IDs")
	if err = ibq.Select(itembatch.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ibq *ItemBatchQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := ibq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ibq *ItemBatchQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, ibq.ctx, "Count")
	if err := ibq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, ibq, querierCount[*ItemBatchQuery](), ibq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (ibq *ItemBatchQuery) CountX(ctx context.Context) int {
	count, err := ibq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ibq *ItemBatchQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, ibq.ctx, "Exist")
	switch _, err := ibq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (ibq *ItemBatchQuery) ExistX(ctx context.Context) bool {
	exist, err := ibq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ItemBatchQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ibq *ItemBatchQuery) Clone() *ItemBatchQuery {
	if ibq == nil {
		return nil
	}
	return &ItemBatchQuery{
		config:      ibq.config,
		ctx:         ibq.ctx.Clone(),
		order:       append([]OrderFunc{}, ibq.order...),
		inters:      append([]Interceptor{}, ibq.inters...),
		predicates:  append([]predicate.ItemBatch{}, ibq.predicates...),
		withCompany: ibq.withCompany.Clone(),
		// clone intermediate query.
		sql:  ibq.sql.Clone(),
		path: ibq.path,
	}
}

// WithCompany tells the query-builder to eager-load the nodes that are connected to
// the "company" edge. The optional arguments are used to configure the query builder of the edge.
func (ibq *ItemBatchQuery) WithCompany(opts ...func(*CompanyQuery)) *ItemBatchQuery {
	query := (&CompanyClient{config: ibq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ibq.withCompany = query
	return ibq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		ItemNumber string `json:"itemNumber,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ItemBatch.Query().
//		GroupBy(itembatch.FieldItemNumber).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (ibq *ItemBatchQuery) GroupBy(field string, fields ...string) *ItemBatchGroupBy {
	ibq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ItemBatchGroupBy{build: ibq}
	grbuild.flds = &ibq.ctx.Fields
	grbuild.label = itembatch.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		ItemNumber string `json:"itemNumber,omitempty"`
//	}
//
//	client.ItemBatch.Query().
//		Select(itembatch.FieldItemNumber).
//		Scan(ctx, &v)
func (ibq *ItemBatchQuery) Select(fields ...string) *ItemBatchSelect {
	ibq.ctx.Fields = append(ibq.ctx.Fields, fields...)
	sbuild := &ItemBatchSelect{ItemBatchQuery: ibq}
	sbuild.label = itembatch.Label
	sbuild.flds, sbuild.scan = &ibq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ItemBatchSelect configured with the given aggregations.
func (ibq *ItemBatchQuery) Aggregate(fns ...AggregateFunc) *ItemBatchSelect {
	return ibq.Select().Aggregate(fns...)
}

func (ibq *ItemBatchQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range ibq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, ibq); err != nil {
				return err
			}
		}
	}
	for _, f := range ibq.ctx.Fields {
		if !itembatch.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ibq.path != nil {
		prev, err := ibq.path(ctx)
		if err != nil {
			return err
		}
		ibq.sql = prev
	}
	return nil
}

func (ibq *ItemBatchQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ItemBatch, error) {
	var (
		nodes       = []*ItemBatch{}
		_spec       = ibq.querySpec()
		loadedTypes = [1]bool{
			ibq.withCompany != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ItemBatch).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ItemBatch{config: ibq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(ibq.modifiers) > 0 {
		_spec.Modifiers = ibq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ibq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := ibq.withCompany; query != nil {
		if err := ibq.loadCompany(ctx, query, nodes, nil,
			func(n *ItemBatch, e *Company) { n.Edges.Company = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (ibq *ItemBatchQuery) loadCompany(ctx context.Context, query *CompanyQuery, nodes []*ItemBatch, init func(*ItemBatch), assign func(*ItemBatch, *Company)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*ItemBatch)
	for i := range nodes {
		fk := nodes[i].CompanyUUID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(company.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "companyUUID" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (ibq *ItemBatchQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ibq.querySpec()
	if len(ibq.modifiers) > 0 {
		_spec.Modifiers = ibq.modifiers
	}
	_spec.Node.Columns = ibq.ctx.Fields
	if len(ibq.ctx.Fields) > 0 {
		_spec.Unique = ibq.ctx.Unique != nil && *ibq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, ibq.driver, _spec)
}

func (ibq *ItemBatchQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(itembatch.Table, itembatch.Columns, sqlgraph.NewFieldSpec(itembatch.FieldID, field.TypeUUID))
	_spec.From = ibq.sql
	if unique := ibq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if ibq.path != nil {
		_spec.Unique = true
	}
	if fields := ibq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, itembatch.FieldID)
		for i := range fields {
			if fields[i] != itembatch.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ibq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ibq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ibq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ibq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ibq *ItemBatchQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ibq.driver.Dialect())
	t1 := builder.Table(itembatch.Table)
	columns := ibq.ctx.Fields
	if len(columns) == 0 {
		columns = itembatch.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ibq.sql != nil {
		selector = ibq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ibq.ctx.Unique != nil && *ibq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range ibq.modifiers {
		m(selector)
	}
	for _, p := range ibq.predicates {
		p(selector)
	}
	for _, p := range ibq.order {
		p(selector)
	}
	if offset := ibq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ibq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (ibq *ItemBatchQuery) ForUpdate(opts ...sql.LockOption) *ItemBatchQuery {
	if ibq.driver.Dialect() == dialect.Postgres {
		ibq.Unique(false)
	}
	ibq.modifiers = append(ibq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return ibq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (ibq *ItemBatchQuery) ForShare(opts ...sql.LockOption) *ItemBatchQuery {
	if ibq.driver.Dialect() == dialect.Postgres {
		ibq.Unique(false)
	}
	ibq.modifiers = append(ibq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return ibq
}

// ItemBatchGroupBy is the group-by builder for ItemBatch entities.
type ItemBatchGroupBy struct {
	selector
	build *ItemBatchQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ibgb *ItemBatchGroupBy) Aggregate(fns ...AggregateFunc) *ItemBatchGroupBy {
	ibgb.fns = append(ibgb.fns, fns...)
	return ibgb
}

// Scan applies the selector query and scans the result into the given value.
func (ibgb *ItemBatchGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ibgb.build.ctx, "GroupBy")
	if err := ibgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ItemBatchQuery, *ItemBatchGroupBy](ctx, ibgb.build, ibgb, ibgb.build.inters, v)
}

func (ibgb *ItemBatchGroupBy) sqlScan(ctx context.Context, root *ItemBatchQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ibgb.fns))
	for _, fn := range ibgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ibgb.flds)+len(ibgb.fns))
		for _, f := range *ibgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ibgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ibgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ItemBatchSelect is the builder for selecting fields of ItemBatch entities.
type ItemBatchSelect struct {
	*ItemBatchQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ibs *ItemBatchSelect) Aggregate(fns ...AggregateFunc) *ItemBatchSelect {
	ibs.fns = append(ibs.fns, fns...)
	return ibs
}

// Scan applies the selector query and scans the result into the given value.
func (ibs *ItemBatchSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ibs.ctx, "Select")
	if err := ibs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ItemBatchQuery, *ItemBatchSelect](ctx, ibs.ItemBatchQuery, ibs, ibs.inters, v)
}

func (ibs *ItemBatchSelect) sqlScan(ctx context.Context, root *ItemBatchQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ibs.fns))
	for _, fn := range ibs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ibs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ibs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
