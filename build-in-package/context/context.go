package context

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Pass a context with a timeout to tell a blocking function that it
	// should abandon its work after the timeout elapses.
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	select {
	case a := <-time.After(1 * time.Second):
		fmt.Println("overslept", a)
	case d := <-ctx.Done():
		fmt.Println(ctx.Err(), d) // prints "context deadline exceeded"
	}

}
