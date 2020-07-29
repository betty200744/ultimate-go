package main

import (
	"fmt"
	"math/rand"
)

// car represents something you drive.
type car struct{}

// String implements the fmt.Stringer interface.
func (car) String() string {
	return "Vroom!"
}

// cloud represents somewhere you store information.
type cloud struct{}

// String implements the fmt.Stringer interface.
func (cloud) String() string {
	return "Big Data!"
}
func main() {
	mvs := []fmt.Stringer{
		car{},
		cloud{},
	}
	for i := 0; i < 10; i++ {
		// Choose a random number from 0 to 1.
		rn := rand.Intn(2)
		if v, ok := mvs[rn].(cloud); ok {
			fmt.Println("Got Lucky:", v)
		} else {
			v, _ := mvs[rn].(car)
			fmt.Println("Got Unlucky", v)
		}
	}
}
