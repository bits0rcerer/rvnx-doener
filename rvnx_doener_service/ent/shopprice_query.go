// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"
	"rvnx_doener_service/ent/kebabshop"
	"rvnx_doener_service/ent/predicate"
	"rvnx_doener_service/ent/shopprice"
	"rvnx_doener_service/ent/twitchuser"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ShopPriceQuery is the builder for querying ShopPrice entities.
type ShopPriceQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.ShopPrice
	// eager-loading edges.
	withShop   *KebabShopQuery
	withAuthor *TwitchUserQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ShopPriceQuery builder.
func (spq *ShopPriceQuery) Where(ps ...predicate.ShopPrice) *ShopPriceQuery {
	spq.predicates = append(spq.predicates, ps...)
	return spq
}

// Limit adds a limit step to the query.
func (spq *ShopPriceQuery) Limit(limit int) *ShopPriceQuery {
	spq.limit = &limit
	return spq
}

// Offset adds an offset step to the query.
func (spq *ShopPriceQuery) Offset(offset int) *ShopPriceQuery {
	spq.offset = &offset
	return spq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (spq *ShopPriceQuery) Unique(unique bool) *ShopPriceQuery {
	spq.unique = &unique
	return spq
}

// Order adds an order step to the query.
func (spq *ShopPriceQuery) Order(o ...OrderFunc) *ShopPriceQuery {
	spq.order = append(spq.order, o...)
	return spq
}

// QueryShop chains the current query on the "shop" edge.
func (spq *ShopPriceQuery) QueryShop() *KebabShopQuery {
	query := &KebabShopQuery{config: spq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := spq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := spq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(shopprice.Table, shopprice.FieldID, selector),
			sqlgraph.To(kebabshop.Table, kebabshop.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, shopprice.ShopTable, shopprice.ShopColumn),
		)
		fromU = sqlgraph.SetNeighbors(spq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryAuthor chains the current query on the "author" edge.
func (spq *ShopPriceQuery) QueryAuthor() *TwitchUserQuery {
	query := &TwitchUserQuery{config: spq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := spq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := spq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(shopprice.Table, shopprice.FieldID, selector),
			sqlgraph.To(twitchuser.Table, twitchuser.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, shopprice.AuthorTable, shopprice.AuthorColumn),
		)
		fromU = sqlgraph.SetNeighbors(spq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ShopPrice entity from the query.
// Returns a *NotFoundError when no ShopPrice was found.
func (spq *ShopPriceQuery) First(ctx context.Context) (*ShopPrice, error) {
	nodes, err := spq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{shopprice.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (spq *ShopPriceQuery) FirstX(ctx context.Context) *ShopPrice {
	node, err := spq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ShopPrice ID from the query.
// Returns a *NotFoundError when no ShopPrice ID was found.
func (spq *ShopPriceQuery) FirstID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = spq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{shopprice.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (spq *ShopPriceQuery) FirstIDX(ctx context.Context) uint64 {
	id, err := spq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ShopPrice entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ShopPrice entity is found.
// Returns a *NotFoundError when no ShopPrice entities are found.
func (spq *ShopPriceQuery) Only(ctx context.Context) (*ShopPrice, error) {
	nodes, err := spq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{shopprice.Label}
	default:
		return nil, &NotSingularError{shopprice.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (spq *ShopPriceQuery) OnlyX(ctx context.Context) *ShopPrice {
	node, err := spq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ShopPrice ID in the query.
// Returns a *NotSingularError when more than one ShopPrice ID is found.
// Returns a *NotFoundError when no entities are found.
func (spq *ShopPriceQuery) OnlyID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = spq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{shopprice.Label}
	default:
		err = &NotSingularError{shopprice.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (spq *ShopPriceQuery) OnlyIDX(ctx context.Context) uint64 {
	id, err := spq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ShopPrices.
func (spq *ShopPriceQuery) All(ctx context.Context) ([]*ShopPrice, error) {
	if err := spq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return spq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (spq *ShopPriceQuery) AllX(ctx context.Context) []*ShopPrice {
	nodes, err := spq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ShopPrice IDs.
func (spq *ShopPriceQuery) IDs(ctx context.Context) ([]uint64, error) {
	var ids []uint64
	if err := spq.Select(shopprice.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (spq *ShopPriceQuery) IDsX(ctx context.Context) []uint64 {
	ids, err := spq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (spq *ShopPriceQuery) Count(ctx context.Context) (int, error) {
	if err := spq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return spq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (spq *ShopPriceQuery) CountX(ctx context.Context) int {
	count, err := spq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (spq *ShopPriceQuery) Exist(ctx context.Context) (bool, error) {
	if err := spq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return spq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (spq *ShopPriceQuery) ExistX(ctx context.Context) bool {
	exist, err := spq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ShopPriceQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (spq *ShopPriceQuery) Clone() *ShopPriceQuery {
	if spq == nil {
		return nil
	}
	return &ShopPriceQuery{
		config:     spq.config,
		limit:      spq.limit,
		offset:     spq.offset,
		order:      append([]OrderFunc{}, spq.order...),
		predicates: append([]predicate.ShopPrice{}, spq.predicates...),
		withShop:   spq.withShop.Clone(),
		withAuthor: spq.withAuthor.Clone(),
		// clone intermediate query.
		sql:    spq.sql.Clone(),
		path:   spq.path,
		unique: spq.unique,
	}
}

// WithShop tells the query-builder to eager-load the nodes that are connected to
// the "shop" edge. The optional arguments are used to configure the query builder of the edge.
func (spq *ShopPriceQuery) WithShop(opts ...func(*KebabShopQuery)) *ShopPriceQuery {
	query := &KebabShopQuery{config: spq.config}
	for _, opt := range opts {
		opt(query)
	}
	spq.withShop = query
	return spq
}

// WithAuthor tells the query-builder to eager-load the nodes that are connected to
// the "author" edge. The optional arguments are used to configure the query builder of the edge.
func (spq *ShopPriceQuery) WithAuthor(opts ...func(*TwitchUserQuery)) *ShopPriceQuery {
	query := &TwitchUserQuery{config: spq.config}
	for _, opt := range opts {
		opt(query)
	}
	spq.withAuthor = query
	return spq
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
//	client.ShopPrice.Query().
//		GroupBy(shopprice.FieldCreated).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (spq *ShopPriceQuery) GroupBy(field string, fields ...string) *ShopPriceGroupBy {
	grbuild := &ShopPriceGroupBy{config: spq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := spq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return spq.sqlQuery(ctx), nil
	}
	grbuild.label = shopprice.Label
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
//	client.ShopPrice.Query().
//		Select(shopprice.FieldCreated).
//		Scan(ctx, &v)
func (spq *ShopPriceQuery) Select(fields ...string) *ShopPriceSelect {
	spq.fields = append(spq.fields, fields...)
	selbuild := &ShopPriceSelect{ShopPriceQuery: spq}
	selbuild.label = shopprice.Label
	selbuild.flds, selbuild.scan = &spq.fields, selbuild.Scan
	return selbuild
}

func (spq *ShopPriceQuery) prepareQuery(ctx context.Context) error {
	for _, f := range spq.fields {
		if !shopprice.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if spq.path != nil {
		prev, err := spq.path(ctx)
		if err != nil {
			return err
		}
		spq.sql = prev
	}
	return nil
}

func (spq *ShopPriceQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ShopPrice, error) {
	var (
		nodes       = []*ShopPrice{}
		withFKs     = spq.withFKs
		_spec       = spq.querySpec()
		loadedTypes = [2]bool{
			spq.withShop != nil,
			spq.withAuthor != nil,
		}
	)
	if spq.withShop != nil || spq.withAuthor != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, shopprice.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*ShopPrice).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &ShopPrice{config: spq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, spq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := spq.withShop; query != nil {
		ids := make([]uint64, 0, len(nodes))
		nodeids := make(map[uint64][]*ShopPrice)
		for i := range nodes {
			if nodes[i].kebab_shop_user_prices == nil {
				continue
			}
			fk := *nodes[i].kebab_shop_user_prices
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
				return nil, fmt.Errorf(`unexpected foreign-key "kebab_shop_user_prices" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Shop = n
			}
		}
	}

	if query := spq.withAuthor; query != nil {
		ids := make([]int64, 0, len(nodes))
		nodeids := make(map[int64][]*ShopPrice)
		for i := range nodes {
			if nodes[i].twitch_user_user_prices == nil {
				continue
			}
			fk := *nodes[i].twitch_user_user_prices
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
				return nil, fmt.Errorf(`unexpected foreign-key "twitch_user_user_prices" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Author = n
			}
		}
	}

	return nodes, nil
}

func (spq *ShopPriceQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := spq.querySpec()
	_spec.Node.Columns = spq.fields
	if len(spq.fields) > 0 {
		_spec.Unique = spq.unique != nil && *spq.unique
	}
	return sqlgraph.CountNodes(ctx, spq.driver, _spec)
}

func (spq *ShopPriceQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := spq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (spq *ShopPriceQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   shopprice.Table,
			Columns: shopprice.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: shopprice.FieldID,
			},
		},
		From:   spq.sql,
		Unique: true,
	}
	if unique := spq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := spq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, shopprice.FieldID)
		for i := range fields {
			if fields[i] != shopprice.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := spq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := spq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := spq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := spq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (spq *ShopPriceQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(spq.driver.Dialect())
	t1 := builder.Table(shopprice.Table)
	columns := spq.fields
	if len(columns) == 0 {
		columns = shopprice.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if spq.sql != nil {
		selector = spq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if spq.unique != nil && *spq.unique {
		selector.Distinct()
	}
	for _, p := range spq.predicates {
		p(selector)
	}
	for _, p := range spq.order {
		p(selector)
	}
	if offset := spq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := spq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ShopPriceGroupBy is the group-by builder for ShopPrice entities.
type ShopPriceGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (spgb *ShopPriceGroupBy) Aggregate(fns ...AggregateFunc) *ShopPriceGroupBy {
	spgb.fns = append(spgb.fns, fns...)
	return spgb
}

// Scan applies the group-by query and scans the result into the given value.
func (spgb *ShopPriceGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := spgb.path(ctx)
	if err != nil {
		return err
	}
	spgb.sql = query
	return spgb.sqlScan(ctx, v)
}

func (spgb *ShopPriceGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range spgb.fields {
		if !shopprice.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := spgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := spgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (spgb *ShopPriceGroupBy) sqlQuery() *sql.Selector {
	selector := spgb.sql.Select()
	aggregation := make([]string, 0, len(spgb.fns))
	for _, fn := range spgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(spgb.fields)+len(spgb.fns))
		for _, f := range spgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(spgb.fields...)...)
}

// ShopPriceSelect is the builder for selecting fields of ShopPrice entities.
type ShopPriceSelect struct {
	*ShopPriceQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (sps *ShopPriceSelect) Scan(ctx context.Context, v interface{}) error {
	if err := sps.prepareQuery(ctx); err != nil {
		return err
	}
	sps.sql = sps.ShopPriceQuery.sqlQuery(ctx)
	return sps.sqlScan(ctx, v)
}

func (sps *ShopPriceSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := sps.sql.Query()
	if err := sps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}