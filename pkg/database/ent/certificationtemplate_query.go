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

	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent/certificationtemplate"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent/company"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent/predicate"
)

// CertificationTemplateQuery is the builder for querying CertificationTemplate entities.
type CertificationTemplateQuery struct {
	config
	ctx         *QueryContext
	order       []OrderFunc
	inters      []Interceptor
	predicates  []predicate.CertificationTemplate
	withCompany *CompanyQuery
	modifiers   []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CertificationTemplateQuery builder.
func (ctq *CertificationTemplateQuery) Where(ps ...predicate.CertificationTemplate) *CertificationTemplateQuery {
	ctq.predicates = append(ctq.predicates, ps...)
	return ctq
}

// Limit the number of records to be returned by this query.
func (ctq *CertificationTemplateQuery) Limit(limit int) *CertificationTemplateQuery {
	ctq.ctx.Limit = &limit
	return ctq
}

// Offset to start from.
func (ctq *CertificationTemplateQuery) Offset(offset int) *CertificationTemplateQuery {
	ctq.ctx.Offset = &offset
	return ctq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ctq *CertificationTemplateQuery) Unique(unique bool) *CertificationTemplateQuery {
	ctq.ctx.Unique = &unique
	return ctq
}

// Order specifies how the records should be ordered.
func (ctq *CertificationTemplateQuery) Order(o ...OrderFunc) *CertificationTemplateQuery {
	ctq.order = append(ctq.order, o...)
	return ctq
}

// QueryCompany chains the current query on the "company" edge.
func (ctq *CertificationTemplateQuery) QueryCompany() *CompanyQuery {
	query := (&CompanyClient{config: ctq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ctq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ctq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(certificationtemplate.Table, certificationtemplate.FieldID, selector),
			sqlgraph.To(company.Table, company.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, certificationtemplate.CompanyTable, certificationtemplate.CompanyColumn),
		)
		fromU = sqlgraph.SetNeighbors(ctq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first CertificationTemplate entity from the query.
// Returns a *NotFoundError when no CertificationTemplate was found.
func (ctq *CertificationTemplateQuery) First(ctx context.Context) (*CertificationTemplate, error) {
	nodes, err := ctq.Limit(1).All(setContextOp(ctx, ctq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{certificationtemplate.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ctq *CertificationTemplateQuery) FirstX(ctx context.Context) *CertificationTemplate {
	node, err := ctq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first CertificationTemplate ID from the query.
// Returns a *NotFoundError when no CertificationTemplate ID was found.
func (ctq *CertificationTemplateQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = ctq.Limit(1).IDs(setContextOp(ctx, ctq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{certificationtemplate.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ctq *CertificationTemplateQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := ctq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single CertificationTemplate entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one CertificationTemplate entity is found.
// Returns a *NotFoundError when no CertificationTemplate entities are found.
func (ctq *CertificationTemplateQuery) Only(ctx context.Context) (*CertificationTemplate, error) {
	nodes, err := ctq.Limit(2).All(setContextOp(ctx, ctq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{certificationtemplate.Label}
	default:
		return nil, &NotSingularError{certificationtemplate.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ctq *CertificationTemplateQuery) OnlyX(ctx context.Context) *CertificationTemplate {
	node, err := ctq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only CertificationTemplate ID in the query.
// Returns a *NotSingularError when more than one CertificationTemplate ID is found.
// Returns a *NotFoundError when no entities are found.
func (ctq *CertificationTemplateQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = ctq.Limit(2).IDs(setContextOp(ctx, ctq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{certificationtemplate.Label}
	default:
		err = &NotSingularError{certificationtemplate.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ctq *CertificationTemplateQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := ctq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of CertificationTemplates.
func (ctq *CertificationTemplateQuery) All(ctx context.Context) ([]*CertificationTemplate, error) {
	ctx = setContextOp(ctx, ctq.ctx, "All")
	if err := ctq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*CertificationTemplate, *CertificationTemplateQuery]()
	return withInterceptors[[]*CertificationTemplate](ctx, ctq, qr, ctq.inters)
}

// AllX is like All, but panics if an error occurs.
func (ctq *CertificationTemplateQuery) AllX(ctx context.Context) []*CertificationTemplate {
	nodes, err := ctq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of CertificationTemplate IDs.
func (ctq *CertificationTemplateQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if ctq.ctx.Unique == nil && ctq.path != nil {
		ctq.Unique(true)
	}
	ctx = setContextOp(ctx, ctq.ctx, "IDs")
	if err = ctq.Select(certificationtemplate.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ctq *CertificationTemplateQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := ctq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ctq *CertificationTemplateQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, ctq.ctx, "Count")
	if err := ctq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, ctq, querierCount[*CertificationTemplateQuery](), ctq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (ctq *CertificationTemplateQuery) CountX(ctx context.Context) int {
	count, err := ctq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ctq *CertificationTemplateQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, ctq.ctx, "Exist")
	switch _, err := ctq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (ctq *CertificationTemplateQuery) ExistX(ctx context.Context) bool {
	exist, err := ctq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CertificationTemplateQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ctq *CertificationTemplateQuery) Clone() *CertificationTemplateQuery {
	if ctq == nil {
		return nil
	}
	return &CertificationTemplateQuery{
		config:      ctq.config,
		ctx:         ctq.ctx.Clone(),
		order:       append([]OrderFunc{}, ctq.order...),
		inters:      append([]Interceptor{}, ctq.inters...),
		predicates:  append([]predicate.CertificationTemplate{}, ctq.predicates...),
		withCompany: ctq.withCompany.Clone(),
		// clone intermediate query.
		sql:  ctq.sql.Clone(),
		path: ctq.path,
	}
}

// WithCompany tells the query-builder to eager-load the nodes that are connected to
// the "company" edge. The optional arguments are used to configure the query builder of the edge.
func (ctq *CertificationTemplateQuery) WithCompany(opts ...func(*CompanyQuery)) *CertificationTemplateQuery {
	query := (&CompanyClient{config: ctq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ctq.withCompany = query
	return ctq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Description string `json:"description,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	httpclient.CertificationTemplate.Query().
//		GroupBy(certificationtemplate.FieldDescription).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (ctq *CertificationTemplateQuery) GroupBy(field string, fields ...string) *CertificationTemplateGroupBy {
	ctq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &CertificationTemplateGroupBy{build: ctq}
	grbuild.flds = &ctq.ctx.Fields
	grbuild.label = certificationtemplate.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Description string `json:"description,omitempty"`
//	}
//
//	httpclient.CertificationTemplate.Query().
//		Select(certificationtemplate.FieldDescription).
//		Scan(ctx, &v)
func (ctq *CertificationTemplateQuery) Select(fields ...string) *CertificationTemplateSelect {
	ctq.ctx.Fields = append(ctq.ctx.Fields, fields...)
	sbuild := &CertificationTemplateSelect{CertificationTemplateQuery: ctq}
	sbuild.label = certificationtemplate.Label
	sbuild.flds, sbuild.scan = &ctq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a CertificationTemplateSelect configured with the given aggregations.
func (ctq *CertificationTemplateQuery) Aggregate(fns ...AggregateFunc) *CertificationTemplateSelect {
	return ctq.Select().Aggregate(fns...)
}

func (ctq *CertificationTemplateQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range ctq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, ctq); err != nil {
				return err
			}
		}
	}
	for _, f := range ctq.ctx.Fields {
		if !certificationtemplate.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ctq.path != nil {
		prev, err := ctq.path(ctx)
		if err != nil {
			return err
		}
		ctq.sql = prev
	}
	return nil
}

func (ctq *CertificationTemplateQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*CertificationTemplate, error) {
	var (
		nodes       = []*CertificationTemplate{}
		_spec       = ctq.querySpec()
		loadedTypes = [1]bool{
			ctq.withCompany != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*CertificationTemplate).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &CertificationTemplate{config: ctq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(ctq.modifiers) > 0 {
		_spec.Modifiers = ctq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ctq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := ctq.withCompany; query != nil {
		if err := ctq.loadCompany(ctx, query, nodes, nil,
			func(n *CertificationTemplate, e *Company) { n.Edges.Company = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (ctq *CertificationTemplateQuery) loadCompany(ctx context.Context, query *CompanyQuery, nodes []*CertificationTemplate, init func(*CertificationTemplate), assign func(*CertificationTemplate, *Company)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*CertificationTemplate)
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

func (ctq *CertificationTemplateQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ctq.querySpec()
	if len(ctq.modifiers) > 0 {
		_spec.Modifiers = ctq.modifiers
	}
	_spec.Node.Columns = ctq.ctx.Fields
	if len(ctq.ctx.Fields) > 0 {
		_spec.Unique = ctq.ctx.Unique != nil && *ctq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, ctq.driver, _spec)
}

func (ctq *CertificationTemplateQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(certificationtemplate.Table, certificationtemplate.Columns, sqlgraph.NewFieldSpec(certificationtemplate.FieldID, field.TypeUUID))
	_spec.From = ctq.sql
	if unique := ctq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if ctq.path != nil {
		_spec.Unique = true
	}
	if fields := ctq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, certificationtemplate.FieldID)
		for i := range fields {
			if fields[i] != certificationtemplate.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ctq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ctq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ctq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ctq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ctq *CertificationTemplateQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ctq.driver.Dialect())
	t1 := builder.Table(certificationtemplate.Table)
	columns := ctq.ctx.Fields
	if len(columns) == 0 {
		columns = certificationtemplate.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ctq.sql != nil {
		selector = ctq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ctq.ctx.Unique != nil && *ctq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range ctq.modifiers {
		m(selector)
	}
	for _, p := range ctq.predicates {
		p(selector)
	}
	for _, p := range ctq.order {
		p(selector)
	}
	if offset := ctq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ctq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (ctq *CertificationTemplateQuery) ForUpdate(opts ...sql.LockOption) *CertificationTemplateQuery {
	if ctq.driver.Dialect() == dialect.Postgres {
		ctq.Unique(false)
	}
	ctq.modifiers = append(ctq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return ctq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (ctq *CertificationTemplateQuery) ForShare(opts ...sql.LockOption) *CertificationTemplateQuery {
	if ctq.driver.Dialect() == dialect.Postgres {
		ctq.Unique(false)
	}
	ctq.modifiers = append(ctq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return ctq
}

// CertificationTemplateGroupBy is the group-by builder for CertificationTemplate entities.
type CertificationTemplateGroupBy struct {
	selector
	build *CertificationTemplateQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ctgb *CertificationTemplateGroupBy) Aggregate(fns ...AggregateFunc) *CertificationTemplateGroupBy {
	ctgb.fns = append(ctgb.fns, fns...)
	return ctgb
}

// Scan applies the selector query and scans the result into the given value.
func (ctgb *CertificationTemplateGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ctgb.build.ctx, "GroupBy")
	if err := ctgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CertificationTemplateQuery, *CertificationTemplateGroupBy](ctx, ctgb.build, ctgb, ctgb.build.inters, v)
}

func (ctgb *CertificationTemplateGroupBy) sqlScan(ctx context.Context, root *CertificationTemplateQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ctgb.fns))
	for _, fn := range ctgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ctgb.flds)+len(ctgb.fns))
		for _, f := range *ctgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ctgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ctgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// CertificationTemplateSelect is the builder for selecting fields of CertificationTemplate entities.
type CertificationTemplateSelect struct {
	*CertificationTemplateQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (cts *CertificationTemplateSelect) Aggregate(fns ...AggregateFunc) *CertificationTemplateSelect {
	cts.fns = append(cts.fns, fns...)
	return cts
}

// Scan applies the selector query and scans the result into the given value.
func (cts *CertificationTemplateSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cts.ctx, "Select")
	if err := cts.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CertificationTemplateQuery, *CertificationTemplateSelect](ctx, cts.CertificationTemplateQuery, cts, cts.inters, v)
}

func (cts *CertificationTemplateSelect) sqlScan(ctx context.Context, root *CertificationTemplateQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(cts.fns))
	for _, fn := range cts.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*cts.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
