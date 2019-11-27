package set

type setObject map[interface{}]string

type Set interface {
	Add(value interface{})
	Contains(value interface{}) bool
	Remove(value interface{})
}
