package crdt

type Counter struct {
	id string
	data map[string]uint
}

func NewCounter() *Counter {
	return &Counter{
		data: make(map[string]uint),
	}
}

func (g *Counter) Inc(incr int) {
	g.counter[g.id] += incr
}

func (g *Counter) Count() (total int) {
	for _, val := range g.counter {
		total += val
	}
	return
}

// Merge provides merging of two counters
func (g *Counter) Merge(c *Counter) uint {
	var total uint
	for ident, val := range c.data {
		if v, ok := g.data[ident]; !ok || v < val {
			g.data[ident] = val
		}
	}

	return total
}
