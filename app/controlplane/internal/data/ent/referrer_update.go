// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/chainloop-dev/chainloop/app/controlplane/internal/data/ent/predicate"
	"github.com/chainloop-dev/chainloop/app/controlplane/internal/data/ent/referrer"
	"github.com/chainloop-dev/chainloop/app/controlplane/internal/data/ent/workflow"
	"github.com/google/uuid"
)

// ReferrerUpdate is the builder for updating Referrer entities.
type ReferrerUpdate struct {
	config
	hooks     []Hook
	mutation  *ReferrerMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the ReferrerUpdate builder.
func (ru *ReferrerUpdate) Where(ps ...predicate.Referrer) *ReferrerUpdate {
	ru.mutation.Where(ps...)
	return ru
}

// AddReferenceIDs adds the "references" edge to the Referrer entity by IDs.
func (ru *ReferrerUpdate) AddReferenceIDs(ids ...uuid.UUID) *ReferrerUpdate {
	ru.mutation.AddReferenceIDs(ids...)
	return ru
}

// AddReferences adds the "references" edges to the Referrer entity.
func (ru *ReferrerUpdate) AddReferences(r ...*Referrer) *ReferrerUpdate {
	ids := make([]uuid.UUID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ru.AddReferenceIDs(ids...)
}

// AddWorkflowIDs adds the "workflows" edge to the Workflow entity by IDs.
func (ru *ReferrerUpdate) AddWorkflowIDs(ids ...uuid.UUID) *ReferrerUpdate {
	ru.mutation.AddWorkflowIDs(ids...)
	return ru
}

// AddWorkflows adds the "workflows" edges to the Workflow entity.
func (ru *ReferrerUpdate) AddWorkflows(w ...*Workflow) *ReferrerUpdate {
	ids := make([]uuid.UUID, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return ru.AddWorkflowIDs(ids...)
}

// Mutation returns the ReferrerMutation object of the builder.
func (ru *ReferrerUpdate) Mutation() *ReferrerMutation {
	return ru.mutation
}

// ClearReferences clears all "references" edges to the Referrer entity.
func (ru *ReferrerUpdate) ClearReferences() *ReferrerUpdate {
	ru.mutation.ClearReferences()
	return ru
}

// RemoveReferenceIDs removes the "references" edge to Referrer entities by IDs.
func (ru *ReferrerUpdate) RemoveReferenceIDs(ids ...uuid.UUID) *ReferrerUpdate {
	ru.mutation.RemoveReferenceIDs(ids...)
	return ru
}

// RemoveReferences removes "references" edges to Referrer entities.
func (ru *ReferrerUpdate) RemoveReferences(r ...*Referrer) *ReferrerUpdate {
	ids := make([]uuid.UUID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ru.RemoveReferenceIDs(ids...)
}

// ClearWorkflows clears all "workflows" edges to the Workflow entity.
func (ru *ReferrerUpdate) ClearWorkflows() *ReferrerUpdate {
	ru.mutation.ClearWorkflows()
	return ru
}

// RemoveWorkflowIDs removes the "workflows" edge to Workflow entities by IDs.
func (ru *ReferrerUpdate) RemoveWorkflowIDs(ids ...uuid.UUID) *ReferrerUpdate {
	ru.mutation.RemoveWorkflowIDs(ids...)
	return ru
}

// RemoveWorkflows removes "workflows" edges to Workflow entities.
func (ru *ReferrerUpdate) RemoveWorkflows(w ...*Workflow) *ReferrerUpdate {
	ids := make([]uuid.UUID, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return ru.RemoveWorkflowIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ru *ReferrerUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ru.sqlSave, ru.mutation, ru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ru *ReferrerUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *ReferrerUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *ReferrerUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (ru *ReferrerUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ReferrerUpdate {
	ru.modifiers = append(ru.modifiers, modifiers...)
	return ru
}

func (ru *ReferrerUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(referrer.Table, referrer.Columns, sqlgraph.NewFieldSpec(referrer.FieldID, field.TypeUUID))
	if ps := ru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if ru.mutation.MetadataCleared() {
		_spec.ClearField(referrer.FieldMetadata, field.TypeJSON)
	}
	if ru.mutation.AnnotationsCleared() {
		_spec.ClearField(referrer.FieldAnnotations, field.TypeJSON)
	}
	if ru.mutation.ReferencesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   referrer.ReferencesTable,
			Columns: referrer.ReferencesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(referrer.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RemovedReferencesIDs(); len(nodes) > 0 && !ru.mutation.ReferencesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   referrer.ReferencesTable,
			Columns: referrer.ReferencesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(referrer.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.ReferencesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   referrer.ReferencesTable,
			Columns: referrer.ReferencesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(referrer.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ru.mutation.WorkflowsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   referrer.WorkflowsTable,
			Columns: referrer.WorkflowsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workflow.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RemovedWorkflowsIDs(); len(nodes) > 0 && !ru.mutation.WorkflowsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   referrer.WorkflowsTable,
			Columns: referrer.WorkflowsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workflow.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.WorkflowsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   referrer.WorkflowsTable,
			Columns: referrer.WorkflowsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workflow.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(ru.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{referrer.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ru.mutation.done = true
	return n, nil
}

// ReferrerUpdateOne is the builder for updating a single Referrer entity.
type ReferrerUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *ReferrerMutation
	modifiers []func(*sql.UpdateBuilder)
}

// AddReferenceIDs adds the "references" edge to the Referrer entity by IDs.
func (ruo *ReferrerUpdateOne) AddReferenceIDs(ids ...uuid.UUID) *ReferrerUpdateOne {
	ruo.mutation.AddReferenceIDs(ids...)
	return ruo
}

// AddReferences adds the "references" edges to the Referrer entity.
func (ruo *ReferrerUpdateOne) AddReferences(r ...*Referrer) *ReferrerUpdateOne {
	ids := make([]uuid.UUID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ruo.AddReferenceIDs(ids...)
}

// AddWorkflowIDs adds the "workflows" edge to the Workflow entity by IDs.
func (ruo *ReferrerUpdateOne) AddWorkflowIDs(ids ...uuid.UUID) *ReferrerUpdateOne {
	ruo.mutation.AddWorkflowIDs(ids...)
	return ruo
}

// AddWorkflows adds the "workflows" edges to the Workflow entity.
func (ruo *ReferrerUpdateOne) AddWorkflows(w ...*Workflow) *ReferrerUpdateOne {
	ids := make([]uuid.UUID, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return ruo.AddWorkflowIDs(ids...)
}

// Mutation returns the ReferrerMutation object of the builder.
func (ruo *ReferrerUpdateOne) Mutation() *ReferrerMutation {
	return ruo.mutation
}

// ClearReferences clears all "references" edges to the Referrer entity.
func (ruo *ReferrerUpdateOne) ClearReferences() *ReferrerUpdateOne {
	ruo.mutation.ClearReferences()
	return ruo
}

// RemoveReferenceIDs removes the "references" edge to Referrer entities by IDs.
func (ruo *ReferrerUpdateOne) RemoveReferenceIDs(ids ...uuid.UUID) *ReferrerUpdateOne {
	ruo.mutation.RemoveReferenceIDs(ids...)
	return ruo
}

// RemoveReferences removes "references" edges to Referrer entities.
func (ruo *ReferrerUpdateOne) RemoveReferences(r ...*Referrer) *ReferrerUpdateOne {
	ids := make([]uuid.UUID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ruo.RemoveReferenceIDs(ids...)
}

// ClearWorkflows clears all "workflows" edges to the Workflow entity.
func (ruo *ReferrerUpdateOne) ClearWorkflows() *ReferrerUpdateOne {
	ruo.mutation.ClearWorkflows()
	return ruo
}

// RemoveWorkflowIDs removes the "workflows" edge to Workflow entities by IDs.
func (ruo *ReferrerUpdateOne) RemoveWorkflowIDs(ids ...uuid.UUID) *ReferrerUpdateOne {
	ruo.mutation.RemoveWorkflowIDs(ids...)
	return ruo
}

// RemoveWorkflows removes "workflows" edges to Workflow entities.
func (ruo *ReferrerUpdateOne) RemoveWorkflows(w ...*Workflow) *ReferrerUpdateOne {
	ids := make([]uuid.UUID, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return ruo.RemoveWorkflowIDs(ids...)
}

// Where appends a list predicates to the ReferrerUpdate builder.
func (ruo *ReferrerUpdateOne) Where(ps ...predicate.Referrer) *ReferrerUpdateOne {
	ruo.mutation.Where(ps...)
	return ruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ruo *ReferrerUpdateOne) Select(field string, fields ...string) *ReferrerUpdateOne {
	ruo.fields = append([]string{field}, fields...)
	return ruo
}

// Save executes the query and returns the updated Referrer entity.
func (ruo *ReferrerUpdateOne) Save(ctx context.Context) (*Referrer, error) {
	return withHooks(ctx, ruo.sqlSave, ruo.mutation, ruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *ReferrerUpdateOne) SaveX(ctx context.Context) *Referrer {
	node, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruo *ReferrerUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *ReferrerUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (ruo *ReferrerUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ReferrerUpdateOne {
	ruo.modifiers = append(ruo.modifiers, modifiers...)
	return ruo
}

func (ruo *ReferrerUpdateOne) sqlSave(ctx context.Context) (_node *Referrer, err error) {
	_spec := sqlgraph.NewUpdateSpec(referrer.Table, referrer.Columns, sqlgraph.NewFieldSpec(referrer.FieldID, field.TypeUUID))
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Referrer.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, referrer.FieldID)
		for _, f := range fields {
			if !referrer.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != referrer.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if ruo.mutation.MetadataCleared() {
		_spec.ClearField(referrer.FieldMetadata, field.TypeJSON)
	}
	if ruo.mutation.AnnotationsCleared() {
		_spec.ClearField(referrer.FieldAnnotations, field.TypeJSON)
	}
	if ruo.mutation.ReferencesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   referrer.ReferencesTable,
			Columns: referrer.ReferencesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(referrer.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RemovedReferencesIDs(); len(nodes) > 0 && !ruo.mutation.ReferencesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   referrer.ReferencesTable,
			Columns: referrer.ReferencesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(referrer.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.ReferencesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   referrer.ReferencesTable,
			Columns: referrer.ReferencesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(referrer.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ruo.mutation.WorkflowsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   referrer.WorkflowsTable,
			Columns: referrer.WorkflowsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workflow.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RemovedWorkflowsIDs(); len(nodes) > 0 && !ruo.mutation.WorkflowsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   referrer.WorkflowsTable,
			Columns: referrer.WorkflowsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workflow.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.WorkflowsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   referrer.WorkflowsTable,
			Columns: referrer.WorkflowsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workflow.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(ruo.modifiers...)
	_node = &Referrer{config: ruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{referrer.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ruo.mutation.done = true
	return _node, nil
}
