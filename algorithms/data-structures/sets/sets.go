package sets

import "ultimate-go/algorithms/data-structures/containers"

// // Reference: https://en.wikipedia.org/wiki/Set_%28abstract_data_type%29
// 数据不可重复， 数据无index
// A list is a data structure that stores values and may have repeated values

type Set interface {
	Add(elements ...interface{})
	Remove(elements ...interface{})
	Contains(elements ...interface{}) bool
	containers.Container
	// Empty() bool
	// Size() int
	// Clear()
	// Values() []interface{}
}
