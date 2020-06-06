package main

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
)

func main() {
	// 注意， NewNode，传參相同， 则生产的id也是相同的
	node0, err := snowflake.NewNode(0)
	node1, err := snowflake.NewNode(1)
	if err != nil {
		panic(err)
	}
	for i := 0; i < 2; i++ {
		fmt.Println(node0.Generate())
	}
	for i := 0; i < 2; i++ {
		fmt.Println(node1.Generate())
	}

}
