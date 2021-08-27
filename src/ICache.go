package src

type Cache interface {
	Size() int
	Set(key interface{}, value interface{})
	Get(key interface{}) interface{}
	Remove(key interface{})
	Clear()
}