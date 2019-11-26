package set

// GSet provides implementation of the G-Set
type GSet struct {
	data setObject
}

func NewGSet() *GSet {
	return &GSet{
		data: make(setObject),
	}
}

func (g *GSet) Data() setObject {
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
