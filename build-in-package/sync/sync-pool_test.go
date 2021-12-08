package sync

import (
	"bytes"
	"io"
	"os"
	"sync"
	"testing"
	"time"
)

var bufPool = sync.Pool{
	New: func() interface{} {
		// The Pool's New function should generally only return pointer
		// types, since a pointer can be put into the return interface
		// value without an allocation:
		return new(bytes.Buffer)
	},
}

// timeNow is a fake version of time.Now for tests.
func timeNow() time.Time {
	return time.Unix(1136214245, 0)
}

func Log(w io.Writer, key, val string) {
	b := bufPool.Get().(*bytes.Buffer)
	b.Reset()
	// Replace this with time.Now() in a real logger.
	b.WriteString(timeNow().UTC().Format(time.RFC3339))
	b.WriteByte(' ')
	b.WriteString(key)
	b.WriteByte('=')
	b.WriteString(val)
	_, _ = w.Write(b.Bytes())
	bufPool.Put(b)
}

func TestPool(t *testing.T) {
	go Log(os.Stdout, "path", "search1 \n")
	go Log(os.Stdout, "path", "search2 \n")
	go Log(os.Stdout, "path", "search3 \n")
	go Log(os.Stdout, "path", "search4 \n")
	time.Sleep(time.Millisecond)
}
