package main

import "fmt"

// declare a named type
type ByteSlice []byte

// method, can be defined for any named type, except pointer and interface
// method, receiver can pointer and values
// method, value receiver, call by value, 不改变receiver本身，需要改变请return, 不需要改变本身的情况下可用
// method, value receiver, call by pointer, 不改变receiver本身，需要改变请return, 不需要改变本身的情况下可用
// method, pointer receiver, call by value, 改变receiver本身
// method, pointer receiver, call by pointer, 改变receiver本身
// method, 传value无side effect, 传pointer有
// method, You should use pointers if you need to mutate the content of the struct, and values otherwise.
func (p ByteSlice) Append(data []byte) []byte { // this is value methods
	p = append(p, data...)
	return p
}

// better
func (p *ByteSlice) Append3(data []byte) { // this is pointer methods
	*p = append(*p, data...)
}

func main() {
	b := ByteSlice{}
	b3 := new(ByteSlice)

	b.Append([]byte(`a`))    // value receive, invoked on value, 不改变value本身
	(&b).Append([]byte(`a`)) // value receive, invoked on pointer， 不改变value本身
	fmt.Println("value receive, b is empty: ", string(b))
	fmt.Println("value receive, return is change a:", string(b.Append([]byte(`a`))))
	fmt.Println("value receive, invoked on pointer, b is still empty: ", string(b), "return is change a: ", string((&b).Append([]byte(`a`))))

	b.Append3([]byte(`c`))  // pointer receive, invoked on value， 改变value本身
	b3.Append3([]byte(`c`)) // pointer receive, invoked on pointer， 改变value本身
	fmt.Println("pointer receive, invoked on pointer, b3 is change c: ", string(*b3))
	fmt.Println("pointer receive, invoked on value, b is change c: ", string(b))
}
