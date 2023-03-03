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

	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent/attributetype"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent/attributetypestotemplates"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent/certificationtemplate"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent/predicate"
)

// AttributeTypesToTemplatesQuery is the builder for querying AttributeTypesToTemplates entities.
type AttributeTypesToTemplatesQuery struct {
	config
	ctx           *QueryContext
	order         []OrderFunc
	inters        []Interceptor
	predicates    []predicate.AttributeTypesToTemplates
	withAttribute *AttributeTypeQuery
	withTemplate  *CertificationTemplateQuery
	modifiers     []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AttributeTypesToTemplatesQuery builder.
func (atttq *AttributeTypesToTemplatesQuery) Where(ps ...predicate.AttributeTypesToTemplates) *AttributeTypesToTemplatesQuery {
	atttq.predicates = append(atttq.predicates, ps...)
	return atttq
}

// Limit the number of records to be returned by this query.
func (atttq *AttributeTypesToTemplatesQuery) Limit(limit int) *AttributeTypesToTemplatesQuery {
	atttq.ctx.Limit = &limit
	return atttq
}

// Offset to start from.
func (atttq *AttributeTypesToTemplatesQuery) Offset(offset int) *AttributeTypesToTemplatesQuery {
	atttq.ctx.Offset = &offset
	return atttq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (atttq *AttributeTypesToTemplatesQuery) Unique(unique bool) *AttributeTypesToTemplatesQuery {
	atttq.ctx.Unique = &unique
	return atttq
}

// Order specifies how the records should be ordered.
func (atttq *AttributeTypesToTemplatesQuery) Order(o ...OrderFunc) *AttributeTypesToTemplatesQuery {
	atttq.order = append(atttq.order, o...)
	return atttq
}

// QueryAttribute chains the current query on the "attribute" edge.
func (atttq *AttributeTypesToTemplatesQuery) QueryAttribute() *AttributeTypeQuery {
	query := (&AttributeTypeClient{config: atttq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := atttq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := atttq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(attributetypestotemplates.Table, attributetypestotemplates.FieldID, selector),
			sqlgraph.To(attributetype.Table, attributetype.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, attributetypestotemplates.AttributeTable, attributetypestotemplates.AttributeColumn),
		)
		fromU = sqlgraph.SetNeighbors(atttq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTemplate chains the current query on the "template" edge.
func (atttq *AttributeTypesToTemplatesQuery) QueryTemplate() *CertificationTemplateQuery {
	query := (&CertificationTemplateClient{config: atttq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := atttq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := atttq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(attributetypestotemplates.Table, attributetypestotemplates.FieldID, selector),
			sqlgraph.To(certificationtemplate.Table, certificationtemplate.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, attributetypestotemplates.TemplateTable, attributetypestotemplates.TemplateColumn),
		)
		fromU = sqlgraph.SetNeighbors(atttq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first AttributeTypesToTemplates entity from the query.
// Returns a *NotFoundError when no AttributeTypesToTemplates was found.
func (atttq *AttributeTypesToTemplatesQuery) First(ctx context.Context) (*AttributeTypesToTemplates, error) {
	nodes, err := atttq.Limit(1).All(setContextOp(ctx, atttq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{attributetypestotemplates.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (atttq *AttributeTypesToTemplatesQuery) FirstX(ctx context.Context) *AttributeTypesToTemplates {
	node, err := atttq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AttributeTypesToTemplates ID from the query.
// Returns a *NotFoundError when no AttributeTypesToTemplates ID was found.
func (atttq *AttributeTypesToTemplatesQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = atttq.Limit(1).IDs(setContextOp(ctx, atttq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{attributetypestotemplates.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (atttq *AttributeTypesToTemplatesQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := atttq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AttributeTypesToTemplates entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one AttributeTypesToTemplates entity is found.
// Returns a *NotFoundError when no AttributeTypesToTemplates entities are found.
func (atttq *AttributeTypesToTemplatesQuery) Only(ctx context.Context) (*AttributeTypesToTemplates, error) {
	nodes, err := atttq.Limit(2).All(setContextOp(ctx, atttq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{attributetypestotemplates.Label}
	default:
		return nil, &NotSingularError{attributetypestotemplates.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (atttq *AttributeTypesToTemplatesQuery) OnlyX(ctx context.Context) *AttributeTypesToTemplates {
	node, err := atttq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AttributeTypesToTemplates ID in the query.
// Returns a *NotSingularError when more than one AttributeTypesToTemplates ID is found.
// Returns a *NotFoundError when no entities are found.
func (atttq *AttributeTypesToTemplatesQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = atttq.Limit(2).IDs(setContextOp(ctx, atttq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{attributetypestotemplates.Label}
	default:
		err = &NotSingularError{attributetypestotemplates.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (atttq *AttributeTypesToTemplatesQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := atttq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AttributeTypesToTemplatesSlice.
func (atttq *AttributeTypesToTemplatesQuery) All(ctx context.Context) ([]*AttributeTypesToTemplates, error) {
	ctx = setContextOp(ctx, atttq.ctx, "All")
	if err := atttq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*AttributeTypesToTemplates, *AttributeTypesToTemplatesQuery]()
	return withInterceptors[[]*AttributeTypesToTemplates](ctx, atttq, qr, atttq.inters)
}

// AllX is like All, but panics if an error occurs.
func (atttq *AttributeTypesToTemplatesQuery) AllX(ctx context.Context) []*AttributeTypesToTemplates {
	nodes, err := atttq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AttributeTypesToTemplates IDs.
func (atttq *AttributeTypesToTemplatesQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if atttq.ctx.Unique == nil && atttq.path != nil {
		atttq.Unique(true)
	}
	ctx = setContextOp(ctx, atttq.ctx, "IDs")
	if err = atttq.Select(attributetypestotemplates.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (atttq *AttributeTypesToTemplatesQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := atttq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (atttq *AttributeTypesToTemplatesQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, atttq.ctx, "Count")
	if err := atttq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, atttq, querierCount[*AttributeTypesToTemplatesQuery](), atttq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (atttq *AttributeTypesToTemplatesQuery) CountX(ctx context.Context) int {
	count, err := atttq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (atttq *AttributeTypesToTemplatesQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, atttq.ctx, "Exist")
	switch _, err := atttq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (atttq *AttributeTypesToTemplatesQuery) ExistX(ctx context.Context) bool {
	exist, err := atttq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AttributeTypesToTemplatesQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (atttq *AttributeTypesToTemplatesQuery) Clone() *AttributeTypesToTemplatesQuery {
	if atttq == nil {
		return nil
	}
	return &AttributeTypesToTemplatesQuery{
		config:        atttq.config,
		ctx:           atttq.ctx.Clone(),
		order:         append([]OrderFunc{}, atttq.order...),
		inters:        append([]Interceptor{}, atttq.inters...),
		predicates:    append([]predicate.AttributeTypesToTemplates{}, atttq.predicates...),
		withAttribute: atttq.withAttribute.Clone(),
		withTemplate:  atttq.withTemplate.Clone(),
		// clone intermediate query.
		sql:  atttq.sql.Clone(),
		path: atttq.path,
	}
}

// WithAttribute tells the query-builder to eager-load the nodes that are connected to
// the "attribute" edge. The optional arguments are used to configure the query builder of the edge.
func (atttq *AttributeTypesToTemplatesQuery) WithAttribute(opts ...func(*AttributeTypeQuery)) *AttributeTypesToTemplatesQuery {
	query := (&AttributeTypeClient{config: atttq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	atttq.withAttribute = query
	return atttq
}

// WithTemplate tells the query-builder to eager-load the nodes that are connected to
// the "template" edge. The optional arguments are used to configure the query builder of the edge.
func (atttq *AttributeTypesToTemplatesQuery) WithTemplate(opts ...func(*CertificationTemplateQuery)) *AttributeTypesToTemplatesQuery {
	query := (&CertificationTemplateClient{config: atttq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	atttq.withTemplate = query
	return atttq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		AttributeTypeUUID uuid.UUID `json:"attributeTypeUUID,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.AttributeTypesToTemplates.Query().
//		GroupBy(attributetypestotemplates.FieldAttributeTypeUUID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (atttq *AttributeTypesToTemplatesQuery) GroupBy(field string, fields ...string) *AttributeTypesToTemplatesGroupBy {
	atttq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &AttributeTypesToTemplatesGroupBy{build: atttq}
	grbuild.flds = &atttq.ctx.Fields
	grbuild.label = attributetypestotemplates.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		AttributeTypeUUID uuid.UUID `json:"attributeTypeUUID,omitempty"`
//	}
//
//	client.AttributeTypesToTemplates.Query().
//		Select(attributetypestotemplates.FieldAttributeTypeUUID).
//		Scan(ctx, &v)
func (atttq *AttributeTypesToTemplatesQuery) Select(fields ...string) *AttributeTypesToTemplatesSelect {
	atttq.ctx.Fields = append(atttq.ctx.Fields, fields...)
	sbuild := &AttributeTypesToTemplatesSelect{AttributeTypesToTemplatesQuery: atttq}
	sbuild.label = attributetypestotemplates.Label
	sbuild.flds, sbuild.scan = &atttq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a AttributeTypesToTemplatesSelect configured with the given aggregations.
func (atttq *AttributeTypesToTemplatesQuery) Aggregate(fns ...AggregateFunc) *AttributeTypesToTemplatesSelect {
	return atttq.Select().Aggregate(fns...)
}

func (atttq *AttributeTypesToTemplatesQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range atttq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, atttq); err != nil {
				return err
			}
		}
	}
	for _, f := range atttq.ctx.Fields {
		if !attributetypestotemplates.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if atttq.path != nil {
		prev, err := atttq.path(ctx)
		if err != nil {
			return err
		}
		atttq.sql = prev
	}
	return nil
}

func (atttq *AttributeTypesToTemplatesQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*AttributeTypesToTemplates, error) {
	var (
		nodes       = []*AttributeTypesToTemplates{}
		_spec       = atttq.querySpec()
		loadedTypes = [2]bool{
			atttq.withAttribute != nil,
			atttq.withTemplate != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*AttributeTypesToTemplates).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &AttributeTypesToTemplates{config: atttq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(atttq.modifiers) > 0 {
		_spec.Modifiers = atttq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, atttq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := atttq.withAttribute; query != nil {
		if err := atttq.loadAttribute(ctx, query, nodes, nil,
			func(n *AttributeTypesToTemplates, e *AttributeType) { n.Edges.Attribute = e }); err != nil {
			return nil, err
		}
	}
	if query := atttq.withTemplate; query != nil {
		if err := atttq.loadTemplate(ctx, query, nodes, nil,
			func(n *AttributeTypesToTemplates, e *CertificationTemplate) { n.Edges.Template = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (atttq *AttributeTypesToTemplatesQuery) loadAttribute(ctx context.Context, query *AttributeTypeQuery, nodes []*AttributeTypesToTemplates, init func(*AttributeTypesToTemplates), assign func(*AttributeTypesToTemplates, *AttributeType)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*AttributeTypesToTemplates)
	for i := range nodes {
		fk := nodes[i].AttributeTypeUUID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(attributetype.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "attributeTypeUUID" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (atttq *AttributeTypesToTemplatesQuery) loadTemplate(ctx context.Context, query *CertificationTemplateQuery, nodes []*AttributeTypesToTemplates, init func(*AttributeTypesToTemplates), assign func(*AttributeTypesToTemplates, *CertificationTemplate)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*AttributeTypesToTemplates)
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

func (atttq *AttributeTypesToTemplatesQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := atttq.querySpec()
	if len(atttq.modifiers) > 0 {
		_spec.Modifiers = atttq.modifiers
	}
	_spec.Node.Columns = atttq.ctx.Fields
	if len(atttq.ctx.Fields) > 0 {
		_spec.Unique = atttq.ctx.Unique != nil && *atttq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, atttq.driver, _spec)
}

func (atttq *AttributeTypesToTemplatesQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(attributetypestotemplates.Table, attributetypestotemplates.Columns, sqlgraph.NewFieldSpec(attributetypestotemplates.FieldID, field.TypeUUID))
	_spec.From = atttq.sql
	if unique := atttq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if atttq.path != nil {
		_spec.Unique = true
	}
	if fields := atttq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, attributetypestotemplates.FieldID)
		for i := range fields {
			if fields[i] != attributetypestotemplates.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := atttq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := atttq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := atttq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := atttq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (atttq *AttributeTypesToTemplatesQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(atttq.driver.Dialect())
	t1 := builder.Table(attributetypestotemplates.Table)
	columns := atttq.ctx.Fields
	if len(columns) == 0 {
		columns = attributetypestotemplates.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if atttq.sql != nil {
		selector = atttq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if atttq.ctx.Unique != nil && *atttq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range atttq.modifiers {
		m(selector)
	}
	for _, p := range atttq.predicates {
		p(selector)
	}
	for _, p := range atttq.order {
		p(selector)
	}
	if offset := atttq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := atttq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (atttq *AttributeTypesToTemplatesQuery) ForUpdate(opts ...sql.LockOption) *AttributeTypesToTemplatesQuery {
	if atttq.driver.Dialect() == dialect.Postgres {
		atttq.Unique(false)
	}
	atttq.modifiers = append(atttq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return atttq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (atttq *AttributeTypesToTemplatesQuery) ForShare(opts ...sql.LockOption) *AttributeTypesToTemplatesQuery {
	if atttq.driver.Dialect() == dialect.Postgres {
		atttq.Unique(false)
	}
	atttq.modifiers = append(atttq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return atttq
}

// AttributeTypesToTemplatesGroupBy is the group-by builder for AttributeTypesToTemplates entities.
type AttributeTypesToTemplatesGroupBy struct {
	selector
	build *AttributeTypesToTemplatesQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (atttgb *AttributeTypesToTemplatesGroupBy) Aggregate(fns ...AggregateFunc) *AttributeTypesToTemplatesGroupBy {
	atttgb.fns = append(atttgb.fns, fns...)
	return atttgb
}

// Scan applies the selector query and scans the result into the given value.
func (atttgb *AttributeTypesToTemplatesGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, atttgb.build.ctx, "GroupBy")
	if err := atttgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AttributeTypesToTemplatesQuery, *AttributeTypesToTemplatesGroupBy](ctx, atttgb.build, atttgb, atttgb.build.inters, v)
}

func (atttgb *AttributeTypesToTemplatesGroupBy) sqlScan(ctx context.Context, root *AttributeTypesToTemplatesQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(atttgb.fns))
	for _, fn := range atttgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*atttgb.flds)+len(atttgb.fns))
		for _, f := range *atttgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*atttgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := atttgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// AttributeTypesToTemplatesSelect is the builder for selecting fields of AttributeTypesToTemplates entities.
type AttributeTypesToTemplatesSelect struct {
	*AttributeTypesToTemplatesQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (attts *AttributeTypesToTemplatesSelect) Aggregate(fns ...AggregateFunc) *AttributeTypesToTemplatesSelect {
	attts.fns = append(attts.fns, fns...)
	return attts
}

// Scan applies the selector query and scans the result into the given value.
func (attts *AttributeTypesToTemplatesSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, attts.ctx, "Select")
	if err := attts.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AttributeTypesToTemplatesQuery, *AttributeTypesToTemplatesSelect](ctx, attts.AttributeTypesToTemplatesQuery, attts, attts.inters, v)
}

func (attts *AttributeTypesToTemplatesSelect) sqlScan(ctx context.Context, root *AttributeTypesToTemplatesQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(attts.fns))
	for _, fn := range attts.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*attts.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := attts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}