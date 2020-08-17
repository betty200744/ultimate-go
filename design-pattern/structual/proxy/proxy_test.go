package proxy

import (
	"gopkg.in/go-playground/assert.v1"
	"testing"
)

func TestProxy_Do(t *testing.T) {
	realSubject := new(RealSubject)
	proxy := new(Proxy)
	proxy.real = realSubject
	res := proxy.Do()
	assert.Equal(t, res, "real, proxy")
}
