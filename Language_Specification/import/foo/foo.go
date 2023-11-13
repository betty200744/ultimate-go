package foo

import (
	"fmt"

	"ultimate-go/Language_Specification/import/barer"
)

type Foo struct {
	b barer.Barer
}

func (b *Foo) Bar() {
	fmt.Println("this bar")
}
