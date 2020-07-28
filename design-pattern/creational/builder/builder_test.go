package builder

import (
	"fmt"
	"testing"
)

func Test_getBuilder(t *testing.T) {
	builder := getBuilder()
	builder.setWindowType("Wooden")
	builder.setDoorType("Wooden")
	builder.setNumFloor(3)
	house := builder.getHouse()
	fmt.Printf("this is house: %#v \n", house)
}
