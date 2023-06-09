// Code generated by entc, DO NOT EDIT.

package ent

import (
	"asr-demo-server/pkg/ent/candidate"
	"asr-demo-server/pkg/ent/wordalignment"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// WordAlignmentCreate is the builder for creating a WordAlignment entity.
type WordAlignmentCreate struct {
	config
	mutation *WordAlignmentMutation
	hooks    []Hook
}

// SetStart sets the "start" field.
func (wac *WordAlignmentCreate) SetStart(f float64) *WordAlignmentCreate {
	wac.mutation.SetStart(f)
	return wac
}

// SetLength sets the "length" field.
func (wac *WordAlignmentCreate) SetLength(f float64) *WordAlignmentCreate {
	wac.mutation.SetLength(f)
	return wac
}

// SetWord sets the "word" field.
func (wac *WordAlignmentCreate) SetWord(s string) *WordAlignmentCreate {
	wac.mutation.SetWord(s)
	return wac
}

// SetToken sets the "token" field.
func (wac *WordAlignmentCreate) SetToken(s string) *WordAlignmentCreate {
	wac.mutation.SetToken(s)
	return wac
}

// SetConfidence sets the "confidence" field.
func (wac *WordAlignmentCreate) SetConfidence(f float64) *WordAlignmentCreate {
	wac.mutation.SetConfidence(f)
	return wac
}

// AddCandidateIDs adds the "candidates" edge to the Candidate entity by IDs.
func (wac *WordAlignmentCreate) AddCandidateIDs(ids ...int) *WordAlignmentCreate {
	wac.mutation.AddCandidateIDs(ids...)
	return wac
}

// AddCandidates adds the "candidates" edges to the Candidate entity.
func (wac *WordAlignmentCreate) AddCandidates(c ...*Candidate) *WordAlignmentCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return wac.AddCandidateIDs(ids...)
}

// Mutation returns the WordAlignmentMutation object of the builder.
func (wac *WordAlignmentCreate) Mutation() *WordAlignmentMutation {
	return wac.mutation
}

// Save creates the WordAlignment in the database.
func (wac *WordAlignmentCreate) Save(ctx context.Context) (*WordAlignment, error) {
	var (
		err  error
		node *WordAlignment
	)
	if len(wac.hooks) == 0 {
		if err = wac.check(); err != nil {
			return nil, err
		}
		node, err = wac.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WordAlignmentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = wac.check(); err != nil {
				return nil, err
			}
			wac.mutation = mutation
			if node, err = wac.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(wac.hooks) - 1; i >= 0; i-- {
			if wac.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = wac.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, wac.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (wac *WordAlignmentCreate) SaveX(ctx context.Context) *WordAlignment {
	v, err := wac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wac *WordAlignmentCreate) Exec(ctx context.Context) error {
	_, err := wac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wac *WordAlignmentCreate) ExecX(ctx context.Context) {
	if err := wac.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wac *WordAlignmentCreate) check() error {
	if _, ok := wac.mutation.Start(); !ok {
		return &ValidationError{Name: "start", err: errors.New(`ent: missing required field "start"`)}
	}
	if v, ok := wac.mutation.Start(); ok {
		if err := wordalignment.StartValidator(v); err != nil {
			return &ValidationError{Name: "start", err: fmt.Errorf(`ent: validator failed for field "start": %w`, err)}
		}
	}
	if _, ok := wac.mutation.Length(); !ok {
		return &ValidationError{Name: "length", err: errors.New(`ent: missing required field "length"`)}
	}
	if v, ok := wac.mutation.Length(); ok {
		if err := wordalignment.LengthValidator(v); err != nil {
			return &ValidationError{Name: "length", err: fmt.Errorf(`ent: validator failed for field "length": %w`, err)}
		}
	}
	if _, ok := wac.mutation.Word(); !ok {
		return &ValidationError{Name: "word", err: errors.New(`ent: missing required field "word"`)}
	}
	if v, ok := wac.mutation.Word(); ok {
		if err := wordalignment.WordValidator(v); err != nil {
			return &ValidationError{Name: "word", err: fmt.Errorf(`ent: validator failed for field "word": %w`, err)}
		}
	}
	if _, ok := wac.mutation.Token(); !ok {
		return &ValidationError{Name: "token", err: errors.New(`ent: missing required field "token"`)}
	}
	if _, ok := wac.mutation.Confidence(); !ok {
		return &ValidationError{Name: "confidence", err: errors.New(`ent: missing required field "confidence"`)}
	}
	if v, ok := wac.mutation.Confidence(); ok {
		if err := wordalignment.ConfidenceValidator(v); err != nil {
			return &ValidationError{Name: "confidence", err: fmt.Errorf(`ent: validator failed for field "confidence": %w`, err)}
		}
	}
	return nil
}

func (wac *WordAlignmentCreate) sqlSave(ctx context.Context) (*WordAlignment, error) {
	_node, _spec := wac.createSpec()
	if err := sqlgraph.CreateNode(ctx, wac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (wac *WordAlignmentCreate) createSpec() (*WordAlignment, *sqlgraph.CreateSpec) {
	var (
		_node = &WordAlignment{config: wac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: wordalignment.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: wordalignment.FieldID,
			},
		}
	)
	if value, ok := wac.mutation.Start(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: wordalignment.FieldStart,
		})
		_node.Start = value
	}
	if value, ok := wac.mutation.Length(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: wordalignment.FieldLength,
		})
		_node.Length = value
	}
	if value, ok := wac.mutation.Word(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: wordalignment.FieldWord,
		})
		_node.Word = value
	}
	if value, ok := wac.mutation.Token(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: wordalignment.FieldToken,
		})
		_node.Token = value
	}
	if value, ok := wac.mutation.Confidence(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: wordalignment.FieldConfidence,
		})
		_node.Confidence = value
	}
	if nodes := wac.mutation.CandidatesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   wordalignment.CandidatesTable,
			Columns: []string{wordalignment.CandidatesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: candidate.FieldID,
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

// WordAlignmentCreateBulk is the builder for creating many WordAlignment entities in bulk.
type WordAlignmentCreateBulk struct {
	config
	builders []*WordAlignmentCreate
}

// Save creates the WordAlignment entities in the database.
func (wacb *WordAlignmentCreateBulk) Save(ctx context.Context) ([]*WordAlignment, error) {
	specs := make([]*sqlgraph.CreateSpec, len(wacb.builders))
	nodes := make([]*WordAlignment, len(wacb.builders))
	mutators := make([]Mutator, len(wacb.builders))
	for i := range wacb.builders {
		func(i int, root context.Context) {
			builder := wacb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*WordAlignmentMutation)
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
					_, err = mutators[i+1].Mutate(root, wacb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, wacb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, wacb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (wacb *WordAlignmentCreateBulk) SaveX(ctx context.Context) []*WordAlignment {
	v, err := wacb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wacb *WordAlignmentCreateBulk) Exec(ctx context.Context) error {
	_, err := wacb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wacb *WordAlignmentCreateBulk) ExecX(ctx context.Context) {
	if err := wacb.Exec(ctx); err != nil {
		panic(err)
	}
}
