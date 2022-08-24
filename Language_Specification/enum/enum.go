package main

import (
	"fmt"
)

type ACTION int32

const (
	EMPTY    ACTION = iota
	PICKUP
	PUTDOWN
	TOUCH
	APPROACH
)

var ACTION_name = map[ACTION]string{
	EMPTY:    "EMPTY",
	PICKUP:   "PICKUP",
	PUTDOWN:  "PUTDOWN",
	TOUCH:    "TOUCH",
	APPROACH: "APPROACH",
}

var ACTION_value = map[string]ACTION{
	"EMPTY":    EMPTY,
	"PICKUP":   PICKUP,
	"PUTDOWN":  PUTDOWN,
	"TOUCH":    TOUCH,
	"APPROACH": APPROACH,
}

func (x ACTION) String() string {
	return ACTION_name[x]
}

func main() {
	fmt.Println("by const:", PICKUP)
	fmt.Println("by key:", ACTION(3))
	fmt.Println("by value", ACTION_value["ACTION_PICKUP"])
	fmt.Printf("enum:  %d \n", PICKUP)
	fmt.Printf("enum string:  %s \n", PICKUP)
}
