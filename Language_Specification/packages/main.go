package main

import (
	"fmt"                                             // import package
	"ultimate-go/Language_Specification/packages/app" // import nested package
)

/*
	package name a.b.c
	main.go is entry
    package compilation: go build -o main ./main.go
    package install: go install ./main.go
	$GOPATH/bin/main
	package init function: first import , first init
    package init function: init first , then main
    package init function, main job is initialize global variables
	import package: use package alias when import package with the some name
	import package: use underscore _ when only initialize package
	import package: use dot . , then all export member can call from local file, not very fond of it
    imported package: is initialized only once per package
*/

/*
Program execution order
go run *.go
├── Main package is executed // 先main
├── All imported packages are initialized // 在import , 注意是递归
|  ├── All imported packages are initialized (recursive definition)
|  ├── All global variables are initialized // 再变量初始化
|  └── init functions are called in lexical file name order // 再init函数初始化
└── Main package is initialized  //再main函数初始化
   ├── All global variables are initialized  //再main变量初始化
   └── init functions are called in lexical file name order //再main.go init函数初始化
*/

func main() {
	fmt.Println("this is main")
	app.App()
}
