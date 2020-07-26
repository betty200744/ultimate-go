package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

/*
	interface: is a collection of method, and only method
	interface: can be implement
	interface: only provide method name, argument, return type, method defined is declare in type
	interface: type is nil, value is nil when defined, only have type and value when implement
	interface: the value and type is dynamic , is the type and value of implement type
	interface: a struct can implement multi interface
*/

// declaring interface type
type RectInterface interface {
	area() float64
	abc() float64
}
type NameInterface interface {
	getName() string
}

// declaring struct type
type Rect struct {
	width  float64
	height float64
	name   string
}
type Circle struct {
	radius float64
}

// declaring struct method
func (r *Rect) area() float64 {
	return r.width * r.height
}
func (r *Rect) abc() float64 {
	return r.width * r.height
}
func (r *Rect) getName() string {
	return r.name
}
func (c *Circle) area() float64 {
	return c.radius * c.radius
}
func (c *Circle) abc() float64 {
	return c.radius * c.radius
}

func main() {
	// interface  implement
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	var ri RectInterface
	var name NameInterface
	fmt.Printf("type of ri: %T and value is %v \n", ri, ri)
	ri = &Rect{width: 3, height: 3}
	name = &Rect{width: 3, height: 3, name: "betty"}
	fmt.Println("this is RectInterface Area method with Rect struct: ", ri.area(), ri.abc())
	fmt.Println("this is NameInterface Name method with Rect struct: ", name.getName())
	fmt.Printf("type of ri: %T and value is %v \n", ri, ri)

	ri = &Circle{radius: 3}
	fmt.Println("this is RectInterface Area method with Circle struct: ", ri.area())
	fmt.Printf("type of ri: %T and value is %v \n", ri, ri)
	time.Sleep(time.Minute * 10)
}
