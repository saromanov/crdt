package set

import uuid "github.com/satori/go.uuid"

type ORSet struct {
	add     map[interface{}]map[string]struct{}
	removed map[interface{}]map[string]struct{}
}

// NewORSet provides creating of the OrSet
func NewORSet() *ORSet {
	return &ORSet{
		add:     make(map[interface{}]map[string]struct{}),
		removed: make(map[interface{}]map[string]struct{}),
	}
}

// Add provides adding element to the set
func (o *ORSet) Add(value interface{}) {
	if m, ok := o.add[value]; ok {
		m[uuid.NewV4().String()] = struct{}{}
		o.add[value] = m
		return
	}

	m := make(map[string]struct{})

	m[uuid.NewV4().String()] = struct{}{}
	o.add[value] = m
}

func (o *ORSet) Remove(value interface{}) {
	r, ok := o.removed[value]
	if !ok {
		r = make(map[string]struct{})
	}

	if m, ok := o.add[value]; ok {
		for uid, _ := range m {
			r[uid] = struct{}{}
		}
	}

	o.removed[value] = r
}

func (o *ORSet) Contains(value interface{}) bool {
	if _, ok := o.add[value]; !ok {
		return false
	}

	if _, ok := o.removed[value]; !ok {
		return false
	}

	for uid := range o.add {
		if _, ok := o.removed[uid]; !ok {
			return true
		}
	}

	return false
}
