package set

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetLen(t *testing.T) {
	s := NewGSet()
	s.Add("1")
	s.Add("2")
	assert.Equal(t, s.Len(), 2)
}

func TestSetLookup(t *testing.T) {
	s := NewGSet()
	s.Add("1")
	s.Add("2")
	assert.Equal(t, s.Lookup("2"), true)
	assert.Equal(t, s.Lookup("3"), false)
}
