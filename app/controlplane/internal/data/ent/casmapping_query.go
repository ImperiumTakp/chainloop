// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/chainloop-dev/chainloop/app/controlplane/internal/data/ent/casbackend"
	"github.com/chainloop-dev/chainloop/app/controlplane/internal/data/ent/casmapping"
	"github.com/chainloop-dev/chainloop/app/controlplane/internal/data/ent/organization"
	"github.com/chainloop-dev/chainloop/app/controlplane/internal/data/ent/predicate"
	"github.com/chainloop-dev/chainloop/app/controlplane/internal/data/ent/workflowrun"
	"github.com/google/uuid"
)

// CASMappingQuery is the builder for querying CASMapping entities.
type CASMappingQuery struct {
	config
	ctx              *QueryContext
	order            []casmapping.OrderOption
	inters           []Interceptor
	predicates       []predicate.CASMapping
	withCasBackend   *CASBackendQuery
	withWorkflowRun  *WorkflowRunQuery
	withOrganization *OrganizationQuery
	withFKs          bool
	modifiers        []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CASMappingQuery builder.
func (cmq *CASMappingQuery) Where(ps ...predicate.CASMapping) *CASMappingQuery {
	cmq.predicates = append(cmq.predicates, ps...)
	return cmq
}

// Limit the number of records to be returned by this query.
func (cmq *CASMappingQuery) Limit(limit int) *CASMappingQuery {
	cmq.ctx.Limit = &limit
	return cmq
}

// Offset to start from.
func (cmq *CASMappingQuery) Offset(offset int) *CASMappingQuery {
	cmq.ctx.Offset = &offset
	return cmq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (cmq *CASMappingQuery) Unique(unique bool) *CASMappingQuery {
	cmq.ctx.Unique = &unique
	return cmq
}

// Order specifies how the records should be ordered.
func (cmq *CASMappingQuery) Order(o ...casmapping.OrderOption) *CASMappingQuery {
	cmq.order = append(cmq.order, o...)
	return cmq
}

// QueryCasBackend chains the current query on the "cas_backend" edge.
func (cmq *CASMappingQuery) QueryCasBackend() *CASBackendQuery {
	query := (&CASBackendClient{config: cmq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cmq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cmq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(casmapping.Table, casmapping.FieldID, selector),
			sqlgraph.To(casbackend.Table, casbackend.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, casmapping.CasBackendTable, casmapping.CasBackendColumn),
		)
		fromU = sqlgraph.SetNeighbors(cmq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryWorkflowRun chains the current query on the "workflow_run" edge.
func (cmq *CASMappingQuery) QueryWorkflowRun() *WorkflowRunQuery {
	query := (&WorkflowRunClient{config: cmq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cmq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cmq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(casmapping.Table, casmapping.FieldID, selector),
			sqlgraph.To(workflowrun.Table, workflowrun.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, casmapping.WorkflowRunTable, casmapping.WorkflowRunColumn),
		)
		fromU = sqlgraph.SetNeighbors(cmq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryOrganization chains the current query on the "organization" edge.
func (cmq *CASMappingQuery) QueryOrganization() *OrganizationQuery {
	query := (&OrganizationClient{config: cmq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cmq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cmq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(casmapping.Table, casmapping.FieldID, selector),
			sqlgraph.To(organization.Table, organization.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, casmapping.OrganizationTable, casmapping.OrganizationColumn),
		)
		fromU = sqlgraph.SetNeighbors(cmq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first CASMapping entity from the query.
// Returns a *NotFoundError when no CASMapping was found.
func (cmq *CASMappingQuery) First(ctx context.Context) (*CASMapping, error) {
	nodes, err := cmq.Limit(1).All(setContextOp(ctx, cmq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{casmapping.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cmq *CASMappingQuery) FirstX(ctx context.Context) *CASMapping {
	node, err := cmq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first CASMapping ID from the query.
// Returns a *NotFoundError when no CASMapping ID was found.
func (cmq *CASMappingQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = cmq.Limit(1).IDs(setContextOp(ctx, cmq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{casmapping.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (cmq *CASMappingQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := cmq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single CASMapping entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one CASMapping entity is found.
// Returns a *NotFoundError when no CASMapping entities are found.
func (cmq *CASMappingQuery) Only(ctx context.Context) (*CASMapping, error) {
	nodes, err := cmq.Limit(2).All(setContextOp(ctx, cmq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{casmapping.Label}
	default:
		return nil, &NotSingularError{casmapping.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cmq *CASMappingQuery) OnlyX(ctx context.Context) *CASMapping {
	node, err := cmq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only CASMapping ID in the query.
// Returns a *NotSingularError when more than one CASMapping ID is found.
// Returns a *NotFoundError when no entities are found.
func (cmq *CASMappingQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = cmq.Limit(2).IDs(setContextOp(ctx, cmq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{casmapping.Label}
	default:
		err = &NotSingularError{casmapping.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (cmq *CASMappingQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := cmq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of CASMappings.
func (cmq *CASMappingQuery) All(ctx context.Context) ([]*CASMapping, error) {
	ctx = setContextOp(ctx, cmq.ctx, "All")
	if err := cmq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*CASMapping, *CASMappingQuery]()
	return withInterceptors[[]*CASMapping](ctx, cmq, qr, cmq.inters)
}

// AllX is like All, but panics if an error occurs.
func (cmq *CASMappingQuery) AllX(ctx context.Context) []*CASMapping {
	nodes, err := cmq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of CASMapping IDs.
func (cmq *CASMappingQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if cmq.ctx.Unique == nil && cmq.path != nil {
		cmq.Unique(true)
	}
	ctx = setContextOp(ctx, cmq.ctx, "IDs")
	if err = cmq.Select(casmapping.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (cmq *CASMappingQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := cmq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (cmq *CASMappingQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, cmq.ctx, "Count")
	if err := cmq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, cmq, querierCount[*CASMappingQuery](), cmq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (cmq *CASMappingQuery) CountX(ctx context.Context) int {
	count, err := cmq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cmq *CASMappingQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, cmq.ctx, "Exist")
	switch _, err := cmq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (cmq *CASMappingQuery) ExistX(ctx context.Context) bool {
	exist, err := cmq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CASMappingQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cmq *CASMappingQuery) Clone() *CASMappingQuery {
	if cmq == nil {
		return nil
	}
	return &CASMappingQuery{
		config:           cmq.config,
		ctx:              cmq.ctx.Clone(),
		order:            append([]casmapping.OrderOption{}, cmq.order...),
		inters:           append([]Interceptor{}, cmq.inters...),
		predicates:       append([]predicate.CASMapping{}, cmq.predicates...),
		withCasBackend:   cmq.withCasBackend.Clone(),
		withWorkflowRun:  cmq.withWorkflowRun.Clone(),
		withOrganization: cmq.withOrganization.Clone(),
		// clone intermediate query.
		sql:  cmq.sql.Clone(),
		path: cmq.path,
	}
}

// WithCasBackend tells the query-builder to eager-load the nodes that are connected to
// the "cas_backend" edge. The optional arguments are used to configure the query builder of the edge.
func (cmq *CASMappingQuery) WithCasBackend(opts ...func(*CASBackendQuery)) *CASMappingQuery {
	query := (&CASBackendClient{config: cmq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cmq.withCasBackend = query
	return cmq
}

// WithWorkflowRun tells the query-builder to eager-load the nodes that are connected to
// the "workflow_run" edge. The optional arguments are used to configure the query builder of the edge.
func (cmq *CASMappingQuery) WithWorkflowRun(opts ...func(*WorkflowRunQuery)) *CASMappingQuery {
	query := (&WorkflowRunClient{config: cmq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cmq.withWorkflowRun = query
	return cmq
}

// WithOrganization tells the query-builder to eager-load the nodes that are connected to
// the "organization" edge. The optional arguments are used to configure the query builder of the edge.
func (cmq *CASMappingQuery) WithOrganization(opts ...func(*OrganizationQuery)) *CASMappingQuery {
	query := (&OrganizationClient{config: cmq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cmq.withOrganization = query
	return cmq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Digest string `json:"digest,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.CASMapping.Query().
//		GroupBy(casmapping.FieldDigest).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (cmq *CASMappingQuery) GroupBy(field string, fields ...string) *CASMappingGroupBy {
	cmq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &CASMappingGroupBy{build: cmq}
	grbuild.flds = &cmq.ctx.Fields
	grbuild.label = casmapping.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Digest string `json:"digest,omitempty"`
//	}
//
//	client.CASMapping.Query().
//		Select(casmapping.FieldDigest).
//		Scan(ctx, &v)
func (cmq *CASMappingQuery) Select(fields ...string) *CASMappingSelect {
	cmq.ctx.Fields = append(cmq.ctx.Fields, fields...)
	sbuild := &CASMappingSelect{CASMappingQuery: cmq}
	sbuild.label = casmapping.Label
	sbuild.flds, sbuild.scan = &cmq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a CASMappingSelect configured with the given aggregations.
func (cmq *CASMappingQuery) Aggregate(fns ...AggregateFunc) *CASMappingSelect {
	return cmq.Select().Aggregate(fns...)
}

func (cmq *CASMappingQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range cmq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, cmq); err != nil {
				return err
			}
		}
	}
	for _, f := range cmq.ctx.Fields {
		if !casmapping.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if cmq.path != nil {
		prev, err := cmq.path(ctx)
		if err != nil {
			return err
		}
		cmq.sql = prev
	}
	return nil
}

func (cmq *CASMappingQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*CASMapping, error) {
	var (
		nodes       = []*CASMapping{}
		withFKs     = cmq.withFKs
		_spec       = cmq.querySpec()
		loadedTypes = [3]bool{
			cmq.withCasBackend != nil,
			cmq.withWorkflowRun != nil,
			cmq.withOrganization != nil,
		}
	)
	if cmq.withCasBackend != nil || cmq.withWorkflowRun != nil || cmq.withOrganization != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, casmapping.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*CASMapping).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &CASMapping{config: cmq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(cmq.modifiers) > 0 {
		_spec.Modifiers = cmq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, cmq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := cmq.withCasBackend; query != nil {
		if err := cmq.loadCasBackend(ctx, query, nodes, nil,
			func(n *CASMapping, e *CASBackend) { n.Edges.CasBackend = e }); err != nil {
			return nil, err
		}
	}
	if query := cmq.withWorkflowRun; query != nil {
		if err := cmq.loadWorkflowRun(ctx, query, nodes, nil,
			func(n *CASMapping, e *WorkflowRun) { n.Edges.WorkflowRun = e }); err != nil {
			return nil, err
		}
	}
	if query := cmq.withOrganization; query != nil {
		if err := cmq.loadOrganization(ctx, query, nodes, nil,
			func(n *CASMapping, e *Organization) { n.Edges.Organization = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (cmq *CASMappingQuery) loadCasBackend(ctx context.Context, query *CASBackendQuery, nodes []*CASMapping, init func(*CASMapping), assign func(*CASMapping, *CASBackend)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*CASMapping)
	for i := range nodes {
		if nodes[i].cas_mapping_cas_backend == nil {
			continue
		}
		fk := *nodes[i].cas_mapping_cas_backend
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(casbackend.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "cas_mapping_cas_backend" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (cmq *CASMappingQuery) loadWorkflowRun(ctx context.Context, query *WorkflowRunQuery, nodes []*CASMapping, init func(*CASMapping), assign func(*CASMapping, *WorkflowRun)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*CASMapping)
	for i := range nodes {
		if nodes[i].cas_mapping_workflow_run == nil {
			continue
		}
		fk := *nodes[i].cas_mapping_workflow_run
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(workflowrun.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "cas_mapping_workflow_run" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (cmq *CASMappingQuery) loadOrganization(ctx context.Context, query *OrganizationQuery, nodes []*CASMapping, init func(*CASMapping), assign func(*CASMapping, *Organization)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*CASMapping)
	for i := range nodes {
		if nodes[i].cas_mapping_organization == nil {
			continue
		}
		fk := *nodes[i].cas_mapping_organization
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(organization.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "cas_mapping_organization" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (cmq *CASMappingQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := cmq.querySpec()
	if len(cmq.modifiers) > 0 {
		_spec.Modifiers = cmq.modifiers
	}
	_spec.Node.Columns = cmq.ctx.Fields
	if len(cmq.ctx.Fields) > 0 {
		_spec.Unique = cmq.ctx.Unique != nil && *cmq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, cmq.driver, _spec)
}

func (cmq *CASMappingQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(casmapping.Table, casmapping.Columns, sqlgraph.NewFieldSpec(casmapping.FieldID, field.TypeUUID))
	_spec.From = cmq.sql
	if unique := cmq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if cmq.path != nil {
		_spec.Unique = true
	}
	if fields := cmq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, casmapping.FieldID)
		for i := range fields {
			if fields[i] != casmapping.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := cmq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := cmq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := cmq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := cmq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (cmq *CASMappingQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(cmq.driver.Dialect())
	t1 := builder.Table(casmapping.Table)
	columns := cmq.ctx.Fields
	if len(columns) == 0 {
		columns = casmapping.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if cmq.sql != nil {
		selector = cmq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if cmq.ctx.Unique != nil && *cmq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range cmq.modifiers {
		m(selector)
	}
	for _, p := range cmq.predicates {
		p(selector)
	}
	for _, p := range cmq.order {
		p(selector)
	}
	if offset := cmq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := cmq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (cmq *CASMappingQuery) Modify(modifiers ...func(s *sql.Selector)) *CASMappingSelect {
	cmq.modifiers = append(cmq.modifiers, modifiers...)
	return cmq.Select()
}

// CASMappingGroupBy is the group-by builder for CASMapping entities.
type CASMappingGroupBy struct {
	selector
	build *CASMappingQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cmgb *CASMappingGroupBy) Aggregate(fns ...AggregateFunc) *CASMappingGroupBy {
	cmgb.fns = append(cmgb.fns, fns...)
	return cmgb
}

// Scan applies the selector query and scans the result into the given value.
func (cmgb *CASMappingGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cmgb.build.ctx, "GroupBy")
	if err := cmgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CASMappingQuery, *CASMappingGroupBy](ctx, cmgb.build, cmgb, cmgb.build.inters, v)
}

func (cmgb *CASMappingGroupBy) sqlScan(ctx context.Context, root *CASMappingQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(cmgb.fns))
	for _, fn := range cmgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*cmgb.flds)+len(cmgb.fns))
		for _, f := range *cmgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*cmgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cmgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// CASMappingSelect is the builder for selecting fields of CASMapping entities.
type CASMappingSelect struct {
	*CASMappingQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (cms *CASMappingSelect) Aggregate(fns ...AggregateFunc) *CASMappingSelect {
	cms.fns = append(cms.fns, fns...)
	return cms
}

// Scan applies the selector query and scans the result into the given value.
func (cms *CASMappingSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cms.ctx, "Select")
	if err := cms.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CASMappingQuery, *CASMappingSelect](ctx, cms.CASMappingQuery, cms, cms.inters, v)
}

func (cms *CASMappingSelect) sqlScan(ctx context.Context, root *CASMappingQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(cms.fns))
	for _, fn := range cms.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*cms.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cms.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (cms *CASMappingSelect) Modify(modifiers ...func(s *sql.Selector)) *CASMappingSelect {
	cms.modifiers = append(cms.modifiers, modifiers...)
	return cms
}
