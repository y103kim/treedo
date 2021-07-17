package database

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setup(t *testing.T) *Database {
	db := &Database{}
	os.Remove("test.db")
	db.Open("test.db")
	return db
}

func teardown(t *testing.T, db *Database) {
	db.Close()
	os.Remove("test.db")
}

func TestVersion(t *testing.T) {
	assert := assert.New(t)
	db := setup(t)

	assert.Nil(db.getVersion())
	assert.Equal(db.version, 0)
	assert.Nil(db.setVersion(2))
	assert.Equal(db.version, 2)

	teardown(t, db)
}

func TestMigrate(t *testing.T) {
	assert := assert.New(t)
	db := setup(t)

	// First migration 0 => 1
	assert.Nil(db.Migrate())
	// Second migration 1 => 1
	assert.Nil(db.Migrate())

	teardown(t, db)
}
