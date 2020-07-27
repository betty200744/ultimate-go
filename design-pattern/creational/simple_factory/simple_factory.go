package simple_factory

import (
	"sync"
	"time"
)

/*
go 语言没有构造函数一说，所以一般会定义NewXXX函数来初始化相关类。 NewXXX 函数返回接口时就是简单工厂模式，也就是说Golang的一般推荐做法就是简单工厂。
*/

// A IdGenerator struct holds the basic information needed for a snowflake generator worker
type IdGenerator struct {
	sync.Mutex
	timestamp int64
	worker    int64
	sequence  int64
}

// NewNode returns a new snowflake worker that can be used to generate snowflake IDs
func NewIdGenerator() (*IdGenerator, error) {
	workerId := 0x3F
	return &IdGenerator{
		timestamp: 0,
		worker:    int64(workerId),
		sequence:  0,
	}, nil
}

// Generate creates and returns a unique snowflake ID
func (s *IdGenerator) Generate() int64 {
	s.Lock()
	defer s.Unlock()
	now := time.Now().UnixNano() / 1000000
	s.timestamp = now
	r := ((now) & 0x01FFFFFFFFFF) | (s.worker) | (s.sequence)
	return r
}
