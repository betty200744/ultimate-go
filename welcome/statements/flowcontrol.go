package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	sum := 0
	summ := 1
	summm := 1
	//for loop
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
	//	for loop 2 , no init and no post
	for summ < 100 {
		summ += summ
	}
	fmt.Println(summ)
	// for is while , so is for {} , is a forever loops
	for summm < 100 {
		summm += summm
	}
	fmt.Println(summm)

	// if
	if sum != 0 {
		fmt.Println("sum is not 1")
	}

	if v := 3; v < sum {
		fmt.Println(v)
	} else {
		fmt.Println(sum)
	}

	// switch
	os := runtime.GOOS
	switch os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("linux")
	default:
		fmt.Printf("%s", os)
	}
	// switch true , 相对于if else， 即简洁的if else
	h := time.Now().Hour()
	today := time.Now().Weekday()
	defer fmt.Println(today) // defer , 即延迟执行， 即最后执行
	switch {
	case h < 12:
		fmt.Println("good morning")
	case h < 17:
		fmt.Println("good afternoon")
	default:
		fmt.Println("good eventing")
	}

	for ii := 0; ii < 5; ii++ {
		defer fmt.Println(ii)
	}

}
