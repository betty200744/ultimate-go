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

// good code
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
