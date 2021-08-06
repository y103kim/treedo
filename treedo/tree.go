package treedo

import (
	"context"

	"github.com/pkg/errors"
	"github.com/y103kim/treedo/database"
	"github.com/y103kim/treedo/ent"
	"github.com/y103kim/treedo/ent/todo"
)

type Tree struct {
	db *database.Database
}

func CreateTree(db *database.Database) *Tree {
	tree := &Tree{db}
	return tree
}

func (t *Tree) CreateTodo(title string) (*ent.Todo, error) {
	var todo *ent.Todo
	err := t.db.Tx(func(ctx context.Context, tx *ent.Tx) error {
		if saved, err := tx.Todo.Create().SetTitle(title).Save(ctx); err != nil {
			return errors.Wrapf(err, "Fail to create todo with title '%s'\n", title)
		} else {
			todo = saved
			return nil
		}
	})
	return todo, err
}

func (t *Tree) GetAllTodos() ([]*ent.Todo, error) {
	var todos []*ent.Todo
	err := t.db.Tx(func(ctx context.Context, tx *ent.Tx) error {
		if saved, err := tx.Todo.Query().All(ctx); err != nil {
			return errors.Wrap(err, "Fail to get all todos")
		} else {
			todos = saved
			return nil
		}
	})
	return todos, err
}

func (t *Tree) LinkTodos(from int, to ...int) error {
	return t.db.Tx(func(ctx context.Context, tx *ent.Tx) error {
		if _, err := tx.Todo.UpdateOneID(from).AddChildIDs(to...).Save(ctx); err != nil {
			return errors.Wrapf(err, "Cannot link todo %d->%v", from, to)
		} else {
			return nil
		}
	})
}

func (t *Tree) QueryChildren(from int) ([]*ent.Todo, error) {
	var todos []*ent.Todo
	err := t.db.Tx(func(ctx context.Context, tx *ent.Tx) error {
		cond := todo.HasParentWith(todo.ID(from))
		if children, err := tx.Todo.Query().Where(cond).All(ctx); err != nil {
			return errors.Wrapf(err, "Cannot query children of todo from %d", from)
		} else {
			todos = children
			return nil
		}
	})
	return todos, err
}
