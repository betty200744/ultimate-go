package containers

// All data structures implement the container interface with the following methods
type Container interface {
	Empty() bool
	Size() int
	Clear()
	Values() []interface{}
}
