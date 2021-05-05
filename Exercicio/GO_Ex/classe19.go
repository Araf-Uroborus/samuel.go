package main

import (
	"fmt"
	"runtime"
)

//https://golang.org/pkg/runtime/

func main() {
	fmt.Println(runtime.GOOS)
	fmt.Print(runtime.GOARCH)
}
