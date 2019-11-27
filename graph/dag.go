package graph

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
