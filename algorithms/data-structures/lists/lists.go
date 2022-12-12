package lists

import "ultimate-go/algorithms/data-structures/containers"

// 数据有index, 数据可重复
type List interface {
	Get(index int) interface{}
	Set(index int, value interface{})
	Add(values ...interface{})
	Remove(index int)
	Insert(index int, value interface{})
	Contains(values ...interface{}) bool
	Swap(index1, index2 int)
	containers.Container
	// Empty() bool
	// Size() int
	// Clear()
	// Values() []interface{}
}
