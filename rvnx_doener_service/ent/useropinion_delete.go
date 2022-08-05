// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"rvnx_doener_service/ent/predicate"
	"rvnx_doener_service/ent/useropinion"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserOpinionDelete is the builder for deleting a UserOpinion entity.
type UserOpinionDelete struct {
	config
	hooks    []Hook
	mutation *UserOpinionMutation
}

// Where appends a list predicates to the UserOpinionDelete builder.
func (uod *UserOpinionDelete) Where(ps ...predicate.UserOpinion) *UserOpinionDelete {
	uod.mutation.Where(ps...)
	return uod
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (uod *UserOpinionDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(uod.hooks) == 0 {
		affected, err = uod.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserOpinionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uod.mutation = mutation
			affected, err = uod.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(uod.hooks) - 1; i >= 0; i-- {
			if uod.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uod.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uod.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (uod *UserOpinionDelete) ExecX(ctx context.Context) int {
	n, err := uod.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (uod *UserOpinionDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: useropinion.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: useropinion.FieldID,
			},
		},
	}
	if ps := uod.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, uod.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// UserOpinionDeleteOne is the builder for deleting a single UserOpinion entity.
type UserOpinionDeleteOne struct {
	uod *UserOpinionDelete
}

// Exec executes the deletion query.
func (uodo *UserOpinionDeleteOne) Exec(ctx context.Context) error {
	n, err := uodo.uod.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{useropinion.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (uodo *UserOpinionDeleteOne) ExecX(ctx context.Context) {
	uodo.uod.ExecX(ctx)
}
