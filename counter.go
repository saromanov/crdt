package crdt

import (
	"fmt"
	"github.com/satori/go.uuid"
)

type Counter struct {
	id   string
	data map[string]uint
}

func NewCounter() *Counter {
	return &Counter{
		data: make(map[string]uint),
		id: uuid.Must(uuid.NewV4()),
	}
}

// Inc provides incrementation of the counter
func (g *Counter) Inc(incr uint) error {
	if incr == 0 {
		return fmt.Errorf("incr value can't be zero")
	}
	g.data[g.id] += incr
	return nil
}

func (g *Counter) Count() uint {
	var total uint
	for _, val := range g.data {
		total += val
	}
	return total
}

// Merge provides merging of two counters
func (g *Counter) Merge(c *Counter){
	for ident, val := range c.data {
		if v, ok := g.data[ident]; !ok || v < val {
			g.data[ident] = val
		}
	}
}
