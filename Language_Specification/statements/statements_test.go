package statements

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestSwitch(t *testing.T) {
	// switch variable
	os := runtime.GOOS
	switch os {
	case "darwin":
		fmt.Println("OS X")
		fallthrough
	case "linux":
		fmt.Println("linux")
	default:
		fmt.Printf("%s", os)
	}
	// switch true
	h := time.Now().Hour()
	switch {
	case h < 12:
		fmt.Println("good morning")
	case h < 17:
		fmt.Println("good afternoon")
	default:
		fmt.Println("good eventing")
	}
}
