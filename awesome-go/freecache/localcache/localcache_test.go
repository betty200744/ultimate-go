package localcache

import (
	"gopkg.in/go-playground/assert.v1"
	"testing"
)

var localcache *LocalCache

func init() {
	localcache = NewLocalCache(512 * 1024)
}

func TestLocalCache_Set(t *testing.T) {
	localcache.Set("a", []byte(`a`))
	res, _ := localcache.Get("a")
	assert.Equal(t, res, []byte(`a`))
}
