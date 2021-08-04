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
	assert.Equal(todos[1].Title, todo_korean.Title)

	teardown(t, db)
}
