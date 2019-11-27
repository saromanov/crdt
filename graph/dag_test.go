package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddEdge(t *testing.T) {
	s := NewDAG()
	s.AddEdge("a", "b")
	s.AddEdge("a", "c")
	assert.Equal(t, s.LookupVertex("a"), true)
}
