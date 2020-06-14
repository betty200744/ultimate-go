package hashset

/*
	Add(elements ...interface{})
	Remove(elements ...interface{})
	Contains(elements ...interface{}) bool
*/
type Set struct {
	elements map[interface{}]struct{}
}

var itemExists = struct{}{}

func New(elements ...interface{}) *Set {
	set := &Set{elements: map[interface{}]struct{}{}}
	if len(elements) > 0 {
		set.Add(elements...)
	}
	return set
}
func (set *Set) Add(elements ...interface{}) {
	for _, item := range elements {
		set.elements[item] = itemExists
	}
}
func (set *Set) Remove(elements ...interface{}) {
	for _, i2 := range elements {
		delete(set.elements, i2)
	}
}
