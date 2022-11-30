package main

import (
	"fmt"
	_ "net/http/pprof"
)

/*
	interface: 最简单的interface就是empty interface{}, 所有类型是继承
	interface: error也是一个interface, 含Error方法
	interface: is a collection of method, and only method
	interface: can be implement
	interface: only provide method name, argument, return type, method defined is declare in type
	interface: type is nil, value is nil when defined, only have type and value when implement
	interface: the value and type is dynamic , is the type and value of implement type
	interface: a struct can implement multi interface
*/

// Declaring interface type
type RectInterface interface {
	area() float64
	abc() float64
}
type NameInterface interface {
	getName() string
}

// Declaring struct type
type Rect struct {
	width  float64
	height float64
	name   string
}
type Circle struct {
	radius float64
}

// Declaring Rect struct method
func (r *Rect) area() float64 {
	return r.width * r.height
}
func (r *Rect) abc() float64 {
	return r.width * r.height
}
func (r *Rect) getName() string {
	return r.name
}

// Declaring Circle struct method
func (c *Circle) area() float64 {
	return c.radius * c.radius
}
func (c *Circle) abc() float64 {
	return c.radius * c.radius
}

// --------------------
// Polymorphic function
// --------------------
func GetArea(r RectInterface) float64 {
	fmt.Printf("Polymorphic function, Type: %T, area: %v \n", r, r.area())
	return r.area()
}

// 充值档位
type ChargeInfo struct {
	Id            int   `json:"id"`
	Price         int   `json:"price"`          // 档位金额，单位分
	DiscountPrice int   `json:"discount_price"` // 支付金额，单位分
	Duration      int   `json:"duration"`       // 产生时间，单位秒
	ExtraDuration int   `json:"extra_duration"` // 额外赠送时间，单位秒
	Status        int   `json:"status"`         // 0：正常； 1：禁用
	Ctime         int64 `json:"ctime,omitempty"`
	Mtime         int64 `json:"mtime,omitempty"`
}

func test() {
	info := &ChargeInfo{
		Id:            1,
		Price:         1,
		DiscountPrice: 1,
		Duration:      1,
		ExtraDuration: 1,
		Status:        1,
		Ctime:         1,
		Mtime:         1,
	}
	i := []interface{}{info}
	fmt.Println(i[0])
}

type FooInterface interface {
	bar() // method receiver is value then interface is value , method receiver is pointer then interface is pointer
}

//Won't throw a compile error. Because the bar method belongs to Foo struct.
type Foo1 struct {
}

func (f Foo1) bar() {
	panic("implement me")
}

//Won't throw a compile error. Because the bar method belongs to Foo struct.
//Will throw a error. Because the bar method is bind to &Foo. And when compile checks whether Foo confirms to Foo interface, it will only check the required method existence.
type Foo2 struct {
}

func (f *Foo2) bar() {
	panic("implement me")
}

func funcName(foo FooInterface) {
	fmt.Println(foo)
}

func main() {
	test()
	r := &Rect{width: 3, height: 3, name: "betty"}
	fmt.Printf("this is Rect, Type: %T, area:  %v, abc: %v \n, getName: %v \n", r, r.area(), r.abc(), r.getName())
	GetArea(r)

	var ri RectInterface
	ri = &Rect{width: 3, height: 3}
	GetArea(ri)
	fmt.Printf("this is RectInterface, Rect Struct,  Type: %T, area: %v, abc: %v , no getName \n", ri, ri.area(), ri.abc())

	ri = &Circle{radius: 3}
	GetArea(ri)
	fmt.Printf("this is RectInterface, Circle struct, Type: %T, area: %v, abc: %v , no getName \n", ri, ri.area(), ri.abc())

	var ni NameInterface
	ni = &Rect{width: 3, height: 3, name: "betty"}
	fmt.Printf("this is NameInterface, Rect Struct,  Type: %T, getName: %v , no area, no abc \n", ni, ni.getName())

	// --------------
	// Interface  Type assertions
	// --------------
	var i interface{} = "hello"
	fmt.Printf("i am string: %T, %v \n", i, i)
	is, ok := i.(string)
	fmt.Printf("i am string after type assertions: %T, %v, %v\n", is, is, ok)
	i = 1
	it, _ := i.(int)
	fmt.Printf("i am int after type assertions:%T, %v \n", it, it)

	funcName(Foo1{})  // Foo1.bar receiver is value , then Foo1 need value
	funcName(&Foo2{}) // Foo2.bar receiver is pointer, then Foo2 need pointer
}
