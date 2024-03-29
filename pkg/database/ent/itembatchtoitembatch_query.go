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

	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent/itembatch"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent/itembatchtoitembatch"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent/predicate"
)

// ItemBatchToItemBatchQuery is the builder for querying ItemBatchToItemBatch entities.
type ItemBatchToItemBatchQuery struct {
	config
	ctx        *QueryContext
	order      []OrderFunc
	inters     []Interceptor
	predicates []predicate.ItemBatchToItemBatch
	withParent *ItemBatchQuery
	withChild  *ItemBatchQuery
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ItemBatchToItemBatchQuery builder.
func (ibtibq *ItemBatchToItemBatchQuery) Where(ps ...predicate.ItemBatchToItemBatch) *ItemBatchToItemBatchQuery {
	ibtibq.predicates = append(ibtibq.predicates, ps...)
	return ibtibq
}

// Limit the number of records to be returned by this query.
func (ibtibq *ItemBatchToItemBatchQuery) Limit(limit int) *ItemBatchToItemBatchQuery {
	ibtibq.ctx.Limit = &limit
	return ibtibq
}

// Offset to start from.
func (ibtibq *ItemBatchToItemBatchQuery) Offset(offset int) *ItemBatchToItemBatchQuery {
	ibtibq.ctx.Offset = &offset
	return ibtibq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ibtibq *ItemBatchToItemBatchQuery) Unique(unique bool) *ItemBatchToItemBatchQuery {
	ibtibq.ctx.Unique = &unique
	return ibtibq
}

// Order specifies how the records should be ordered.
func (ibtibq *ItemBatchToItemBatchQuery) Order(o ...OrderFunc) *ItemBatchToItemBatchQuery {
	ibtibq.order = append(ibtibq.order, o...)
	return ibtibq
}

// QueryParent chains the current query on the "parent" edge.
func (ibtibq *ItemBatchToItemBatchQuery) QueryParent() *ItemBatchQuery {
	query := (&ItemBatchClient{config: ibtibq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ibtibq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ibtibq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(itembatchtoitembatch.Table, itembatchtoitembatch.FieldID, selector),
			sqlgraph.To(itembatch.Table, itembatch.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, itembatchtoitembatch.ParentTable, itembatchtoitembatch.ParentColumn),
		)
		fromU = sqlgraph.SetNeighbors(ibtibq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryChild chains the current query on the "child" edge.
func (ibtibq *ItemBatchToItemBatchQuery) QueryChild() *ItemBatchQuery {
	query := (&ItemBatchClient{config: ibtibq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ibtibq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ibtibq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(itembatchtoitembatch.Table, itembatchtoitembatch.FieldID, selector),
			sqlgraph.To(itembatch.Table, itembatch.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, itembatchtoitembatch.ChildTable, itembatchtoitembatch.ChildColumn),
		)
		fromU = sqlgraph.SetNeighbors(ibtibq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ItemBatchToItemBatch entity from the query.
// Returns a *NotFoundError when no ItemBatchToItemBatch was found.
func (ibtibq *ItemBatchToItemBatchQuery) First(ctx context.Context) (*ItemBatchToItemBatch, error) {
	nodes, err := ibtibq.Limit(1).All(setContextOp(ctx, ibtibq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{itembatchtoitembatch.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ibtibq *ItemBatchToItemBatchQuery) FirstX(ctx context.Context) *ItemBatchToItemBatch {
	node, err := ibtibq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ItemBatchToItemBatch ID from the query.
// Returns a *NotFoundError when no ItemBatchToItemBatch ID was found.
func (ibtibq *ItemBatchToItemBatchQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = ibtibq.Limit(1).IDs(setContextOp(ctx, ibtibq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{itembatchtoitembatch.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ibtibq *ItemBatchToItemBatchQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := ibtibq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ItemBatchToItemBatch entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ItemBatchToItemBatch entity is found.
// Returns a *NotFoundError when no ItemBatchToItemBatch entities are found.
func (ibtibq *ItemBatchToItemBatchQuery) Only(ctx context.Context) (*ItemBatchToItemBatch, error) {
	nodes, err := ibtibq.Limit(2).All(setContextOp(ctx, ibtibq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{itembatchtoitembatch.Label}
	default:
		return nil, &NotSingularError{itembatchtoitembatch.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ibtibq *ItemBatchToItemBatchQuery) OnlyX(ctx context.Context) *ItemBatchToItemBatch {
	node, err := ibtibq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ItemBatchToItemBatch ID in the query.
// Returns a *NotSingularError when more than one ItemBatchToItemBatch ID is found.
// Returns a *NotFoundError when no entities are found.
func (ibtibq *ItemBatchToItemBatchQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = ibtibq.Limit(2).IDs(setContextOp(ctx, ibtibq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{itembatchtoitembatch.Label}
	default:
		err = &NotSingularError{itembatchtoitembatch.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ibtibq *ItemBatchToItemBatchQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := ibtibq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ItemBatchToItemBatches.
func (ibtibq *ItemBatchToItemBatchQuery) All(ctx context.Context) ([]*ItemBatchToItemBatch, error) {
	ctx = setContextOp(ctx, ibtibq.ctx, "All")
	if err := ibtibq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ItemBatchToItemBatch, *ItemBatchToItemBatchQuery]()
	return withInterceptors[[]*ItemBatchToItemBatch](ctx, ibtibq, qr, ibtibq.inters)
}

// AllX is like All, but panics if an error occurs.
func (ibtibq *ItemBatchToItemBatchQuery) AllX(ctx context.Context) []*ItemBatchToItemBatch {
	nodes, err := ibtibq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ItemBatchToItemBatch IDs.
func (ibtibq *ItemBatchToItemBatchQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if ibtibq.ctx.Unique == nil && ibtibq.path != nil {
		ibtibq.Unique(true)
	}
	ctx = setContextOp(ctx, ibtibq.ctx, "IDs")
	if err = ibtibq.Select(itembatchtoitembatch.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ibtibq *ItemBatchToItemBatchQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := ibtibq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ibtibq *ItemBatchToItemBatchQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, ibtibq.ctx, "Count")
	if err := ibtibq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, ibtibq, querierCount[*ItemBatchToItemBatchQuery](), ibtibq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (ibtibq *ItemBatchToItemBatchQuery) CountX(ctx context.Context) int {
	count, err := ibtibq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ibtibq *ItemBatchToItemBatchQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, ibtibq.ctx, "Exist")
	switch _, err := ibtibq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (ibtibq *ItemBatchToItemBatchQuery) ExistX(ctx context.Context) bool {
	exist, err := ibtibq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ItemBatchToItemBatchQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ibtibq *ItemBatchToItemBatchQuery) Clone() *ItemBatchToItemBatchQuery {
	if ibtibq == nil {
		return nil
	}
	return &ItemBatchToItemBatchQuery{
		config:     ibtibq.config,
		ctx:        ibtibq.ctx.Clone(),
		order:      append([]OrderFunc{}, ibtibq.order...),
		inters:     append([]Interceptor{}, ibtibq.inters...),
		predicates: append([]predicate.ItemBatchToItemBatch{}, ibtibq.predicates...),
		withParent: ibtibq.withParent.Clone(),
		withChild:  ibtibq.withChild.Clone(),
		// clone intermediate query.
		sql:  ibtibq.sql.Clone(),
		path: ibtibq.path,
	}
}

// WithParent tells the query-builder to eager-load the nodes that are connected to
// the "parent" edge. The optional arguments are used to configure the query builder of the edge.
func (ibtibq *ItemBatchToItemBatchQuery) WithParent(opts ...func(*ItemBatchQuery)) *ItemBatchToItemBatchQuery {
	query := (&ItemBatchClient{config: ibtibq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ibtibq.withParent = query
	return ibtibq
}

// WithChild tells the query-builder to eager-load the nodes that are connected to
// the "child" edge. The optional arguments are used to configure the query builder of the edge.
func (ibtibq *ItemBatchToItemBatchQuery) WithChild(opts ...func(*ItemBatchQuery)) *ItemBatchToItemBatchQuery {
	query := (&ItemBatchClient{config: ibtibq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ibtibq.withChild = query
	return ibtibq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		ChildUUID uuid.UUID `json:"childUUID,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ItemBatchToItemBatch.Query().
//		GroupBy(itembatchtoitembatch.FieldChildUUID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (ibtibq *ItemBatchToItemBatchQuery) GroupBy(field string, fields ...string) *ItemBatchToItemBatchGroupBy {
	ibtibq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ItemBatchToItemBatchGroupBy{build: ibtibq}
	grbuild.flds = &ibtibq.ctx.Fields
	grbuild.label = itembatchtoitembatch.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		ChildUUID uuid.UUID `json:"childUUID,omitempty"`
//	}
//
//	client.ItemBatchToItemBatch.Query().
//		Select(itembatchtoitembatch.FieldChildUUID).
//		Scan(ctx, &v)
func (ibtibq *ItemBatchToItemBatchQuery) Select(fields ...string) *ItemBatchToItemBatchSelect {
	ibtibq.ctx.Fields = append(ibtibq.ctx.Fields, fields...)
	sbuild := &ItemBatchToItemBatchSelect{ItemBatchToItemBatchQuery: ibtibq}
	sbuild.label = itembatchtoitembatch.Label
	sbuild.flds, sbuild.scan = &ibtibq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ItemBatchToItemBatchSelect configured with the given aggregations.
func (ibtibq *ItemBatchToItemBatchQuery) Aggregate(fns ...AggregateFunc) *ItemBatchToItemBatchSelect {
	return ibtibq.Select().Aggregate(fns...)
}

func (ibtibq *ItemBatchToItemBatchQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range ibtibq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, ibtibq); err != nil {
				return err
			}
		}
	}
	for _, f := range ibtibq.ctx.Fields {
		if !itembatchtoitembatch.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ibtibq.path != nil {
		prev, err := ibtibq.path(ctx)
		if err != nil {
			return err
		}
		ibtibq.sql = prev
	}
	return nil
}

func (ibtibq *ItemBatchToItemBatchQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ItemBatchToItemBatch, error) {
	var (
		nodes       = []*ItemBatchToItemBatch{}
		_spec       = ibtibq.querySpec()
		loadedTypes = [2]bool{
			ibtibq.withParent != nil,
			ibtibq.withChild != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ItemBatchToItemBatch).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ItemBatchToItemBatch{config: ibtibq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(ibtibq.modifiers) > 0 {
		_spec.Modifiers = ibtibq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ibtibq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := ibtibq.withParent; query != nil {
		if err := ibtibq.loadParent(ctx, query, nodes, nil,
			func(n *ItemBatchToItemBatch, e *ItemBatch) { n.Edges.Parent = e }); err != nil {
			return nil, err
		}
	}
	if query := ibtibq.withChild; query != nil {
		if err := ibtibq.loadChild(ctx, query, nodes, nil,
			func(n *ItemBatchToItemBatch, e *ItemBatch) { n.Edges.Child = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (ibtibq *ItemBatchToItemBatchQuery) loadParent(ctx context.Context, query *ItemBatchQuery, nodes []*ItemBatchToItemBatch, init func(*ItemBatchToItemBatch), assign func(*ItemBatchToItemBatch, *ItemBatch)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*ItemBatchToItemBatch)
	for i := range nodes {
		fk := nodes[i].ParentUUID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(itembatch.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "parentUUID" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (ibtibq *ItemBatchToItemBatchQuery) loadChild(ctx context.Context, query *ItemBatchQuery, nodes []*ItemBatchToItemBatch, init func(*ItemBatchToItemBatch), assign func(*ItemBatchToItemBatch, *ItemBatch)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*ItemBatchToItemBatch)
	for i := range nodes {
		fk := nodes[i].ChildUUID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(itembatch.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "childUUID" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (ibtibq *ItemBatchToItemBatchQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ibtibq.querySpec()
	if len(ibtibq.modifiers) > 0 {
		_spec.Modifiers = ibtibq.modifiers
	}
	_spec.Node.Columns = ibtibq.ctx.Fields
	if len(ibtibq.ctx.Fields) > 0 {
		_spec.Unique = ibtibq.ctx.Unique != nil && *ibtibq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, ibtibq.driver, _spec)
}

func (ibtibq *ItemBatchToItemBatchQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(itembatchtoitembatch.Table, itembatchtoitembatch.Columns, sqlgraph.NewFieldSpec(itembatchtoitembatch.FieldID, field.TypeUUID))
	_spec.From = ibtibq.sql
	if unique := ibtibq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if ibtibq.path != nil {
		_spec.Unique = true
	}
	if fields := ibtibq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, itembatchtoitembatch.FieldID)
		for i := range fields {
			if fields[i] != itembatchtoitembatch.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ibtibq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ibtibq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ibtibq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ibtibq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ibtibq *ItemBatchToItemBatchQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ibtibq.driver.Dialect())
	t1 := builder.Table(itembatchtoitembatch.Table)
	columns := ibtibq.ctx.Fields
	if len(columns) == 0 {
		columns = itembatchtoitembatch.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ibtibq.sql != nil {
		selector = ibtibq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ibtibq.ctx.Unique != nil && *ibtibq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range ibtibq.modifiers {
		m(selector)
	}
	for _, p := range ibtibq.predicates {
		p(selector)
	}
	for _, p := range ibtibq.order {
		p(selector)
	}
	if offset := ibtibq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ibtibq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (ibtibq *ItemBatchToItemBatchQuery) ForUpdate(opts ...sql.LockOption) *ItemBatchToItemBatchQuery {
	if ibtibq.driver.Dialect() == dialect.Postgres {
		ibtibq.Unique(false)
	}
	ibtibq.modifiers = append(ibtibq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return ibtibq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (ibtibq *ItemBatchToItemBatchQuery) ForShare(opts ...sql.LockOption) *ItemBatchToItemBatchQuery {
	if ibtibq.driver.Dialect() == dialect.Postgres {
		ibtibq.Unique(false)
	}
	ibtibq.modifiers = append(ibtibq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return ibtibq
}

// ItemBatchToItemBatchGroupBy is the group-by builder for ItemBatchToItemBatch entities.
type ItemBatchToItemBatchGroupBy struct {
	selector
	build *ItemBatchToItemBatchQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ibtibgb *ItemBatchToItemBatchGroupBy) Aggregate(fns ...AggregateFunc) *ItemBatchToItemBatchGroupBy {
	ibtibgb.fns = append(ibtibgb.fns, fns...)
	return ibtibgb
}

// Scan applies the selector query and scans the result into the given value.
func (ibtibgb *ItemBatchToItemBatchGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ibtibgb.build.ctx, "GroupBy")
	if err := ibtibgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ItemBatchToItemBatchQuery, *ItemBatchToItemBatchGroupBy](ctx, ibtibgb.build, ibtibgb, ibtibgb.build.inters, v)
}

func (ibtibgb *ItemBatchToItemBatchGroupBy) sqlScan(ctx context.Context, root *ItemBatchToItemBatchQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ibtibgb.fns))
	for _, fn := range ibtibgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ibtibgb.flds)+len(ibtibgb.fns))
		for _, f := range *ibtibgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ibtibgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ibtibgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ItemBatchToItemBatchSelect is the builder for selecting fields of ItemBatchToItemBatch entities.
type ItemBatchToItemBatchSelect struct {
	*ItemBatchToItemBatchQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ibtibs *ItemBatchToItemBatchSelect) Aggregate(fns ...AggregateFunc) *ItemBatchToItemBatchSelect {
	ibtibs.fns = append(ibtibs.fns, fns...)
	return ibtibs
}

// Scan applies the selector query and scans the result into the given value.
func (ibtibs *ItemBatchToItemBatchSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ibtibs.ctx, "Select")
	if err := ibtibs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ItemBatchToItemBatchQuery, *ItemBatchToItemBatchSelect](ctx, ibtibs.ItemBatchToItemBatchQuery, ibtibs, ibtibs.inters, v)
}

func (ibtibs *ItemBatchToItemBatchSelect) sqlScan(ctx context.Context, root *ItemBatchToItemBatchQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ibtibs.fns))
	for _, fn := range ibtibs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ibtibs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ibtibs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
