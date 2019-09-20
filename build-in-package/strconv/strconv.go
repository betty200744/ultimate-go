package strconv

import (
	"fmt"
	"strconv"
)

func main() {
	// string to int
	i, _ := strconv.Atoi("32")
	i64, _ := strconv.ParseInt("3232", 10, 64)
	i32, _ := strconv.ParseInt("3232", 10, 32)
	fmt.Printf("%T \n", i)
	fmt.Printf("%T \n", int32(i32))
	fmt.Printf("%T \n", i64)

	//int to string
	s := strconv.Itoa(33)
	s1 := int32(33)
	s2 := strconv.FormatInt(int64(222), 10)
	fmt.Printf("%T \n", s)
	fmt.Printf("%T \n", s1)
	fmt.Printf("%T \n", fmt.Sprint(s1))
	fmt.Printf("%T \n", s2)

	// string to utf8, ascii,
	q := strconv.Quote("Hello, 世界")
	q1 := strconv.QuoteToASCII("Hello, 世界")
	fmt.Println(q)
	fmt.Println(q1)

}
