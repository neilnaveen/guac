// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/predicate"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/sourcetype"
)

// SourceTypeDelete is the builder for deleting a SourceType entity.
type SourceTypeDelete struct {
	config
	hooks    []Hook
	mutation *SourceTypeMutation
}

// Where appends a list predicates to the SourceTypeDelete builder.
func (std *SourceTypeDelete) Where(ps ...predicate.SourceType) *SourceTypeDelete {
	std.mutation.Where(ps...)
	return std
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (std *SourceTypeDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, std.sqlExec, std.mutation, std.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (std *SourceTypeDelete) ExecX(ctx context.Context) int {
	n, err := std.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (std *SourceTypeDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(sourcetype.Table, sqlgraph.NewFieldSpec(sourcetype.FieldID, field.TypeInt))
	if ps := std.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, std.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	std.mutation.done = true
	return affected, err
}

// SourceTypeDeleteOne is the builder for deleting a single SourceType entity.
type SourceTypeDeleteOne struct {
	std *SourceTypeDelete
}

// Where appends a list predicates to the SourceTypeDelete builder.
func (stdo *SourceTypeDeleteOne) Where(ps ...predicate.SourceType) *SourceTypeDeleteOne {
	stdo.std.mutation.Where(ps...)
	return stdo
}

// Exec executes the deletion query.
func (stdo *SourceTypeDeleteOne) Exec(ctx context.Context) error {
	n, err := stdo.std.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{sourcetype.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (stdo *SourceTypeDeleteOne) ExecX(ctx context.Context) {
	if err := stdo.Exec(ctx); err != nil {
		panic(err)
	}
}