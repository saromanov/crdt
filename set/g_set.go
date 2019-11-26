package set

// GSet provides implementation of the G-Set
type GSet struct {
	data map[interface{}]string
}

func NewGSet() *GSet {
	return &GSet{
		data: make(map[interface{}]string),
	}
}

func (g *GSet) Data() map[interface{}]string {
	return g.data
}

func (g *GSet) Add(value interface{}) {
	g.data[value] = "data"
}

func (g *GSet) Lookup(value interface{}) bool {
	_, ok := g.data[value]
	return ok
}

func (g *GSet) Merge(gs *GSet) {
	for _, d := range gs.Data() {
		if ok := g.Lookup(d); !ok {
			g.Add(d)
		}
	}
}

func (g *GSet) Len() int {
	return len(g.data)
}
