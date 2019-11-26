package crdt

// GSet provides implementation of the G-Set
type GSet struct {
	data map[interface{}] string
}

func NewGSet() *GSet {
	return &GSet {
		data: make(map[interface{}]string),
	}
}

func (g*GSet) Add(value interface{}) {
	g.data[value] = "data"
}

func (g*GSet) Lookup(value interface{}) bool {
	_, ok := g.data[value]
	return ok
}

