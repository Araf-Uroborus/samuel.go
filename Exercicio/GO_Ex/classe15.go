package main

import (
	"fmt"
)

var x bool

func main() {
	a := 7
	b := 25

	x = (a > b)
	fmt.Println(x)

	x = (a < b)
	fmt.Println(x)

	x = (a == b)
	fmt.Println(x)

	x = (a != b)
	fmt.Println(x)

	x = (a >= b)
	fmt.Println(x)

	x = (a <= b)
	fmt.Println(x)
}
