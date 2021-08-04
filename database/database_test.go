package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransaction(t *testing.T) {
	assert := assert.New(t)
	database := CreateDatabase("test.sqlite")
	assert.Nil(database.Close())
}
