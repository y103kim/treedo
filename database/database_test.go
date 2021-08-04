package database

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/y103kim/treedo/ent"
)

func TestTransaction(t *testing.T) {
	assert := assert.New(t)
	database := CreateDatabase("test.sqlite")
	assert.Nil(database.Tx(func(ctx context.Context, tx *ent.Tx) error {
		assert.Zero(tx.Todo.Query().Count(ctx))
		return nil
	}))
	assert.Nil(database.Close())
}
