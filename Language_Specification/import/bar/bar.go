package bar

import (
	"fmt"
)

type Bar struct {
}

func (b *Bar) Bar() {
	fmt.Println("this bar")
}
