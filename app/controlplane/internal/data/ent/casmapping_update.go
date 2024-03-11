// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/chainloop-dev/chainloop/app/controlplane/internal/data/ent/casmapping"
	"github.com/chainloop-dev/chainloop/app/controlplane/internal/data/ent/predicate"
)

// CASMappingUpdate is the builder for updating CASMapping entities.
type CASMappingUpdate struct {
	config
	hooks     []Hook
	mutation  *CASMappingMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the CASMappingUpdate builder.
func (cmu *CASMappingUpdate) Where(ps ...predicate.CASMapping) *CASMappingUpdate {
	cmu.mutation.Where(ps...)
	return cmu
}

// Mutation returns the CASMappingMutation object of the builder.
func (cmu *CASMappingUpdate) Mutation() *CASMappingMutation {
	return cmu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cmu *CASMappingUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, cmu.sqlSave, cmu.mutation, cmu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cmu *CASMappingUpdate) SaveX(ctx context.Context) int {
	affected, err := cmu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cmu *CASMappingUpdate) Exec(ctx context.Context) error {
	_, err := cmu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cmu *CASMappingUpdate) ExecX(ctx context.Context) {
	if err := cmu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cmu *CASMappingUpdate) check() error {
	if _, ok := cmu.mutation.CasBackendID(); cmu.mutation.CasBackendCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "CASMapping.cas_backend"`)
	}
	if _, ok := cmu.mutation.OrganizationID(); cmu.mutation.OrganizationCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "CASMapping.organization"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (cmu *CASMappingUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *CASMappingUpdate {
	cmu.modifiers = append(cmu.modifiers, modifiers...)
	return cmu
}

func (cmu *CASMappingUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := cmu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(casmapping.Table, casmapping.Columns, sqlgraph.NewFieldSpec(casmapping.FieldID, field.TypeUUID))
	if ps := cmu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	_spec.AddModifiers(cmu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, cmu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{casmapping.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cmu.mutation.done = true
	return n, nil
}

// CASMappingUpdateOne is the builder for updating a single CASMapping entity.
type CASMappingUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *CASMappingMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Mutation returns the CASMappingMutation object of the builder.
func (cmuo *CASMappingUpdateOne) Mutation() *CASMappingMutation {
	return cmuo.mutation
}

// Where appends a list predicates to the CASMappingUpdate builder.
func (cmuo *CASMappingUpdateOne) Where(ps ...predicate.CASMapping) *CASMappingUpdateOne {
	cmuo.mutation.Where(ps...)
	return cmuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cmuo *CASMappingUpdateOne) Select(field string, fields ...string) *CASMappingUpdateOne {
	cmuo.fields = append([]string{field}, fields...)
	return cmuo
}

// Save executes the query and returns the updated CASMapping entity.
func (cmuo *CASMappingUpdateOne) Save(ctx context.Context) (*CASMapping, error) {
	return withHooks(ctx, cmuo.sqlSave, cmuo.mutation, cmuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cmuo *CASMappingUpdateOne) SaveX(ctx context.Context) *CASMapping {
	node, err := cmuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cmuo *CASMappingUpdateOne) Exec(ctx context.Context) error {
	_, err := cmuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cmuo *CASMappingUpdateOne) ExecX(ctx context.Context) {
	if err := cmuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cmuo *CASMappingUpdateOne) check() error {
	if _, ok := cmuo.mutation.CasBackendID(); cmuo.mutation.CasBackendCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "CASMapping.cas_backend"`)
	}
	if _, ok := cmuo.mutation.OrganizationID(); cmuo.mutation.OrganizationCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "CASMapping.organization"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (cmuo *CASMappingUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *CASMappingUpdateOne {
	cmuo.modifiers = append(cmuo.modifiers, modifiers...)
	return cmuo
}

func (cmuo *CASMappingUpdateOne) sqlSave(ctx context.Context) (_node *CASMapping, err error) {
	if err := cmuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(casmapping.Table, casmapping.Columns, sqlgraph.NewFieldSpec(casmapping.FieldID, field.TypeUUID))
	id, ok := cmuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "CASMapping.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cmuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, casmapping.FieldID)
		for _, f := range fields {
			if !casmapping.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != casmapping.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cmuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	_spec.AddModifiers(cmuo.modifiers...)
	_node = &CASMapping{config: cmuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cmuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{casmapping.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cmuo.mutation.done = true
	return _node, nil
}
