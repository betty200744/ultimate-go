package foo

import (
	"fmt"

	"gobyexample/Language_Specification/import/barer"
)

type Foo struct {
	b barer.Barer
}

func (b *Foo) Bar() {
	fmt.Println("this bar")
}
