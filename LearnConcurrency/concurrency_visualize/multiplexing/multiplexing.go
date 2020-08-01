package multiplexing

import "fmt"

func FanIn(inputs ...<-chan string) <-chan string {
	// multiplexing , 一个chan接收多个chan的数据
	out := make(chan string)
	for _, i2 := range inputs {
		go func(ch <-chan string) {
			for {
				out <- fmt.Sprintf("receive %s, send pong!", <-ch)
			}
		}(i2)
	}
	return out
}
