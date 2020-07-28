package object_pool

import (
	"fmt"
	"sync"
)

/*
The Object Pool Design Pattern is a creational design pattern,
in which a Pool of objects is initialized and created beforehand and kept in a Pool.
As and when needed, a client can request an object from the Pool, use it, and return it to the Pool.
The object in the Pool is never destroyed.

*/

type Pool struct {
	idle     []IConnection
	active   []IConnection
	capacity int
	mulock   *sync.Mutex
}

func initPool(conns []IConnection) (*Pool, error) {
	if len(conns) == 0 {
		return nil, fmt.Errorf("cannot create a pool of 0 length")
	}
	active := make([]IConnection, 0)
	return &Pool{
		idle:     conns,
		active:   active,
		capacity: len(conns),
		mulock:   new(sync.Mutex),
	}, nil
}

func (p *Pool) Get() IConnection {
	if len(p.idle) == 0 {
		conn := NewConnection()
		_ = p.Put(conn)
	}
	c := p.idle[0]
	p.idle = p.idle[1:]
	return c
}
func (p *Pool) Put(conn IConnection) error {
	p.idle = append(p.idle, conn)
	return nil
}
