package treedo

import (
	"context"

	"github.com/pkg/errors"
	"github.com/y103kim/treedo/database"
	"github.com/y103kim/treedo/ent"
)

type Tree struct {
	db      *database.Database
	Nodes   map[int]*ent.Todo
	Parents map[int][]int
}

func CreateTree(db *database.Database) *Tree {
	tree := &Tree{db, nil, nil}
	tree.Nodes = make(map[int]*ent.Todo)
	tree.Parents = make(map[int][]int)
	return tree
}

func (t *Tree) CreateTodo(title string) error {
	return t.db.Tx(func(ctx context.Context, tx *ent.Tx) error {
		if res, err := tx.Todo.Create().SetTitle(title).Save(ctx); err != nil {
			return errors.Wrapf(err, "Fail to create todo with title '%s'\n", title)
		} else {
			t.Nodes[res.ID] = res
			return nil
		}
	})
}

func (t *Tree) FetchAll() error {
	return t.db.Tx(func(ctx context.Context, tx *ent.Tx) error {
		if results, err := tx.Todo.Query().WithChild().All(ctx); err != nil {
			return errors.Wrap(err, "Fail to get all todos")
		} else {
			for _, res := range results {
				t.Nodes[res.ID] = res
			}
			return nil
		}
	})
}

func (t *Tree) LinkTodos(from int, to ...int) error {
	return t.db.Tx(func(ctx context.Context, tx *ent.Tx) error {
		if res, err := tx.Todo.UpdateOneID(from).AddChildIDs(to...).Save(ctx); err != nil {
			return errors.Wrapf(err, "Cannot link todo %d->%v", from, to)
		} else {
			for _, todo_id := range to {
				res.Edges.Child = append(res.Edges.Child, t.Nodes[todo_id])
			}
			t.Nodes[res.ID] = res
			return nil
		}
	})
}
