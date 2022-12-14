// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"rvnx_doener_service/ent/predicate"
	"rvnx_doener_service/ent/twitchuser"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TwitchUserDelete is the builder for deleting a TwitchUser entity.
type TwitchUserDelete struct {
	config
	hooks    []Hook
	mutation *TwitchUserMutation
}

// Where appends a list predicates to the TwitchUserDelete builder.
func (tud *TwitchUserDelete) Where(ps ...predicate.TwitchUser) *TwitchUserDelete {
	tud.mutation.Where(ps...)
	return tud
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (tud *TwitchUserDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(tud.hooks) == 0 {
		affected, err = tud.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TwitchUserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tud.mutation = mutation
			affected, err = tud.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tud.hooks) - 1; i >= 0; i-- {
			if tud.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tud.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tud.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (tud *TwitchUserDelete) ExecX(ctx context.Context) int {
	n, err := tud.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (tud *TwitchUserDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: twitchuser.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: twitchuser.FieldID,
			},
		},
	}
	if ps := tud.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, tud.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// TwitchUserDeleteOne is the builder for deleting a single TwitchUser entity.
type TwitchUserDeleteOne struct {
	tud *TwitchUserDelete
}

// Exec executes the deletion query.
func (tudo *TwitchUserDeleteOne) Exec(ctx context.Context) error {
	n, err := tudo.tud.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{twitchuser.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (tudo *TwitchUserDeleteOne) ExecX(ctx context.Context) {
	tudo.tud.ExecX(ctx)
}
