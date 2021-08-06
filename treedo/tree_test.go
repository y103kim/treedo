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
	todo1, err := tree.CreateTodo("Todo 1")
	assert.Nil(err)
	assert.Equal("Todo 1", todo1.Title)

	todo_korean, err := tree.CreateTodo("한글 테스트")
	assert.Nil(err)
	assert.Equal("한글 테스트", todo_korean.Title)

	todos, err := tree.GetAllTodos()
	assert.Nil(err)
	assert.Len(todos, 2)
	assert.Equal(todos[0].Title, todo1.Title)
	assert.Equal(todos[0].ID, 1)
	assert.Equal(todos[1].Title, todo_korean.Title)
	assert.Equal(todos[1].ID, 2)

	assert.Nil(tree.LinkTodos(1, 2))
	children, err := tree.QueryChildren(1)
	assert.Equal(children[0].Title, todo_korean.Title)
	assert.Equal(children[0].ID, 2)

	_, err = tree.CreateTodo("Todo 3")
	_, err = tree.CreateTodo("Todo 4")
	assert.Nil(tree.LinkTodos(2, 3, 4))
	children, err = tree.QueryChildren(2)
	assert.Equal(children[0].ID, 3)
	assert.Equal(children[1].ID, 4)

	teardown(t, db)
}
