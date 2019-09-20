package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(strings.NewReader("adfadafdafdafdasfdsabc"))
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
