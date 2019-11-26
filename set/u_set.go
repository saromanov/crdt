package set

type USet struct {
	data setObject
}

func (u *USet) Add(value interface{}) {
	u.data[value] = "data"
}
