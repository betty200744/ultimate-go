package singleton

import (
	"fmt"
	"sync"
)

/*
DB instance – we only want to create only one instance of DB object and that instance will be used throughout the application.
Logger instance – again only one instance of the logger should be created and it should be used throughout the application.
*/

var (
	lock   = &sync.Mutex{}
	once   = &sync.Once{}
	single *Single
)

type Single struct {
}

func GetInstance1() *Single {
	if single == nil {
		lock.Lock()
		defer lock.Unlock()
		single = &Single{}
		fmt.Println("Creating Single Instance Now")
	} else {
		fmt.Println("Single Instance already created")
	}
	return single
}
func GetInstance2() *Single {
	if single == nil {
		once.Do(func() {
			single = &Single{}
		})
		fmt.Println("Creating Single Instance Now")
	} else {
		fmt.Println("Single Instance already created")
	}
	return single
}
