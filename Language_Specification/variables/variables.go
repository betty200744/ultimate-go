package main

import "fmt"

type AA struct {
	Id string `json:"id"`
}

type AI struct {
	Id string `json:"id"`
}

func (ai *AI) Foo() {
	fmt.Println(ai.Id)
}

// a bool is 1 byte, int16 is 2 bytes, float32 is 4 bytes, int64, 8 bytes
var (
	d   bool              // false
	bp  *bool             // nil
	a   int               // 0
	ip  *int              // nil
	b   string            // ""
	c   float64           // 0
	sl  []string          // []
	sa  AA                // {}
	sap *AA               // nil
	ai  AI                // {}
	mm  map[string]string // map[]
	cc  chan string       // nil
)

func main() {
	// Declare and initialize
	// Zero value concept
	fmt.Printf("a int:  %T , value is : [%v]\n", a, a)
	fmt.Printf("b string:  %T , value is : [%v]\n", b, b)
	fmt.Printf("c float64: %T , value is : [%v]\n", c, c)
	fmt.Printf("d bool: %T , value is : [%v]\n", d, d)

	fmt.Printf("this is bp: %T , value is : [%v] \n", bp, bp)
	fmt.Printf("this is ip: %T , value is : [%v] \n", ip, ip)
	fmt.Printf("this is sl: %T , value is : [%v] \n", sl, sl)
	fmt.Printf("this is sa: %T , value is : [%v] \n", sa, sa)
	fmt.Printf("this is sap: %T , value is : [%v] \n", sap, sap)
	fmt.Printf("this is ai: %T , value is : [%v] \n", ai, ai)
	fmt.Printf("this is mm: %T , value is : [%v] \n", mm, mm)
	fmt.Printf("this is mm key: %T , value is : [%v] \n", mm["aaa"], mm["aaa"])
	fmt.Printf("this is cc: %T, , value is : [%v] \n", cc, cc)
	fmt.Printf("this is ai.Id: %T, , value is : [%v] \n", ai.Id, ai.Id)
	fmt.Printf("this is sa.Id: %T, , value is : [%v] \n", sa.Id, sa.Id)

	// Short Variable Declare
	aa := 10
	bb := "hello"
	co := 3.145
	dd := true

	fmt.Printf("aa : %T  , value is : [%v] \n", aa, aa)
	fmt.Printf("bb : %T  , value is : [%v] \n", bb, bb)
	fmt.Printf("co : %T  , value is : [%v] \n", co, co)
	fmt.Printf("dd : %T  , value is : [%v] \n", dd, dd)

}
