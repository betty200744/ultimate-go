package time

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	go func() {
		for t := range ticker.C {
			fmt.Println(t)
		}
	}()
	time.Sleep(time.Second * 5)
}
