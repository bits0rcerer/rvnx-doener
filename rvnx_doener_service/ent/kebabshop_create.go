// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"rvnx_doener_service/ent/kebabshop"
	"rvnx_doener_service/ent/scorerating"
	"rvnx_doener_service/ent/shopprice"
	"rvnx_doener_service/ent/useropinion"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// KebabShopCreate is the builder for creating a KebabShop entity.
type KebabShopCreate struct {
	config
	mutation *KebabShopMutation
	hooks    []Hook
}

// SetOsmID sets the "osm_id" field.
func (ksc *KebabShopCreate) SetOsmID(i int) *KebabShopCreate {
	ksc.mutation.SetOsmID(i)
	return ksc
}

// SetNillableOsmID sets the "osm_id" field if the given value is not nil.
func (ksc *KebabShopCreate) SetNillableOsmID(i *int) *KebabShopCreate {
	if i != nil {
		ksc.SetOsmID(*i)
	}
	return ksc
}

// SetName sets the "name" field.
func (ksc *KebabShopCreate) SetName(s string) *KebabShopCreate {
	ksc.mutation.SetName(s)
	return ksc
}

// SetCreated sets the "created" field.
func (ksc *KebabShopCreate) SetCreated(t time.Time) *KebabShopCreate {
	ksc.mutation.SetCreated(t)
	return ksc
}

// SetNillableCreated sets the "created" field if the given value is not nil.
func (ksc *KebabShopCreate) SetNillableCreated(t *time.Time) *KebabShopCreate {
	if t != nil {
		ksc.SetCreated(*t)
	}
	return ksc
}

// SetLat sets the "lat" field.
func (ksc *KebabShopCreate) SetLat(f float64) *KebabShopCreate {
	ksc.mutation.SetLat(f)
	return ksc
}

// SetLng sets the "lng" field.
func (ksc *KebabShopCreate) SetLng(f float64) *KebabShopCreate {
	ksc.mutation.SetLng(f)
	return ksc
}

// SetVisible sets the "visible" field.
func (ksc *KebabShopCreate) SetVisible(b bool) *KebabShopCreate {
	ksc.mutation.SetVisible(b)
	return ksc
}

// SetNillableVisible sets the "visible" field if the given value is not nil.
func (ksc *KebabShopCreate) SetNillableVisible(b *bool) *KebabShopCreate {
	if b != nil {
		ksc.SetVisible(*b)
	}
	return ksc
}

// SetID sets the "id" field.
func (ksc *KebabShopCreate) SetID(u uint64) *KebabShopCreate {
	ksc.mutation.SetID(u)
	return ksc
}

// AddUserScoreIDs adds the "user_scores" edge to the ScoreRating entity by IDs.
func (ksc *KebabShopCreate) AddUserScoreIDs(ids ...uint64) *KebabShopCreate {
	ksc.mutation.AddUserScoreIDs(ids...)
	return ksc
}

