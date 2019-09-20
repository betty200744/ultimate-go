package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(strings.NewReader("abc"))
	for scanner.Scan() {
		fmt.Println(string(scanner.Bytes()))
	}
}
