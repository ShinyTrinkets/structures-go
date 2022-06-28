package structs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTable(t *testing.T) {
	assert := assert.New(t)

	table := NewCache()
	assert.Equal(0, table.Size())
	assert.False(table.Exists("x"))
	assert.False(table.Exists("y"))

	table.Set("x", 0.0)
	assert.Equal(1, table.Size())
	assert.True(table.Exists("x"))

	v, ok := table.Get("x")
	assert.True(ok)
	assert.Equal(0.0, v)

	table.Set("y", 987)
	assert.Equal(2, table.Size())
	assert.True(table.Exists("y"))
}
