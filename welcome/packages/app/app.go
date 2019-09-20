package app

import "fmt"

/*
	app.go is archive
    go build -o app ./app.go
    go install ./app.go
    $GOPATH/pkg/darwin_amd64/gobyexample/welcome/packages/app.a
	import package
	export Variable
*/

const APPNAME = "whale"

func App() {
	fmt.Println("this is App")
}
