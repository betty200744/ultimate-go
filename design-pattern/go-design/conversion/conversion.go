package main

import "fmt"

type Move interface {
	Move()
}
type Lock interface {
	Lock()
}
type MoveLock interface {
	Move
	Lock
}
type bike struct{}

func (b *bike) Move() {
	fmt.Println("bike move")
}
func (b *bike) Lock() {
	fmt.Println("bike lock")
}

func main() {
	var m Move
	var ml MoveLock
	ml = &bike{}
	ml.Move()
	m = ml
	m.Move()
	b, _ := m.(*bike)
	ml = b
	ml.Move()
	ml.Lock()
}
