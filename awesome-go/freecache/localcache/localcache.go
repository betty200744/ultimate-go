package localcache

import (
	"github.com/coocood/freecache"
	"sync"
)

const (
	minCacheSize = 512 * 1024       // 512kb
	maxCacheSize = 20 * 1024 * 1024 // 20m
)

var (
	once sync.Once
	ins  *LocalCache
)

type LocalCache struct {
	c *freecache.Cache
}

func NewLocalCache(size int) *LocalCache {
	c := freecache.NewCache(size)
	return &LocalCache{c: c}
}

func LoadLocalCache(size int) *LocalCache {
	once.Do(func() {
		ins = NewLocalCache(size)
	})
	return ins
}

func (lc *LocalCache) Get(key string) (value []byte, err error) {
	return lc.c.Get([]byte(key))
}
func (lc *LocalCache) Set(key string, val []byte) error {
	return lc.SetWithExpire(key, val, 0)
}
func (lc *LocalCache) SetWithExpire(key string, val []byte, expireSeconds int) error {
	return lc.c.Set([]byte(key), val, expireSeconds)
}
func (lc *LocalCache) Del(key string) (affected bool) {
	return lc.c.Del([]byte(key))
}
