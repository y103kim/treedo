package tools

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntSets(t *testing.T) {
	assert := assert.New(t)
	s := CreateIntSets()
	s.Add(10, 1, 2, 3, 4, 5)
	s.Add(10, 1, 3, 5, 8)
	s.Erase(10, 3, 4, 5, 6, 7)
	s.Add(11, 1, 2, 3, 4, 5)
	s.Add(11, 1, 3, 5, 8)
	s.Erase(11, 3, 4, 5, 6, 7)
	assert.Equal(true, s.Check(10, 1))
	assert.Equal(false, s.Check(10, 3))
	assert.Equal(false, s.Check(10, 10))
	assert.Equal(true, s.Check(10, 8))
	assert.ElementsMatch([]int{1, 2, 8}, s.List(10))

	assert.Equal(true, s.Check(11, 1))
	assert.Equal(false, s.Check(11, 3))
	assert.Equal(false, s.Check(11, 10))
	assert.Equal(true, s.Check(11, 8))
	assert.ElementsMatch([]int{1, 2, 8}, s.List(11))
}
