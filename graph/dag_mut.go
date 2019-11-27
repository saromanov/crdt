package graph

import (
	"github.com/saromanov/crdt/set"
)

// DAGMut provides implementation of the add-removing dag
type DAGMut struct {
	added   *set.TwoPSet
	removed *set.TwoPSet
	edges   *set.GSet
}

func NewDAGMut() *DAGMut {
	return &DAGMut{
		added:   set.NewTwoPSet(),
		removed: set.NewTwoPSet(),
		edges:   set.NewGSet(),
	}
}
