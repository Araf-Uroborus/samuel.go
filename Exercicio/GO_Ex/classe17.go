package main

import (
	"fmt"
)

//https://golang.org/ref/spec#Numeric_types

var x int8 = -128
var y int8 = 127

var z uint = 256
var w int = 10000000
var a byte = 254

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Println(z)
	fmt.Printf("%T\n", z)
	fmt.Println(w)
	fmt.Printf("%T\n", w)
	fmt.Println(a)
	fmt.Printf("%T\n", a)
}
