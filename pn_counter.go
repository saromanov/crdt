package crdt

// PnCounter defines implementation of the PN-COUNTER
type PnCounter struct {
	pos *Counter
	neg *Counter
}

// NewPnCounter initializes pn counter
func NewPnCounter()*PnCounter {
	return &PnCounter{
		pos: NewCounter(),
		neg: NewCounter(),
	}
}

func (p*PnCounter) Inc(num uint) {
	p.pos.Inc(num)
}

func (p*PnCounter) Dec() {
	p.neg.Inc(1)
}

func (p*PnCounter) Value() int{
	return int(p.pos.Count()) - int(p.neg.Count())
}

func(p *PnCounter) Merge(n*PnCounter) {
	p.pos.Merge(n.pos)
	p.neg.Merge(n.neg)
}
