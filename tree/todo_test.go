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
	assert.Equal(int64(1), node.Id)

	copied := &Todo{}
	db.Get(copied, 1)
	assert.Equal(int64(1), copied.Id)
	assert.Equal("Test Todo", copied.Title)
	assert.Equal("Not Started", copied.Status)
	assert.NotEqual(int64(0), copied.CreatedAt)
	assert.NotEqual(int64(0), copied.UpdatedAt)

	copied.Title = "Updated Title"
	copied.Status = "Updated Status"
	copied.Hidden = 1
	copied.UpdatedAt += 1
	assert.Nil(db.Update(copied, []string{"Title", "Status", "Hidden", "UpdatedAt"}))

	updated := &Todo{}
	db.Get(updated, 1)
	assert.Equal(int64(1), copied.Id)
	assert.Equal("Updated Title", updated.Title)
	assert.Equal("Updated Status", updated.Status)
	assert.Equal(int64(1), copied.Hidden)
	assert.Equal(copied.UpdatedAt, updated.UpdatedAt)

	teardown(t, db)
}
