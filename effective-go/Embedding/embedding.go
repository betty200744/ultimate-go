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
type ReadWriterN struct {
	N int `json:"n"`
}

func (rw *ReadWriterN) Read(p []byte) (n int, err error) {
	return rw.N, nil
}
func (rw *ReadWriterN) Write(p []byte) (n int, err error) {
	rw.N = 3
	return rw.N, nil
}

// embed struct
type Job struct {
	Id   int64 `json:"id"`
	Type int32 `json:"type"`
}
type User struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Job  Job
}

func main() {
	// ReadWriterN implements embed interface ReaderWriter
	rwN := ReadWriterN{N: 2}
	rwN.Read([]byte{})
	rwN.Write([]byte{})
	fmt.Println(rwN.N)
	// embed struct
	job := &User{Id: int64(0), Name: "u1", Job: Job{Type: int32(2)}}
	fmt.Println(job)
}
