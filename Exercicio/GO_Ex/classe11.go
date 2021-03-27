package main

import (
	"fmt"
)

var a int

type cerveja int

var b cerveja

func main() {
	a = 42
	fmt.Println(a)
	fmt.Printf("%T", a)

	b = 43
	fmt.Println(b)
	fmt.Printf("%T", b)
}
