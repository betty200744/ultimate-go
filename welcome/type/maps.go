package main

import (
	"fmt"
)

func main() {
	// make, init, set, get, len , delete, present, range, keys , with if
	m := make(map[string]string, 0)
	mm := map[string]string{"a": "a", "b": "b"}
	m["a"] = "a"
	m["b"] = "b"
	delete(m, "b")
	_, present := m["b"]
	fmt.Println(m, m["a"], len(m), present, mm)
	if _, prs := m["b"]; !prs {
		fmt.Println("m['b'] not present")
	}
}
