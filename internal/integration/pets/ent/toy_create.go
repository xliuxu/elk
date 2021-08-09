// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/masseelch/elk/internal/integration/pets/ent/pet"
	"github.com/masseelch/elk/internal/integration/pets/ent/toy"
)

// ToyCreate is the builder for creating a Toy entity.
type ToyCreate struct {
	config
	mutation *ToyMutation
	hooks    []Hook
}

// SetColor sets the "color" field.
func (tc *ToyCreate) SetColor(t toy.Color) *ToyCreate {
	tc.mutation.SetColor(t)
	return tc
}

// SetMaterial sets the "material" field.
func (tc *ToyCreate) SetMaterial(t toy.Material) *ToyCreate {
	tc.mutation.SetMaterial(t)
	return tc
}

// SetTitle sets the "title" field.
func (tc *ToyCreate) SetTitle(s string) *ToyCreate {
	tc.mutation.SetTitle(s)
	return tc
}

// SetOwnerID sets the "owner" edge to the Pet entity by ID.
func (tc *ToyCreate) SetOwnerID(id int) *ToyCreate {
	tc.mutation.SetOwnerID(id)
	return tc
}

// SetNillableOwnerID sets the "owner" edge to the Pet entity by ID if the given value is not nil.
func (tc *ToyCreate) SetNillableOwnerID(id *int) *ToyCreate {
	if id != nil {
		tc = tc.SetOwnerID(*id)
	}
	return tc
}

// SetOwner sets the "owner" edge to the Pet entity.
func (tc *ToyCreate) SetOwner(p *Pet) *ToyCreate {
	return tc.SetOwnerID(p.ID)
}

// Mutation returns the ToyMutation object of the builder.
func (tc *ToyCreate) Mutation() *ToyMutation {
	return tc.mutation
}

// Save creates the Toy in the database.
func (tc *ToyCreate) Save(ctx context.Context) (*Toy, error) {
	var (
		err  error
		node *Toy
	)
	if len(tc.hooks) == 0 {
		if err = tc.check(); err != nil {
			return nil, err
		}
		node, err = tc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ToyMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tc.check(); err != nil {
				return nil, err
			}
			tc.mutation = mutation
			if node, err = tc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(tc.hooks) - 1; i >= 0; i-- {
			if tc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (tc *ToyCreate) SaveX(ctx context.Context) *Toy {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *ToyCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *ToyCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *ToyCreate) check() error {
	if _, ok := tc.mutation.Color(); !ok {
		return &ValidationError{Name: "color", err: errors.New(`ent: missing required field "color"`)}
	}
	if v, ok := tc.mutation.Color(); ok {
		if err := toy.ColorValidator(v); err != nil {
			return &ValidationError{Name: "color", err: fmt.Errorf(`ent: validator failed for field "color": %w`, err)}
		}
	}
	if _, ok := tc.mutation.Material(); !ok {
		return &ValidationError{Name: "material", err: errors.New(`ent: missing required field "material"`)}
	}
	if v, ok := tc.mutation.Material(); ok {
		if err := toy.MaterialValidator(v); err != nil {
			return &ValidationError{Name: "material", err: fmt.Errorf(`ent: validator failed for field "material": %w`, err)}
		}
	}
	if _, ok := tc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "title"`)}
	}
	return nil
}

func (tc *ToyCreate) sqlSave(ctx context.Context) (*Toy, error) {
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (tc *ToyCreate) createSpec() (*Toy, *sqlgraph.CreateSpec) {
	var (
		_node = &Toy{config: tc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: toy.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: toy.FieldID,
			},
		}
	)
	if value, ok := tc.mutation.Color(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: toy.FieldColor,
		})
		_node.Color = value
	}
	if value, ok := tc.mutation.Material(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: toy.FieldMaterial,
		})
		_node.Material = value
	}
	if value, ok := tc.mutation.Title(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: toy.FieldTitle,
		})
		_node.Title = value
	}
	if nodes := tc.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   toy.OwnerTable,
			Columns: []string{toy.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: pet.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.pet_toys = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ToyCreateBulk is the builder for creating many Toy entities in bulk.
type ToyCreateBulk struct {
	config
	builders []*ToyCreate
}

// Save creates the Toy entities in the database.
func (tcb *ToyCreateBulk) Save(ctx context.Context) ([]*Toy, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Toy, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ToyMutation)
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
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *ToyCreateBulk) SaveX(ctx context.Context) []*Toy {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *ToyCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *ToyCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}
