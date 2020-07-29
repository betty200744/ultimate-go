package error_handling

import (
	"fmt"
	"testing"
)

func TestWebCall1(t *testing.T) {
	err := WebCall1()
	fmt.Println(err)
}
func TestWebCall2(t *testing.T) {
	err := WebCall2()
	fmt.Println(err)
}

func TestUnmarshal(t *testing.T) {
	var u struct {
		Name int
	}
	err := Unmarshal([]byte(`{"name":"bill"}`), u) // Run with a value and pointer.
	if err != nil {
		// This is a special type assertion that only works on the switch.
		switch e := err.(type) {
		case *UnmarshalTypeError:
			fmt.Printf("UnmarshalTypeError: Value[%s] Type[%v]\n", e.Value, e.Type)
		case *InvalidUnmarshalError:
			fmt.Printf("InvalidUnmarshalError: Type[%v]\n", e.Type)
		default:
			fmt.Println(err)
		}
	}
	err1 := Unmarshal([]byte(``), &u) // Run with a value and pointer.
	if err1 != nil {
		// This is a special type assertion that only works on the switch.
		switch e := err1.(type) {
		case *UnmarshalTypeError:
			fmt.Printf("UnmarshalTypeError: Value[%s] Type[%v]\n", e.Value, e.Type)
		case *InvalidUnmarshalError:
			fmt.Printf("InvalidUnmarshalError: Type[%v]\n", e.Type)
		}
	}

	fmt.Println("Name:", u.Name)
}
