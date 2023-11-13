package object_pool

import "ultimate-go/design-pattern/creational/simple_factory"

type IConnection interface {
	getId() int64
}

type connection struct {
	id int64
}

func (c *connection) getId() int64 {
	return c.id
}

func NewConnection() *connection {
	s, _ := simple_factory.NewIdGenerator()
	id := s.Generate()
	return &connection{id: id}
}
