package treedo

import (
	"context"

	"github.com/pkg/errors"
	"github.com/y103kim/treedo/database"
	"github.com/y103kim/treedo/ent"
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
