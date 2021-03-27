package main

import (
	"fmt"
)

var z int
var w float64

func main() {
	x := 28
	y := 28.34561

	fmt.Println(x)
	fmt.Println(y)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)

	z = 124535
	w = 12.5686

	fmt.Println(z)
	fmt.Println(w)
	fmt.Printf("%T\n", z)
	fmt.Printf("%T\n", w)

}
