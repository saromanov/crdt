package crdt

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestIncCounter(t *testing.T) {
	c := NewCounter()
	c.Inc(1)
	c.Inc(1)
	assert.Equal(t, c.Count(), uint(2))
}
