package main

import "fmt"

type Pointer struct {
	Weight int `json:"weight"`
	Height int `json:"height"`
}

// bad code
func IsPositiveBad(p *Pointer) bool {
	if p.Weight > 0 && p.Height > 0 {
		return true
	} else {
		return false
	}
}

// good code, indicate message struct has specific behavior
// good code, 这样的话， IsPositive名字可以用于其他struct, 而不会重名
func (p *Pointer) IsPositive() bool {
	if p.Weight > 0 && p.Height > 0 {
		return true
	} else {
		return false
	}
}

func main() {
	pp := &Pointer{Weight: 1, Height: 1}
	fmt.Println(pp.IsPositive())
	fmt.Println(IsPositiveBad(pp))
}
