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
	assert.Nil(db.GetVersion())
	assert.Equal(db.Version, 0)
	assert.Nil(db.SetVersion(2))
	assert.Equal(db.Version, 2)

	os.Remove(filename)
}
