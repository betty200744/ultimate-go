package main

import (
	"fmt"
	"github.com/lib/pq"
)

func unique(stringSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range stringSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func UniqueStringArray(stringSlice *pq.StringArray) *pq.StringArray {
	keys := make(map[string]bool)
	list := new(pq.StringArray)
	for _, entry := range *stringSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			*list = append(*list, entry)
		}
	}
	return list
}

func main() {
	stringSlice := []string{"1", "5", "3", "6", "9", "9", "4", "2", "3", "1", "5"}
	stringSlice = append(stringSlice, stringSlice...)
	fmt.Println(stringSlice)
	uniqueSlice := unique(stringSlice)
	fmt.Println(uniqueSlice)

	sa := &pq.StringArray{"a", "b"}

	*sa = append(*sa, *sa...)
	fmt.Println(sa)
	su := UniqueStringArray(sa)
	fmt.Println(su)

}
