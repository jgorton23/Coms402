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

	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent/certification"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent/certificationtemplate"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent/company"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent/itembatch"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent/predicate"
)

// CertificationQuery is the builder for querying Certification entities.
type CertificationQuery struct {
	config
	ctx           *QueryContext
	order         []OrderFunc
	inters        []Interceptor
	predicates    []predicate.Certification
	withCompany   *CompanyQuery
	withItemBatch *ItemBatchQuery
	withTemplate  *CertificationTemplateQuery
	modifiers     []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CertificationQuery builder.
func (cq *CertificationQuery) Where(ps ...predicate.Certification) *CertificationQuery {
	cq.predicates = append(cq.predicates, ps...)
	return cq
}

// Limit the number of records to be returned by this query.
func (cq *CertificationQuery) Limit(limit int) *CertificationQuery {
	cq.ctx.Limit = &limit
	return cq
}

// Offset to start from.
func (cq *CertificationQuery) Offset(offset int) *CertificationQuery {
	cq.ctx.Offset = &offset
	return cq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (cq *CertificationQuery) Unique(unique bool) *CertificationQuery {
	cq.ctx.Unique = &unique
	return cq
}

// Order specifies how the records should be ordered.
func (cq *CertificationQuery) Order(o ...OrderFunc) *CertificationQuery {
	cq.order = append(cq.order, o...)
	return cq
}

// QueryCompany chains the current query on the "company" edge.
func (cq *CertificationQuery) QueryCompany() *CompanyQuery {
	query := (&CompanyClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(certification.Table, certification.FieldID, selector),
			sqlgraph.To(company.Table, company.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, certification.CompanyTable, certification.CompanyColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryItemBatch chains the current query on the "itemBatch" edge.
func (cq *CertificationQuery) QueryItemBatch() *ItemBatchQuery {
	query := (&ItemBatchClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(certification.Table, certification.FieldID, selector),
			sqlgraph.To(itembatch.Table, itembatch.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, certification.ItemBatchTable, certification.ItemBatchColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTemplate chains the current query on the "template" edge.
func (cq *CertificationQuery) QueryTemplate() *CertificationTemplateQuery {
	query := (&CertificationTemplateClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(certification.Table, certification.FieldID, selector),
			sqlgraph.To(certificationtemplate.Table, certificationtemplate.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, certification.TemplateTable, certification.TemplateColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Certification entity from the query.
// Returns a *NotFoundError when no Certification was found.
func (cq *CertificationQuery) First(ctx context.Context) (*Certification, error) {
	nodes, err := cq.Limit(1).All(setContextOp(ctx, cq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{certification.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cq *CertificationQuery) FirstX(ctx context.Context) *Certification {
	node, err := cq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Certification ID from the query.
// Returns a *NotFoundError when no Certification ID was found.
func (cq *CertificationQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = cq.Limit(1).IDs(setContextOp(ctx, cq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{certification.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (cq *CertificationQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := cq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Certification entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Certification entity is found.
// Returns a *NotFoundError when no Certification entities are found.
func (cq *CertificationQuery) Only(ctx context.Context) (*Certification, error) {
	nodes, err := cq.Limit(2).All(setContextOp(ctx, cq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{certification.Label}
	default:
		return nil, &NotSingularError{certification.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cq *CertificationQuery) OnlyX(ctx context.Context) *Certification {
	node, err := cq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Certification ID in the query.
// Returns a *NotSingularError when more than one Certification ID is found.
// Returns a *NotFoundError when no entities are found.
func (cq *CertificationQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = cq.Limit(2).IDs(setContextOp(ctx, cq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{certification.Label}
	default:
		err = &NotSingularError{certification.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (cq *CertificationQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := cq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Certifications.
func (cq *CertificationQuery) All(ctx context.Context) ([]*Certification, error) {
	ctx = setContextOp(ctx, cq.ctx, "All")
	if err := cq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Certification, *CertificationQuery]()
	return withInterceptors[[]*Certification](ctx, cq, qr, cq.inters)
}

// AllX is like All, but panics if an error occurs.
func (cq *CertificationQuery) AllX(ctx context.Context) []*Certification {
	nodes, err := cq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Certification IDs.
func (cq *CertificationQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if cq.ctx.Unique == nil && cq.path != nil {
		cq.Unique(true)
	}
	ctx = setContextOp(ctx, cq.ctx, "IDs")
	if err = cq.Select(certification.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (cq *CertificationQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := cq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (cq *CertificationQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, cq.ctx, "Count")
	if err := cq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, cq, querierCount[*CertificationQuery](), cq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (cq *CertificationQuery) CountX(ctx context.Context) int {
	count, err := cq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cq *CertificationQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, cq.ctx, "Exist")
	switch _, err := cq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (cq *CertificationQuery) ExistX(ctx context.Context) bool {
	exist, err := cq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CertificationQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cq *CertificationQuery) Clone() *CertificationQuery {
	if cq == nil {
		return nil
	}
	return &CertificationQuery{
		config:        cq.config,
		ctx:           cq.ctx.Clone(),
		order:         append([]OrderFunc{}, cq.order...),
		inters:        append([]Interceptor{}, cq.inters...),
		predicates:    append([]predicate.Certification{}, cq.predicates...),
		withCompany:   cq.withCompany.Clone(),
		withItemBatch: cq.withItemBatch.Clone(),
		withTemplate:  cq.withTemplate.Clone(),
		// clone intermediate query.
		sql:  cq.sql.Clone(),
		path: cq.path,
	}
}

// WithCompany tells the query-builder to eager-load the nodes that are connected to
// the "company" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *CertificationQuery) WithCompany(opts ...func(*CompanyQuery)) *CertificationQuery {
	query := (&CompanyClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withCompany = query
	return cq
}

// WithItemBatch tells the query-builder to eager-load the nodes that are connected to
// the "itemBatch" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *CertificationQuery) WithItemBatch(opts ...func(*ItemBatchQuery)) *CertificationQuery {
	query := (&ItemBatchClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withItemBatch = query
	return cq
}

// WithTemplate tells the query-builder to eager-load the nodes that are connected to
// the "template" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *CertificationQuery) WithTemplate(opts ...func(*CertificationTemplateQuery)) *CertificationQuery {
	query := (&CertificationTemplateClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withTemplate = query
	return cq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		PrimaryAttribute string `json:"primaryAttribute,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Certification.Query().
//		GroupBy(certification.FieldPrimaryAttribute).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (cq *CertificationQuery) GroupBy(field string, fields ...string) *CertificationGroupBy {
	cq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &CertificationGroupBy{build: cq}
	grbuild.flds = &cq.ctx.Fields
	grbuild.label = certification.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		PrimaryAttribute string `json:"primaryAttribute,omitempty"`
//	}
//
//	client.Certification.Query().
//		Select(certification.FieldPrimaryAttribute).
//		Scan(ctx, &v)
func (cq *CertificationQuery) Select(fields ...string) *CertificationSelect {
	cq.ctx.Fields = append(cq.ctx.Fields, fields...)
	sbuild := &CertificationSelect{CertificationQuery: cq}
	sbuild.label = certification.Label
	sbuild.flds, sbuild.scan = &cq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a CertificationSelect configured with the given aggregations.
func (cq *CertificationQuery) Aggregate(fns ...AggregateFunc) *CertificationSelect {
	return cq.Select().Aggregate(fns...)
}

func (cq *CertificationQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range cq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, cq); err != nil {
				return err
			}
		}
	}
	for _, f := range cq.ctx.Fields {
		if !certification.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if cq.path != nil {
		prev, err := cq.path(ctx)
		if err != nil {
			return err
		}
		cq.sql = prev
	}
	return nil
}

func (cq *CertificationQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Certification, error) {
	var (
		nodes       = []*Certification{}
		_spec       = cq.querySpec()
		loadedTypes = [3]bool{
			cq.withCompany != nil,
			cq.withItemBatch != nil,
			cq.withTemplate != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Certification).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Certification{config: cq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(cq.modifiers) > 0 {
		_spec.Modifiers = cq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, cq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := cq.withCompany; query != nil {
		if err := cq.loadCompany(ctx, query, nodes, nil,
			func(n *Certification, e *Company) { n.Edges.Company = e }); err != nil {
			return nil, err
		}
	}
	if query := cq.withItemBatch; query != nil {
		if err := cq.loadItemBatch(ctx, query, nodes, nil,
			func(n *Certification, e *ItemBatch) { n.Edges.ItemBatch = e }); err != nil {
			return nil, err
		}
	}
	if query := cq.withTemplate; query != nil {
		if err := cq.loadTemplate(ctx, query, nodes, nil,
			func(n *Certification, e *CertificationTemplate) { n.Edges.Template = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (cq *CertificationQuery) loadCompany(ctx context.Context, query *CompanyQuery, nodes []*Certification, init func(*Certification), assign func(*Certification, *Company)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Certification)
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
func (cq *CertificationQuery) loadItemBatch(ctx context.Context, query *ItemBatchQuery, nodes []*Certification, init func(*Certification), assign func(*Certification, *ItemBatch)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Certification)
	for i := range nodes {
		fk := nodes[i].ItemBatchUUID
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
			return fmt.Errorf(`unexpected foreign-key "itemBatchUUID" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (cq *CertificationQuery) loadTemplate(ctx context.Context, query *CertificationTemplateQuery, nodes []*Certification, init func(*Certification), assign func(*Certification, *CertificationTemplate)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Certification)
	for i := range nodes {
		fk := nodes[i].TemplateUUID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(certificationtemplate.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "templateUUID" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (cq *CertificationQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := cq.querySpec()
	if len(cq.modifiers) > 0 {
		_spec.Modifiers = cq.modifiers
	}
	_spec.Node.Columns = cq.ctx.Fields
	if len(cq.ctx.Fields) > 0 {
		_spec.Unique = cq.ctx.Unique != nil && *cq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, cq.driver, _spec)
}

func (cq *CertificationQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(certification.Table, certification.Columns, sqlgraph.NewFieldSpec(certification.FieldID, field.TypeUUID))
	_spec.From = cq.sql
	if unique := cq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if cq.path != nil {
		_spec.Unique = true
	}
	if fields := cq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, certification.FieldID)
		for i := range fields {
			if fields[i] != certification.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := cq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := cq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := cq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := cq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (cq *CertificationQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(cq.driver.Dialect())
	t1 := builder.Table(certification.Table)
	columns := cq.ctx.Fields
	if len(columns) == 0 {
		columns = certification.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if cq.sql != nil {
		selector = cq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if cq.ctx.Unique != nil && *cq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range cq.modifiers {
		m(selector)
	}
	for _, p := range cq.predicates {
		p(selector)
	}
	for _, p := range cq.order {
		p(selector)
	}
	if offset := cq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := cq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (cq *CertificationQuery) ForUpdate(opts ...sql.LockOption) *CertificationQuery {
	if cq.driver.Dialect() == dialect.Postgres {
		cq.Unique(false)
	}
	cq.modifiers = append(cq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return cq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (cq *CertificationQuery) ForShare(opts ...sql.LockOption) *CertificationQuery {
	if cq.driver.Dialect() == dialect.Postgres {
		cq.Unique(false)
	}
	cq.modifiers = append(cq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return cq
}

// CertificationGroupBy is the group-by builder for Certification entities.
type CertificationGroupBy struct {
	selector
	build *CertificationQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cgb *CertificationGroupBy) Aggregate(fns ...AggregateFunc) *CertificationGroupBy {
	cgb.fns = append(cgb.fns, fns...)
	return cgb
}

// Scan applies the selector query and scans the result into the given value.
func (cgb *CertificationGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cgb.build.ctx, "GroupBy")
	if err := cgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CertificationQuery, *CertificationGroupBy](ctx, cgb.build, cgb, cgb.build.inters, v)
}

func (cgb *CertificationGroupBy) sqlScan(ctx context.Context, root *CertificationQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(cgb.fns))
	for _, fn := range cgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*cgb.flds)+len(cgb.fns))
		for _, f := range *cgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*cgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// CertificationSelect is the builder for selecting fields of Certification entities.
type CertificationSelect struct {
	*CertificationQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (cs *CertificationSelect) Aggregate(fns ...AggregateFunc) *CertificationSelect {
	cs.fns = append(cs.fns, fns...)
	return cs
}

// Scan applies the selector query and scans the result into the given value.
func (cs *CertificationSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cs.ctx, "Select")
	if err := cs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CertificationQuery, *CertificationSelect](ctx, cs.CertificationQuery, cs, cs.inters, v)
}

func (cs *CertificationSelect) sqlScan(ctx context.Context, root *CertificationQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(cs.fns))
	for _, fn := range cs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*cs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}