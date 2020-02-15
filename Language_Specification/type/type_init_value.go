package main

import "fmt"

type Aer interface {
	Foo()
}
type AA struct {
	Id string `json:"id"`
}

type AI struct {
	Id string `json:"id"`
}

func (ai *AI) Foo() {
	fmt.Println(ai.Id)
}

var (
	b   bool              // false
	bp  *bool             // nil
	i   int               // 0
	ip  *int              // nil
	s   string            // ""
	sl  []string          // []
	sa  AA                // {}
	sap *AA               // nil
	ai  AI                // {}
	mm  map[string]string // map[]
	cc  chan string       // nil
)

func main() {
	fmt.Println("this is b: ", b)
	fmt.Println("this is bp: ", bp)
	fmt.Println("this is i: ", i)
	fmt.Println("this is ip: ", ip)
	fmt.Println("this is s: ", s)
	fmt.Println("this is sl: ", sl)
	fmt.Println("this is sa: ", sa)
	fmt.Println("this is sap: ", sap)
	fmt.Println("this is ai: ", ai)
	fmt.Println("this is mm: ", mm)
	fmt.Println("this is mm key: ", mm["aaa"])
	fmt.Println("this is cc: ", cc)
	fmt.Println("this is ai.Id: ", ai.Id)
	fmt.Println("this is sa.Id: ", sa.Id)
	fmt.Println("this is sap.Id: ", sap.Id)
}
