package graph

import "fmt"

// DAG provides implementation of add only graph
type DAG struct {
	vertices []interface{}
	edges    map[interface{}][]interface{}
}

func NewDAG() *DAG {
	return &DAG{
		vertices: []interface{}{},
		edges:    make(map[interface{}][]interface{}),
	}
}

func (d *DAG) LookupVertex(v interface{}) bool {
	if _, ok := d.edges[v]; ok {
		return true
	}

	return false
}

func (d *DAG) LookupEdge(a, b interface{}) bool {
	verts, ok := d.edges[a]
	if !ok {
		return false
	}

	for _, v := range verts {
		if v == b {
			return true
		}
	}

	return false
}

func (d *DAG) AddEdge(a, b interface{}) error {
	if ok := d.LookupVertex(a); ok {
		return fmt.Errorf("vertex is exist")
	}

	d.edges[a] = append(d.edges[a], b)
	return nil
}
