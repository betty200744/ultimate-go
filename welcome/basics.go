// package
package main

// import package
import (
	"fmt"
	"math"
	"math/rand" //package name is the last name of the path
	"time"
)

var c, python, java = true, 2, "java test" // variables with init value ,auto type
var node, goo int                          // variables , shard type

func Add(x, y int) int { // func take 0 or more arguments and type int , and return int , export
	return x + y
}

func swap(x, y string) (string, string) { // two argument shard type string , and return multi string
	return y, x
}

func split(sum int) (x, y int) {
	x = sum + 5
	y = sum - x
	// naked return , auto return x, y
	return
}

func Joins(a string, paths ...string) { // 同js
	fmt.Println(a)
	fmt.Println(paths)
}

// basic type
var (
	IsPay  bool   = false
	MaxInt uint64 = 1<<64 - 1
	i      int
	f      float32
	s      string
)

func main() {
	// type bool
	var b bool
	// type byte array
	barr := []byte(`{"name"}`)
	fmt.Println(barr)

	fmt.Printf("hello world")
	fmt.Println(time.Now())
	fmt.Println("this is rand int", rand.Intn(10))
	fmt.Println(math.Sqrt(7))
	fmt.Println(math.Pi) // math export name , 大写默认被export
	fmt.Println(Add(1, 2))
	k := 3 // shot variable declarations with implicit type, 省略var， 隐式类型int， 类型不可变

	y, x := swap("x", "y")
	fmt.Println(swap("a", "b"))
	fmt.Println(y, x)
	fmt.Println(split(4))
	fmt.Println(i, c, python, java, node, goo)
	fmt.Println(k)
	fmt.Println(IsPay, MaxInt, i, f, b, s) // basic type

	// type conversions, 类型转换
	fmt.Println(float32(i))
	fmt.Println(string(i))
	fmt.Println(uint(i))

	//	type interface
	j := k // j is implicit int type
	fmt.Println(j)
	//	const variables
	const World = "world"
	// passing arguments  to ...parameters
	Joins("a", "b", "c")
}
