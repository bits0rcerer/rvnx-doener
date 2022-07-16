// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"rvnx_doener_service/ent/kebabshop"
	"rvnx_doener_service/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/jackc/pgtype"
)

// KebabShopUpdate is the builder for updating KebabShop entities.
type KebabShopUpdate struct {
	config
	hooks    []Hook
	mutation *KebabShopMutation
}

// Where appends a list predicates to the KebabShopUpdate builder.
func (ksu *KebabShopUpdate) Where(ps ...predicate.KebabShop) *KebabShopUpdate {
	ksu.mutation.Where(ps...)
	return ksu
}

// SetOsmID sets the "osm_id" field.
func (ksu *KebabShopUpdate) SetOsmID(i int) *KebabShopUpdate {
	ksu.mutation.ResetOsmID()
	ksu.mutation.SetOsmID(i)
	return ksu
}

// SetNillableOsmID sets the "osm_id" field if the given value is not nil.
func (ksu *KebabShopUpdate) SetNillableOsmID(i *int) *KebabShopUpdate {
	if i != nil {
		ksu.SetOsmID(*i)
	}
	return ksu
}

// AddOsmID adds i to the "osm_id" field.
func (ksu *KebabShopUpdate) AddOsmID(i int) *KebabShopUpdate {
	ksu.mutation.AddOsmID(i)
	return ksu
}

// ClearOsmID clears the value of the "osm_id" field.
func (ksu *KebabShopUpdate) ClearOsmID() *KebabShopUpdate {
	ksu.mutation.ClearOsmID()
	return ksu
}

// SetName sets the "name" field.
func (ksu *KebabShopUpdate) SetName(s string) *KebabShopUpdate {
	ksu.mutation.SetName(s)
	return ksu
}

// SetPoint sets the "point" field.
func (ksu *KebabShopUpdate) SetPoint(pg *pgtype.Point) *KebabShopUpdate {
	ksu.mutation.SetPoint(pg)
	return ksu
}

// Mutation returns the KebabShopMutation object of the builder.
func (ksu *KebabShopUpdate) Mutation() *KebabShopMutation {
	return ksu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ksu *KebabShopUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ksu.hooks) == 0 {
		affected, err = ksu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*KebabShopMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ksu.mutation = mutation
			affected, err = ksu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ksu.hooks) - 1; i >= 0; i-- {
			if ksu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ksu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ksu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ksu *KebabShopUpdate) SaveX(ctx context.Context) int {
	affected, err := ksu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ksu *KebabShopUpdate) Exec(ctx context.Context) error {
	_, err := ksu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ksu *KebabShopUpdate) ExecX(ctx context.Context) {
	if err := ksu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ksu *KebabShopUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   kebabshop.Table,
			Columns: kebabshop.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: kebabshop.FieldID,
			},
		},
	}
	if ps := ksu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ksu.mutation.OsmID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: kebabshop.FieldOsmID,
		})
	}
	if value, ok := ksu.mutation.AddedOsmID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: kebabshop.FieldOsmID,
		})
	}
	if ksu.mutation.OsmIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: kebabshop.FieldOsmID,
		})
	}
	if value, ok := ksu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: kebabshop.FieldName,
		})
	}
	if value, ok := ksu.mutation.Point(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: kebabshop.FieldPoint,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ksu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{kebabshop.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// KebabShopUpdateOne is the builder for updating a single KebabShop entity.
type KebabShopUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *KebabShopMutation
}

// SetOsmID sets the "osm_id" field.
func (ksuo *KebabShopUpdateOne) SetOsmID(i int) *KebabShopUpdateOne {
	ksuo.mutation.ResetOsmID()
	ksuo.mutation.SetOsmID(i)
	return ksuo
}

// SetNillableOsmID sets the "osm_id" field if the given value is not nil.
func (ksuo *KebabShopUpdateOne) SetNillableOsmID(i *int) *KebabShopUpdateOne {
	if i != nil {
		ksuo.SetOsmID(*i)
	}
	return ksuo
}

// AddOsmID adds i to the "osm_id" field.
func (ksuo *KebabShopUpdateOne) AddOsmID(i int) *KebabShopUpdateOne {
	ksuo.mutation.AddOsmID(i)
	return ksuo
}

// ClearOsmID clears the value of the "osm_id" field.
func (ksuo *KebabShopUpdateOne) ClearOsmID() *KebabShopUpdateOne {
	ksuo.mutation.ClearOsmID()
	return ksuo
}

// SetName sets the "name" field.
func (ksuo *KebabShopUpdateOne) SetName(s string) *KebabShopUpdateOne {
	ksuo.mutation.SetName(s)
	return ksuo
}

// SetPoint sets the "point" field.
func (ksuo *KebabShopUpdateOne) SetPoint(pg *pgtype.Point) *KebabShopUpdateOne {
	ksuo.mutation.SetPoint(pg)
	return ksuo
}

// Mutation returns the KebabShopMutation object of the builder.
func (ksuo *KebabShopUpdateOne) Mutation() *KebabShopMutation {
	return ksuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ksuo *KebabShopUpdateOne) Select(field string, fields ...string) *KebabShopUpdateOne {
	ksuo.fields = append([]string{field}, fields...)
	return ksuo
}

// Save executes the query and returns the updated KebabShop entity.
func (ksuo *KebabShopUpdateOne) Save(ctx context.Context) (*KebabShop, error) {
	var (
		err  error
		node *KebabShop
	)
	if len(ksuo.hooks) == 0 {
		node, err = ksuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*KebabShopMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ksuo.mutation = mutation
			node, err = ksuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ksuo.hooks) - 1; i >= 0; i-- {
			if ksuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ksuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ksuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*KebabShop)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from KebabShopMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ksuo *KebabShopUpdateOne) SaveX(ctx context.Context) *KebabShop {
	node, err := ksuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ksuo *KebabShopUpdateOne) Exec(ctx context.Context) error {
	_, err := ksuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ksuo *KebabShopUpdateOne) ExecX(ctx context.Context) {
	if err := ksuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ksuo *KebabShopUpdateOne) sqlSave(ctx context.Context) (_node *KebabShop, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   kebabshop.Table,
			Columns: kebabshop.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: kebabshop.FieldID,
			},
		},
	}
	id, ok := ksuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "KebabShop.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ksuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, kebabshop.FieldID)
		for _, f := range fields {
			if !kebabshop.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != kebabshop.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ksuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ksuo.mutation.OsmID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: kebabshop.FieldOsmID,
		})
	}
	if value, ok := ksuo.mutation.AddedOsmID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: kebabshop.FieldOsmID,
		})
	}
	if ksuo.mutation.OsmIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: kebabshop.FieldOsmID,
		})
	}
	if value, ok := ksuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: kebabshop.FieldName,
		})
	}
	if value, ok := ksuo.mutation.Point(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: kebabshop.FieldPoint,
		})
	}
	_node = &KebabShop{config: ksuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ksuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{kebabshop.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
