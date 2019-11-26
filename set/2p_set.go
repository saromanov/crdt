package set

type TwoPSet struct {
	added   setObject
	removed setObject
}

func NewTwoPSet() *TwoPSet {
	return &TwoPSet{
		added:   make(setObject),
		removed: make(setObject),
	}
}

func (g *TwoPSet) Add(value interface{}) {
	g.added[value] = "data"
}

func (g *TwoPSet) Remove(value interface{}) {
	g.removed[value] = "data"
}

func (g *TwoPSet) Lookup(value interface{}) bool {
	_, ok := g.added[value]
	_, okRem := g.removed[value]
	return ok && !okRem
}
