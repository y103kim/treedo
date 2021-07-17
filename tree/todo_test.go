package tree

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/y103kim/treedo/database"
)

func setup(t *testing.T) *database.Database {
	db := &database.Database{}
	os.Remove("test.db")
	db.Open("test.db")
	return db
}

func teardown(t *testing.T, db *database.Database) {
	db.Close()
	os.Remove("test.db")
}

func TestCRUD(t *testing.T) {
	assert := assert.New(t)
	db := setup(t)
	assert.Nil(db.Migrate())

	node := CreateTodo("Test Todo")
	assert.Nil(db.Insert(node))
	assert.Equal(node.id, int64(1))

	teardown(t, db)
}
