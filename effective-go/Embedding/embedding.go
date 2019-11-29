package main

import "fmt"

// embed interface
type Reader interface {
	Read(p []byte) (n int, err error)
}
type Writer interface {
	Write(p []byte) (n int, err error)
}
type ReaderWriter interface {
	Reader
	Writer
}

// embed struct
type ReadWriter struct {
	reader Reader
	writer Writer
}

func (rw *ReadWriter) Read(p []byte) (n int, err error) {
	return rw.reader.Read(p)
}

// embed struct
type User struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}
type Job struct {
	Id   int64 `json:"id"`
	Type int32 `json:"type"`
	User User  `json:"user"`
}

func main() {
	fmt.Println("embedding~")
}
