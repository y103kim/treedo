package treedo

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/y103kim/treedo/database"
)

func setup(t *testing.T) *database.Database {
	os.Remove("test.sqlite")
	db := database.CreateDatabase("test.sqlite")
	return db
}

func teardown(t *testing.T, db *database.Database) {
	db.Close()
	os.Remove("test.sqlite")
}

func TestTransaction(t *testing.T) {
	assert := assert.New(t)
	db := setup(t)
	tree := CreateTree(db)
	assert.Nil(tree.CreateTodo("Todo 1"))
	assert.Equal("Todo 1", tree.Nodes[1].Title)

	assert.Nil(tree.CreateTodo("할일 2"))
	assert.Equal("할일 2", tree.Nodes[2].Title)

	assert.Nil(tree.LinkTodos(1, 2))
	children := tree.Nodes[1].Edges.Child
	assert.Equal(children[0].ID, 2)

	assert.Nil(tree.CreateTodo("Todo 3"))
	assert.Nil(tree.CreateTodo("Todo 4"))
	assert.Nil(tree.LinkTodos(2, 3, 4))
	children = tree.Nodes[2].Edges.Child
	assert.Equal(children[0].ID, 3)
	assert.Equal(children[1].ID, 4)

	teardown(t, db)
}
