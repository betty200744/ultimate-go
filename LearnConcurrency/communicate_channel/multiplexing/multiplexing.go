package multiplexing

import "fmt"

func FanIn(inputs ...<-chan string) <-chan string {
	// multiplexing , 一个chan接收多个chan的数据
	out := make(chan string)
	for _, i2 := range inputs {
		go func(ch <-chan string) {
			for i3 := range ch {
				out <- fmt.Sprintf("receive %s, send pong!", i3)
			}
		}(i2)
	}
	return out
}
