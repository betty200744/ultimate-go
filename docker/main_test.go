package main

import "testing"

func TestFoo(t *testing.T) {
	t.Run("foo", func(t *testing.T) {
		Foo()
	})
}