// AddUserScores adds the "user_scores" edges to the ScoreRating entity.
func (ksc *KebabShopCreate) AddUserScores(s ...*ScoreRating) *KebabShopCreate {
	ids := make([]uint64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return ksc.AddUserScoreIDs(ids...)
}

// AddUserPriceIDs adds the "user_prices" edge to the ShopPrice entity by IDs.
func (ksc *KebabShopCreate) AddUserPriceIDs(ids ...uint64) *KebabShopCreate {
	ksc.mutation.AddUserPriceIDs(ids...)
	return ksc
}

// AddUserPrices adds the "user_prices" edges to the ShopPrice entity.
func (ksc *KebabShopCreate) AddUserPrices(s ...*ShopPrice) *KebabShopCreate {
	ids := make([]uint64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return ksc.AddUserPriceIDs(ids...)
}

// AddUserOpinionIDs adds the "user_opinions" edge to the UserOpinion entity by IDs.
func (ksc *KebabShopCreate) AddUserOpinionIDs(ids ...uint64) *KebabShopCreate {
	ksc.mutation.AddUserOpinionIDs(ids...)
	return ksc
}

// AddUserOpinions adds the "user_opinions" edges to the UserOpinion entity.
func (ksc *KebabShopCreate) AddUserOpinions(u ...*UserOpinion) *KebabShopCreate {
	ids := make([]uint64, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return ksc.AddUserOpinionIDs(ids...)
}

// Mutation returns the KebabShopMutation object of the builder.
func (ksc *KebabShopCreate) Mutation() *KebabShopMutation {
	return ksc.mutation
}

// Save creates the KebabShop in the database.
func (ksc *KebabShopCreate) Save(ctx context.Context) (*KebabShop, error) {
	var (
		err  error
		node *KebabShop
	)
	if err := ksc.defaults(); err != nil {
		return nil, err
	}
	if len(ksc.hooks) == 0 {
		if err = ksc.check(); err != nil {
			return nil, err
		}
		node, err = ksc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*KebabShopMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ksc.check(); err != nil {
				return nil, err
			}
			ksc.mutation = mutation
			if node, err = ksc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ksc.hooks) - 1; i >= 0; i-- {
			if ksc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ksc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ksc.mutation)
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

// SaveX calls Save and panics if Save returns an error.
func (ksc *KebabShopCreate) SaveX(ctx context.Context) *KebabShop {
	v, err := ksc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ksc *KebabShopCreate) Exec(ctx context.Context) error {
	_, err := ksc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ksc *KebabShopCreate) ExecX(ctx context.Context) {
	if err := ksc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ksc *KebabShopCreate) defaults() error {
	if _, ok := ksc.mutation.Created(); !ok {
		if kebabshop.DefaultCreated == nil {
			return fmt.Errorf("ent: uninitialized kebabshop.DefaultCreated (forgotten import ent/runtime?)")
		}
		v := kebabshop.DefaultCreated()
		ksc.mutation.SetCreated(v)
	}
	if _, ok := ksc.mutation.Visible(); !ok {
		v := kebabshop.DefaultVisible
		ksc.mutation.SetVisible(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (ksc *KebabShopCreate) check() error {
	if _, ok := ksc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "KebabShop.name"`)}
	}
	if _, ok := ksc.mutation.Created(); !ok {
		return &ValidationError{Name: "created", err: errors.New(`ent: missing required field "KebabShop.created"`)}
	}
	if _, ok := ksc.mutation.Lat(); !ok {
		return &ValidationError{Name: "lat", err: errors.New(`ent: missing required field "KebabShop.lat"`)}
	}
	if _, ok := ksc.mutation.Lng(); !ok {
		return &ValidationError{Name: "lng", err: errors.New(`ent: missing required field "KebabShop.lng"`)}
	}
	if _, ok := ksc.mutation.Visible(); !ok {
		return &ValidationError{Name: "visible", err: errors.New(`ent: missing required field "KebabShop.visible"`)}
	}
	return nil
}

func (ksc *KebabShopCreate) sqlSave(ctx context.Context) (*KebabShop, error) {
	_node, _spec := ksc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ksc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint64(id)
	}
	return _node, nil
}

func (ksc *KebabShopCreate) createSpec() (*KebabShop, *sqlgraph.CreateSpec) {
	var (
		_node = &KebabShop{config: ksc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: kebabshop.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: kebabshop.FieldID,
			},
		}
	)
	if id, ok := ksc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ksc.mutation.OsmID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: kebabshop.FieldOsmID,
		})
		_node.OsmID = &value
	}
	if value, ok := ksc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: kebabshop.FieldName,
		})
		_node.Name = value
	}
	if value, ok := ksc.mutation.Created(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: kebabshop.FieldCreated,
		})
		_node.Created = value
	}
	if value, ok := ksc.mutation.Lat(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: kebabshop.FieldLat,
		})
		_node.Lat = value
	}
	if value, ok := ksc.mutation.Lng(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: kebabshop.FieldLng,
		})
		_node.Lng = value
	}
	if value, ok := ksc.mutation.Visible(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: kebabshop.FieldVisible,
		})
		_node.Visible = value
	}
	if nodes := ksc.mutation.UserScoresIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   kebabshop.UserScoresTable,
			Columns: []string{kebabshop.UserScoresColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: scorerating.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ksc.mutation.UserPricesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   kebabshop.UserPricesTable,
			Columns: []string{kebabshop.UserPricesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: shopprice.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ksc.mutation.UserOpinionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   kebabshop.UserOpinionsTable,
			Columns: []string{kebabshop.UserOpinionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: useropinion.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// KebabShopCreateBulk is the builder for creating many KebabShop entities in bulk.
type KebabShopCreateBulk struct {
	config
	builders []*KebabShopCreate
}

// Save creates the KebabShop entities in the database.
func (kscb *KebabShopCreateBulk) Save(ctx context.Context) ([]*KebabShop, error) {
	specs := make([]*sqlgraph.CreateSpec, len(kscb.builders))
	nodes := make([]*KebabShop, len(kscb.builders))
	mutators := make([]Mutator, len(kscb.builders))
	for i := range kscb.builders {
		func(i int, root context.Context) {
			builder := kscb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*KebabShopMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, kscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, kscb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = uint64(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, kscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (kscb *KebabShopCreateBulk) SaveX(ctx context.Context) []*KebabShop {
	v, err := kscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (kscb *KebabShopCreateBulk) Exec(ctx context.Context) error {
	_, err := kscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (kscb *KebabShopCreateBulk) ExecX(ctx context.Context) {
	if err := kscb.Exec(ctx); err != nil {
		panic(err)
	}
}
