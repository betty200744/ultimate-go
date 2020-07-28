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
	fmt.Printf("Builder pattern: this is house,  %#v \n", house)
}
func TestUserBuilder_Build(t *testing.T) {
	ub := &UserBuilder{}
	user := ub.
		Name("Michael Scott").
		Role("manager").
		Build()
	fmt.Printf("Builder pattern: this is user,  %#v \n", user)
}
