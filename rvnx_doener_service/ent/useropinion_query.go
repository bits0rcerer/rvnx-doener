// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"
	"rvnx_doener_service/ent/kebabshop"
	"rvnx_doener_service/ent/predicate"
	"rvnx_doener_service/ent/twitchuser"
	"rvnx_doener_service/ent/useropinion"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserOpinionQuery is the builder for querying UserOpinion entities.
type UserOpinionQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.UserOpinion
	// eager-loading edges.
	withShop   *KebabShopQuery
	withAuthor *TwitchUserQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the UserOpinionQuery builder.
func (uoq *UserOpinionQuery) Where(ps ...predicate.UserOpinion) *UserOpinionQuery {
	uoq.predicates = append(uoq.predicates, ps...)
	return uoq
}

// Limit adds a limit step to the query.
func (uoq *UserOpinionQuery) Limit(limit int) *UserOpinionQuery {
	uoq.limit = &limit
	return uoq
}

// Offset adds an offset step to the query.
func (uoq *UserOpinionQuery) Offset(offset int) *UserOpinionQuery {
	uoq.offset = &offset
	return uoq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (uoq *UserOpinionQuery) Unique(unique bool) *UserOpinionQuery {
	uoq.unique = &unique
	return uoq
}

// Order adds an order step to the query.
func (uoq *UserOpinionQuery) Order(o ...OrderFunc) *UserOpinionQuery {
	uoq.order = append(uoq.order, o...)
	return uoq
}

// QueryShop chains the current query on the "shop" edge.
func (uoq *UserOpinionQuery) QueryShop() *KebabShopQuery {
	query := &KebabShopQuery{config: uoq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := uoq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := uoq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(useropinion.Table, useropinion.FieldID, selector),
			sqlgraph.To(kebabshop.Table, kebabshop.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, useropinion.ShopTable, useropinion.ShopColumn),
		)
		fromU = sqlgraph.SetNeighbors(uoq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryAuthor chains the current query on the "author" edge.
func (uoq *UserOpinionQuery) QueryAuthor() *TwitchUserQuery {
	query := &TwitchUserQuery{config: uoq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := uoq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := uoq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(useropinion.Table, useropinion.FieldID, selector),
			sqlgraph.To(twitchuser.Table, twitchuser.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, useropinion.AuthorTable, useropinion.AuthorColumn),
		)
		fromU = sqlgraph.SetNeighbors(uoq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first UserOpinion entity from the query.
// Returns a *NotFoundError when no UserOpinion was found.
func (uoq *UserOpinionQuery) First(ctx context.Context) (*UserOpinion, error) {
	nodes, err := uoq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{useropinion.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (uoq *UserOpinionQuery) FirstX(ctx context.Context) *UserOpinion {
	node, err := uoq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first UserOpinion ID from the query.
// Returns a *NotFoundError when no UserOpinion ID was found.
func (uoq *UserOpinionQuery) FirstID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = uoq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{useropinion.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (uoq *UserOpinionQuery) FirstIDX(ctx context.Context) uint64 {
	id, err := uoq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single UserOpinion entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one UserOpinion entity is found.
// Returns a *NotFoundError when no UserOpinion entities are found.
func (uoq *UserOpinionQuery) Only(ctx context.Context) (*UserOpinion, error) {
	nodes, err := uoq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{useropinion.Label}
	default:
		return nil, &NotSingularError{useropinion.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (uoq *UserOpinionQuery) OnlyX(ctx context.Context) *UserOpinion {
	node, err := uoq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only UserOpinion ID in the query.
// Returns a *NotSingularError when more than one UserOpinion ID is found.
// Returns a *NotFoundError when no entities are found.
func (uoq *UserOpinionQuery) OnlyID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = uoq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{useropinion.Label}
	default:
		err = &NotSingularError{useropinion.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (uoq *UserOpinionQuery) OnlyIDX(ctx context.Context) uint64 {
	id, err := uoq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of UserOpinions.
func (uoq *UserOpinionQuery) All(ctx context.Context) ([]*UserOpinion, error) {
	if err := uoq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return uoq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (uoq *UserOpinionQuery) AllX(ctx context.Context) []*UserOpinion {
	nodes, err := uoq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of UserOpinion IDs.
func (uoq *UserOpinionQuery) IDs(ctx context.Context) ([]uint64, error) {
	var ids []uint64
	if err := uoq.Select(useropinion.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (uoq *UserOpinionQuery) IDsX(ctx context.Context) []uint64 {
	ids, err := uoq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (uoq *UserOpinionQuery) Count(ctx context.Context) (int, error) {
	if err := uoq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return uoq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (uoq *UserOpinionQuery) CountX(ctx context.Context) int {
	count, err := uoq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (uoq *UserOpinionQuery) Exist(ctx context.Context) (bool, error) {
	if err := uoq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return uoq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (uoq *UserOpinionQuery) ExistX(ctx context.Context) bool {
	exist, err := uoq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UserOpinionQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (uoq *UserOpinionQuery) Clone() *UserOpinionQuery {
	if uoq == nil {
		return nil
	}
	return &UserOpinionQuery{
		config:     uoq.config,
		limit:      uoq.limit,
		offset:     uoq.offset,
		order:      append([]OrderFunc{}, uoq.order...),
		predicates: append([]predicate.UserOpinion{}, uoq.predicates...),
		withShop:   uoq.withShop.Clone(),
		withAuthor: uoq.withAuthor.Clone(),
		// clone intermediate query.
		sql:    uoq.sql.Clone(),
		path:   uoq.path,
		unique: uoq.unique,
	}
}

// WithShop tells the query-builder to eager-load the nodes that are connected to
// the "shop" edge. The optional arguments are used to configure the query builder of the edge.
func (uoq *UserOpinionQuery) WithShop(opts ...func(*KebabShopQuery)) *UserOpinionQuery {
	query := &KebabShopQuery{config: uoq.config}
	for _, opt := range opts {
		opt(query)
	}
	uoq.withShop = query
	return uoq
}

// WithAuthor tells the query-builder to eager-load the nodes that are connected to
// the "author" edge. The optional arguments are used to configure the query builder of the edge.
func (uoq *UserOpinionQuery) WithAuthor(opts ...func(*TwitchUserQuery)) *UserOpinionQuery {
	query := &TwitchUserQuery{config: uoq.config}
	for _, opt := range opts {
		opt(query)
	}
	uoq.withAuthor = query
	return uoq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Created time.Time `json:"created,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.UserOpinion.Query().
//		GroupBy(useropinion.FieldCreated).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (uoq *UserOpinionQuery) GroupBy(field string, fields ...string) *UserOpinionGroupBy {
	grbuild := &UserOpinionGroupBy{config: uoq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := uoq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return uoq.sqlQuery(ctx), nil
	}
	grbuild.label = useropinion.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Created time.Time `json:"created,omitempty"`
//	}
//
//	client.UserOpinion.Query().
//		Select(useropinion.FieldCreated).
//		Scan(ctx, &v)
func (uoq *UserOpinionQuery) Select(fields ...string) *UserOpinionSelect {
	uoq.fields = append(uoq.fields, fields...)
	selbuild := &UserOpinionSelect{UserOpinionQuery: uoq}
	selbuild.label = useropinion.Label
	selbuild.flds, selbuild.scan = &uoq.fields, selbuild.Scan
	return selbuild
}

func (uoq *UserOpinionQuery) prepareQuery(ctx context.Context) error {
	for _, f := range uoq.fields {
		if !useropinion.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if uoq.path != nil {
		prev, err := uoq.path(ctx)
		if err != nil {
			return err
		}
		uoq.sql = prev
	}
	return nil
}

func (uoq *UserOpinionQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*UserOpinion, error) {
	var (
		nodes       = []*UserOpinion{}
		withFKs     = uoq.withFKs
		_spec       = uoq.querySpec()
		loadedTypes = [2]bool{
			uoq.withShop != nil,
			uoq.withAuthor != nil,
		}
	)
	if uoq.withShop != nil || uoq.withAuthor != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, useropinion.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*UserOpinion).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &UserOpinion{config: uoq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, uoq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := uoq.withShop; query != nil {
		ids := make([]uint64, 0, len(nodes))
		nodeids := make(map[uint64][]*UserOpinion)
		for i := range nodes {
			if nodes[i].kebab_shop_user_opinions == nil {
				continue
			}
			fk := *nodes[i].kebab_shop_user_opinions
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(kebabshop.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "kebab_shop_user_opinions" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Shop = n
			}
		}
	}

	if query := uoq.withAuthor; query != nil {
		ids := make([]int64, 0, len(nodes))
		nodeids := make(map[int64][]*UserOpinion)
		for i := range nodes {
			if nodes[i].twitch_user_user_opinions == nil {
				continue
			}
			fk := *nodes[i].twitch_user_user_opinions
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(twitchuser.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "twitch_user_user_opinions" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Author = n
			}
		}
	}

	return nodes, nil
}

func (uoq *UserOpinionQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := uoq.querySpec()
	_spec.Node.Columns = uoq.fields
	if len(uoq.fields) > 0 {
		_spec.Unique = uoq.unique != nil && *uoq.unique
	}
	return sqlgraph.CountNodes(ctx, uoq.driver, _spec)
}

func (uoq *UserOpinionQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := uoq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (uoq *UserOpinionQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   useropinion.Table,
			Columns: useropinion.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: useropinion.FieldID,
			},
		},
		From:   uoq.sql,
		Unique: true,
	}
	if unique := uoq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := uoq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, useropinion.FieldID)
		for i := range fields {
			if fields[i] != useropinion.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := uoq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := uoq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := uoq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := uoq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (uoq *UserOpinionQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(uoq.driver.Dialect())
	t1 := builder.Table(useropinion.Table)
	columns := uoq.fields
	if len(columns) == 0 {
		columns = useropinion.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if uoq.sql != nil {
		selector = uoq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if uoq.unique != nil && *uoq.unique {
		selector.Distinct()
	}
	for _, p := range uoq.predicates {
		p(selector)
	}
	for _, p := range uoq.order {
		p(selector)
	}
	if offset := uoq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := uoq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// UserOpinionGroupBy is the group-by builder for UserOpinion entities.
type UserOpinionGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (uogb *UserOpinionGroupBy) Aggregate(fns ...AggregateFunc) *UserOpinionGroupBy {
	uogb.fns = append(uogb.fns, fns...)
	return uogb
}

// Scan applies the group-by query and scans the result into the given value.
func (uogb *UserOpinionGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := uogb.path(ctx)
	if err != nil {
		return err
	}
	uogb.sql = query
	return uogb.sqlScan(ctx, v)
}

func (uogb *UserOpinionGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range uogb.fields {
		if !useropinion.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := uogb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := uogb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (uogb *UserOpinionGroupBy) sqlQuery() *sql.Selector {
	selector := uogb.sql.Select()
	aggregation := make([]string, 0, len(uogb.fns))
	for _, fn := range uogb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(uogb.fields)+len(uogb.fns))
		for _, f := range uogb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(uogb.fields...)...)
}

// UserOpinionSelect is the builder for selecting fields of UserOpinion entities.
type UserOpinionSelect struct {
	*UserOpinionQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (uos *UserOpinionSelect) Scan(ctx context.Context, v interface{}) error {
	if err := uos.prepareQuery(ctx); err != nil {
		return err
	}
	uos.sql = uos.UserOpinionQuery.sqlQuery(ctx)
	return uos.sqlScan(ctx, v)
}

func (uos *UserOpinionSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := uos.sql.Query()
	if err := uos.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
