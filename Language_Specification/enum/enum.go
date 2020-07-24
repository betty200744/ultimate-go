package main

import (
	"fmt"
)

type ACTION_TYPE int32

const (
	ACTION_TYPE_ACTION_EMPTY    ACTION_TYPE = 0
	ACTION_TYPE_ACTION_PICKUP   ACTION_TYPE = 1
	ACTION_TYPE_ACTION_PUTDOWN  ACTION_TYPE = 2
	ACTION_TYPE_ACTION_TOUCH    ACTION_TYPE = 3
	ACTION_TYPE_ACTION_APPROACH ACTION_TYPE = 4
)

var ACTION_TYPE_name = map[int32]string{
	0: "ACTION_EMPTY",
	1: "ACTION_PICKUP",
	2: "ACTION_PUTDOWN",
	3: "ACTION_TOUCH",
	4: "ACTION_APPROACH",
}

var ACTION_TYPE_value = map[string]int32{
	"ACTION_EMPTY":    0,
	"ACTION_PICKUP":   1,
	"ACTION_PUTDOWN":  2,
	"ACTION_TOUCH":    3,
	"ACTION_APPROACH": 4,
}

func (x ACTION_TYPE) String() string {
	return ACTION_TYPE_name[int32(x)]
}

func main() {
	fmt.Println("by const:", ACTION_TYPE_ACTION_PICKUP)
	fmt.Println("by key:", ACTION_TYPE(3))
	fmt.Println("by value", ACTION_TYPE_value["ACTION_PICKUP"])
	fmt.Printf("enum:  %d \n", ACTION_TYPE_ACTION_PICKUP)
	fmt.Printf("enum string:  %v \n", ACTION_TYPE_ACTION_PICKUP)
}
