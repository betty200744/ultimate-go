package main

import (
	"fmt"
	"strings"
)

func CamelCase(s string) string {
	fields := strings.Fields(s)
	ups := []string{}
	for _, value := range fields {
		ups = append(ups, strings.ToUpper(value))
	}
	return strings.Join(ups, " ")
}

/*

 */
func StringsReader() {
	a := "a"
	b := "b"
	c := "a b c"
	m := "i am betty"
	fmt.Println("string compare: a > b or a = b or a < b ?", strings.Compare(b, a))
	fmt.Println("'abc' contain a ? ", strings.Contains(c, "a"))
	fmt.Println("'abc' contain a or b or c ? ", strings.ContainsAny(c, "ade"))
	fmt.Println("'abc' count is ", strings.Count(c, "a"))
	fmt.Println("abc field[0] is ", strings.Fields(c)[0])
	fmt.Println("this is CamelCase", CamelCase(m))
	r := strings.NewReader("this is read")

	fmt.Println(r)
}
