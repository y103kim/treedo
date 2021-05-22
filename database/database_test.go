package database

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVersion(t *testing.T) {
	assert := assert.New(t)

	db := &Database{}
	filename := "test.db"

	assert.Nil(db.Open(filename))
	assert.Nil(db.getVersion())
	assert.Equal(db.version, 0)
	assert.Nil(db.setVersion(2))
	assert.Equal(db.version, 2)

	os.Remove(filename)
}
