package object_pool

import (
	"fmt"
	"testing"
)

func TestPool_Get(t *testing.T) {
	conns := make([]IConnection, 0)
	conns = append(conns, NewConnection())
	conns = append(conns, NewConnection())
	conns = append(conns, NewConnection())
	pool, _ := initPool(conns)
	fmt.Println(pool.idle)
	c := pool.Get()
	fmt.Println(c.getId())
}
func TestNewConnection(t *testing.T) {
	conn := NewConnection()
	fmt.Println(conn.getId())
}
