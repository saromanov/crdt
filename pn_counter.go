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

func (p*PnCounter) Inc() {
	p.pos.Inc()
}

func (p*PnCounter) Dec() {
	p.neg.Inc()
}

func (p*PnCounter) Value() int{
	return p.pos.Count() - p.neg.Count()
}

func(p *PnCounter) Merge(n*PnCounter) {
	p.pos.Merge(n.pos)
	p.neg.Merge(n.neg)
}
