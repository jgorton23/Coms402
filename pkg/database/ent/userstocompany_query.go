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
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent/predicate"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent/user"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent/userstocompany"
)

// UsersToCompanyQuery is the builder for querying UsersToCompany entities.
type UsersToCompanyQuery struct {
	config
	ctx         *QueryContext
	order       []OrderFunc
	inters      []Interceptor
	predicates  []predicate.UsersToCompany
	withUser    *UserQuery
	withCompany *CompanyQuery
	modifiers   []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the UsersToCompanyQuery builder.
func (utcq *UsersToCompanyQuery) Where(ps ...predicate.UsersToCompany) *UsersToCompanyQuery {
	utcq.predicates = append(utcq.predicates, ps...)
	return utcq
}

// Limit the number of records to be returned by this query.
func (utcq *UsersToCompanyQuery) Limit(limit int) *UsersToCompanyQuery {
	utcq.ctx.Limit = &limit
	return utcq
}

// Offset to start from.
func (utcq *UsersToCompanyQuery) Offset(offset int) *UsersToCompanyQuery {
	utcq.ctx.Offset = &offset
	return utcq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (utcq *UsersToCompanyQuery) Unique(unique bool) *UsersToCompanyQuery {
	utcq.ctx.Unique = &unique
	return utcq
}

// Order specifies how the records should be ordered.
func (utcq *UsersToCompanyQuery) Order(o ...OrderFunc) *UsersToCompanyQuery {
	utcq.order = append(utcq.order, o...)
	return utcq
}

// QueryUser chains the current query on the "user" edge.
func (utcq *UsersToCompanyQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: utcq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := utcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := utcq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(userstocompany.Table, userstocompany.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, userstocompany.UserTable, userstocompany.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(utcq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCompany chains the current query on the "company" edge.
func (utcq *UsersToCompanyQuery) QueryCompany() *CompanyQuery {
	query := (&CompanyClient{config: utcq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := utcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := utcq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(userstocompany.Table, userstocompany.FieldID, selector),
			sqlgraph.To(company.Table, company.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, userstocompany.CompanyTable, userstocompany.CompanyColumn),
		)
		fromU = sqlgraph.SetNeighbors(utcq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first UsersToCompany entity from the query.
// Returns a *NotFoundError when no UsersToCompany was found.
func (utcq *UsersToCompanyQuery) First(ctx context.Context) (*UsersToCompany, error) {
	nodes, err := utcq.Limit(1).All(setContextOp(ctx, utcq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{userstocompany.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (utcq *UsersToCompanyQuery) FirstX(ctx context.Context) *UsersToCompany {
	node, err := utcq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first UsersToCompany ID from the query.
// Returns a *NotFoundError when no UsersToCompany ID was found.
func (utcq *UsersToCompanyQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = utcq.Limit(1).IDs(setContextOp(ctx, utcq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{userstocompany.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (utcq *UsersToCompanyQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := utcq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single UsersToCompany entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one UsersToCompany entity is found.
// Returns a *NotFoundError when no UsersToCompany entities are found.
func (utcq *UsersToCompanyQuery) Only(ctx context.Context) (*UsersToCompany, error) {
	nodes, err := utcq.Limit(2).All(setContextOp(ctx, utcq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{userstocompany.Label}
	default:
		return nil, &NotSingularError{userstocompany.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (utcq *UsersToCompanyQuery) OnlyX(ctx context.Context) *UsersToCompany {
	node, err := utcq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only UsersToCompany ID in the query.
// Returns a *NotSingularError when more than one UsersToCompany ID is found.
// Returns a *NotFoundError when no entities are found.
func (utcq *UsersToCompanyQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = utcq.Limit(2).IDs(setContextOp(ctx, utcq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{userstocompany.Label}
	default:
		err = &NotSingularError{userstocompany.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (utcq *UsersToCompanyQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := utcq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of UsersToCompanies.
func (utcq *UsersToCompanyQuery) All(ctx context.Context) ([]*UsersToCompany, error) {
	ctx = setContextOp(ctx, utcq.ctx, "All")
	if err := utcq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*UsersToCompany, *UsersToCompanyQuery]()
	return withInterceptors[[]*UsersToCompany](ctx, utcq, qr, utcq.inters)
}

// AllX is like All, but panics if an error occurs.
func (utcq *UsersToCompanyQuery) AllX(ctx context.Context) []*UsersToCompany {
	nodes, err := utcq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of UsersToCompany IDs.
func (utcq *UsersToCompanyQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if utcq.ctx.Unique == nil && utcq.path != nil {
		utcq.Unique(true)
	}
	ctx = setContextOp(ctx, utcq.ctx, "IDs")
	if err = utcq.Select(userstocompany.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (utcq *UsersToCompanyQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := utcq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (utcq *UsersToCompanyQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, utcq.ctx, "Count")
	if err := utcq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, utcq, querierCount[*UsersToCompanyQuery](), utcq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (utcq *UsersToCompanyQuery) CountX(ctx context.Context) int {
	count, err := utcq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (utcq *UsersToCompanyQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, utcq.ctx, "Exist")
	switch _, err := utcq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (utcq *UsersToCompanyQuery) ExistX(ctx context.Context) bool {
	exist, err := utcq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UsersToCompanyQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (utcq *UsersToCompanyQuery) Clone() *UsersToCompanyQuery {
	if utcq == nil {
		return nil
	}
	return &UsersToCompanyQuery{
		config:      utcq.config,
		ctx:         utcq.ctx.Clone(),
		order:       append([]OrderFunc{}, utcq.order...),
		inters:      append([]Interceptor{}, utcq.inters...),
		predicates:  append([]predicate.UsersToCompany{}, utcq.predicates...),
		withUser:    utcq.withUser.Clone(),
		withCompany: utcq.withCompany.Clone(),
		// clone intermediate query.
		sql:  utcq.sql.Clone(),
		path: utcq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (utcq *UsersToCompanyQuery) WithUser(opts ...func(*UserQuery)) *UsersToCompanyQuery {
	query := (&UserClient{config: utcq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	utcq.withUser = query
	return utcq
}

// WithCompany tells the query-builder to eager-load the nodes that are connected to
// the "company" edge. The optional arguments are used to configure the query builder of the edge.
func (utcq *UsersToCompanyQuery) WithCompany(opts ...func(*CompanyQuery)) *UsersToCompanyQuery {
	query := (&CompanyClient{config: utcq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	utcq.withCompany = query
	return utcq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CompanyUUID uuid.UUID `json:"companyUUID,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.UsersToCompany.Query().
//		GroupBy(userstocompany.FieldCompanyUUID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (utcq *UsersToCompanyQuery) GroupBy(field string, fields ...string) *UsersToCompanyGroupBy {
	utcq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &UsersToCompanyGroupBy{build: utcq}
	grbuild.flds = &utcq.ctx.Fields
	grbuild.label = userstocompany.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CompanyUUID uuid.UUID `json:"companyUUID,omitempty"`
//	}
//
//	client.UsersToCompany.Query().
//		Select(userstocompany.FieldCompanyUUID).
//		Scan(ctx, &v)
func (utcq *UsersToCompanyQuery) Select(fields ...string) *UsersToCompanySelect {
	utcq.ctx.Fields = append(utcq.ctx.Fields, fields...)
	sbuild := &UsersToCompanySelect{UsersToCompanyQuery: utcq}
	sbuild.label = userstocompany.Label
	sbuild.flds, sbuild.scan = &utcq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a UsersToCompanySelect configured with the given aggregations.
func (utcq *UsersToCompanyQuery) Aggregate(fns ...AggregateFunc) *UsersToCompanySelect {
	return utcq.Select().Aggregate(fns...)
}

func (utcq *UsersToCompanyQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range utcq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, utcq); err != nil {
				return err
			}
		}
	}
	for _, f := range utcq.ctx.Fields {
		if !userstocompany.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if utcq.path != nil {
		prev, err := utcq.path(ctx)
		if err != nil {
			return err
		}
		utcq.sql = prev
	}
	return nil
}

func (utcq *UsersToCompanyQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*UsersToCompany, error) {
	var (
		nodes       = []*UsersToCompany{}
		_spec       = utcq.querySpec()
		loadedTypes = [2]bool{
			utcq.withUser != nil,
			utcq.withCompany != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*UsersToCompany).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &UsersToCompany{config: utcq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(utcq.modifiers) > 0 {
		_spec.Modifiers = utcq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, utcq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := utcq.withUser; query != nil {
		if err := utcq.loadUser(ctx, query, nodes, nil,
			func(n *UsersToCompany, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	if query := utcq.withCompany; query != nil {
		if err := utcq.loadCompany(ctx, query, nodes, nil,
			func(n *UsersToCompany, e *Company) { n.Edges.Company = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (utcq *UsersToCompanyQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*UsersToCompany, init func(*UsersToCompany), assign func(*UsersToCompany, *User)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*UsersToCompany)
	for i := range nodes {
		fk := nodes[i].UserID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "userID" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (utcq *UsersToCompanyQuery) loadCompany(ctx context.Context, query *CompanyQuery, nodes []*UsersToCompany, init func(*UsersToCompany), assign func(*UsersToCompany, *Company)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*UsersToCompany)
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

func (utcq *UsersToCompanyQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := utcq.querySpec()
	if len(utcq.modifiers) > 0 {
		_spec.Modifiers = utcq.modifiers
	}
	_spec.Node.Columns = utcq.ctx.Fields
	if len(utcq.ctx.Fields) > 0 {
		_spec.Unique = utcq.ctx.Unique != nil && *utcq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, utcq.driver, _spec)
}

func (utcq *UsersToCompanyQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(userstocompany.Table, userstocompany.Columns, sqlgraph.NewFieldSpec(userstocompany.FieldID, field.TypeUUID))
	_spec.From = utcq.sql
	if unique := utcq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if utcq.path != nil {
		_spec.Unique = true
	}
	if fields := utcq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, userstocompany.FieldID)
		for i := range fields {
			if fields[i] != userstocompany.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := utcq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := utcq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := utcq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := utcq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (utcq *UsersToCompanyQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(utcq.driver.Dialect())
	t1 := builder.Table(userstocompany.Table)
	columns := utcq.ctx.Fields
	if len(columns) == 0 {
		columns = userstocompany.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if utcq.sql != nil {
		selector = utcq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if utcq.ctx.Unique != nil && *utcq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range utcq.modifiers {
		m(selector)
	}
	for _, p := range utcq.predicates {
		p(selector)
	}
	for _, p := range utcq.order {
		p(selector)
	}
	if offset := utcq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := utcq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (utcq *UsersToCompanyQuery) ForUpdate(opts ...sql.LockOption) *UsersToCompanyQuery {
	if utcq.driver.Dialect() == dialect.Postgres {
		utcq.Unique(false)
	}
	utcq.modifiers = append(utcq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return utcq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (utcq *UsersToCompanyQuery) ForShare(opts ...sql.LockOption) *UsersToCompanyQuery {
	if utcq.driver.Dialect() == dialect.Postgres {
		utcq.Unique(false)
	}
	utcq.modifiers = append(utcq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return utcq
}

// UsersToCompanyGroupBy is the group-by builder for UsersToCompany entities.
type UsersToCompanyGroupBy struct {
	selector
	build *UsersToCompanyQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (utcgb *UsersToCompanyGroupBy) Aggregate(fns ...AggregateFunc) *UsersToCompanyGroupBy {
	utcgb.fns = append(utcgb.fns, fns...)
	return utcgb
}

// Scan applies the selector query and scans the result into the given value.
func (utcgb *UsersToCompanyGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, utcgb.build.ctx, "GroupBy")
	if err := utcgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UsersToCompanyQuery, *UsersToCompanyGroupBy](ctx, utcgb.build, utcgb, utcgb.build.inters, v)
}

func (utcgb *UsersToCompanyGroupBy) sqlScan(ctx context.Context, root *UsersToCompanyQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(utcgb.fns))
	for _, fn := range utcgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*utcgb.flds)+len(utcgb.fns))
		for _, f := range *utcgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*utcgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := utcgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// UsersToCompanySelect is the builder for selecting fields of UsersToCompany entities.
type UsersToCompanySelect struct {
	*UsersToCompanyQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (utcs *UsersToCompanySelect) Aggregate(fns ...AggregateFunc) *UsersToCompanySelect {
	utcs.fns = append(utcs.fns, fns...)
	return utcs
}

// Scan applies the selector query and scans the result into the given value.
func (utcs *UsersToCompanySelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, utcs.ctx, "Select")
	if err := utcs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UsersToCompanyQuery, *UsersToCompanySelect](ctx, utcs.UsersToCompanyQuery, utcs, utcs.inters, v)
}

func (utcs *UsersToCompanySelect) sqlScan(ctx context.Context, root *UsersToCompanyQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(utcs.fns))
	for _, fn := range utcs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*utcs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := utcs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
